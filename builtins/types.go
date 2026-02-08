package builtins

import (
	"fmt"
	"math"
	"strconv"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("types",
		// type(x) — returns the type name as a string
		fn("type", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("type() expects 1 argument, got %d", len(args))
			}
			return environment.NewString(args[0].TypeName()), nil
		}),

		// isNumber(x)
		fn("isNumber", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isNumber() expects 1 argument, got %d", len(args))
			}
			return environment.NewBool(args[0].Type == environment.NumberType), nil
		}),

		// isString(x)
		fn("isString", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isString() expects 1 argument, got %d", len(args))
			}
			return environment.NewBool(args[0].Type == environment.StringType), nil
		}),

		// isBool(x)
		fn("isBool", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isBool() expects 1 argument, got %d", len(args))
			}
			return environment.NewBool(args[0].Type == environment.BooleanType), nil
		}),

		// isArray(x)
		fn("isArray", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isArray() expects 1 argument, got %d", len(args))
			}
			return environment.NewBool(args[0].Type == environment.ArrayType), nil
		}),

		// isObject(x)
		fn("isObject", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isObject() expects 1 argument, got %d", len(args))
			}
			return environment.NewBool(args[0].Type == environment.ObjectType), nil
		}),

		// isNil(x)
		fn("isNil", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isNil() expects 1 argument, got %d", len(args))
			}
			return environment.NewBool(args[0].Type == environment.NilType), nil
		}),

		// isFunction(x)
		fn("isFunction", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isFunction() expects 1 argument, got %d", len(args))
			}
			return environment.NewBool(args[0].Type == environment.FunctionType || args[0].Type == environment.BuiltinFnType), nil
		}),

		// toInt(x) — truncates number or parses string to integer
		fn("toInt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("toInt() expects 1 argument, got %d", len(args))
			}
			switch args[0].Type {
			case environment.NumberType:
				return environment.NewNumber(math.Trunc(args[0].Num)), nil
			case environment.StringType:
				n, err := strconv.ParseFloat(args[0].Str, 64)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("toInt() cannot convert %q to integer", args[0].Str)
				}
				return environment.NewNumber(math.Trunc(n)), nil
			case environment.BooleanType:
				if args[0].Bool {
					return environment.NewNumber(1), nil
				}
				return environment.NewNumber(0), nil
			default:
				return environment.NewNil(), fmt.Errorf("toInt() cannot convert %s to integer", args[0].TypeName())
			}
		}),

		// toFloat(x) — converts to float
		fn("toFloat", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("toFloat() expects 1 argument, got %d", len(args))
			}
			switch args[0].Type {
			case environment.NumberType:
				return args[0], nil
			case environment.StringType:
				n, err := strconv.ParseFloat(args[0].Str, 64)
				if err != nil {
					return environment.NewNil(), fmt.Errorf("toFloat() cannot convert %q to float", args[0].Str)
				}
				return environment.NewNumber(n), nil
			case environment.BooleanType:
				if args[0].Bool {
					return environment.NewNumber(1), nil
				}
				return environment.NewNumber(0), nil
			default:
				return environment.NewNil(), fmt.Errorf("toFloat() cannot convert %s to float", args[0].TypeName())
			}
		}),

		// toString(x) — converts any value to string
		fn("toString", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("toString() expects 1 argument, got %d", len(args))
			}
			return environment.NewString(args[0].String()), nil
		}),

		// toBool(x) — converts to boolean (falsy: 0, "", null, false; truthy: everything else)
		fn("toBool", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("toBool() expects 1 argument, got %d", len(args))
			}
			switch args[0].Type {
			case environment.BooleanType:
				return args[0], nil
			case environment.NilType:
				return environment.NewBool(false), nil
			case environment.NumberType:
				return environment.NewBool(args[0].Num != 0), nil
			case environment.StringType:
				return environment.NewBool(args[0].Str != ""), nil
			default:
				return environment.NewBool(true), nil
			}
		}),
	))
}
