package environment

import (
	"fmt"
	"math"
	"strings"
)

type ValueType int

const (
	NilType ValueType = iota
	StringType
	BooleanType
	NumberType
	FunctionType
	BuiltinFnType
	ArrayType
	ObjectType
	StructDefType
	InstanceType
	EnumType
	EnumMemberType
)

// Param describes a single function parameter (name + optional/default metadata).
type Param struct {
	Name       string      // identifier
	HasDefault bool        // true when a default expression was provided (name = expr)
	Default    interface{} // AST node for the default expression (stored as interface{} to avoid parser import)
	Optional   bool        // true when declared with `name?` (implicit default `null`)
}

// FuncDef holds the information needed to call a user-defined function.
type FuncDef struct {
	Name       string
	Params     []Param
	Body       interface{} // *parser.BlockContext — kept as interface{} to avoid import cycle
	ClosureEnv *Env        // the environment where the function was defined (lexical closure)
	DefFile    string      // file where function was defined (for stack traces)
	DefLine    int         // line where function was defined (for stack traces)
}

// BuiltinFn is the signature for native Go functions callable from Fig.
type BuiltinFn func(args []Value) (Value, error)

// StructField describes a field in a struct definition.
type StructField struct {
	Name    string
	Default Value // default value (NilType if none)
}

// StructDef holds a struct blueprint: its name, fields, and methods.
type StructDef struct {
	Name    string
	Fields  []StructField
	Methods map[string]*FuncDef
}

// Instance is a live struct instance with its own field values.
type Instance struct {
	Def    *StructDef
	Fields *ObjData // field values (pointer for mutability)
}

// EnumDef describes an enum type declaration.
type EnumDef struct {
	Name    string
	Members []string // ordered member names
}

// EnumMember represents a single enum variant value.
type EnumMember struct {
	EnumName string
	Name     string
	Ordinal  int
}

type Value struct {
	Type    ValueType
	Num     float64
	Str     string
	Bool    bool
	Func    *FuncDef    // non-nil when Type == FunctionType
	Builtin BuiltinFn   // non-nil when Type == BuiltinFnType
	BName   string      // builtin function name (for error messages)
	Arr     *[]Value    // non-nil when Type == ArrayType (pointer for mutability)
	Obj     *ObjData    // non-nil when Type == ObjectType (pointer for mutability)
	Struct  *StructDef  // non-nil when Type == StructDefType
	Inst    *Instance   // non-nil when Type == InstanceType
	EnumDef *EnumDef    // non-nil when Type == EnumType
	EnumMem *EnumMember // non-nil when Type == EnumMemberType
}

func NewEnumDef(ed *EnumDef) Value { return Value{Type: EnumType, EnumDef: ed} }
func NewEnumMember(enumName, name string, ord int) Value {
	return Value{Type: EnumMemberType, EnumMem: &EnumMember{EnumName: enumName, Name: name, Ordinal: ord}}
}

// ObjData stores object key-value pairs with insertion order.
type ObjData struct {
	Entries map[string]Value
	Keys    []string // insertion order
}

// Constructors
func NewNumber(n float64) Value     { return Value{Type: NumberType, Num: n} }
func NewString(s string) Value      { return Value{Type: StringType, Str: s} }
func NewBool(b bool) Value          { return Value{Type: BooleanType, Bool: b} }
func NewNil() Value                 { return Value{Type: NilType} }
func NewFunction(fd *FuncDef) Value { return Value{Type: FunctionType, Func: fd} }
func NewBuiltinFn(name string, fn BuiltinFn) Value {
	return Value{Type: BuiltinFnType, Builtin: fn, BName: name}
}
func NewArray(elems []Value) Value {
	arr := make([]Value, len(elems))
	copy(arr, elems)
	return Value{Type: ArrayType, Arr: &arr}
}
func NewObject(entries map[string]Value, keys []string) Value {
	return Value{Type: ObjectType, Obj: &ObjData{Entries: entries, Keys: keys}}
}
func NewStructDef(sd *StructDef) Value {
	return Value{Type: StructDefType, Struct: sd}
}
func NewInstance(inst *Instance) Value {
	return Value{Type: InstanceType, Inst: inst}
}

func (v Value) String() string {
	switch v.Type {
	case NumberType:
		// Format integers without scientific notation
		if v.Num == math.Trunc(v.Num) && !math.IsInf(v.Num, 0) && !math.IsNaN(v.Num) {
			return fmt.Sprintf("%.0f", v.Num)
		}
		return fmt.Sprintf("%g", v.Num)
	case BooleanType:
		return fmt.Sprintf("%t", v.Bool)
	case StringType:
		return v.Str
	case NilType:
		return "null"
	case FunctionType:
		if v.Func != nil {
			return fmt.Sprintf("<fn %s>", v.Func.Name)
		}
		return "<fn>"
	case BuiltinFnType:
		return fmt.Sprintf("<builtin %s>", v.BName)
	case ArrayType:
		if v.Arr == nil {
			return "[]"
		}
		parts := make([]string, len(*v.Arr))
		for i, e := range *v.Arr {
			if e.Type == StringType {
				parts[i] = fmt.Sprintf("%q", e.Str)
			} else {
				parts[i] = e.String()
			}
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case ObjectType:
		if v.Obj == nil {
			return "{}"
		}
		parts := make([]string, 0, len(v.Obj.Keys))
		for _, k := range v.Obj.Keys {
			val := v.Obj.Entries[k]
			if val.Type == StringType {
				parts = append(parts, fmt.Sprintf("%s: %q", k, val.Str))
			} else {
				parts = append(parts, fmt.Sprintf("%s: %s", k, val.String()))
			}
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case StructDefType:
		if v.Struct != nil {
			return fmt.Sprintf("<struct %s>", v.Struct.Name)
		}
		return "<struct>"
	case EnumType:
		if v.EnumDef != nil {
			return fmt.Sprintf("<enum %s>", v.EnumDef.Name)
		}
		return "<enum>"
	case EnumMemberType:
		if v.EnumMem != nil {
			return fmt.Sprintf("%s.%s", v.EnumMem.EnumName, v.EnumMem.Name)
		}
		return "<enum member>"
	case InstanceType:
		if v.Inst != nil {
			parts := make([]string, 0, len(v.Inst.Fields.Keys))
			for _, k := range v.Inst.Fields.Keys {
				val := v.Inst.Fields.Entries[k]
				if val.Type == StringType {
					parts = append(parts, fmt.Sprintf("%s: %q", k, val.Str))
				} else {
					parts = append(parts, fmt.Sprintf("%s: %s", k, val.String()))
				}
			}
			return v.Inst.Def.Name + "{" + strings.Join(parts, ", ") + "}"
		}
		return "<instance>"
	}

	return ""
}

// IsTruthy returns whether the value should be considered truthy in conditions
func (v Value) IsTruthy() bool {
	switch v.Type {
	case NilType:
		return false
	case BooleanType:
		return v.Bool
	case ArrayType:
		if v.Arr == nil {
			return false
		}
		return len(*v.Arr) > 0
	case ObjectType:
		if v.Obj == nil {
			return false
		}
		return len(v.Obj.Entries) > 0
	case InstanceType:
		return true
	case StructDefType:
		return true
	default:
		return true
	}
}

// Typed accessors that return errors on mismatch
func (v Value) AsNumber() (float64, error) {
	if v.Type != NumberType {
		return 0, fmt.Errorf("not a number: %s", v.TypeName())
	}
	return v.Num, nil
}

func (v Value) AsString() (string, error) {
	if v.Type != StringType {
		return "", fmt.Errorf("not a string: %s", v.TypeName())
	}
	return v.Str, nil
}

func (v Value) AsBool() (bool, error) {
	if v.Type != BooleanType {
		return false, fmt.Errorf("not a boolean: %s", v.TypeName())
	}
	return v.Bool, nil
}

func (v Value) TypeName() string {
	switch v.Type {
	case NumberType:
		return "number"
	case StringType:
		return "string"
	case BooleanType:
		return "boolean"
	case NilType:
		return "null"
	case FunctionType:
		return "function"
	case BuiltinFnType:
		return "function"
	case ArrayType:
		return "array"
	case ObjectType:
		return "object"
	case StructDefType:
		return "struct"
	case EnumMemberType:
		if v.EnumMem != nil {
			return v.EnumMem.EnumName
		}
		return "enum"
	case InstanceType:
		if v.Inst != nil {
			return v.Inst.Def.Name
		}
		return "instance"
	default:
		return "unknown"
	}
}
