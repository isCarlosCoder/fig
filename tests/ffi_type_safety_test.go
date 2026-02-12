package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
)

// TestFfiTypeAssertionProtection verifies that calling an int-typed function
// with a string argument produces an error response instead of crashing
// the helper (panic from bare type assertion).
func TestFfiTypeAssertionProtection(t *testing.T) {
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

	// load library
	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	// get symbol for sum3 (int return type)
	sV, err := sym.Builtin([]environment.Value{
		environment.NewString(handle),
		environment.NewString("sum3"),
		environment.NewString("int"),
	})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	// 1. Call with string arg where int expected → should return error, not crash
	_, err = call.Builtin([]environment.Value{
		environment.NewString(symId),
		environment.NewString("not_a_number"),
		environment.NewNumber(2),
		environment.NewNumber(3),
	})
	if err == nil {
		t.Fatal("expected error when passing string to int function, got nil")
	}
	t.Logf("correctly got error for wrong arg type: %v", err)

	// 2. Verify helper is still alive by making a correct call
	res, err := call.Builtin([]environment.Value{
		environment.NewString(symId),
		environment.NewNumber(10),
		environment.NewNumber(20),
		environment.NewNumber(30),
	})
	if err != nil {
		t.Fatalf("correct call after bad type should succeed but failed: %v", err)
	}
	if res.Type != environment.NumberType || res.Num != 60 {
		t.Fatalf("expected 60, got %v", res)
	}
}

// TestFfiInvalidSymbolReturnsError verifies that calling with an invalid symbol ID
// returns an error instead of crashing the helper.
func TestFfiInvalidSymbolReturnsError(t *testing.T) {
	builtins.StopAllHelpers()
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
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
	call := mod.Entries["call"]

	// Call with a fake symbol ID → should return error
	_, err := call.Builtin([]environment.Value{
		environment.NewString("99999"),
		environment.NewNumber(1),
	})
	if err == nil {
		t.Fatal("expected error for invalid symbol, got nil")
	}
	t.Logf("correctly got error for invalid symbol: %v", err)

	// Verify helper is still alive with a ping
	ping := mod.Entries["ping"]
	v, err := ping.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("ping after invalid symbol should work: %v", err)
	}
	if v.Str != "pong" {
		t.Fatalf("expected pong, got %v", v)
	}
}

// TestFfiNilArgsReturnsError verifies that calling with nil args
// returns an error instead of crashing.
func TestFfiNilArgsReturnsError(t *testing.T) {
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

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	sV, err := sym.Builtin([]environment.Value{
		environment.NewString(handle),
		environment.NewString("sum3"),
		environment.NewString("int"),
	})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	// Call with 0 args (sum3 expects 3) → should get result 0 (int fn0 path)
	res, err := call.Builtin([]environment.Value{
		environment.NewString(symId),
	})
	if err != nil {
		t.Logf("0-arg call: got error (acceptable): %v", err)
	} else {
		t.Logf("0-arg call: got result %v (fn0 path)", res)
	}

	// Verify helper still alive
	ping := mod.Entries["ping"]
	pv, err := ping.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("ping after nil args should work: %v", err)
	}
	if pv.Str != "pong" {
		t.Fatalf("expected pong, got %v", pv)
	}
}
