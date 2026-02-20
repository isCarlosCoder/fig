package builtins

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("path",
		// join(...segments)
		fn("join", func(args []environment.Value) (environment.Value, error) {
			if len(args) < 1 {
				return environment.NewNil(), fmt.Errorf("join() expects at least 1 argument, got %d", len(args))
			}
			parts := make([]string, len(args))
			for i, a := range args {
				if a.Type != environment.StringType {
					return environment.NewNil(), fmt.Errorf("join() argument #%d must be a string", i+1)
				}
				parts[i] = a.Str
			}
			return environment.NewString(filepath.Join(parts...)), nil
		}),

		// base(p)
		fn("base", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("base() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("base() argument must be a string")
			}
			return environment.NewString(filepath.Base(args[0].Str)), nil
		}),

		// dir(p)
		fn("dir", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("dir() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("dir() argument must be a string")
			}
			return environment.NewString(filepath.Dir(args[0].Str)), nil
		}),

		// ext(p)
		fn("ext", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ext() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("ext() argument must be a string")
			}
			return environment.NewString(filepath.Ext(args[0].Str)), nil
		}),

		// abs(p)
		fn("abs", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("abs() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("abs() argument must be a string")
			}
			p := args[0].Str
			a, err := filepath.Abs(p)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("abs(): %v", err)
			}
			return environment.NewString(a), nil
		}),

		// clean(p)
		fn("clean", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("clean() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("clean() argument must be a string")
			}
			return environment.NewString(filepath.Clean(args[0].Str)), nil
		}),

		// isAbs(p)
		fn("isAbs", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isAbs() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("isAbs() argument must be a string")
			}
			return environment.NewBool(filepath.IsAbs(args[0].Str)), nil
		}),

		// rel(base, target)
		fn("rel", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("rel() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType || args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("rel() arguments must be strings")
			}
			rel, err := filepath.Rel(args[0].Str, args[1].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("rel(): %v", err)
			}
			return environment.NewString(rel), nil
		}),

		// split(p)
		fn("split", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("split() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("split() argument must be a string")
			}
			dir, file := filepath.Split(args[0].Str)
			return environment.NewArray([]environment.Value{environment.NewString(dir), environment.NewString(file)}), nil
		}),

		// splitExt(p)
		fn("splitExt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("splitExt() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("splitExt() argument must be a string")
			}
			p := args[0].Str
			ext := filepath.Ext(p)
			root := strings.TrimSuffix(p, ext)
			return environment.NewArray([]environment.Value{environment.NewString(root), environment.NewString(ext)}), nil
		}),

		// exists(p) - same semantics as io.exists but placed here for convenience
		fn("exists", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("exists() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("exists() path must be a string")
			}
			_, err := os.Stat(args[0].Str)
			return environment.NewBool(!os.IsNotExist(err)), nil
		}),

		// real(p) - evaluate symlinks and return absolute cleaned path
		fn("real", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("real() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("real() argument must be a string")
			}
			r, err := filepath.EvalSymlinks(args[0].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("real(): %v", err)
			}
			return environment.NewString(r), nil
		}),
	))
}
