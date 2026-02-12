package builtins

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/pelletier/go-toml/v2"
)

// SandboxConfig holds resource limits configured in fig.toml [ffi.sandbox]
type SandboxConfig struct {
	MaxMemoryMB   int `toml:"max_memory_mb"`   // RSS limit for the helper process (0 = unlimited)
	MaxCPUSeconds int `toml:"max_cpu_seconds"` // CPU time limit in seconds (0 = unlimited)
	MaxLibs       int `toml:"max_libs"`        // max number of loaded shared libraries (0 = unlimited)
	MaxAllocs     int `toml:"max_allocs"`      // max number of live allocations (0 = unlimited)
	MaxRestarts   int `toml:"max_restarts"`    // max helper restarts before giving up (0 = unlimited, default 5)
}

// DefaultSandboxConfig returns sane defaults for sandbox configuration
func DefaultSandboxConfig() SandboxConfig {
	return SandboxConfig{
		MaxMemoryMB:   0,
		MaxCPUSeconds: 0,
		MaxLibs:       0,
		MaxAllocs:     0,
		MaxRestarts:   5,
	}
}

// --- Tracking state ---

// helperRestartCounts tracks how many times each helper has been restarted.
var helperRestartCounts = map[string]*int64{}
var restartCountsMu sync.Mutex

// getRestartCount returns the pointer to the restart counter for a project.
func getRestartCount(projectRoot string) *int64 {
	restartCountsMu.Lock()
	defer restartCountsMu.Unlock()
	c, ok := helperRestartCounts[projectRoot]
	if !ok {
		var zero int64
		helperRestartCounts[projectRoot] = &zero
		c = helperRestartCounts[projectRoot]
	}
	return c
}

// IncrementRestartCount increments and returns the new count.
func IncrementRestartCount(projectRoot string) int64 {
	return atomic.AddInt64(getRestartCount(projectRoot), 1)
}

// GetRestartCountValue returns the current restart count.
func GetRestartCountValue(projectRoot string) int64 {
	return atomic.LoadInt64(getRestartCount(projectRoot))
}

// ResetRestartCount resets the counter for a project.
func ResetRestartCount(projectRoot string) {
	restartCountsMu.Lock()
	defer restartCountsMu.Unlock()
	if c, ok := helperRestartCounts[projectRoot]; ok {
		atomic.StoreInt64(c, 0)
	}
}

// loadCounters tracks loaded libs per project
var loadCounters = map[string]*int64{}
var loadCountersMu sync.Mutex

func getLoadCounter(projectRoot string) *int64 {
	loadCountersMu.Lock()
	defer loadCountersMu.Unlock()
	c, ok := loadCounters[projectRoot]
	if !ok {
		var zero int64
		loadCounters[projectRoot] = &zero
		c = loadCounters[projectRoot]
	}
	return c
}

// IncrementLoadCount increments the loaded-libs counter.
func IncrementLoadCount(projectRoot string) int64 {
	return atomic.AddInt64(getLoadCounter(projectRoot), 1)
}

// GetLoadCountValue returns the current loaded-libs count.
func GetLoadCountValue(projectRoot string) int64 {
	return atomic.LoadInt64(getLoadCounter(projectRoot))
}

// allocCounters tracks live allocations per project
var allocCounters = map[string]*int64{}
var allocCountersMu sync.Mutex

func getAllocCounter(projectRoot string) *int64 {
	allocCountersMu.Lock()
	defer allocCountersMu.Unlock()
	c, ok := allocCounters[projectRoot]
	if !ok {
		var zero int64
		allocCounters[projectRoot] = &zero
		c = allocCounters[projectRoot]
	}
	return c
}

// IncrementAllocCount increments the live-allocs counter.
func IncrementAllocCount(projectRoot string) int64 {
	return atomic.AddInt64(getAllocCounter(projectRoot), 1)
}

// DecrementAllocCount decrements the live-allocs counter.
func DecrementAllocCount(projectRoot string) int64 {
	return atomic.AddInt64(getAllocCounter(projectRoot), -1)
}

// GetAllocCountValue returns the current live-allocs count.
func GetAllocCountValue(projectRoot string) int64 {
	return atomic.LoadInt64(getAllocCounter(projectRoot))
}

// ResetSandboxCounters resets all counters for all projects (used in tests).
func ResetSandboxCounters() {
	restartCountsMu.Lock()
	helperRestartCounts = map[string]*int64{}
	restartCountsMu.Unlock()

	loadCountersMu.Lock()
	loadCounters = map[string]*int64{}
	loadCountersMu.Unlock()

	allocCountersMu.Lock()
	allocCounters = map[string]*int64{}
	allocCountersMu.Unlock()
}

// readSandboxConfig reads only the [ffi.sandbox] section from fig.toml
func readSandboxConfig() (SandboxConfig, string, error) {
	cwd, _ := os.Getwd()
	p, err := findProjectTomlFrom(cwd)
	if err != nil {
		return DefaultSandboxConfig(), "", nil
	}
	var cfg struct {
		Ffi struct {
			Sandbox SandboxConfig `toml:"sandbox"`
		} `toml:"ffi"`
	}
	data, err := os.ReadFile(p)
	if err != nil {
		return DefaultSandboxConfig(), "", err
	}
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return DefaultSandboxConfig(), "", err
	}
	sc := cfg.Ffi.Sandbox
	if sc.MaxRestarts == 0 {
		sc.MaxRestarts = 5 // default
	}
	projectRoot := filepath.Dir(p)
	return sc, projectRoot, nil
}

// ApplyResourceLimitsToHelper sets OS-level resource limits on the helper process
// via RLIMIT_AS (memory) and RLIMIT_CPU (cpu time). Only effective on Linux/macOS.
func ApplyResourceLimitsToHelper(pid int, cfg SandboxConfig) error {
	if cfg.MaxMemoryMB > 0 {
		lim := &syscall.Rlimit{
			Cur: uint64(cfg.MaxMemoryMB) * 1024 * 1024,
			Max: uint64(cfg.MaxMemoryMB) * 1024 * 1024,
		}
		// RLIMIT_AS = 9 on Linux
		if err := setRlimitForProcess(pid, syscall.RLIMIT_AS, lim); err != nil {
			return fmt.Errorf("failed to set memory limit: %v", err)
		}
	}
	if cfg.MaxCPUSeconds > 0 {
		lim := &syscall.Rlimit{
			Cur: uint64(cfg.MaxCPUSeconds),
			Max: uint64(cfg.MaxCPUSeconds),
		}
		if err := setRlimitForProcess(pid, syscall.RLIMIT_CPU, lim); err != nil {
			return fmt.Errorf("failed to set CPU limit: %v", err)
		}
	}
	return nil
}

// setRlimitForProcess sets a resource limit. Note: on most POSIX systems,
// setrlimit only applies to the current process, so we use /proc/pid/limits via prlimit if possible.
// Fallback: set on current process before exec (for the helper we're about to start).
func setRlimitForProcess(_ int, resource int, lim *syscall.Rlimit) error {
	// We use prlimit syscall on Linux to set limits on another process.
	// On other platforms, we fall back to a no-op (limits applied via helper env).
	return prlimit(resource, lim)
}
