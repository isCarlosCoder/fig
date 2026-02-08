package builtins

import (
	"fmt"
	"regexp"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("regex",
		// match(s, pattern) — returns true if pattern matches s
		fn("match", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("match() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("match() first argument must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("match() second argument must be a string (pattern)")
			}
			re, err := regexp.Compile(args[1].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("match() invalid pattern: %v", err)
			}
			return environment.NewBool(re.MatchString(args[0].Str)), nil
		}),

		// findAll(s, pattern) — returns array of all matches
		fn("findAll", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("findAll() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("findAll() first argument must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("findAll() second argument must be a string (pattern)")
			}
			re, err := regexp.Compile(args[1].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("findAll() invalid pattern: %v", err)
			}
			matches := re.FindAllString(args[0].Str, -1)
			result := make([]environment.Value, len(matches))
			for i, m := range matches {
				result[i] = environment.NewString(m)
			}
			return environment.NewArray(result), nil
		}),

		// replaceRegex(s, pattern, repl) — replaces all matches with repl
		fn("replaceRegex", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("replaceRegex() expects 3 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("replaceRegex() first argument must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("replaceRegex() second argument must be a string (pattern)")
			}
			if args[2].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("replaceRegex() third argument must be a string")
			}
			re, err := regexp.Compile(args[1].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("replaceRegex() invalid pattern: %v", err)
			}
			result := re.ReplaceAllString(args[0].Str, args[2].Str)
			return environment.NewString(result), nil
		}),

		// splitRegex(s, pattern) — splits string by regex pattern
		fn("splitRegex", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("splitRegex() expects 2 arguments, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("splitRegex() first argument must be a string")
			}
			if args[1].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("splitRegex() second argument must be a string (pattern)")
			}
			re, err := regexp.Compile(args[1].Str)
			if err != nil {
				return environment.NewNil(), fmt.Errorf("splitRegex() invalid pattern: %v", err)
			}
			parts := re.Split(args[0].Str, -1)
			result := make([]environment.Value, len(parts))
			for i, p := range parts {
				result[i] = environment.NewString(p)
			}
			return environment.NewArray(result), nil
		}),
	))
}
