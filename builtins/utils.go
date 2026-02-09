package builtins

import (
	"fmt"
	"unicode"

	"github.com/iscarloscoder/fig/environment"
	"golang.org/x/text/unicode/norm"
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

		// fromCodePoint(n) alias to chr
		fn("fromCodePoint", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("fromCodePoint() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("fromCodePoint() argument must be a number")
			}
			ni := int(n)
			if float64(ni) != n {
				return environment.NewNil(), fmt.Errorf("fromCodePoint() argument must be an integer code point")
			}
			if ni < 0 || ni > 0x10FFFF {
				return environment.NewNil(), fmt.Errorf("fromCodePoint() code point out of range")
			}
			return environment.NewString(string(rune(ni))), nil
		}),

		// codePointAt(s, index) -> returns code point (number) at rune index
		fn("codePointAt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("codePointAt() expects 2 arguments, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("codePointAt() first argument must be a string")
			}
			i, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("codePointAt() second argument must be a number")
			}
			ii := int(i)
			if float64(ii) != i {
				return environment.NewNil(), fmt.Errorf("codePointAt() index must be an integer")
			}
			r := []rune(s)
			if ii < 0 || ii >= len(r) {
				return environment.NewNil(), fmt.Errorf("codePointAt() index out of range")
			}
			return environment.NewNumber(float64(r[ii])), nil
		}),

		// codePoints(s) -> returns array of code points
		fn("codePoints", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("codePoints() expects 1 argument, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("codePoints() argument must be a string")
			}
			r := []rune(s)
			arr := make([]environment.Value, len(r))
			for i, ru := range r {
				arr[i] = environment.NewNumber(float64(ru))
			}
			return environment.NewArray(arr), nil
		}),

		// normalize(form, s) -> NFC/NFD/NFKC/NFKD
		fn("normalize", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("normalize() expects 2 arguments, got %d", len(args))
			}
			form, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("normalize() first argument must be a string (form)")
			}
			s, err := args[1].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("normalize() second argument must be a string")
			}
			switch form {
			case "NFC":
				return environment.NewString(norm.NFC.String(s)), nil
			case "NFD":
				return environment.NewString(norm.NFD.String(s)), nil
			case "NFKC":
				return environment.NewString(norm.NFKC.String(s)), nil
			case "NFKD":
				return environment.NewString(norm.NFKD.String(s)), nil
			default:
				return environment.NewNil(), fmt.Errorf("normalize() unknown form: %s", form)
			}
		}),

		// isLetter(s) and isDigit(s)
		fn("isLetter", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isLetter() expects 1 argument, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("isLetter() argument must be a string")
			}
			r := []rune(s)
			if len(r) == 0 {
				return environment.NewBool(false), nil
			}
			return environment.NewBool(unicode.IsLetter(r[0])), nil
		}),

		fn("isDigit", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("isDigit() expects 1 argument, got %d", len(args))
			}
			s, err := args[0].AsString()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("isDigit() argument must be a string")
			}
			r := []rune(s)
			if len(r) == 0 {
				return environment.NewBool(false), nil
			}
			return environment.NewBool(unicode.IsDigit(r[0])), nil
		}),
	))
}
