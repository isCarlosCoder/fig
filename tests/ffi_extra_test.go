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

// setupExtraLib builds the helper + lib_extra.so and sets up a project dir.
// Returns the ffi module, the lib handle string, and a cleanup function.
func setupExtraLib(t *testing.T) (*builtins.Module, string) {
	t.Helper()
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
	libPath := filepath.Join(libDir, "libextra.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib_extra.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	old, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(old) })
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}

	// load lib
	load := mod.Entries["load"]
	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	return mod, handle
}

func findRepoRoot(t *testing.T) string {
	t.Helper()
	p, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot get cwd: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(p, "go.mod")); err == nil {
			return p
		}
		parent := filepath.Dir(p)
		if parent == p {
			t.Fatalf("repo root (go.mod) not found")
		}
		p = parent
	}
}

// --- 0-arg tests ---

func TestFfiIntZeroArgs(t *testing.T) {
	mod, handle := setupExtraLib(t)
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// get_counter: int ()
	sV, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("get_counter"), environment.NewString("int")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(symId)})
	if err != nil {
		t.Fatalf("call get_counter() failed: %v", err)
	}
	if res.Type != environment.NumberType || res.Num != 0 {
		t.Fatalf("expected 0 (initial counter), got %v", res)
	}
}

func TestFfiIntOneArg(t *testing.T) {
	mod, handle := setupExtraLib(t)
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// double_int: int (int)
	sV, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("double_int"), environment.NewString("int")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(symId), environment.NewNumber(7)})
	if err != nil {
		t.Fatalf("call double_int(7) failed: %v", err)
	}
	if res.Type != environment.NumberType || res.Num != 14 {
		t.Fatalf("expected 14, got %v", res)
	}
}

func TestFfiDoubleZeroArgs(t *testing.T) {
	mod, handle := setupExtraLib(t)
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// get_pi: double ()
	sV, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("get_pi"), environment.NewString("double")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(symId)})
	if err != nil {
		t.Fatalf("call get_pi() failed: %v", err)
	}
	if res.Type != environment.NumberType {
		t.Fatalf("expected number, got %v", res.Type)
	}
	if res.Num < 3.14 || res.Num > 3.15 {
		t.Fatalf("expected ~3.14159, got %f", res.Num)
	}
}

func TestFfiDoubleOneArg(t *testing.T) {
	mod, handle := setupExtraLib(t)
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// square: double (double)
	sV, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("square"), environment.NewString("double")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(symId), environment.NewNumber(5)})
	if err != nil {
		t.Fatalf("call square(5) failed: %v", err)
	}
	if res.Type != environment.NumberType || res.Num != 25 {
		t.Fatalf("expected 25, got %v", res)
	}
}

func TestFfiStringZeroArgs(t *testing.T) {
	mod, handle := setupExtraLib(t)
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// hello: string ()
	sV, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("hello"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(symId)})
	if err != nil {
		t.Fatalf("call hello() failed: %v", err)
	}
	if res.Type != environment.StringType || res.Str != "hello from C" {
		t.Fatalf("expected 'hello from C', got %v", res)
	}
}

func TestFfiStringOneArg(t *testing.T) {
	mod, handle := setupExtraLib(t)
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// shout: string (string)
	sV, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("shout"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(symId), environment.NewString("hey")})
	if err != nil {
		t.Fatalf("call shout('hey') failed: %v", err)
	}
	if res.Type != environment.StringType || res.Str != "hey!" {
		t.Fatalf("expected 'hey!', got %v", res)
	}
}

func TestFfiVoidZeroArgs(t *testing.T) {
	mod, handle := setupExtraLib(t)
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// inc_counter: void ()
	sV, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("inc_counter"), environment.NewString("void")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(symId)})
	if err != nil {
		t.Fatalf("call inc_counter() failed: %v", err)
	}
	// void returns nil
	if res.Type != environment.NilType {
		t.Fatalf("expected nil (void return), got %v", res)
	}
}

func TestFfiVoidThenRead(t *testing.T) {
	mod, handle := setupExtraLib(t)
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	// set_counter: void (int)
	setSym, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("set_counter"), environment.NewString("void")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	setId := setSym.Str

	// get_counter: int ()
	getSym, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("get_counter"), environment.NewString("int")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	getId := getSym.Str

	// set counter to 42
	_, err = call.Builtin([]environment.Value{environment.NewString(setId), environment.NewNumber(42)})
	if err != nil {
		t.Fatalf("call set_counter(42) failed: %v", err)
	}

	// read counter
	res, err := call.Builtin([]environment.Value{environment.NewString(getId)})
	if err != nil {
		t.Fatalf("call get_counter() failed: %v", err)
	}
	if res.Type != environment.NumberType || res.Num != 42 {
		t.Fatalf("expected 42, got %v", res)
	}
}
