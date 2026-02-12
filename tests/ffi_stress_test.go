package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
)

// TestFfiStressConcurrentCalls issues 200 concurrent call()s to sum3
// and verifies all return the correct result.
func TestFfiStressConcurrentCalls(t *testing.T) {
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

	const N = 200
	var wg sync.WaitGroup
	errs := make([]error, N)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			a := float64(i)
			b := float64(i + 1)
			c := float64(i + 2)
			res, err := call.Builtin([]environment.Value{
				environment.NewString(symId),
				environment.NewNumber(a),
				environment.NewNumber(b),
				environment.NewNumber(c),
			})
			if err != nil {
				errs[i] = fmt.Errorf("call %d failed: %v", i, err)
				return
			}
			expected := 3*float64(i) + 3
			if res.Type != environment.NumberType || res.Num != expected {
				errs[i] = fmt.Errorf("call %d: expected %v, got %v", i, expected, res)
			}
		}(i)
	}
	wg.Wait()

	for _, e := range errs {
		if e != nil {
			t.Fatal(e)
		}
	}
}

// TestFfiStressCallbackStorm fires 50 sequential calls that each invoke a callback.
// Note: callback-invoking C calls are serialized in the helper (CGo blocking),
// so this test verifies robustness over many callback round-trips rather than concurrency.
func TestFfiStressCallbackStorm(t *testing.T) {
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
	libPath := filepath.Join(libDir, "libcb.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib_cb.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\ncall_timeout = 15000\n", bin)
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
	regCb := mod.Entries["register_callback"]

	cbV, err := regCb.Builtin([]environment.Value{
		environment.NewBuiltinFn("stress_cb", func(args []environment.Value) (environment.Value, error) {
			if len(args) > 0 && args[0].Type == environment.StringType {
				return environment.NewString("!" + args[0].Str), nil
			}
			return environment.NewString("!"), nil
		}),
	})
	if err != nil {
		t.Fatalf("register_callback failed: %v", err)
	}
	cbId := ""
	if cbV.Type == environment.ObjectType && cbV.Obj != nil {
		if v, ok := cbV.Obj.Entries["__cb__"]; ok {
			cbId = v.Str
		}
	}
	if cbId == "" {
		t.Fatal("callback id not found")
	}

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	sV, err := sym.Builtin([]environment.Value{
		environment.NewString(handle),
		environment.NewString("call_then_prefix"),
		environment.NewString("string"),
	})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	const N = 50
	errs := make([]error, N)
	for i := 0; i < N; i++ {
		arg := fmt.Sprintf("msg%d", i)
		res, err := call.Builtin([]environment.Value{
			environment.NewString(symId),
			environment.NewString(cbId),
			environment.NewString(arg),
		})
		if err != nil {
			errs[i] = fmt.Errorf("call %d failed: %v", i, err)
			continue
		}
		// call_then_prefix(cbid, prefix) calls callback with "world", then prepends prefix
		// callback returns "!" + "world" = "!world"
		// so result = arg + "!world"
		expected := arg + "!world"
		if res.Type != environment.StringType || res.Str != expected {
			errs[i] = fmt.Errorf("call %d: expected %q, got %v", i, expected, res)
		}
	}

	for _, e := range errs {
		if e != nil {
			t.Fatal(e)
		}
	}
}
