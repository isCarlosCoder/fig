package builtins

import (
	"fmt"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("objects",
		// keys(obj) — returns array of keys
		fn("keys", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("keys() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("keys() argument must be an object")
			}
			keys := args[0].Obj.Keys
			result := make([]environment.Value, len(keys))
			for i, k := range keys {
				result[i] = environment.NewString(k)
			}
			return environment.NewArray(result), nil
		}),

		// values(obj) — returns array of values
		fn("values", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("values() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("values() argument must be an object")
			}
			obj := args[0].Obj
			result := make([]environment.Value, len(obj.Keys))
			for i, k := range obj.Keys {
				result[i] = obj.Entries[k]
			}
			return environment.NewArray(result), nil
		}),

		// entries(obj) — returns array of [key, value] pairs
		fn("entries", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("entries() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("entries() argument must be an object")
			}
			obj := args[0].Obj
			result := make([]environment.Value, len(obj.Keys))
			for i, k := range obj.Keys {
				pair := []environment.Value{environment.NewString(k), obj.Entries[k]}
				result[i] = environment.NewArray(pair)
			}
			return environment.NewArray(result), nil
		}),

		// hasKey(obj, key) — returns bool
		fn("hasKey", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("hasKey() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("hasKey() first argument must be an object")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("hasKey() second argument must be a string")
			}
			_, ok := args[0].Obj.Entries[args[1].Str]
			return environment.NewBool(ok), nil
		}),

		// deleteKey(obj, key) — removes key, returns deleted value (or nil)
		fn("deleteKey", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("deleteKey() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("deleteKey() first argument must be an object")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("deleteKey() second argument must be a string")
			}
			obj := args[0].Obj
			key := args[1].Str
			val, ok := obj.Entries[key]
			if !ok {
				return environment.NewNil(), nil
			}
			delete(obj.Entries, key)
			newKeys := make([]string, 0, len(obj.Keys)-1)
			for _, k := range obj.Keys {
				if k != key {
					newKeys = append(newKeys, k)
				}
			}
			obj.Keys = newKeys
			return val, nil
		}),

		// merge(a, b) — returns new object with keys from both (b overwrites a)
		fn("merge", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("merge() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("merge() first argument must be an object")
			}
			if args[1].Type != environment.ObjectType || args[1].Obj == nil {
				return environment.NewNil(), fmt.Errorf("merge() second argument must be an object")
			}
			a, b := args[0].Obj, args[1].Obj
			entries := make(map[string]environment.Value)
			var keys []string
			for _, k := range a.Keys {
				entries[k] = a.Entries[k]
				keys = append(keys, k)
			}
			for _, k := range b.Keys {
				if _, exists := entries[k]; !exists {
					keys = append(keys, k)
				}
				entries[k] = b.Entries[k]
			}
			return environment.NewObject(entries, keys), nil
		}),

		// clone(obj) — returns a shallow copy
		fn("clone", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("clone() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("clone() argument must be an object")
			}
			obj := args[0].Obj
			entries := make(map[string]environment.Value, len(obj.Entries))
			keys := make([]string, len(obj.Keys))
			copy(keys, obj.Keys)
			for k, v := range obj.Entries {
				entries[k] = v
			}
			return environment.NewObject(entries, keys), nil
		}),

		// size(obj) — returns number of keys
		fn("size", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("size() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("size() argument must be an object")
			}
			return environment.NewNumber(float64(len(args[0].Obj.Keys))), nil
		}),

		// clear(obj) — removes all keys, returns the object
		fn("clear", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("clear() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.ObjectType || args[0].Obj == nil {
				return environment.NewNil(), fmt.Errorf("clear() argument must be an object")
			}
			obj := args[0].Obj
			obj.Entries = make(map[string]environment.Value)
			obj.Keys = nil
			return args[0], nil
		}),
	))
}
