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

// TestFfiCoercionIntTruncatesFloat verifies that when argType is "int" and
// the user passes 42.7, the value is truncated to 42 before sending to helper.
func TestFfiCoercionIntTruncatesFloat(t *testing.T) {
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

	// Load lib and resolve sum3 with argTypes ["int", "int", "int"]
	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	argTypesArr := environment.NewArray([]environment.Value{
		environment.NewString("int"),
		environment.NewString("int"),
		environment.NewString("int"),
	})
	sV, err := sym.Builtin([]environment.Value{
		environment.NewString(handle),
		environment.NewString("sum3"),
		environment.NewString("int"),
		argTypesArr,
	})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	// Call with float values that should be truncated: 10.7 + 20.9 + 30.3 → 10 + 20 + 30 = 60
	res, err := call.Builtin([]environment.Value{
		environment.NewString(symId),
		environment.NewNumber(10.7),
		environment.NewNumber(20.9),
		environment.NewNumber(30.3),
	})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	// After truncation: sum3(10, 20, 30) = 60
	if res.Type != environment.NumberType || res.Num != 60 {
		t.Fatalf("expected 60 (truncated), got %v", res)
	}
}

// TestFfiCoercionIntFromStringError verifies that when argType is "int" and
// user passes a non-numeric string, a clear error is returned.
func TestFfiCoercionIntFromStringError(t *testing.T) {
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

	argTypesArr := environment.NewArray([]environment.Value{
		environment.NewString("int"),
		environment.NewString("int"),
		environment.NewString("int"),
	})
	sV, err := sym.Builtin([]environment.Value{
		environment.NewString(handle),
		environment.NewString("sum3"),
		environment.NewString("int"),
		argTypesArr,
	})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	// Call with a string "abc" where int is expected — should get a clear error
	_, err = call.Builtin([]environment.Value{
		environment.NewString(symId),
		environment.NewString("abc"),
		environment.NewNumber(2),
		environment.NewNumber(3),
	})
	if err == nil {
		t.Fatal("expected error when passing string where int expected")
	}
	if !strings.Contains(err.Error(), "expects int") {
		t.Fatalf("expected 'expects int' error, got: %v", err)
	}
}

// TestFfiCoercionNumberToString verifies that when argType is "string"
// and user passes a number, it's automatically converted to string.
func TestFfiCoercionNumberToString(t *testing.T) {
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

	// Load lib and resolve dupstr (returns copy of first arg) with argTypes ["string"]
	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	argTypesArr := environment.NewArray([]environment.Value{
		environment.NewString("string"),
	})
	sV, err := sym.Builtin([]environment.Value{
		environment.NewString(handle),
		environment.NewString("dupstr"),
		environment.NewString("string"),
		argTypesArr,
	})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	// Call with a number 42 where string is expected — should be coerced to "42"
	res, err := call.Builtin([]environment.Value{
		environment.NewString(symId),
		environment.NewNumber(42),
	})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.StringType || res.Str != "42" {
		t.Fatalf("expected string \"42\", got %v", res)
	}
}

// TestFfiCallWithoutArgTypesNoRegression verifies that call() without argTypes
// works the same as before (no regression from adding coercion logic).
func TestFfiCallWithoutArgTypesNoRegression(t *testing.T) {
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

	// Resolve without argTypes — use 2 args (only name and rtype)
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

	// Call without argTypes — values pass through directly
	res, err := call.Builtin([]environment.Value{
		environment.NewString(symId),
		environment.NewNumber(10),
		environment.NewNumber(20),
		environment.NewNumber(30),
	})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.NumberType || res.Num != 60 {
		t.Fatalf("expected 60, got %v", res)
	}
}
