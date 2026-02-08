package builtins

import (
	"fmt"
	"strings"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("strings",
		// len(s)
		fn("len", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("len() expects 1 argument, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("len() argument must be a string")
			}
			return environment.NewNumber(float64(len([]rune(s)))), nil
		}),

		// upper(s)
		fn("upper", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("upper() expects 1 argument, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("upper() argument must be a string")
			}
			return environment.NewString(strings.ToUpper(s)), nil
		}),

		// lower(s)
		fn("lower", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("lower() expects 1 argument, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("lower() argument must be a string")
			}
			return environment.NewString(strings.ToLower(s)), nil
		}),

		// trim(s)
		fn("trim", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("trim() expects 1 argument, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("trim() argument must be a string")
			}
			return environment.NewString(strings.TrimSpace(s)), nil
		}),

		// split(s, sep)
		fn("split", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("split() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("split() first argument must be a string")
			}
			sep, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("split() second argument must be a string")
			}
			parts := strings.Split(s, sep)
			elems := make([]environment.Value, len(parts))
			for i, p := range parts {
				elems[i] = environment.NewString(p)
			}
			return environment.NewArray(elems), nil
		}),

		// join(arr, sep)
		fn("join", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("join() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("join() first argument must be an array")
			}
			sep, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("join() second argument must be a string")
			}
			parts := make([]string, len(*args[0].Arr))
			for i, v := range *args[0].Arr {
				parts[i] = v.String()
			}
			return environment.NewString(strings.Join(parts, sep)), nil
		}),

		// replace(s, old, new)
		fn("replace", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("replace() expects 3 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("replace() first argument must be a string")
			}
			old, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("replace() second argument must be a string")
			}
			newStr, err := args[2].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("replace() third argument must be a string")
			}
			return environment.NewString(strings.ReplaceAll(s, old, newStr)), nil
		}),

		// contains(s, sub)
		fn("contains", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("contains() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("contains() first argument must be a string")
			}
			sub, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("contains() second argument must be a string")
			}
			return environment.NewBool(strings.Contains(s, sub)), nil
		}),

		// startsWith(s, prefix)
		fn("startsWith", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("startsWith() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("startsWith() first argument must be a string")
			}
			prefix, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("startsWith() second argument must be a string")
			}
			return environment.NewBool(strings.HasPrefix(s, prefix)), nil
		}),

		// endsWith(s, suffix)
		fn("endsWith", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("endsWith() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("endsWith() first argument must be a string")
			}
			suffix, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("endsWith() second argument must be a string")
			}
			return environment.NewBool(strings.HasSuffix(s, suffix)), nil
		}),

		// indexOf(s, sub)
		fn("indexOf", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("indexOf() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("indexOf() first argument must be a string")
			}
			sub, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("indexOf() second argument must be a string")
			}
			return environment.NewNumber(float64(strings.Index(s, sub))), nil
		}),

		// lastIndexOf(s, sub)
		fn("lastIndexOf", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("lastIndexOf() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("lastIndexOf() first argument must be a string")
			}
			sub, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("lastIndexOf() second argument must be a string")
			}
			return environment.NewNumber(float64(strings.LastIndex(s, sub))), nil
		}),

		// substring(s, start, end)
		fn("substring", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("substring() expects 3 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("substring() first argument must be a string")
			}
			start, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("substring() second argument must be a number")
			}
			end, err := args[2].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("substring() third argument must be a number")
			}
			runes := []rune(s)
			si := int(start)
			ei := int(end)
			if si < 0 {
				si = 0
			}
			if ei > len(runes) {
				ei = len(runes)
			}
			if si > ei {
				return environment.NewString(""), nil
			}
			return environment.NewString(string(runes[si:ei])), nil
		}),

		// charAt(s, i)
		fn("charAt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("charAt() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("charAt() first argument must be a string")
			}
			idx, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("charAt() second argument must be a number")
			}
			runes := []rune(s)
			i := int(idx)
			if i < 0 || i >= len(runes) {
				return environment.NewNil(), fmt.Errorf("charAt() index %d out of range (length %d)", i, len(runes))
			}
			return environment.NewString(string(runes[i])), nil
		}),

		// repeat(s, n)
		fn("repeat", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("repeat() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("repeat() first argument must be a string")
			}
			n, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("repeat() second argument must be a number")
			}
			return environment.NewString(strings.Repeat(s, int(n))), nil
		}),
	))
}
