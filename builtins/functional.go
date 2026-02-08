package builtins

import (
	"fmt"
	"sync"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("functional",
		// call(fn, ...args) — calls a builtin function with given args
		fn("call", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("call() expects at least 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.BuiltinFnType {
				return environment.NewNil(), fmt.Errorf("call() first argument must be a builtin function")
			}
			return args[0].Builtin(args[1:])
		}),

		// apply(fn, arr) — calls a builtin function with array as arguments
		fn("apply", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("apply() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.BuiltinFnType {
				return environment.NewNil(), fmt.Errorf("apply() first argument must be a builtin function")
			}
			if args[1].Type != environment.ArrayType || args[1].Arr == nil {
				return environment.NewNil(), fmt.Errorf("apply() second argument must be an array")
			}
			return args[0].Builtin(*args[1].Arr)
		}),

		// partial(fn, ...bound) — returns a new function with the first args pre-filled
		fn("partial", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 2 {
				return environment.NewNil(), fmt.Errorf("partial() expects at least 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.BuiltinFnType {
				return environment.NewNil(), fmt.Errorf("partial() first argument must be a builtin function")
			}
			origFn := args[0].Builtin
			bound := make([]environment.Value, len(args)-1)
			copy(bound, args[1:])
			wrapped := func(rest []environment.Value) (environment.Value, error) {
				all := make([]environment.Value, len(bound)+len(rest))
				copy(all, bound)
				copy(all[len(bound):], rest)
				return origFn(all)
			}
			return environment.NewBuiltinFn("partial("+args[0].BName+")", environment.BuiltinFn(wrapped)), nil
		}),

		// once(fn) — returns a function that only executes fn once; subsequent calls return the first result
		fn("once", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("once() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.BuiltinFnType {
				return environment.NewNil(), fmt.Errorf("once() argument must be a builtin function")
			}
			origFn := args[0].Builtin
			var onceCtrl sync.Once
			var result environment.Value
			var resultErr error
			wrapped := func(callArgs []environment.Value) (environment.Value, error) {
				onceCtrl.Do(func() {
					result, resultErr = origFn(callArgs)
				})
				return result, resultErr
			}
			return environment.NewBuiltinFn("once("+args[0].BName+")", environment.BuiltinFn(wrapped)), nil
		}),

		// memo(fn) — returns a memoized version (caches results by string key of args)
		fn("memo", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("memo() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.BuiltinFnType {
				return environment.NewNil(), fmt.Errorf("memo() argument must be a builtin function")
			}
			origFn := args[0].Builtin
			cache := make(map[string]environment.Value)
			wrapped := func(callArgs []environment.Value) (environment.Value, error) {
				key := ""
				for i, a := range callArgs {
					if i > 0 {
						key += ","
					}
					key += a.String()
				}
				if val, ok := cache[key]; ok {
					return val, nil
				}
				val, err := origFn(callArgs)
				if err != nil {
					return val, err
				}
				cache[key] = val
				return val, nil
			}
			return environment.NewBuiltinFn("memo("+args[0].BName+")", environment.BuiltinFn(wrapped)), nil
		}),
	))
}
