package builtins

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/iscarloscoder/fig/environment"
)

func init() {
	register(newModule("system",
		// now() — returns current Unix timestamp in milliseconds
		fn("now", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("now() expects 0 arguments, got %d", len(args))
			}
			ms := float64(time.Now().UnixMilli())
			return environment.NewNumber(ms), nil
		}),

		// sleep(ms) — pauses execution for ms milliseconds
		fn("sleep", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("sleep() expects 1 argument, got %d", len(args))
			}
			ms, err := args[0].AsNumber()
			if err != nil {
				return environment.NewNil(), fmt.Errorf("sleep() argument must be a number")
			}
			time.Sleep(time.Duration(ms) * time.Millisecond)
			return environment.NewNil(), nil
		}),

		// clock() — returns high-resolution time in seconds (float)
		fn("clock", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("clock() expects 0 arguments, got %d", len(args))
			}
			secs := float64(time.Now().UnixNano()) / 1e9
			return environment.NewNumber(secs), nil
		}),

		// env(name) — returns environment variable or null
		fn("env", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("env() expects 1 argument, got %d", len(args))
			}
			if args[0].Type != environment.StringType {
				return environment.NewNil(), fmt.Errorf("env() argument must be a string")
			}
			val, ok := os.LookupEnv(args[0].Str)
			if !ok {
				return environment.NewNil(), nil
			}
			return environment.NewString(val), nil
		}),

		// exit(code) — exits the process with given code
		// Uses panic(ExitSignal) so the signal can be caught by the test runner.
		fn("exit", func(args []environment.Value) (environment.Value, error) {
			code := 0
			if len(args) == 1 {
				n, err := args[0].AsNumber()
				if err != nil {
					return environment.NewNil(), fmt.Errorf("exit() argument must be a number")
				}
				code = int(n)
			}
			panic(environment.ExitSignal{Code: code})
		}),

		// args() — returns command-line arguments as array
		fn("args", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("args() expects 0 arguments, got %d", len(args))
			}
			osArgs := os.Args
			result := make([]environment.Value, len(osArgs))
			for i, a := range osArgs {
				result[i] = environment.NewString(a)
			}
			return environment.NewArray(result), nil
		}),

		// platform() — returns the OS name
		fn("platform", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("platform() expects 0 arguments, got %d", len(args))
			}
			return environment.NewString(runtime.GOOS), nil
		}),

		// version() — returns the Fig language version
		fn("version", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("version() expects 0 arguments, got %d", len(args))
			}
			return environment.NewString("0.1.0"), nil
		}),
	))
}
