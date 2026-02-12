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

// TestFfiNestedStructExpansion tests that nested structs are recursively expanded.
// Person { name: string, age: int, addr: Address }
// Address { city: string, zip: int }
// Flattened call: person_with_addr(name, age, city, zip)
func TestFfiNestedStructExpansion(t *testing.T) {
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
	libPath := filepath.Join(libDir, "libnested.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib_nested.c")
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
	sym := mod.Entries["sym"]
	call := mod.Entries["call"]
	defineStruct := mod.Entries["define_struct"]

	// define Address struct
	addrFields := []environment.Value{
		environment.NewObject(map[string]environment.Value{
			"name": environment.NewString("city"),
			"type": environment.NewString("string"),
		}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{
			"name": environment.NewString("zip"),
			"type": environment.NewString("int"),
		}, []string{"name", "type"}),
	}
	_, err := defineStruct.Builtin([]environment.Value{
		environment.NewString("Address"),
		environment.NewArray(addrFields),
	})
	if err != nil {
		t.Fatalf("define_struct Address failed: %v", err)
	}

	// define Person struct with nested Address
	personFields := []environment.Value{
		environment.NewObject(map[string]environment.Value{
			"name": environment.NewString("name"),
			"type": environment.NewString("string"),
		}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{
			"name": environment.NewString("age"),
			"type": environment.NewString("int"),
		}, []string{"name", "type"}),
		environment.NewObject(map[string]environment.Value{
			"name": environment.NewString("addr"),
			"type": environment.NewString("struct:Address"),
		}, []string{"name", "type"}),
	}
	_, err = defineStruct.Builtin([]environment.Value{
		environment.NewString("Person"),
		environment.NewArray(personFields),
	})
	if err != nil {
		t.Fatalf("define_struct Person failed: %v", err)
	}

	// load library
	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	// sym with arg_types ["struct:Person"]
	sV, err := sym.Builtin([]environment.Value{
		environment.NewString(handle),
		environment.NewString("person_with_addr"),
		environment.NewString("string"),
		environment.NewArray([]environment.Value{environment.NewString("struct:Person")}),
	})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	// create Address object
	addr := environment.NewObject(map[string]environment.Value{
		"city": environment.NewString("Curitiba"),
		"zip":  environment.NewNumber(80000),
	}, []string{"city", "zip"})

	// create Person object with nested Address
	person := environment.NewObject(map[string]environment.Value{
		"__struct__": environment.NewString("Person"),
		"name":       environment.NewString("Carlos"),
		"age":        environment.NewNumber(30),
		"addr":       addr,
	}, []string{"__struct__", "name", "age", "addr"})

	// call: should flatten to person_with_addr("Carlos", 30, "Curitiba", 80000)
	res, err := call.Builtin([]environment.Value{environment.NewString(symId), person})
	if err != nil {
		t.Fatalf("call failed: %v", err)
	}
	expected := "Carlos age=30 city=Curitiba zip=80000"
	if res.Type != environment.StringType || res.Str != expected {
		t.Fatalf("expected %q, got %v", expected, res)
	}
}
