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

func TestFfiTypedStructExpansion(t *testing.T) {
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
	def := mod.Entries["define_struct"]

	// define struct Person
	_, err := def.Builtin([]environment.Value{environment.NewString("Person"), environment.NewArray([]environment.Value{
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("name"), "type": environment.NewString("string")}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("age"), "type": environment.NewString("int")}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("score"), "type": environment.NewString("number")}, []string{"name", "type"}),
	})})
	if err != nil {
		t.Fatalf("define_struct failed: %v", err)
	}

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	// register symbol with argTypes indicating struct expansion
	argTypes := environment.NewArray([]environment.Value{environment.NewString("struct:Person")})
	symv, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("person_summary"), environment.NewString("string"), argTypes})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sid := symv.Str

	// create struct object (Person)
	obj := environment.NewObject(map[string]environment.Value{"__struct__": environment.NewString("Person"), "name": environment.NewString("alice"), "age": environment.NewNumber(30), "score": environment.NewNumber(12.5)}, []string{"__struct__", "name", "age", "score"})

	// first, try calling with explicit typed args to validate wrapper (register a fresh symbol without argTypes)
	symv2, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("person_summary"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	_ = symv2.Str
	// explicit call to person_summary_ints using a symbol registered for that name
	symv3, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("person_summary_ints"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sid3 := symv3.Str
	res2, err := call.Builtin([]environment.Value{environment.NewString(sid3), environment.NewString("alice"), environment.NewNumber(30), environment.NewNumber(12)})
	if err != nil {
		t.Fatalf("call failed (explicit args): %v", err)
	}
	if res2.Type != environment.StringType {
		t.Fatalf("expected string response, got %v", res2)
	}
	if !(contains(res2.Str, "Name:alice") && contains(res2.Str, "Age:30") && contains(res2.Str, "Score:12")) {
		t.Fatalf("unexpected response from explicit-arg call: %q", res2.Str)
	}

	// now call with struct expansion
	res, err := call.Builtin([]environment.Value{environment.NewString(sid), obj})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.StringType {
		t.Fatalf("expected string response, got %v", res)
	}
	if !(contains(res.Str, "Name:alice") && contains(res.Str, "Age:30") && contains(res.Str, "Score:12.50")) {
		t.Fatalf("unexpected response from typed struct call: %q", res.Str)
	}
}
