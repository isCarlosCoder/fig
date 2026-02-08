package builtins

import (
	"fmt"
	"strings"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("debug",
		// panic(msg) — stops execution with an error
		fn("panic", func(args []environment.Value) (environment.Value, error) {
			msg := "panic!"
			if len(args) >= 1 {
				msg = args[0].String()
			}
			return environment.NewNil(), fmt.Errorf("panic: %s", msg)
		}),

		// assert(cond, msg) — if cond is falsy, raises error with msg
		fn("assert", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("assert() expects at least 1 argument, got %d", len(args))
			}
			ok := isTruthy(args[0])
			if !ok {
				msg := "assertion failed"
				if len(args) >= 2 {
					msg = args[1].String()
				}
				return environment.NewNil(), fmt.Errorf("assert: %s", msg)
			}
			return environment.NewBool(true), nil
		}),

		// dump(x) — returns a detailed string representation of a value
		fn("dump", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("dump() expects 1 argument, got %d", len(args))
			}
			v := args[0]
			detail := dumpValue(v, 0)
			return environment.NewString(detail), nil
		}),

		// inspect(x) — returns type and value as string
		fn("inspect", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("inspect() expects 1 argument, got %d", len(args))
			}
			v := args[0]
			result := fmt.Sprintf("<%s: %s>", v.TypeName(), v.String())
			return environment.NewString(result), nil
		}),

		// type(x) — alias, returns type name
		fn("type", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("type() expects 1 argument, got %d", len(args))
			}
			return environment.NewString(args[0].TypeName()), nil
		}),
	))
}

func isTruthy(v environment.Value) bool {
	switch v.Type {
	case environment.BooleanType:
		return v.Bool
	case environment.NilType:
		return false
	case environment.NumberType:
		return v.Num != 0
	case environment.StringType:
		return v.Str != ""
	default:
		return true
	}
}

func dumpValue(v environment.Value, indent int) string {
	prefix := strings.Repeat("  ", indent)
	switch v.Type {
	case environment.ArrayType:
		if v.Arr == nil || len(*v.Arr) == 0 {
			return prefix + "array(0) []"
		}
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%sarray(%d) [\n", prefix, len(*v.Arr)))
		for i, elem := range *v.Arr {
			sb.WriteString(fmt.Sprintf("%s  [%d] %s\n", prefix, i, dumpValue(elem, 0)))
		}
		sb.WriteString(prefix + "]")
		return sb.String()
	case environment.ObjectType:
		if v.Obj == nil || len(v.Obj.Keys) == 0 {
			return prefix + "object(0) {}"
		}
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("%sobject(%d) {\n", prefix, len(v.Obj.Keys)))
		for _, k := range v.Obj.Keys {
			sb.WriteString(fmt.Sprintf("%s  %s: %s\n", prefix, k, dumpValue(v.Obj.Entries[k], 0)))
		}
		sb.WriteString(prefix + "}")
		return sb.String()
	case environment.NumberType:
		return fmt.Sprintf("%s(%s)", v.TypeName(), v.String())
	case environment.StringType:
		return fmt.Sprintf("%s(%q)", v.TypeName(), v.Str)
	default:
		return fmt.Sprintf("%s(%s)", v.TypeName(), v.String())
	}
}
