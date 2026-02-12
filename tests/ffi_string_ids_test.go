package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
)

// TestFfiStringIDsNoCorruption loads a library and resolves 200 symbols
// (repeatedly the same function) verifying that returned IDs are
// well-formed strings like "sym-N" and that calling through each one
// still returns the correct result.
func TestFfiStringIDsNoCorruption(t *testing.T) {
	builtins.StopAllHelpers()
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\ncall_timeout = 10000\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	load := mod.Entries["load"]
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// Load the library once
	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	// Verify handle is a well-formed string ID
	if !strings.HasPrefix(handle, "lib-") {
		t.Fatalf("expected handle format 'lib-N', got %q", handle)
	}

	// Resolve 200 symbols and verify each ID format + correctness
	const N = 200
	seen := map[string]bool{}
	for i := 0; i < N; i++ {
		sV, err := sym.Builtin([]environment.Value{
			environment.NewString(handle),
			environment.NewString("sum3"),
			environment.NewString("int"),
		})
		if err != nil {
			t.Fatalf("sym %d failed: %v", i, err)
		}
		sid := sV.Str

		// Verify string format
		if !strings.HasPrefix(sid, "sym-") {
			t.Fatalf("sym %d: expected format 'sym-N', got %q", i, sid)
		}

		// Each resolution creates a NEW symbol ID
		if seen[sid] {
			t.Fatalf("sym %d: duplicate symbol ID %q", i, sid)
		}
		seen[sid] = true

		// Verify the symbol actually works
		a := float64(i)
		b := float64(i + 1)
		c := float64(i + 2)
		res, err := call.Builtin([]environment.Value{
			environment.NewString(sid),
			environment.NewNumber(a),
			environment.NewNumber(b),
			environment.NewNumber(c),
		})
		if err != nil {
			t.Fatalf("call %d (sym=%s) failed: %v", i, sid, err)
		}
		expected := a + b + c
		if res.Type != environment.NumberType || res.Num != expected {
			t.Fatalf("call %d: expected %v, got %v", i, expected, res)
		}
	}
}

// TestFfiMultipleLibsDistinctHandles loads the same .so multiple times
// and verifies that each gets a distinct lib-N handle string, and that
// symbols resolved from different handles work independently.
func TestFfiMultipleLibsDistinctHandles(t *testing.T) {
	builtins.StopAllHelpers()
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\ncall_timeout = 10000\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	load := mod.Entries["load"]
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// Load the library 10 times — each should get a unique handle
	const NLibs = 10
	handles := make([]string, NLibs)
	handleSet := map[string]bool{}

	for i := 0; i < NLibs; i++ {
		v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
		if err != nil {
			t.Fatalf("load %d failed: %v", i, err)
		}
		h := v.Str
		if !strings.HasPrefix(h, "lib-") {
			t.Fatalf("load %d: expected format 'lib-N', got %q", i, h)
		}
		if handleSet[h] {
			t.Fatalf("load %d: duplicate handle %q", i, h)
		}
		handleSet[h] = true
		handles[i] = h
	}

	// Resolve sum3 from each handle and verify calling works
	for i, h := range handles {
		sV, err := sym.Builtin([]environment.Value{
			environment.NewString(h),
			environment.NewString("sum3"),
			environment.NewString("int"),
		})
		if err != nil {
			t.Fatalf("sym on handle %s failed: %v", h, err)
		}
		sid := sV.Str
		if !strings.HasPrefix(sid, "sym-") {
			t.Fatalf("sym on handle %s: expected 'sym-N', got %q", h, sid)
		}

		a := float64(i * 10)
		b := float64(i*10 + 1)
		c := float64(i*10 + 2)
		res, err := call.Builtin([]environment.Value{
			environment.NewString(sid),
			environment.NewNumber(a),
			environment.NewNumber(b),
			environment.NewNumber(c),
		})
		if err != nil {
			t.Fatalf("call via handle %s sym %s failed: %v", h, sid, err)
		}
		expected := a + b + c
		if res.Type != environment.NumberType || res.Num != expected {
			t.Fatalf("call %d: expected %v, got %v", i, expected, res)
		}
	}
}
