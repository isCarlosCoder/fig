package builtins

import (
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/iscarloscoder/fig/environment"
)

// step limit control (atomic)
var stepLimitDisabled int32 = 0

func DisableStepLimit() {
	atomic.StoreInt32(&stepLimitDisabled, 1)
}

func EnableStepLimit() {
	atomic.StoreInt32(&stepLimitDisabled, 0)
}

func IsStepLimitDisabled() bool {
	return atomic.LoadInt32(&stepLimitDisabled) != 0
}

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

		// args() — returns command-line arguments as array (OS args)
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

		// disableStepLimit() — disable evaluation step counting (process-wide)
		fn("disableStepLimit", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("disableStepLimit() expects 0 arguments")
			}
			DisableStepLimit()
			return environment.NewNil(), nil
		}),

		// enableStepLimit() — re-enable evaluation step counting
		fn("enableStepLimit", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("enableStepLimit() expects 0 arguments")
			}
			EnableStepLimit()
			return environment.NewNil(), nil
		}),

		// isStepLimitDisabled() -> boolean
		fn("isStepLimitDisabled", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("isStepLimitDisabled() expects 0 args")
			}
			return environment.NewBool(IsStepLimitDisabled()), nil
		}),

		// withoutStepLimit(fn) — execute fn with the step limit temporarily disabled
		// Requires `task` module to be available so the function runs in interpreter goroutine.
		fn("withoutStepLimit", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 1 {
				return environment.NewNil(), fmt.Errorf("withoutStepLimit() expects 1 argument: function")
			}
			if !isCallable(args[0]) {
				return environment.NewNil(), fmt.Errorf("withoutStepLimit() argument must be a function")
			}
			if TaskSpawner == nil {
				return environment.NewNil(), fmt.Errorf("withoutStepLimit() requires the 'task' module to be loaded")
			}

			cb := args[0]
			resultCh := make(chan TaskResult, 1)
			// set disabled, spawn, then wait and re-enable
			DisableStepLimit()
			TaskSpawner(cb, resultCh)
			res := <-resultCh
			EnableStepLimit()
			if res.Err != nil {
				return environment.NewNil(), res.Err
			}
			return res.Value, nil
		}),

		// argv() — script args injected by the CLI when running a script
		fn("argv", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("argv() expects 0 arguments, got %d", len(args))
			}
			arr := make([]environment.Value, len(ScriptArgs))
			for i, s := range ScriptArgs {
				arr[i] = environment.NewString(s)
			}
			return environment.NewArray(arr), nil
		}),

		// cwd() — current working directory when script was executed
		fn("cwd", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("cwd() expects 0 arguments, got %d", len(args))
			}
			return environment.NewString(ScriptCwd), nil
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
