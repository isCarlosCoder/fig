package builtins

import (
	"fmt"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("utils",
		// ord(s) -> code point of single character
		fn("ord", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ord() expects 1 argument, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("ord() argument must be a string")
			}
			r := []rune(s)
			if len(r) != 1 {
				return environment.NewNil(), fmt.Errorf("ord() expects a single character string")
			}
			return environment.NewNumber(float64(r[0])), nil
		}),

		// chr(n) -> character from code point
		fn("chr", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("chr() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("chr() argument must be a number")
			}
			ni := int(n)
			if float64(ni) != n {
				return environment.NewNil(), fmt.Errorf("chr() argument must be an integer code point")
			}
			// Validate range
			if ni < 0 || ni > 0x10FFFF {
				return environment.NewNil(), fmt.Errorf("chr() code point out of range")
			}
			return environment.NewString(string(rune(ni))), nil
		}),
	))
}
