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

func TestFfiStructPointerReturn(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and library
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libstruct.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib_struct.c")
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
	def := mod.Entries["define_struct"]

	// define struct Foo
	_, err := def.Builtin([]environment.Value{environment.NewString("Foo"), environment.NewArray([]environment.Value{
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("foo"), "type": environment.NewString("int")}, []string{"name", "type"}),
	})})
	if err != nil {
		t.Fatalf("define_struct failed: %v", err)
	}

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	// register make_foo (returns struct pointer)
	symv, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("make_foo"), environment.NewString("struct:Foo"), environment.NewArray([]environment.Value{environment.NewString("int")})})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sidMake := symv.Str

	// register get_foo (takes struct pointer)
	symv2, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("get_foo"), environment.NewString("int"), environment.NewArray([]environment.Value{environment.NewString("struct:Foo")})})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sidGet := symv2.Str

	// call make_foo
	res, err := call.Builtin([]environment.Value{environment.NewString(sidMake), environment.NewNumber(42)})
	if err != nil {
		t.Fatalf("call make_foo failed: %v", err)
	}
	if res.Type != environment.ObjectType || res.Obj == nil {
		t.Fatalf("expected object pointer marker, got %v", res)
	}
	pidV, ok := res.Obj.Entries["__ptrid__"]
	if !ok || pidV.Type != environment.StringType {
		t.Fatalf("missing ptrid in result: %v", res.Obj)
	}
	if s, ok := res.Obj.Entries["__struct__"]; !ok || s.Str != "Foo" {
		t.Fatalf("invalid struct name in marker: %v", res.Obj)
	}

	// now call get_foo using the marker
	res2, err := call.Builtin([]environment.Value{environment.NewString(sidGet), res})
	if err != nil {
		t.Fatalf("call get_foo failed: %v", err)
	}
	if res2.Type != environment.NumberType {
		t.Fatalf("expected numeric result, got %v", res2)
	}
	if int(res2.Num) != 42 {
		t.Fatalf("unexpected value from get_foo: %v", res2.Num)
	}
}
