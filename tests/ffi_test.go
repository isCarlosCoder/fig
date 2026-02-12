package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
)

func TestFfiEnabledDefaultFalse(t *testing.T) {
	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	en := mod.Entries["enabled"]
	v, err := en.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("enabled() returned error: %v", err)
	}
	if v.Type != environment.BooleanType || v.Bool != false {
		t.Fatalf("expected enabled=false by default, got %v", v)
	}
}

func TestFfiHelperPing(t *testing.T) {
	builtins.StopAllHelpers()
	// helper to find repo root (where go.mod exists)
	repoRoot := func(t *testing.T) string {
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
	// build helper binary into temp path
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	// build helper from repo root
	root := repoRoot(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	// create temp project dir with fig.toml enabling ffi and helper path
	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}

	// run test from proj dir
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	ping := mod.Entries["ping"]
	if ping.Type != environment.BuiltinFnType {
		t.Fatalf("expected builtin ping, got %v", ping.Type)
	}
	v, err := ping.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("ping failed: %v", err)
	}
	if v.Type != environment.StringType || v.Str != "pong" {
		t.Fatalf("expected pong, got %v", v)
	}
}

func TestFfiConcurrentPing(t *testing.T) {
	builtins.StopAllHelpers()
	// same setup as TestFfiHelperPing but issue many concurrent pings
	repoRoot := func(t *testing.T) string {
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
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRoot(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n", bin)
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
	ping := mod.Entries["ping"]
	if ping.Type != environment.BuiltinFnType {
		t.Fatalf("expected builtin ping, got %v", ping.Type)
	}

	const concurrency = 50
	errs := make(chan error, concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			v, err := ping.Builtin([]environment.Value{})
			if err != nil {
				errs <- err
				return
			}
			if v.Type != environment.StringType || v.Str != "pong" {
				errs <- fmt.Errorf("expected pong, got %v", v)
				return
			}
			errs <- nil
		}()
	}
	for i := 0; i < concurrency; i++ {
		if err := <-errs; err != nil {
			// log helper stderr for debugging
			if s := builtins.HelperStderrFor(proj); s != "" {
				t.Logf("helper stderr:\n%s", s)
			}
			t.Fatalf("concurrent ping failed: %v", err)
		}
	}
}

func TestFfiLoadAndCall(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper binary into temp path
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	// build helper from repo root
	root := func() string {
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
	}()
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	// build test C library (libtest.so)
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
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
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	load := mod.Entries["load"]
	call := mod.Entries["call"]

	if load.Type != environment.BuiltinFnType {
		t.Fatalf("expected builtin load, got %v", load.Type)
	}
	if call.Type != environment.BuiltinFnType {
		t.Fatalf("expected builtin call, got %v", call.Type)
	}

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if v.Type != environment.StringType {
		t.Fatalf("expected handle string, got %v", v)
	}
	handle := v.Str

	// register symbol "sum" with rtype int
	sym, err := mod.Entries["sym"].Builtin([]environment.Value{environment.NewString(handle), environment.NewString("sum3"), environment.NewString("int")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	if sym.Type != environment.StringType {
		t.Fatalf("expected sym to return symbol id string, got %v", sym)
	}
	symId := sym.Str

	// call sum 1+2+3 = 6
	res, err := call.Builtin([]environment.Value{environment.NewString(symId), environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.NumberType || res.Num != 6 {
		t.Fatalf("expected 6, got %v", res)
	}
}

// helper to find repo root (go.mod) reused by tests
func repoRootForTest(t *testing.T) string {
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

func TestFfiCallDouble(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and lib
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	// build lib
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
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
	defer os.Chdir(old)
	os.Chdir(proj)

	// Also add string test for dupstr
	mod := builtins.Get("ffi")
	load := mod.Entries["load"]
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	// Test dupstr (returns a duplicated string)
	sdup, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("dupstr"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym dupstr failed: %v", err)
	}
	siddup := sdup.Str

	resdup, err := call.Builtin([]environment.Value{environment.NewString(siddup), environment.NewString("hello")})
	if err != nil {
		t.Fatalf("dupstr call failed: %v", err)
	}
	if resdup.Type != environment.StringType || resdup.Str != "hello" {
		t.Fatalf("expected dupstr to return \"hello\", got %v", resdup)
	}

	s, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("mul2"), environment.NewString("double")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sid := s.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(sid), environment.NewNumber(2.5), environment.NewNumber(4.0)})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.NumberType || res.Num != 10.0 {
		t.Fatalf("expected 10.0, got %v", res)
	}
}

func TestFfiConcat(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and lib
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	// build lib
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
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
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	load := mod.Entries["load"]
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	sc, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("concat"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym concat failed: %v", err)
	}
	sid := sc.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(sid), environment.NewString("hello"), environment.NewString(" world")})
	if err != nil {
		t.Fatalf("concat call failed: %v", err)
	}
	if res.Type != environment.StringType || res.Str != "hello world" {
		t.Fatalf("expected concat to return \"hello world\", got %v", res)
	}
}

func TestFfiArrayRoundtrip(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and project
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n", bin)
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
	callRaw := mod.Entries["call_raw"]
	if callRaw.Type != environment.BuiltinFnType {
		t.Fatalf("expected builtin call_raw, got %v", callRaw.Type)
	}
	argArray := environment.NewArray([]environment.Value{environment.NewNumber(1), environment.NewNumber(2), environment.NewNumber(3)})
	v, err := callRaw.Builtin([]environment.Value{argArray})
	if err != nil {
		t.Fatalf("call_raw failed: %v", err)
	}
	if v.Type != environment.ArrayType {
		t.Fatalf("expected array type from call_raw, got %v", v)
	}
	arr := *v.Arr
	if len(arr) != 1 {
		t.Fatalf("expected outer array length 1 (args list), got %d", len(arr))
	}
	if arr[0].Type != environment.ArrayType {
		t.Fatalf("expected nested array as first element, got %v", arr[0])
	}
	nested := *arr[0].Arr
	if len(nested) != 3 {
		t.Fatalf("expected nested array length 3, got %d", len(nested))
	}
	if nested[0].Num != 1 || nested[1].Num != 2 || nested[2].Num != 3 {
		t.Fatalf("unexpected nested array contents: %v", nested)
	}
}

func TestFfiCallback(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and lib
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	// build lib
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libcb.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib_cb.c")
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
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	load := mod.Entries["load"]
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]
	reg := mod.Entries["register_callback"]
	unreg := mod.Entries["unregister_callback"]

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	cbIdObj, err := reg.Builtin([]environment.Value{environment.NewBuiltinFn("cb", func(a []environment.Value) (environment.Value, error) {
		if len(a) != 1 {
			return environment.NewNil(), fmt.Errorf("cb: expects 1 arg")
		}
		if s, err := a[0].AsString(); err == nil {
			return environment.NewString(s + "!"), nil
		}
		return environment.NewNil(), fmt.Errorf("cb: invalid arg")
	})})
	if err != nil {
		t.Fatalf("register_callback failed: %v", err)
	}
	if cbIdObj.Type != environment.ObjectType {
		t.Fatalf("expected callback object, got %v", cbIdObj)
	}

	symv, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("call_then_prefix"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sid := symv.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(sid), cbIdObj, environment.NewString("hello ")})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.StringType || res.Str != "hello world!" {
		t.Fatalf("expected 'hello world!', got %v", res)
	}

	// cleanup
	_, _ = unreg.Builtin([]environment.Value{cbIdObj})
}

func TestFfiMemoryOwnership(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and lib
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	// build lib
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libmem.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib_mem.c")
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
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	load := mod.Entries["load"]
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]
	alloc := mod.Entries["alloc"]
	free := mod.Entries["free"]
	memWrite := mod.Entries["mem_write"]

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	// resolve symbol
	symv, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("echo_mem"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sid := symv.Str

	// allocate memory and write 'hello' into it
	mobj, err := alloc.Builtin([]environment.Value{environment.NewNumber(16)})
	if err != nil {
		t.Fatalf("alloc failed: %v", err)
	}
	if mobj.Type != environment.ObjectType {
		t.Fatalf("expected mem object, got %v", mobj)
	}

	// write
	bobj, err := builtins.Get("ffi").Entries["bytes_from_string"].Builtin([]environment.Value{environment.NewString("hello")})
	if err != nil {
		t.Fatalf("bytes_from_string failed: %v", err)
	}
	t.Logf("alloc returned mem object: %+v", mobj)
	_, err = memWrite.Builtin([]environment.Value{mobj, bobj, environment.NewNumber(0)})
	if err != nil {
		t.Fatalf("mem_write failed: %v", err)
	}

	// call echo_mem with mem object
	res, err := call.Builtin([]environment.Value{environment.NewString(sid), mobj})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.StringType || res.Str != "hello" {
		t.Fatalf("expected 'hello', got %v", res)
	}

	// cleanup
	_, _ = free.Builtin([]environment.Value{mobj})
}

func TestFfiBadArgType(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and lib
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	// build lib
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
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
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	load := mod.Entries["load"]
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	s, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("sum3"), environment.NewString("int")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sid := s.Str

	_, err = call.Builtin([]environment.Value{environment.NewString(sid), environment.NewNumber(1), environment.NewString("x"), environment.NewNumber(3)})
	if err == nil {
		t.Fatalf("expected error for bad arg type, got nil")
	}
}

func TestFfiMissingSymbol(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and lib
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	// build lib
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
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
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	load := mod.Entries["load"]
	sym := mod.Entries["sym"]

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	_, err = sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("nope"), environment.NewString("int")})
	if err == nil {
		t.Fatalf("expected error for missing symbol, got nil")
	}
}

func TestFfiCallTimeout(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and project
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	helpCmd := mod.Entries["helper_cmd"]
	if helpCmd.Type != environment.BuiltinFnType {
		t.Fatalf("expected helper_cmd builtin")
	}
	_, err := helpCmd.Builtin([]environment.Value{environment.NewString("sleep"), environment.NewNumber(5000)})
	if err == nil {
		t.Fatalf("expected timeout error from helper sleep, got nil")
	}
}

func TestFfiHelperCrashRestart(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and project
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("cannot write fig.toml: %v", err)
	}
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(proj)

	mod := builtins.Get("ffi")
	helpCmd := mod.Entries["helper_cmd"]
	ping := mod.Entries["ping"]

	// ensure helper started
	if _, err := ping.Builtin([]environment.Value{}); err != nil {
		t.Fatalf("initial ping failed: %v", err)
	}
	// crash helper
	_, _ = helpCmd.Builtin([]environment.Value{environment.NewString("crash")})
	// allow short time for restart
	time.Sleep(100 * time.Millisecond)
	v, err := ping.Builtin([]environment.Value{})
	if err != nil || v.Type != environment.StringType || v.Str != "pong" {
		t.Fatalf("expected ping after crash/restart, got err=%v v=%v", err, v)
	}
}
