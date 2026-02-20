package builtins

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/iscarloscoder/fig/environment"
)

// Script runtime state injected by CLI before running a script.
// These are set by `main.runFile` and read by builtins (e.g., system.argv()/cwd()).
// ScriptFile holds the path of the currently executing .fig source. It is
// updated by the interpreter when running or importing files, and is used by
// runtime.file() / runtime.dir().
var (
	ScriptArgs []string
	ScriptCwd  string
	ScriptFile string
)

func init() {
	register(newModule("runtime",
		// gc() — forces a garbage collection
		fn("gc", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("gc() expects 0 arguments, got %d", len(args))
			}
			runtime.GC()
			return environment.NewNil(), nil
		}),

		// memUsage() — returns an object with memory stats in bytes
		fn("memUsage", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("memUsage() expects 0 arguments, got %d", len(args))
			}
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			entries := map[string]environment.Value{
				"alloc":      environment.NewNumber(float64(m.Alloc)),
				"totalAlloc": environment.NewNumber(float64(m.TotalAlloc)),
				"sys":        environment.NewNumber(float64(m.Sys)),
				"numGC":      environment.NewNumber(float64(m.NumGC)),
			}
			keys := []string{"alloc", "totalAlloc", "sys", "numGC"}
			return environment.NewObject(entries, keys), nil
		}),

		// version() — returns the Go runtime version
		fn("version", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("version() expects 0 arguments, got %d", len(args))
			}
			return environment.NewString(runtime.Version()), nil
		}),

		// platform() — returns GOOS/GOARCH
		fn("platform", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("platform() expects 0 arguments, got %d", len(args))
			}
			return environment.NewString(runtime.GOOS + "/" + runtime.GOARCH), nil
		}),

		// numCPU() — returns the number of logical CPUs
		fn("numCPU", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("numCPU() expects 0 arguments, got %d", len(args))
			}
			return environment.NewNumber(float64(runtime.NumCPU())), nil
		}),

		// file() — absolute path of the currently executing source file
		fn("file", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("file() expects 0 arguments, got %d", len(args))
			}
			return environment.NewString(ScriptFile), nil
		}),

		// dir() — directory containing the currently executing source file
		fn("dir", func(args []environment.Value) (environment.Value, error) {
			if len(args) != 0 {
				return environment.NewNil(), fmt.Errorf("dir() expects 0 arguments, got %d", len(args))
			}
			if ScriptFile == "" {
				return environment.NewString(""), nil
			}
			return environment.NewString(filepath.Dir(ScriptFile)), nil
		}),
	))
}
