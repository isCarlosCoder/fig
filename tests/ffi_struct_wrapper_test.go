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

// basic wrapper functionality (construction/validation/flattening)
func TestFfiStructWrapperBasics(t *testing.T) {
	mod := builtins.Get("ffi")
	structFn := mod.Entries["struct"]

	// create descriptor Point{x:int,y:int}
	fields := environment.NewArray([]environment.Value{
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("x"), "type": environment.NewString("int")}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("y"), "type": environment.NewString("int")}, []string{"name", "type"}),
	})
	descV, err := structFn.Builtin([]environment.Value{environment.NewString("Point"), fields})
	if err != nil {
		t.Fatalf("struct creation failed: %v", err)
	}
	if descV.Type != environment.ObjectType || descV.Obj == nil {
		t.Fatalf("descriptor not object: %v", descV)
	}
	desc := descV.Obj
	if nm, ok := desc.Entries["name"]; !ok || nm.Str != "Point" {
		t.Fatal("descriptor missing name")
	}
	if _, ok := desc.Entries["new"]; !ok {
		t.Fatal("descriptor missing new method")
	}
	if _, ok := desc.Entries["validate"]; !ok {
		t.Fatal("descriptor missing validate method")
	}
	if _, ok := desc.Entries["flatten"]; !ok {
		t.Fatal("descriptor missing flatten method")
	}

	newFn := desc.Entries["new"]
	validateFn := desc.Entries["validate"]
	flattenFn := desc.Entries["flatten"]

	// positional construction should work
	instV, err := newFn.Builtin([]environment.Value{environment.NewNumber(3), environment.NewNumber(4)})
	if err != nil {
		t.Fatalf("Point.new positional failed: %v", err)
	}
	if instV.Type != environment.ObjectType {
		t.Fatalf("instance not object: %v", instV)
	}
	inst := instV.Obj
	if s, ok := inst.Entries["__struct__"]; !ok || s.Str != "Point" {
		t.Fatalf("instance missing __struct__: %v", instV)
	}
	if x, ok := inst.Entries["x"]; !ok || x.Num != 3 {
		t.Fatalf("instance.x wrong: %v", x)
	}
	if y, ok := inst.Entries["y"]; !ok || y.Num != 4 {
		t.Fatalf("instance.y wrong: %v", y)
	}

	// named construction
	argMap := environment.NewObject(map[string]environment.Value{"x": environment.NewNumber(7), "y": environment.NewNumber(8)}, []string{"x", "y"})
	instV2, err := newFn.Builtin([]environment.Value{argMap})
	if err != nil {
		t.Fatalf("Point.new map failed: %v", err)
	}
	if instV2.Obj.Entries["x"].Num != 7 {
		t.Fatal("map instance.x wrong")
	}

	// validation
	_, err = validateFn.Builtin([]environment.Value{instV2})
	if err != nil {
		t.Fatalf("validate should succeed: %v", err)
	}
	// invalid value
	bad := environment.NewObject(map[string]environment.Value{"x": environment.NewString("foo"), "y": environment.NewNumber(1)}, []string{"x", "y"})
	_, err = validateFn.Builtin([]environment.Value{bad})
	if err == nil {
		t.Fatal("validate did not catch bad type")
	}

	// flatten returns values/types
	flatV, err := flattenFn.Builtin([]environment.Value{instV})
	if err != nil {
		t.Fatalf("flatten failed: %v", err)
	}
	if flatV.Type != environment.ObjectType {
		t.Fatalf("flatten result not object: %v", flatV)
	}
	flatObj := flatV.Obj
	vals := flatObj.Entries["values"].Arr
	typesArr := flatObj.Entries["types"].Arr
	if len(*vals) != 2 || len(*typesArr) != 2 {
		t.Fatalf("flatten arrays wrong length")
	}
}

// integration test: use wrapper with ffi.sym and ffi.call
func TestFfiStructWrapperIntegration(t *testing.T) {
	builtins.StopAllHelpers()
	// build helper and library as in typed struct test
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	root := repoRootForTest(t)
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtyped.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib_typed.c")
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
	structFn := mod.Entries["struct"]

	// define wrapper Person via high-level API
	fields := environment.NewArray([]environment.Value{
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("name"), "type": environment.NewString("string")}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("age"), "type": environment.NewString("int")}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("score"), "type": environment.NewString("number")}, []string{"name", "type"}),
	})
	descV, err := structFn.Builtin([]environment.Value{environment.NewString("Person"), fields})
	if err != nil {
		t.Fatalf("struct wrapper failed: %v", err)
	}

	// create instance using .new
	descObj := descV.Obj
	newFn := descObj.Entries["new"]
	instV, err := newFn.Builtin([]environment.Value{environment.NewString("alice"), environment.NewNumber(30), environment.NewNumber(12.5)})
	if err != nil {
		t.Fatalf("constructor failed: %v", err)
	}

	// load lib and register symbol using wrapper descriptor as argType
	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str
	argArr := environment.NewArray([]environment.Value{descV})
	symv, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("person_summary"), environment.NewString("string"), argArr})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sid := symv.Str

	res, err := call.Builtin([]environment.Value{environment.NewString(sid), instV})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.StringType {
		t.Fatalf("expected string response, got %v", res)
	}
	if !(contains(res.Str, "Name:alice") && contains(res.Str, "Age:30") && contains(res.Str, "Score:12.50")) {
		t.Fatalf("unexpected response: %q", res.Str)
	}
}
