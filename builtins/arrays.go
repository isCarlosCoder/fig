package builtins

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("arrays",
		// push(arr, v) — appends v, returns the array
		fn("push", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("push() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("push() first argument must be an array")
			}
			*args[0].Arr = append(*args[0].Arr, args[1])
			return args[0], nil
		}),

		// pop(arr) — removes and returns last element
		fn("pop", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("pop() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("pop() argument must be an array")
			}
			arr := *args[0].Arr
			if len(arr) == 0 {
				return environment.NewNil(), fmt.Errorf("pop() on empty array")
			}
			last := arr[len(arr)-1]
			*args[0].Arr = arr[:len(arr)-1]
			return last, nil
		}),

		// shift(arr) — removes and returns first element
		fn("shift", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("shift() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("shift() argument must be an array")
			}
			arr := *args[0].Arr
			if len(arr) == 0 {
				return environment.NewNil(), fmt.Errorf("shift() on empty array")
			}
			first := arr[0]
			*args[0].Arr = arr[1:]
			return first, nil
		}),

		// unshift(arr, v) — prepends v, returns the array
		fn("unshift", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("unshift() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("unshift() first argument must be an array")
			}
			*args[0].Arr = append([]environment.Value{args[1]}, *args[0].Arr...)
			return args[0], nil
		}),

		// insert(arr, i, v) — inserts v at index i
		fn("insert", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("insert() expects 3 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("insert() first argument must be an array")
			}
			idx, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("insert() second argument must be a number")
			}
			i := int(idx)
			arr := *args[0].Arr
			if i < 0 || i > len(arr) {
				return environment.NewNil(), fmt.Errorf("insert() index %d out of range (length %d)", i, len(arr))
			}
			newArr := make([]environment.Value, len(arr)+1)
			copy(newArr, arr[:i])
			newArr[i] = args[2]
			copy(newArr[i+1:], arr[i:])
			*args[0].Arr = newArr
			return args[0], nil
		}),

		// remove(arr, i) — removes element at index i, returns it
		fn("remove", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("remove() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("remove() first argument must be an array")
			}
			idx, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("remove() second argument must be a number")
			}
			i := int(idx)
			arr := *args[0].Arr
			if i < 0 || i >= len(arr) {
				return environment.NewNil(), fmt.Errorf("remove() index %d out of range (length %d)", i, len(arr))
			}
			removed := arr[i]
			*args[0].Arr = append(arr[:i], arr[i+1:]...)
			return removed, nil
		}),

		// slice(arr, start, end) — returns a new sub-array
		fn("slice", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("slice() expects 3 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("slice() first argument must be an array")
			}
			start, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("slice() second argument must be a number")
			}
			end, err := args[2].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("slice() third argument must be a number")
			}
			arr := *args[0].Arr
			si, ei := int(start), int(end)
			if si < 0 {
				si = 0
			}
			if ei > len(arr) {
				ei = len(arr)
			}
			if si > ei {
				return environment.NewArray([]environment.Value{}), nil
			}
			result := make([]environment.Value, ei-si)
			copy(result, arr[si:ei])
			return environment.NewArray(result), nil
		}),

		// concat(a, b) — returns a new array combining both
		fn("concat", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("concat() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("concat() first argument must be an array")
			}
			if args[1].Type != environment.ArrayType || args[1].Arr == nil {
				return environment.NewNil(), fmt.Errorf("concat() second argument must be an array")
			}
			a := *args[0].Arr
			b := *args[1].Arr
			result := make([]environment.Value, len(a)+len(b))
			copy(result, a)
			copy(result[len(a):], b)
			return environment.NewArray(result), nil
		}),

		// reverse(arr) — reverses in place, returns the array
		fn("reverse", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("reverse() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("reverse() argument must be an array")
			}
			arr := *args[0].Arr
			for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
				arr[i], arr[j] = arr[j], arr[i]
			}
			return args[0], nil
		}),

		// sort(arr) — sorts numbers/strings in place, returns the array
		fn("sort", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sort() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("sort() argument must be an array")
			}
			arr := *args[0].Arr
			sort.SliceStable(arr, func(i, j int) bool {
				a, b := arr[i], arr[j]
				if a.Type == environment.NumberType && b.Type == environment.NumberType {
					return a.Num < b.Num
				}
				return a.String() < b.String()
			})
			return args[0], nil
		}),

		// map(arr, fn) — returns new array with fn applied to each element
		fn("map", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("map() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("map() first argument must be an array")
			}
			if args[1].Type != environment.BuiltinFnType && args[1].Type != environment.FunctionType {
				return environment.NewNil(), fmt.Errorf("map() second argument must be a function")
			}
			// map needs the interpreter to call functions — store callback for later
			return environment.NewNil(), fmt.Errorf("map() requires interpreter callback (use for loop instead for now)")
		}),

		// filter(arr, fn)
		fn("filter", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("filter() expects 2 arguments, got %d", len(args))
			}
			return environment.NewNil(), fmt.Errorf("filter() requires interpreter callback (use for loop instead for now)")
		}),

		// reduce(arr, fn, init)
		fn("reduce", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("reduce() expects 3 arguments, got %d", len(args))
			}
			return environment.NewNil(), fmt.Errorf("reduce() requires interpreter callback (use for loop instead for now)")
		}),

		// find(arr, fn)
		fn("find", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("find() expects 2 arguments, got %d", len(args))
			}
			return environment.NewNil(), fmt.Errorf("find() requires interpreter callback (use for loop instead for now)")
		}),

		// index(arr, v) — returns index of v or -1
		fn("index", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("index() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("index() first argument must be an array")
			}
			for i, v := range *args[0].Arr {
				if valuesEqualBuiltin(v, args[1]) {
					return environment.NewNumber(float64(i)), nil
				}
			}
			return environment.NewNumber(-1), nil
		}),

		// contains(arr, v) — returns bool
		fn("contains", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("contains() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("contains() first argument must be an array")
			}
			for _, v := range *args[0].Arr {
				if valuesEqualBuiltin(v, args[1]) {
					return environment.NewBool(true), nil
				}
			}
			return environment.NewBool(false), nil
		}),

		// unique(arr) — returns new array with duplicates removed
		fn("unique", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("unique() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("unique() argument must be an array")
			}
			var result []environment.Value
			for _, v := range *args[0].Arr {
				found := false
				for _, r := range result {
					if valuesEqualBuiltin(v, r) {
						found = true
						break
					}
				}
				if !found {
					result = append(result, v)
				}
			}
			return environment.NewArray(result), nil
		}),

		// shuffle(arr) — shuffles in place, returns the array
		fn("shuffle", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("shuffle() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("shuffle() argument must be an array")
			}
			arr := *args[0].Arr
			rand.Shuffle(len(arr), func(i, j int) {
				arr[i], arr[j] = arr[j], arr[i]
			})
			return args[0], nil
		}),

		// len(arr) — returns length
		fn("len", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("len() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ArrayType || args[0].Arr == nil {
				return environment.NewNil(), fmt.Errorf("len() argument must be an array")
			}
			return environment.NewNumber(float64(len(*args[0].Arr))), nil
		}),
	))
}

// valuesEqualBuiltin compares two Values for equality (used by index, contains, unique).
func valuesEqualBuiltin(a, b environment.Value) bool {
	if a.Type != b.Type {
		return false
	}
	switch a.Type {
	case environment.NumberType:
		return a.Num == b.Num
	case environment.StringType:
		return a.Str == b.Str
	case environment.BooleanType:
		return a.Bool == b.Bool
	case environment.NilType:
		return true
	default:
		return a.String() == b.String()
	}
}
