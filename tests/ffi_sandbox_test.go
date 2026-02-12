package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
)

// TestSandboxMaxLibsLimit verifies that the max_libs sandbox policy blocks excessive loads.
func TestSandboxMaxLibsLimit(t *testing.T) {
	builtins.StopAllHelpers()

	root := repoRootForTest(t)
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	// Build test lib
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	// Set max_libs = 2
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n\n[ffi.sandbox]\nmax_libs = 2\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	load := mod.Entries["load"]

	// First two loads should succeed
	_, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("first load failed: %v", err)
	}
	_, err = load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("second load failed: %v", err)
	}

	// Third load should be blocked by sandbox
	_, err = load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err == nil {
		t.Fatal("expected sandbox error for third load, got nil")
	}
	if !strings.Contains(err.Error(), "sandbox: max loaded libraries limit") {
		t.Fatalf("expected sandbox limit error, got: %v", err)
	}
}

// TestSandboxMaxAllocsLimit verifies that the max_allocs sandbox policy blocks excessive allocations.
func TestSandboxMaxAllocsLimit(t *testing.T) {
	builtins.StopAllHelpers()

	root := repoRootForTest(t)
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	// Set max_allocs = 2
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n\n[ffi.sandbox]\nmax_allocs = 2\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	alloc := mod.Entries["alloc"]
	free := mod.Entries["free"]

	// Alloc 1
	m1, err := alloc.Builtin([]environment.Value{environment.NewNumber(64)})
	if err != nil {
		t.Fatalf("alloc 1 failed: %v", err)
	}
	// Alloc 2
	m2, err := alloc.Builtin([]environment.Value{environment.NewNumber(64)})
	if err != nil {
		t.Fatalf("alloc 2 failed: %v", err)
	}
	// Alloc 3 should fail
	_, err = alloc.Builtin([]environment.Value{environment.NewNumber(64)})
	if err == nil {
		t.Fatal("expected sandbox error for third alloc, got nil")
	}
	if !strings.Contains(err.Error(), "sandbox: max live allocations limit") {
		t.Fatalf("expected sandbox limit error, got: %v", err)
	}

	// Free one, then alloc should succeed again
	_, err = free.Builtin([]environment.Value{m1})
	if err != nil {
		t.Fatalf("free failed: %v", err)
	}
	_, err = alloc.Builtin([]environment.Value{environment.NewNumber(64)})
	if err != nil {
		t.Fatalf("alloc after free should succeed: %v", err)
	}

	// Cleanup
	_, _ = free.Builtin([]environment.Value{m2})
}

// TestSandboxMaxRestartsLimit verifies the restart limit policy.
func TestSandboxMaxRestartsLimit(t *testing.T) {
	builtins.StopAllHelpers()

	root := repoRootForTest(t)
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	// Set max_restarts = 2
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n\n[ffi.sandbox]\nmax_restarts = 2\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	ping := mod.Entries["ping"]
	helpCmd := mod.Entries["helper_cmd"]

	// Initial ping should work
	_, err := ping.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("initial ping failed: %v", err)
	}

	// Crash and restart 1
	_, _ = helpCmd.Builtin([]environment.Value{environment.NewString("crash")})
	time.Sleep(100 * time.Millisecond)
	_, err = ping.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("ping after restart 1 should work: %v", err)
	}

	// Crash and restart 2
	_, _ = helpCmd.Builtin([]environment.Value{environment.NewString("crash")})
	time.Sleep(100 * time.Millisecond)
	_, err = ping.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("ping after restart 2 should work: %v", err)
	}

	// Crash 3 — should hit max_restarts limit
	_, _ = helpCmd.Builtin([]environment.Value{environment.NewString("crash")})
	time.Sleep(100 * time.Millisecond)
	_, err = ping.Builtin([]environment.Value{})
	if err == nil {
		t.Fatal("expected restart limit error, got nil")
	}
	if !strings.Contains(err.Error(), "restart limit exceeded") {
		t.Fatalf("expected restart limit error, got: %v", err)
	}
}

// TestSandboxStatusBuiltin verifies the sandbox_status() builtin returns correct data.
func TestSandboxStatusBuiltin(t *testing.T) {
	builtins.StopAllHelpers()

	root := repoRootForTest(t)
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n\n[ffi.sandbox]\nmax_libs = 10\nmax_allocs = 20\nmax_memory_mb = 256\nmax_cpu_seconds = 30\nmax_restarts = 3\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	status := mod.Entries["sandbox_status"]

	res, err := status.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("sandbox_status failed: %v", err)
	}
	if res.Type != environment.ObjectType || res.Obj == nil {
		t.Fatal("expected object from sandbox_status")
	}

	// Check configured values
	checkNum := func(key string, expected float64) {
		v, ok := res.Obj.Entries[key]
		if !ok {
			t.Errorf("missing key %s", key)
			return
		}
		if v.Type != environment.NumberType || v.Num != expected {
			t.Errorf("expected %s=%v, got %v", key, expected, v.Num)
		}
	}
	checkNum("max_libs", 10)
	checkNum("max_allocs", 20)
	checkNum("max_memory_mb", 256)
	checkNum("max_cpu_seconds", 30)
	checkNum("max_restarts", 3)
	checkNum("loaded_libs", 0)
	checkNum("live_allocs", 0)
	checkNum("restarts", 0)
}

// TestSandboxDefaultMaxRestarts verifies defaults when no sandbox config is set.
func TestSandboxDefaultMaxRestarts(t *testing.T) {
	builtins.StopAllHelpers()

	root := repoRootForTest(t)
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	// No sandbox section — defaults apply
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	status := mod.Entries["sandbox_status"]

	res, err := status.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("sandbox_status failed: %v", err)
	}

	// Default max_restarts should be 5
	v, ok := res.Obj.Entries["max_restarts"]
	if !ok || v.Type != environment.NumberType || v.Num != 5 {
		t.Errorf("expected default max_restarts=5, got %v", v)
	}

	// max_libs/max_allocs should be 0 (unlimited)
	v2 := res.Obj.Entries["max_libs"]
	if v2.Num != 0 {
		t.Errorf("expected default max_libs=0, got %v", v2.Num)
	}
}

// TestSandboxStrdupAllocCountTracking verifies strdup increments and free_string decrements alloc counter.
func TestSandboxStrdupAllocCountTracking(t *testing.T) {
	builtins.StopAllHelpers()

	root := repoRootForTest(t)
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n\n[ffi.sandbox]\nmax_allocs = 3\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	strdup := mod.Entries["strdup"]
	freeStr := mod.Entries["free_string"]
	status := mod.Entries["sandbox_status"]

	// Strdup 1
	m1, err := strdup.Builtin([]environment.Value{environment.NewString("hello")})
	if err != nil {
		t.Fatalf("strdup 1 failed: %v", err)
	}

	// Check counter
	s, _ := status.Builtin([]environment.Value{})
	if a := s.Obj.Entries["live_allocs"]; a.Num != 1 {
		t.Errorf("expected live_allocs=1, got %v", a.Num)
	}

	// Free and check counter decremented
	_, err = freeStr.Builtin([]environment.Value{m1})
	if err != nil {
		t.Fatalf("free_string failed: %v", err)
	}
	s, _ = status.Builtin([]environment.Value{})
	if a := s.Obj.Entries["live_allocs"]; a.Num != 0 {
		t.Errorf("expected live_allocs=0 after free, got %v", a.Num)
	}
}
