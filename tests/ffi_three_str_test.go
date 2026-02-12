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

func TestFfiJoin3Works(t *testing.T) {
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
	libPath := filepath.Join(libDir, "libjoin.so")
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

	symvJoin, err := sym.Builtin([]environment.Value{environment.NewString(handle), environment.NewString("join3"), environment.NewString("string")})
	if err != nil {
		t.Fatalf("sym join3 failed: %v", err)
	}
	sidJoin := symvJoin.Str
	resJoin, err := call.Builtin([]environment.Value{environment.NewString(sidJoin), environment.NewString("a"), environment.NewString("b"), environment.NewString("c")})
	if err != nil {
		t.Fatalf("join3 call failed: %v", err)
	}
	if resJoin.Type != environment.StringType || resJoin.Str != "abc" {
		t.Fatalf("join3 unexpected: %v", resJoin)
	}
}
