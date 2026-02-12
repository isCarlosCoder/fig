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

func TestFfiStructMarshalling(t *testing.T) {
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

	// define struct Person
	_, err := def.Builtin([]environment.Value{environment.NewString("Person"), environment.NewArray([]environment.Value{
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("name"), "type": environment.NewString("string")}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{"name": environment.NewString("age"), "type": environment.NewString("int")}, []string{"name", "type"}),
	})})
	if err != nil {
		t.Fatalf("define_struct failed: %v", err)
	}

	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	symv, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("struct_prefix"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	sid := symv.Str

	// create struct object (Person)
	obj := environment.NewObject(map[string]environment.Value{"__struct__": environment.NewString("Person"), "name": environment.NewString("alice"), "age": environment.NewNumber(30)}, []string{"__struct__", "name", "age"})

	res, err := call.Builtin([]environment.Value{environment.NewString(sid), obj, environment.NewString("hi ")})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	if res.Type != environment.StringType {
		t.Fatalf("expected string response, got %v", res)
	}
	if len(res.Str) < 3 || res.Str[:3] != "hi " {
		t.Fatalf("expected prefix 'hi ', got %q", res.Str)
	}
	if !(contains(res.Str, "\"name\":\"alice\"") && contains(res.Str, "\"age\":30")) {
		t.Fatalf("expected struct json inside result, got %q", res.Str)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (len(sub) == 0 || (len(s) > 0 && (indexOf(s, sub) >= 0)))
}

func indexOf(s, sub string) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
