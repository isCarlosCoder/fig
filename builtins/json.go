package builtins

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("json",
		// jsonParse(s) — parses a JSON string into a Fig value
		fn("parse", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("parse() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("parse() argument must be a string")
			}
			var raw interface{}
			if err := json.Unmarshal([]byte(args[0].Str), &raw); err != nil {
				return environment.NewNil(), fmt.Errorf("parse() invalid JSON: %v", err)
			}
			return goToFig(raw), nil
		}),

		// jsonStringify(x) — converts a Fig value to a JSON string
		fn("stringify", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("stringify() expects 1 argument, got %d", len(args))
			}
			goVal := figToGo(args[0])
			data, err := json.Marshal(goVal)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("stringify() error: %v", err)
			}
			return environment.NewString(string(data)), nil
		}),

		// serialize(x) — same as stringify but pretty-printed
		fn("serialize", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("serialize() expects 1 argument, got %d", len(args))
			}
			goVal := figToGo(args[0])
			data, err := json.MarshalIndent(goVal, "", "  ")
			if err != nil {
				return environment.NewNil(), fmt.Errorf("serialize() error: %v", err)
			}
			return environment.NewString(string(data)), nil
		}),

		// deserialize(s) — alias for parse
		fn("deserialize", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("deserialize() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("deserialize() argument must be a string")
			}
			var raw interface{}
			if err := json.Unmarshal([]byte(args[0].Str), &raw); err != nil {
				return environment.NewNil(), fmt.Errorf("deserialize() invalid JSON: %v", err)
			}
			return goToFig(raw), nil
		}),
	))
}

// goToFig converts a Go interface{} (from json.Unmarshal) to a Fig Value.
func goToFig(v interface{}) environment.Value {
	if v == nil {
		return environment.NewNil()
	}
	switch val := v.(type) {
	case float64:
		return environment.NewNumber(val)
	case bool:
		return environment.NewBool(val)
	case string:
		return environment.NewString(val)
	case []interface{}:
		elems := make([]environment.Value, len(val))
		for i, e := range val {
			elems[i] = goToFig(e)
		}
		return environment.NewArray(elems)
	case map[string]interface{}:
		entries := make(map[string]environment.Value, len(val))
		keys := make([]string, 0, len(val))
		for k := range val {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			entries[k] = goToFig(val[k])
		}
		return environment.NewObject(entries, keys)
	default:
		return environment.NewString(fmt.Sprintf("%v", val))
	}
}

// figToGo converts a Fig Value to a Go interface{} suitable for json.Marshal.
func figToGo(v environment.Value) interface{} {
	switch v.Type {
	case environment.NumberType:
		if v.Num == math.Trunc(v.Num) && !math.IsInf(v.Num, 0) {
			return int64(v.Num)
		}
		return v.Num
	case environment.BooleanType:
		return v.Bool
	case environment.StringType:
		return v.Str
	case environment.NilType:
		return nil
	case environment.ArrayType:
		if v.Arr == nil {
			return []interface{}{}
		}
		arr := make([]interface{}, len(*v.Arr))
		for i, e := range *v.Arr {
			arr[i] = figToGo(e)
		}
		return arr
	case environment.ObjectType:
		if v.Obj == nil {
			return map[string]interface{}{}
		}
		m := make(map[string]interface{}, len(v.Obj.Keys))
		for _, k := range v.Obj.Keys {
			m[k] = figToGo(v.Obj.Entries[k])
		}
		return m
	default:
		return v.String()
	}
}

// jsonString produces a proper JSON representation of a Fig Value (used internally).
func jsonString(v environment.Value) string {
	goVal := figToGo(v)
	data, err := json.Marshal(goVal)
	if err != nil {
		return "null"
	}
	// For objects, maintain key order
	if v.Type == environment.ObjectType && v.Obj != nil {
		var sb strings.Builder
		sb.WriteByte('{')
		for i, k := range v.Obj.Keys {
			if i > 0 {
				sb.WriteByte(',')
			}
			kb, _ := json.Marshal(k)
			vb, _ := json.Marshal(figToGo(v.Obj.Entries[k]))
			sb.Write(kb)
			sb.WriteByte(':')
			sb.Write(vb)
		}
		sb.WriteByte('}')
		return sb.String()
	}
	return string(data)
}
