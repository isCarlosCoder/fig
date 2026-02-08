package builtins

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("math",
		// abs(x)
		fn("abs", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("abs() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("abs() argument must be a number")
			}
			return environment.NewNumber(math.Abs(n)), nil
		}),

		// pow(a, b)
		fn("pow", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("pow() expects 2 arguments, got %d", len(args))
			}
			a, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("pow() first argument must be a number")
			}
			b, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("pow() second argument must be a number")
			}
			return environment.NewNumber(math.Pow(a, b)), nil
		}),

		// sqrt(x)
		fn("sqrt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sqrt() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("sqrt() argument must be a number")
			}
			return environment.NewNumber(math.Sqrt(n)), nil
		}),

		// cbrt(x)
		fn("cbrt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("cbrt() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cbrt() argument must be a number")
			}
			return environment.NewNumber(math.Cbrt(n)), nil
		}),

		// floor(x)
		fn("floor", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("floor() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("floor() argument must be a number")
			}
			return environment.NewNumber(math.Floor(n)), nil
		}),

		// ceil(x)
		fn("ceil", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("ceil() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("ceil() argument must be a number")
			}
			return environment.NewNumber(math.Ceil(n)), nil
		}),

		// round(x)
		fn("round", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("round() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("round() argument must be a number")
			}
			return environment.NewNumber(math.Round(n)), nil
		}),

		// min(a, b)
		fn("min", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("min() expects 2 arguments, got %d", len(args))
			}
			a, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("min() first argument must be a number")
			}
			b, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("min() second argument must be a number")
			}
			return environment.NewNumber(math.Min(a, b)), nil
		}),

		// max(a, b)
		fn("max", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("max() expects 2 arguments, got %d", len(args))
			}
			a, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("max() first argument must be a number")
			}
			b, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("max() second argument must be a number")
			}
			return environment.NewNumber(math.Max(a, b)), nil
		}),

		// clamp(x, min, max)
		fn("clamp", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 3 {
				return environment.NewNil(), fmt.Errorf("clamp() expects 3 arguments, got %d", len(args))
			}
			x, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("clamp() first argument must be a number")
			}
			lo, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("clamp() second argument must be a number")
			}
			hi, err := args[2].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("clamp() third argument must be a number")
			}
			return environment.NewNumber(math.Max(lo, math.Min(hi, x))), nil
		}),

		// rand() -> [0, 1)
		fn("rand", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("rand() expects 0 arguments, got %d", len(args))
			}
			return environment.NewNumber(rand.Float64()), nil
		}),

		// randInt(min, max)
		fn("randInt", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 2 {
				return environment.NewNil(), fmt.Errorf("randInt() expects 2 arguments, got %d", len(args))
			}
			lo, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("randInt() first argument must be a number")
			}
			hi, err := args[1].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("randInt() second argument must be a number")
			}
			loI := int(lo)
			hiI := int(hi)
			if hiI <= loI {
				return environment.NewNil(), fmt.Errorf("randInt() max must be greater than min")
			}
			return environment.NewNumber(float64(loI + rand.Intn(hiI-loI))), nil
		}),

		// sin(x)
		fn("sin", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sin() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("sin() argument must be a number")
			}
			return environment.NewNumber(math.Sin(n)), nil
		}),

		// cos(x)
		fn("cos", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("cos() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("cos() argument must be a number")
			}
			return environment.NewNumber(math.Cos(n)), nil
		}),

		// tan(x)
		fn("tan", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("tan() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("tan() argument must be a number")
			}
			return environment.NewNumber(math.Tan(n)), nil
		}),

		// log(x) - natural log
		fn("log", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("log() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("log() argument must be a number")
			}
			return environment.NewNumber(math.Log(n)), nil
		}),

		// log10(x)
		fn("log10", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("log10() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("log10() argument must be a number")
			}
			return environment.NewNumber(math.Log10(n)), nil
		}),

		// exp(x)
		fn("exp", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("exp() expects 1 argument, got %d", len(args))
			}
			n, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("exp() argument must be a number")
			}
			return environment.NewNumber(math.Exp(n)), nil
		}),

		// Constants
		num("PI", math.Pi),
		num("E", math.E),
		num("INF", math.Inf(1)),
	))
}
