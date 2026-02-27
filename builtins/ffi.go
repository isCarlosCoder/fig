package builtins

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/iscarloscoder/fig/environment"
	"github.com/pelletier/go-toml/v2"
)

var cbMu sync.Mutex
var nextCb uint64
var callbacks = map[string]environment.Value{}

// struct schemas: name -> ordered list of fields (name,type)
var structSchemasMu sync.Mutex
var structSchemas = map[string][]struct {
	Field string
	Type  string
}{}

// annotation keys for wrapper objects
const (
	structDefKey      = "__ffi_struct_def__"
	structInstanceKey = "__struct__"
	// when a struct is returned from the helper we actually keep the C
	// pointer inside the helper; the object delivered to Fig only contains
	// this id so that subsequent calls can reuse it without knowing the
	// raw pointer value.
	ptrInstanceKey = "__ptrid__"
)

// helper that mirrors ffi-gen's IsSupportedType (runtime version)
func isSupportedType(t string) bool {
	// accept common synonyms used throughout code
	switch t {
	case "int", "integer", "double", "float", "number", "string", "void":
		return true
	}
	if strings.HasPrefix(t, "struct:") && len(t) > 7 {
		return true
	}
	return false
}

// symbol argument type hints: symbolId -> []argType
var symbolArgTypesMu sync.Mutex
var symbolArgTypes = map[string][]string{}

// store declared return types for symbols (used during call-time coercion)
var symbolRetTypesMu sync.Mutex
var symbolRetTypes = map[string]string{}

// ResetFfiState clears all FFI global state (struct schemas, symbol arg types, callbacks).
// Useful between test runs to avoid cross-test contamination.
func ResetFfiState() {
	structSchemasMu.Lock()
	structSchemas = map[string][]struct {
		Field string
		Type  string
	}{}
	structSchemasMu.Unlock()

	symbolArgTypesMu.Lock()
	symbolArgTypes = map[string][]string{}
	symbolArgTypesMu.Unlock()

	cbMu.Lock()
	callbacks = map[string]environment.Value{}
	atomic.StoreUint64(&nextCb, 0)
	cbMu.Unlock()
}

// expandStructFields recursively expands a struct value's fields into flat arg lists.
// When a field has type "struct:X", it recurses into that nested struct.
func expandStructFields(schemaName string, val environment.Value, mar *[]interface{}, types *[]string) error {
	// pointer markers cannot be expanded
	if val.Type == environment.ObjectType && val.Obj != nil {
		if _, ok := val.Obj.Entries[ptrInstanceKey]; ok {
			return fmt.Errorf("cannot expand pointer marker for struct %s", schemaName)
		}
	}
	structSchemasMu.Lock()
	fields, has := structSchemas[schemaName]
	structSchemasMu.Unlock()
	if !has {
		return fmt.Errorf("call: unknown struct schema: %s", schemaName)
	}
	for _, f := range fields {
		fv, ok := val.Obj.Entries[f.Field]
		if !ok {
			return fmt.Errorf("missing struct field: %s.%s", schemaName, f.Field)
		}
		if len(f.Type) > 7 && f.Type[:7] == "struct:" {
			// nested struct: recurse
			nestedName := f.Type[7:]
			if fv.Type != environment.ObjectType || fv.Obj == nil {
				return fmt.Errorf("call: expected struct object for nested %s.%s", schemaName, f.Field)
			}
			if err := expandStructFields(nestedName, fv, mar, types); err != nil {
				return err
			}
		} else {
			x, err := fromEnvironmentValue(fv)
			if err != nil {
				return fmt.Errorf("cannot convert struct field %s.%s: %v", schemaName, f.Field, err)
			}
			*mar = append(*mar, x)
			*types = append(*types, f.Type)
		}
	}
	return nil
}

func findProjectTomlFrom(startDir string) (string, error) {
	dir := startDir
	for {
		candidate := filepath.Join(dir, "fig.toml")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("fig.toml not found")
		}
		dir = parent
	}
}

func readFfiEnabled() (bool, string, error) {
	// find project toml relative to cwd
	cwd, _ := os.Getwd()
	p, err := findProjectTomlFrom(cwd)
	if err != nil {
		return false, "", nil
	}
	var cfg struct {
		Ffi struct {
			Enabled bool   `toml:"enabled"`
			Helper  string `toml:"helper"`
		} `toml:"ffi"`
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return false, "", err
	}
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return false, "", err
	}
	return cfg.Ffi.Enabled, cfg.Ffi.Helper, nil
}

func readFfiConfig() (bool, string, string, int, error) {
	// find project toml relative to cwd
	cwd, _ := os.Getwd()
	p, err := findProjectTomlFrom(cwd)
	if err != nil {
		return false, "", "", 0, nil
	}
	var cfg struct {
		Ffi struct {
			Enabled     bool   `toml:"enabled"`
			Helper      string `toml:"helper"`
			CallTimeout *int   `toml:"call_timeout"` // milliseconds: 0 = unlimited, omitted = default (3000)
		} `toml:"ffi"`
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return false, "", "", 0, err
	}
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return false, "", "", 0, err
	}
	projectRoot := filepath.Dir(p)
	var callTimeoutMs int = -1 // sentinel meaning "not specified, use runtime default"
	if cfg.Ffi.CallTimeout != nil {
		callTimeoutMs = *cfg.Ffi.CallTimeout
	}
	return cfg.Ffi.Enabled, cfg.Ffi.Helper, projectRoot, callTimeoutMs, nil
}

// toEnvironmentValue converts a JSON-like response from the helper into an environment.Value
func toEnvironmentValue(v interface{}) (environment.Value, error) {
	if v == nil {
		return environment.NewNil(), nil
	}
	switch t := v.(type) {
	case float64:
		return environment.NewNumber(t), nil
	case string:
		return environment.NewString(t), nil
	case bool:
		return environment.NewBool(t), nil
	case []interface{}:
		res := make([]environment.Value, 0, len(t))
		for _, el := range t {
			val, err := toEnvironmentValue(el)
			if err != nil {
				return environment.NewNil(), err
			}
			res = append(res, val)
		}
		return environment.NewArray(res), nil
	case map[string]interface{}:
		// detect bytes object
		m := t
		if b64, ok := m["__bytes__"].(string); ok {
			obj := map[string]environment.Value{"__bytes__": environment.NewString(b64)}
			return environment.NewObject(obj, []string{"__bytes__"}), nil
		}
		if cbid, ok := m["__cb__"].(string); ok {
			obj := map[string]environment.Value{"__cb__": environment.NewString(cbid)}
			return environment.NewObject(obj, []string{"__cb__"}), nil
		}
		// generic object
		entries := make(map[string]environment.Value)
		keys := make([]string, 0, len(m))
		for k, vv := range m {
			val, err := toEnvironmentValue(vv)
			if err != nil {
				return environment.NewNil(), err
			}
			entries[k] = val
			keys = append(keys, k)
		}
		return environment.NewObject(entries, keys), nil

	default:
		return environment.NewNil(), fmt.Errorf("unsupported response type: %T", v)
	}
}

// fromEnvironmentValue converts environment.Value into a JSON-serializable Go value
func fromEnvironmentValue(v environment.Value) (interface{}, error) {
	switch v.Type {
	case environment.NilType:
		return nil, nil
	case environment.NumberType:
		return v.Num, nil
	case environment.StringType:
		return v.Str, nil
	case environment.BooleanType:
		return v.Bool, nil
	case environment.ArrayType:
		if v.Arr == nil {
			return []interface{}{}, nil
		}
		res := make([]interface{}, 0, len(*v.Arr))
		for _, el := range *v.Arr {
			x, err := fromEnvironmentValue(el)
			if err != nil {
				return nil, err
			}
			res = append(res, x)
		}
		return res, nil
	case environment.ObjectType:
		if v.Obj == nil {
			return map[string]interface{}{}, nil
		}
		// special bytes or cb object
		if b, ok := v.Obj.Entries["__bytes__"]; ok && b.Type == environment.StringType {
			return map[string]interface{}{"__bytes__": b.Str}, nil
		}
		if cb, ok := v.Obj.Entries["__cb__"]; ok && cb.Type == environment.StringType {
			return map[string]interface{}{"__cb__": cb.Str}, nil
		}
		m := make(map[string]interface{})
		for k, vv := range v.Obj.Entries {
			x, err := fromEnvironmentValue(vv)
			if err != nil {
				return nil, err
			}
			m[k] = x
		}
		return m, nil
	default:
		return nil, fmt.Errorf("unsupported return type for callback: %v", v.Type)
	}
}

// runCallback executes a registered callback by id with provided json-serializable args
func runCallback(cbId string, args []interface{}) (interface{}, error) {
	cbMu.Lock()
	cb, ok := callbacks[cbId]
	cbMu.Unlock()
	if !ok {
		return nil, fmt.Errorf("unknown callback id: %s", cbId)
	}
	// convert args
	var envArgs []environment.Value
	for _, a := range args {
		v, err := toEnvironmentValue(a)
		if err != nil {
			return nil, fmt.Errorf("cannot convert arg: %v", err)
		}
		envArgs = append(envArgs, v)
	}
	// invoke the function with timeout and panic safety
	type result struct {
		v   environment.Value
		err error
	}
	ch := make(chan result, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				ch <- result{environment.NewNil(), fmt.Errorf("panic in callback: %v", r)}
			}
		}()
		res, err := invokeFn(cb, envArgs)
		ch <- result{res, err}
	}()
	select {
	case r := <-ch:
		if r.err != nil {
			return nil, r.err
		}
		return fromEnvironmentValue(r.v)
	case <-time.After(2 * time.Second):
		return nil, fmt.Errorf("callback timed out")
	}
}

func init() {
	register(newModule("ffi",
		fn("enabled", func(args []environment.Value) (environment.Value, error) {
			en, _, err := readFfiEnabled()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			return environment.NewBool(en), nil
		}),

		// lib_ext() -> returns the shared library extension for the current OS
		fn("lib_ext", func(args []environment.Value) (environment.Value, error) {
			return environment.NewString(LibExt()), nil
		}),

		// lib_name(base) -> returns a conventional shared library filename
		fn("lib_name", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("lib_name(base) expects 1 argument")
			}
			base, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("lib_name(base) expects a string")
			}
			return environment.NewString(LibName(base)), nil
		}),

		// sandbox_status() -> returns object with current sandbox counters and config
		fn("sandbox_status", func(args []environment.Value) (environment.Value, error) {
			sbCfg, projectRoot, _ := readSandboxConfig()
			entries := map[string]environment.Value{
				"max_memory_mb":   environment.NewNumber(float64(sbCfg.MaxMemoryMB)),
				"max_cpu_seconds": environment.NewNumber(float64(sbCfg.MaxCPUSeconds)),
				"max_libs":        environment.NewNumber(float64(sbCfg.MaxLibs)),
				"max_allocs":      environment.NewNumber(float64(sbCfg.MaxAllocs)),
				"max_restarts":    environment.NewNumber(float64(sbCfg.MaxRestarts)),
				"loaded_libs":     environment.NewNumber(float64(GetLoadCountValue(projectRoot))),
				"live_allocs":     environment.NewNumber(float64(GetAllocCountValue(projectRoot))),
				"restarts":        environment.NewNumber(float64(GetRestartCountValue(projectRoot))),
			}
			keys := []string{"max_memory_mb", "max_cpu_seconds", "max_libs", "max_allocs", "max_restarts", "loaded_libs", "live_allocs", "restarts"}
			return environment.NewObject(entries, keys), nil
		}),

		// register_callback(fn) -> returns object {"__cb__": id}
		fn("register_callback", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("register_callback(fn) expects 1 arg")
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("register_callback: arg must be a function")
			}
			id := fmt.Sprintf("cb-%d", atomic.AddUint64(&nextCb, 1))
			cbMu.Lock()
			callbacks[id] = args[0]
			cbMu.Unlock()
			obj := map[string]environment.Value{"__cb__": environment.NewString(id)}
			return environment.NewObject(obj, []string{"__cb__"}), nil
		}),

		// unregister_callback(cbObj) -> removes registered callback
		fn("unregister_callback", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("unregister_callback expects 1 arg")
			}
			var id string
			if args[0].Type == environment.StringType {
				id = args[0].Str
			} else if args[0].Type == environment.ObjectType && args[0].Obj != nil {
				if v, ok := args[0].Obj.Entries["__cb__"]; ok && v.Type == environment.StringType {
					id = v.Str
				}
			}
			if id == "" {
				return environment.NewNil(), fmt.Errorf("unregister_callback: expected callback id or object")
			}
			cbMu.Lock()
			delete(callbacks, id)
			cbMu.Unlock()
			return environment.NewNil(), nil
		}),

		fn("load", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("load(path) expects 1 argument")
			}
			path, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("load(path) expects a string")
			}
			en, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			if !en {
				return environment.NewNil(), fmt.Errorf("FFI is not enabled for this project; run 'fig setup-ffi' and enable in fig.toml")
			}
			if helper == "" {
				return environment.NewNil(), fmt.Errorf("no ffi.helper configured in fig.toml; run 'fig setup-ffi'")
			}
			if _, statErr := os.Stat(helper); statErr != nil {
				return environment.NewNil(), fmt.Errorf("ffi helper not found at '%s' (run 'fig setup-ffi')", helper)
			}
			// sandbox: check max_libs limit
			sbCfg, _, _ := readSandboxConfig()
			if sbCfg.MaxLibs > 0 {
				cur := GetLoadCountValue(projectRoot)
				if cur >= int64(sbCfg.MaxLibs) {
					return environment.NewNil(), fmt.Errorf("sandbox: max loaded libraries limit reached (%d)", sbCfg.MaxLibs)
				}
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			handle, err := hc.Load(path)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("load failed: %v; stderr: %s", err, hc.stderrString())
			}
			IncrementLoadCount(projectRoot)
			return environment.NewString(handle), nil
		}),

		fn("sym", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 && len(args) != 3 && len(args) != 4 {
				return environment.NewNil(), fmt.Errorf("sym(handle, name[, rtype[, argTypes]]) expects 2..4 args")
			}
			handle, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("sym: first arg must be handle string")
			}
			name, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("sym: second arg must be symbol name")
			}
			rtype := "int"
			if len(args) >= 3 {
				rtype, _ = args[2].AsString()
			}
			var argTypes []string
			if len(args) == 4 {
				// expect array of strings or struct descriptor objects
				if args[3].Type != environment.ArrayType || args[3].Arr == nil {
					return environment.NewNil(), fmt.Errorf("sym: fourth arg must be array")
				}
				for _, v := range *args[3].Arr {
					switch v.Type {
					case environment.StringType:
						argTypes = append(argTypes, v.Str)
					case environment.ObjectType:
						// wrapper descriptor? look for structDefKey or name
						if def, ok := v.Obj.Entries[structDefKey]; ok && def.Type == environment.StringType {
							argTypes = append(argTypes, "struct:"+def.Str)
						} else if nm, ok := v.Obj.Entries["name"]; ok && nm.Type == environment.StringType {
							argTypes = append(argTypes, "struct:"+nm.Str)
						} else {
							return environment.NewNil(), fmt.Errorf("sym: argTypes array contains unsupported object")
						}
					default:
						return environment.NewNil(), fmt.Errorf("sym: argTypes must be strings or struct descriptors")
					}
				}
			}
			en, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			if !en {
				return environment.NewNil(), fmt.Errorf("FFI is not enabled for this project; run 'fig setup-ffi' and enable in fig.toml")
			}
			if helper == "" {
				return environment.NewNil(), fmt.Errorf("no ffi.helper configured in fig.toml; run 'fig setup-ffi'")
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			sym, err := hc.Sym(handle, name, rtype)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("sym failed: %v; stderr: %s", err, hc.stderrString())
			}
			// store argTypes for this symbol id if provided
			if len(argTypes) > 0 {
				symbolArgTypesMu.Lock()
				symbolArgTypes[sym] = argTypes
				symbolArgTypesMu.Unlock()
			}
			// also remember declared return type to help coercion in call()
			symbolRetTypesMu.Lock()
			symbolRetTypes[sym] = rtype
			symbolRetTypesMu.Unlock()
			return environment.NewString(sym), nil
		}),

		fn("call", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("call(symbolId, ...args) expects symbol id and zero or more args")
			}
			// first arg: symbol id string
			symbolId, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("call(symbolId, ...): first arg must be symbol id string")
			}
			en, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			if !en {
				return environment.NewNil(), fmt.Errorf("FFI is not enabled for this project; run 'fig setup-ffi' and enable in fig.toml")
			}
			if helper == "" {
				return environment.NewNil(), fmt.Errorf("no ffi.helper configured in fig.toml; run 'fig setup-ffi'")
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			// apply configurable call timeout if provided in fig.toml
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			// marshal args after symbol id
			var mar []interface{}
			var expandedArgTypes []string
			// if argTypes declared for this symbol, expand accordingly
			symbolArgTypesMu.Lock()
			argTypesForSym, hasArgTypes := symbolArgTypes[symbolId]
			symbolArgTypesMu.Unlock()
			if hasArgTypes {
				// consume input args by index
				ai := 1
				for _, at := range argTypesForSym {
					if ai >= len(args) {
						return environment.NewNil(), fmt.Errorf("call: missing argument for declared argTypes")
					}
					in := args[ai]
					ai++
					switch {
					case len(at) > 7 && at[:7] == "struct:":
						name := at[7:]
						// pointer marker check first
						if in.Type == environment.ObjectType && in.Obj != nil {
							if pidV, ok := in.Obj.Entries[ptrInstanceKey]; ok && pidV.Type == environment.StringType {
								// optional sanity check on struct name
								if so, ok := in.Obj.Entries["__struct__"]; ok && so.Type == environment.StringType {
									if so.Str != name {
										return environment.NewNil(), fmt.Errorf("call: struct pointer name mismatch: expected %s", name)
									}
								}
								x, err := fromEnvironmentValue(in)
								if err != nil {
									return environment.NewNil(), err
								}
								mar = append(mar, x)
								expandedArgTypes = append(expandedArgTypes, at)
								continue
							}
						}
						// expect object or struct object
						if in.Type != environment.ObjectType || in.Obj == nil {
							return environment.NewNil(), fmt.Errorf("call: expected struct object for %s", name)
						}
						// ensure struct name matches if provided
						if so, ok := in.Obj.Entries["__struct__"]; ok && so.Type == environment.StringType {
							if so.Str != name {
								return environment.NewNil(), fmt.Errorf("call: struct object name mismatch: expected %s", name)
							}
						}
						// recursively expand struct fields (supports nested structs)
						if err := expandStructFields(name, in, &mar, &expandedArgTypes); err != nil {
							return environment.NewNil(), err
						}
					default:
						// simple types with coercion based on declared argType
						isIntType := at == "int" || at == "integer"
						isDoubleType := at == "double" || at == "float" || at == "number"
						isStringType := at == "string" || at == "str"

						switch {
						case isIntType && in.Type == environment.NumberType:
							// coerce float64 → int (truncation), verify safe range
							iv := int(in.Num)
							if math.Abs(in.Num) > math.MaxInt32 {
								return environment.NewNil(), fmt.Errorf("call: arg %d value %v overflows int32 range", ai-1, in.Num)
							}
							mar = append(mar, float64(iv))
							expandedArgTypes = append(expandedArgTypes, at)
						case isIntType && in.Type == environment.StringType:
							// try string → int conversion
							iv, convErr := strconv.Atoi(in.Str)
							if convErr != nil {
								return environment.NewNil(), fmt.Errorf("call: arg %d expects int, got string %q (not a valid integer)", ai-1, in.Str)
							}
							mar = append(mar, float64(iv))
							expandedArgTypes = append(expandedArgTypes, at)
						case isDoubleType && in.Type == environment.NumberType:
							// pass-through
							mar = append(mar, in.Num)
							expandedArgTypes = append(expandedArgTypes, at)
						case isDoubleType && in.Type == environment.StringType:
							fv, convErr := strconv.ParseFloat(in.Str, 64)
							if convErr != nil {
								return environment.NewNil(), fmt.Errorf("call: arg %d expects double, got string %q (not a valid number)", ai-1, in.Str)
							}
							mar = append(mar, fv)
							expandedArgTypes = append(expandedArgTypes, at)
						case isStringType && in.Type == environment.NumberType:
							// coerce number → string
							mar = append(mar, fmt.Sprintf("%g", in.Num))
							expandedArgTypes = append(expandedArgTypes, at)
						case isStringType && in.Type == environment.StringType:
							mar = append(mar, in.Str)
							expandedArgTypes = append(expandedArgTypes, at)
						case in.Type == environment.NumberType:
							mar = append(mar, in.Num)
							expandedArgTypes = append(expandedArgTypes, at)
						case in.Type == environment.StringType:
							mar = append(mar, in.Str)
							expandedArgTypes = append(expandedArgTypes, at)
						case in.Type == environment.BooleanType:
							mar = append(mar, in.Bool)
							expandedArgTypes = append(expandedArgTypes, at)
						case in.Type == environment.ObjectType:
							x, err := fromEnvironmentValue(in)
							if err != nil {
								return environment.NewNil(), fmt.Errorf("cannot convert object arg: %v", err)
							}
							mar = append(mar, x)
							expandedArgTypes = append(expandedArgTypes, at)
						default:
							return environment.NewNil(), fmt.Errorf("call: arg %d expects %s, got %v", ai-1, at, in.Type)
						}
					}
				}
			} else {
				for i := 1; i < len(args); i++ {
					// if it's an object, convert generically to JSON serializable structure
					if args[i].Type == environment.ObjectType {
						x, err := fromEnvironmentValue(args[i])
						if err != nil {
							return environment.NewNil(), fmt.Errorf("cannot convert object arg: %v", err)
						}
						mar = append(mar, x)
						continue
					}
					switch args[i].Type {
					case environment.NumberType:
						mar = append(mar, args[i].Num)
					case environment.StringType:
						mar = append(mar, args[i].Str)
					case environment.BooleanType:
						mar = append(mar, args[i].Bool)
					default:
						return environment.NewNil(), fmt.Errorf("unsupported argument type for call: %v", args[i].Type)
					}
				}
			}
			res, err := hc.CallSymbol(symbolId, mar, expandedArgTypes)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("call failed: %v; stderr: %s", err, hc.stderrString())
			}
			val, err := toEnvironmentValue(res)
			if err != nil {
				return environment.NewNil(), err
			}
			return val, nil
		}),

		// struct_(name, fieldsArr) -> high-level wrapper API
		// returns a descriptor object with methods: new, validate, flatten
		fn("struct_", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("struct(name, fields)")
			}
			name, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("struct: name must be string")
			}
			if args[1].Type != environment.ArrayType || args[1].Arr == nil {
				return environment.NewNil(), fmt.Errorf("struct: fields must be array")
			}
			var fields []struct {
				Field string
				Type  string
			}
			for _, el := range *args[1].Arr {
				if el.Type != environment.ObjectType || el.Obj == nil {
					return environment.NewNil(), fmt.Errorf("struct: each field must be object")
				}
				nV, okN := el.Obj.Entries["name"]
				tV, okT := el.Obj.Entries["type"]
				if !okN || !okT || nV.Type != environment.StringType || tV.Type != environment.StringType {
					return environment.NewNil(), fmt.Errorf("struct: field object must have 'name' and 'type' strings")
				}
				tt := tV.Str
				if !isSupportedType(tt) {
					return environment.NewNil(), fmt.Errorf("struct %s, field %s: unknown type: %s", name, nV.Str, tt)
				}
				fields = append(fields, struct {
					Field string
					Type  string
				}{Field: nV.Str, Type: tt})
			}
			// register same as define_struct
			structSchemasMu.Lock()
			structSchemas[name] = fields
			structSchemasMu.Unlock()

			coerce := func(v environment.Value, t string) (environment.Value, error) {
				isInt := t == "int" || t == "integer"
				isDouble := t == "double" || t == "float" || t == "number"
				isString := t == "string" || t == "str"
				switch {
				case isInt && v.Type == environment.NumberType:
					iv := int(v.Num)
					if math.Abs(v.Num) > math.MaxInt32 {
						return environment.NewNil(), fmt.Errorf("field %s expects int, value %v overflows int32", t, v.Num)
					}
					return environment.NewNumber(float64(iv)), nil
				case isInt && v.Type == environment.StringType:
					iv, convErr := strconv.Atoi(v.Str)
					if convErr != nil {
						return environment.NewNil(), fmt.Errorf("field expects int, got string %q", v.Str)
					}
					return environment.NewNumber(float64(iv)), nil
				case isDouble && v.Type == environment.NumberType:
					return environment.NewNumber(v.Num), nil
				case isDouble && v.Type == environment.StringType:
					fv, convErr := strconv.ParseFloat(v.Str, 64)
					if convErr != nil {
						return environment.NewNil(), fmt.Errorf("field expects double, got string %q", v.Str)
					}
					return environment.NewNumber(fv), nil
				case isString && v.Type == environment.NumberType:
					return environment.NewString(fmt.Sprintf("%g", v.Num)), nil
				case isString && v.Type == environment.StringType:
					return environment.NewString(v.Str), nil
				case v.Type == environment.NumberType:
					return environment.NewNumber(v.Num), nil
				case v.Type == environment.StringType:
					return environment.NewString(v.Str), nil
				case v.Type == environment.BooleanType:
					return environment.NewBool(v.Bool), nil
				case v.Type == environment.ObjectType:
					x, err := toEnvironmentValue(v)
					if err != nil {
						return environment.NewNil(), err
					}
					return x, nil
				default:
					return environment.NewNil(), fmt.Errorf("unsupported field type %v", v.Type)
				}
			}

			entries := map[string]environment.Value{
				structDefKey: environment.NewString(name),
				"name":       environment.NewString(name),
				"fields":     args[1],
			}
			keys := []string{structDefKey, "name", "fields"}

			// new constructor
			newFn := func(args []environment.Value) (environment.Value, error) {
				inst := make(map[string]environment.Value)
				inst["__struct__"] = environment.NewString(name)
				if len(args) == 1 && args[0].Type == environment.ObjectType && args[0].Obj != nil {
					for _, f := range fields {
						v, ok := args[0].Obj.Entries[f.Field]
						if !ok {
							return environment.NewNil(), fmt.Errorf("missing field %s", f.Field)
						}
						cv, err := coerce(v, f.Type)
						if err != nil {
							return environment.NewNil(), err
						}
						inst[f.Field] = cv
					}
				} else {
					if len(args) != len(fields) {
						return environment.NewNil(), fmt.Errorf("expected %d args, got %d", len(fields), len(args))
					}
					for i, f := range fields {
						cv, err := coerce(args[i], f.Type)
						if err != nil {
							return environment.NewNil(), err
						}
						inst[f.Field] = cv
					}
				}
				return environment.NewObject(inst, nil), nil
			}
			entries["new"] = environment.NewBuiltinFn("new", newFn)
			keys = append(keys, "new")

			validateFn := func(args []environment.Value) (environment.Value, error) {
				if len(args) != 1 {
					return environment.NewNil(), fmt.Errorf("validate expects 1 arg")
				}
				obj := args[0]
				if obj.Type != environment.ObjectType || obj.Obj == nil {
					return environment.NewNil(), fmt.Errorf("validate: argument must be object")
				}
				for _, f := range fields {
					v, ok := obj.Obj.Entries[f.Field]
					if !ok {
						return environment.NewNil(), fmt.Errorf("missing field %s", f.Field)
					}
					if _, err := coerce(v, f.Type); err != nil {
						return environment.NewNil(), err
					}
				}
				return environment.NewBool(true), nil
			}
			entries["validate"] = environment.NewBuiltinFn("validate", validateFn)
			keys = append(keys, "validate")

			flattenFn := func(args []environment.Value) (environment.Value, error) {
				if len(args) != 1 {
					return environment.NewNil(), fmt.Errorf("flatten expects 1 arg")
				}
				obj := args[0]
				if obj.Type != environment.ObjectType || obj.Obj == nil {
					return environment.NewNil(), fmt.Errorf("flatten: arg must be object")
				}
				var mar []interface{}
				var types []string
				if err := expandStructFields(name, obj, &mar, &types); err != nil {
					return environment.NewNil(), err
				}
				vals := make([]environment.Value, len(mar))
				for i, mv := range mar {
					vv, err := toEnvironmentValue(mv)
					if err != nil {
						return environment.NewNil(), err
					}
					vals[i] = vv
				}
				tvals := make([]environment.Value, len(types))
				for i, tv := range types {
					tvals[i] = environment.NewString(tv)
				}
				resp := map[string]environment.Value{"values": environment.NewArray(vals), "types": environment.NewArray(tvals)}
				return environment.NewObject(resp, []string{"values", "types"}), nil
			}
			entries["flatten"] = environment.NewBuiltinFn("flatten", flattenFn)
			keys = append(keys, "flatten")

			return environment.NewObject(entries, keys), nil
		}),

		// define_struct(name, fieldsArr) -> registers a struct schema for marshalling
		// Note: wrapper API available as `struct_` builtin (see above).
		// fieldsArr is an array of objects: [{name: "field1", type: "string"}, ...]
		fn("define_struct", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("define_struct(name, fieldsArr) expects 2 args")
			}
			name, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("define_struct: first arg must be string name")
			}
			if args[1].Type != environment.ArrayType || args[1].Arr == nil {
				return environment.NewNil(), fmt.Errorf("define_struct: fields must be an array")
			}
			var fields []struct {
				Field string
				Type  string
			}
			for _, el := range *args[1].Arr {
				if el.Type != environment.ObjectType || el.Obj == nil {
					return environment.NewNil(), fmt.Errorf("define_struct: each field must be object with 'name' and 'type'")
				}
				nV, okN := el.Obj.Entries["name"]
				tV, okT := el.Obj.Entries["type"]
				if !okN || !okT || nV.Type != environment.StringType || tV.Type != environment.StringType {
					return environment.NewNil(), fmt.Errorf("define_struct: field object must have 'name' and 'type' strings")
				}
				fields = append(fields, struct {
					Field string
					Type  string
				}{Field: nV.Str, Type: tV.Str})
			}
			structSchemasMu.Lock()
			structSchemas[name] = fields
			structSchemasMu.Unlock()
			return environment.NewNil(), nil
		}),

		// helper debug command - send a raw command to helper (testing only)
		fn("helper_cmd", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("helper_cmd(cmd, ...args)")
			}
			cmdStr, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("helper_cmd: cmd must be string")
			}
			en, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			if !en {
				return environment.NewNil(), fmt.Errorf("FFI is not enabled for this project; run 'fig setup-ffi' and enable in fig.toml")
			}
			if helper == "" {
				return environment.NewNil(), fmt.Errorf("no ffi.helper configured in fig.toml; run 'fig setup-ffi'")
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			req := map[string]interface{}{"cmd": cmdStr}
			if len(args) >= 2 {
				// pass single numeric ms parameter if provided
				if args[1].Type == environment.NumberType {
					req["ms"] = args[1].Num
				}
			}
			resp, err := hc.call(req)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("helper_cmd failed: %v", err)
			}
			if ok, _ := resp["ok"].(bool); !ok {
				return environment.NewNil(), fmt.Errorf("helper_cmd: %v; stderr: %s", ffiParseError(resp, "helper_cmd"), hc.stderrString())
			}
			val, err := toEnvironmentValue(resp["result"])
			if err != nil {
				// retrocompat: try old "resp" field
				if oldResp, ok := resp["resp"]; ok {
					val, err = toEnvironmentValue(oldResp)
				}
				if err != nil {
					return environment.NewNil(), err
				}
			}
			return val, nil
		}),

		// bytes_from_string(s) -> special object representing bytes (base64)
		fn("bytes_from_string", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("bytes_from_string(s) expects 1 arg")
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("bytes_from_string: arg must be string")
			}
			enc := base64.StdEncoding.EncodeToString([]byte(s))
			obj := map[string]environment.Value{"__bytes__": environment.NewString(enc)}
			return environment.NewObject(obj, []string{"__bytes__"}), nil
		}),

		// bytes_to_string(obj) -> decodes a bytes object (created via bytes_from_string) back to string
		fn("bytes_to_string", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("bytes_to_string(obj) expects 1 arg")
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("bytes_to_string: expected bytes object")
			}
			if v, ok := args[0].Obj.Entries["__bytes__"]; ok && v.Type == environment.StringType {
				b, err := base64.StdEncoding.DecodeString(v.Str)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("invalid base64 in bytes object")
				}
				return environment.NewString(string(b)), nil
			}
			return environment.NewNil(), fmt.Errorf("bytes_to_string: invalid bytes object")
		}),

		// bytes_from_array(arr) -> special bytes object (base64) from numeric array (0-255)
		fn("bytes_from_array", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("bytes_from_array(arr) expects 1 arg")
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("bytes_from_array: expected array of numbers")
			}
			arr := *args[0].Arr
			b := make([]byte, len(arr))
			for i, el := range arr {
				if el.Type != environment.NumberType {
					return environment.NewNil(), fmt.Errorf("bytes_from_array: array elements must be numbers (0-255)")
				}
				n := int(el.Num)
				if n < 0 || n > 255 {
					return environment.NewNil(), fmt.Errorf("bytes_from_array: array elements must be in 0..255")
				}
				b[i] = byte(n)
			}
			enc := base64.StdEncoding.EncodeToString(b)
			obj := map[string]environment.Value{"__bytes__": environment.NewString(enc)}
			return environment.NewObject(obj, []string{"__bytes__"}), nil
		}),

		// bytes_to_array(obj) -> decodes bytes object into array of numbers (0-255)
		fn("bytes_to_array", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("bytes_to_array(obj) expects 1 arg")
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("bytes_to_array: expected bytes object")
			}
			if v, ok := args[0].Obj.Entries["__bytes__"]; ok && v.Type == environment.StringType {
				b, err := base64.StdEncoding.DecodeString(v.Str)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("invalid base64 in bytes object")
				}
				elems := make([]environment.Value, len(b))
				for i, bb := range b {
					elems[i] = environment.NewNumber(float64(bb))
				}
				return environment.NewArray(elems), nil
			}
			return environment.NewNil(), fmt.Errorf("bytes_to_array: invalid bytes object")
		}),

		// alloc(size) -> returns memory object {"__mem__": id}
		fn("alloc", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("alloc(size) expects 1 arg")
			}
			if args[0].Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("alloc: size must be number")
			}
			sz := int(args[0].Num)
			if sz <= 0 {
				return environment.NewNil(), fmt.Errorf("alloc: size must be > 0")
			}
			_, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			// sandbox: check max_allocs limit
			sbCfg, _, _ := readSandboxConfig()
			if sbCfg.MaxAllocs > 0 {
				cur := GetAllocCountValue(projectRoot)
				if cur >= int64(sbCfg.MaxAllocs) {
					return environment.NewNil(), fmt.Errorf("sandbox: max live allocations limit reached (%d)", sbCfg.MaxAllocs)
				}
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			mid, err := hc.Alloc(sz)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("alloc failed: %v; stderr: %s", err, hc.stderrString())
			}
			IncrementAllocCount(projectRoot)
			obj := map[string]environment.Value{"__mem__": environment.NewString(mid), "size": environment.NewNumber(float64(sz))}
			return environment.NewObject(obj, []string{"__mem__", "size"}), nil
		}),

		// free(memObj) -> frees memory in helper
		fn("free", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("free(mem) expects 1 arg")
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("free: expected mem object")
			}
			v, ok := args[0].Obj.Entries["__mem__"]
			if !ok || v.Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("free: invalid mem object")
			}
			mid := v.Str
			_, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			if err := hc.Free(mid); err != nil {
				return environment.NewNil(), fmt.Errorf("free failed: %v; stderr: %s", err, hc.stderrString())
			}
			DecrementAllocCount(projectRoot)
			return environment.NewNil(), nil
		}),

		// strdup(str) -> duplicates string as C-allocated memory, returns mem object
		fn("strdup", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("strdup(str) expects 1 arg")
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("strdup: arg must be string")
			}
			_, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			// sandbox: check max_allocs limit
			sbCfg, _, _ := readSandboxConfig()
			if sbCfg.MaxAllocs > 0 {
				cur := GetAllocCountValue(projectRoot)
				if cur >= int64(sbCfg.MaxAllocs) {
					return environment.NewNil(), fmt.Errorf("sandbox: max live allocations limit reached (%d)", sbCfg.MaxAllocs)
				}
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			mid, err := hc.Strdup(s)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("strdup failed: %v; stderr: %s", err, hc.stderrString())
			}
			IncrementAllocCount(projectRoot)
			sz := len(s) + 1
			obj := map[string]environment.Value{"__mem__": environment.NewString(mid), "size": environment.NewNumber(float64(sz))}
			return environment.NewObject(obj, []string{"__mem__", "size"}), nil
		}),

		// free_string(memObj) -> convenience alias for free, frees a strdup'd string
		fn("free_string", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("free_string(mem) expects 1 arg")
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("free_string: expected mem object")
			}
			v, ok := args[0].Obj.Entries["__mem__"]
			if !ok || v.Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("free_string: invalid mem object")
			}
			mid := v.Str
			_, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			if err := hc.Free(mid); err != nil {
				return environment.NewNil(), fmt.Errorf("free_string failed: %v; stderr: %s", err, hc.stderrString())
			}
			DecrementAllocCount(projectRoot)
			return environment.NewNil(), nil
		}),

		// mem_write(memObj, bytesObj, offset)
		fn("mem_write", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("mem_write(mem, bytes, offset) expects 3 args")
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("mem_write: expected mem object")
			}
			midV, ok := args[0].Obj.Entries["__mem__"]
			if !ok || midV.Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("mem_write: invalid mem object")
			}
			mid := midV.Str
			// bytes argument
			if args[1].Type != environment.ObjectType || args[1].Obj == nil {
				return environment.NewNil(), fmt.Errorf("mem_write: expected bytes object for data")
			}
			b64V, ok := args[1].Obj.Entries["__bytes__"]
			if !ok || b64V.Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("mem_write: expected bytes object")
			}
			// offset
			off, err := args[2].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("mem_write: offset must be number")
			}
			_, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			if err := hc.MemWrite(mid, b64V.Str, int(off)); err != nil {
				return environment.NewNil(), fmt.Errorf("mem_write failed: %v; stderr: %s", err, hc.stderrString())
			}
			return environment.NewNil(), nil
		}),

		// mem_read(memObj, offset, len) -> bytes object
		fn("mem_read", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("mem_read(mem, offset, len) expects 3 args")
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("mem_read: expected mem object")
			}
			midV, ok := args[0].Obj.Entries["__mem__"]
			if !ok || midV.Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("mem_read: invalid mem object")
			}
			mid := midV.Str
			offf, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("mem_read: offset must be number")
			}
			lnf, err := args[2].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("mem_read: len must be number")
			}
			_, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			b64, err := hc.MemRead(mid, int(offf), int(lnf))
			if err != nil {
				return environment.NewNil(), fmt.Errorf("mem_read failed: %v; stderr: %s", err, hc.stderrString())
			}
			obj := map[string]environment.Value{"__bytes__": environment.NewString(b64)}
			return environment.NewObject(obj, []string{"__bytes__"}), nil
		}),

		// runCallback called by helper client when helper asks to invoke cb id
		// returns a JSON-serializable response
		fn("_run_callback_internal", func(args []environment.Value) (environment.Value, error) {
			// internal exposed for tests; not meant for public use
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("_run_callback_internal(cbId, argsArr)")
			}
			cbid, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("first arg must be cb id string")
			}
			if args[1].Type != environment.ArrayType || args[1].Arr == nil {
				return environment.NewNil(), fmt.Errorf("second arg must be array of args")
			}
			// convert to []interface{}
			var jargs []interface{}
			for _, el := range *args[1].Arr {
				x, err := fromEnvironmentValue(el)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("cannot convert arg: %v", err)
				}
				jargs = append(jargs, x)
			}
			res, err := runCallback(cbid, jargs)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("callback failed: %v", err)
			}
			// wrap response as environment value using toEnvironmentValue via json marshal/unmarshal to ensure consistent types
			b, _ := json.Marshal(res)
			var any interface{}
			_ = json.Unmarshal(b, &any)
			v, err := toEnvironmentValue(any)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot convert callback response: %v", err)
			}
			return v, nil
		}),

		// call_raw(args...) - call helper without symbol and return raw response (array/string/num)
		fn("call_raw", func(args []environment.Value) (environment.Value, error) {
			// marshal args
			var mar []interface{}
			for i := 0; i < len(args); i++ {
				switch args[i].Type {
				case environment.NumberType:
					mar = append(mar, args[i].Num)
				case environment.StringType:
					mar = append(mar, args[i].Str)
				case environment.BooleanType:
					mar = append(mar, args[i].Bool)
				case environment.ArrayType:
					if args[i].Arr != nil {
						arr := *args[i].Arr
						var jr []interface{}
						for _, el := range arr {
							switch el.Type {
							case environment.NumberType:
								jr = append(jr, el.Num)
							case environment.StringType:
								jr = append(jr, el.Str)
							case environment.BooleanType:
								jr = append(jr, el.Bool)
							case environment.ObjectType:
								if el.Obj != nil {
									if v, ok := el.Obj.Entries["__bytes__"]; ok && v.Type == environment.StringType {
										jr = append(jr, map[string]interface{}{"__bytes__": v.Str})
										continue
									}
								}
								return environment.NewNil(), fmt.Errorf("unsupported nested object type for call_raw")
							default:
								return environment.NewNil(), fmt.Errorf("unsupported nested arg type for call_raw: %v", el.Type)
							}
						}
						mar = append(mar, jr)
					}
				case environment.ObjectType:
					// support top-level bytes object
					if args[i].Obj != nil {
						if v, ok := args[i].Obj.Entries["__bytes__"]; ok && v.Type == environment.StringType {
							mar = append(mar, map[string]interface{}{"__bytes__": v.Str})
							continue
						}
					}
					return environment.NewNil(), fmt.Errorf("unsupported arg type for call_raw: %v", args[i].Type)
				default:
					return environment.NewNil(), fmt.Errorf("unsupported arg type for call_raw: %v", args[i].Type)
				}
			}
			en, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			if !en {
				return environment.NewNil(), fmt.Errorf("FFI is not enabled for this project; run 'fig setup-ffi' and enable in fig.toml")
			}
			if helper == "" {
				return environment.NewNil(), fmt.Errorf("no ffi.helper configured in fig.toml; run 'fig setup-ffi'")
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			res, err := hc.Call(mar)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("call_raw failed: %v; stderr: %s", err, hc.stderrString())
			}
			val, err := toEnvironmentValue(res)
			if err != nil {
				return environment.NewNil(), err
			}
			return val, nil
		}),

		// ping helper (starts it transiently and asks for pong)
		fn("ping", func(args []environment.Value) (environment.Value, error) {
			en, helper, projectRoot, callTimeoutMs, err := readFfiConfig()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot read fig.toml: %v", err)
			}
			if !en {
				return environment.NewNil(), fmt.Errorf("FFI is not enabled for this project; run 'fig setup-ffi' and enable in fig.toml")
			}
			if helper == "" {
				return environment.NewNil(), fmt.Errorf("no ffi.helper configured in fig.toml; run 'fig setup-ffi'")
			}
			hc, err := getHelperForProject(projectRoot, helper)
			// apply timeout config if present
			if callTimeoutMs >= 0 {
				if callTimeoutMs == 0 {
					hc.callTimeout = 0
				} else {
					hc.callTimeout = time.Duration(callTimeoutMs) * time.Millisecond
				}
			}
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cannot start helper: %v", err)
			}
			resp, err := hc.call(map[string]interface{}{"cmd": "ping"})
			if err != nil {
				// try restart once
				_ = hc.Stop()
				helpersMu.Lock()
				delete(helpers, projectRoot)
				helpersMu.Unlock()
				hc2, err2 := getHelperForProject(projectRoot, helper)
				if err2 != nil {
					return environment.NewNil(), fmt.Errorf("helper error: %v; restart failed: %v", err, err2)
				}
				resp, err = hc2.call(map[string]interface{}{"cmd": "ping"})
				if err != nil {
					return environment.NewNil(), fmt.Errorf("helper error after restart: %v", err)
				}
			}
			if ok, _ := resp["ok"].(bool); !ok {
				return environment.NewNil(), fmt.Errorf("ping: %v; stderr: %s", ffiParseError(resp, "ping"), hc.stderrString())
			}
			// new envelope: result; old: resp
			if r, _ := resp["result"].(string); r == "pong" {
				return environment.NewString("pong"), nil
			}
			if r, _ := resp["resp"].(string); r == "pong" {
				return environment.NewString("pong"), nil
			}
			return environment.NewNil(), fmt.Errorf("unexpected helper response: %v", resp)
		}),
	))
}
