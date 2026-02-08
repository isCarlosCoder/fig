package builtins

import (
	"fmt"
	"sync"

	"github.com/iscarloscoder/fig/environment"
)

// FnCaller is set by the interpreter so that functional helpers can invoke
// user-defined Fig functions (FunctionType) in addition to builtins.
// It receives the function value and arguments, returns the result and any error.
var FnCaller func(fn environment.Value, args []environment.Value) (environment.Value, error)

// isCallable checks whether a value is a function (builtin or user-defined).
func isCallable(v environment.Value) bool {
	return v.Type == environment.BuiltinFnType || v.Type == environment.FunctionType
}

// invokeFn calls a function value (builtin or user-defined) with the given args.
func invokeFn(fnVal environment.Value, args []environment.Value) (environment.Value, error) {
	if fnVal.Type == environment.BuiltinFnType {
		return fnVal.Builtin(args)
	}
	if fnVal.Type == environment.FunctionType {
		if FnCaller == nil {
			return environment.NewNil(), fmt.Errorf("interpreter callback not set for user functions")
		}
		return FnCaller(fnVal, args)
	}
	return environment.NewNil(), fmt.Errorf("value is not a function")
}

// fnLabel returns a display name for a function value.
func fnLabel(v environment.Value) string {
	if v.Type == environment.BuiltinFnType {
		return v.BName
	}
	if v.Type == environment.FunctionType && v.Func != nil {
		if v.Func.Name != "" {
			return v.Func.Name
		}
	}
	return "<fn>"
}

func init() {
	register(newModule("functional",
		// call(fn, ...args) — calls a function with given args
		fn("call", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("call() expects at least 1 argument, got %d", len(args))
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("call() first argument must be a function")
			}
			return invokeFn(args[0], args[1:])
		}),

		// apply(fn, arr) — calls a function with array as arguments
		fn("apply", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("apply() expects 2 arguments, got %d", len(args))
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("apply() first argument must be a function")
			}
			if args[1].Type != environment.ArrayType || args[1].Arr == nil {
				return environment.NewNil(), fmt.Errorf("apply() second argument must be an array")
			}
			return invokeFn(args[0], *args[1].Arr)
		}),

		// partial(fn, ...bound) — returns a new function with the first args pre-filled
		fn("partial", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 2 {
				return environment.NewNil(), fmt.Errorf("partial() expects at least 2 arguments, got %d", len(args))
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("partial() first argument must be a function")
			}
			origFn := args[0]
			bound := make([]environment.Value, len(args)-1)
			copy(bound, args[1:])
			wrapped := func(rest []environment.Value) (environment.Value, error) {
				all := make([]environment.Value, len(bound)+len(rest))
				copy(all, bound)
				copy(all[len(bound):], rest)
				return invokeFn(origFn, all)
			}
			return environment.NewBuiltinFn("partial("+fnLabel(args[0])+")", environment.BuiltinFn(wrapped)), nil
		}),

		// once(fn) — returns a function that only executes fn once; subsequent calls return the first result
		fn("once", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("once() expects 1 argument, got %d", len(args))
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("once() argument must be a function")
			}
			origFn := args[0]
			var onceCtrl sync.Once
			var result environment.Value
			var resultErr error
			wrapped := func(callArgs []environment.Value) (environment.Value, error) {
				onceCtrl.Do(func() {
					result, resultErr = invokeFn(origFn, callArgs)
				})
				return result, resultErr
			}
			return environment.NewBuiltinFn("once("+fnLabel(args[0])+")", environment.BuiltinFn(wrapped)), nil
		}),

		// memo(fn) — returns a memoized version (caches results by string key of args)
		fn("memo", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("memo() expects 1 argument, got %d", len(args))
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("memo() argument must be a function")
			}
			origFn := args[0]
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
				val, err := invokeFn(origFn, callArgs)
				if err != nil {
					return val, err
				}
				cache[key] = val
				return val, nil
			}
			return environment.NewBuiltinFn("memo("+fnLabel(args[0])+")", environment.BuiltinFn(wrapped)), nil
		}),
	))
}
