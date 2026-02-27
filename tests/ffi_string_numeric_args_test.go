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

// verify that string-return functions can be called with numeric arguments and
// the values are converted to strings automatically (builtins + helper fix).
func TestFfiStringReturnNumericArgs(t *testing.T) {
	builtins.StopAllHelpers()
	root := repoRootForTest(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	// write small C file
	csrc := `#include <stdlib.h>
#include <stdio.h>
#include <string.h>

char* num_to_str(double x) {
    char buf[64];
    snprintf(buf, sizeof(buf), "num:%g", x);
    return strdup(buf);
}
`
	libDir := t.TempDir()
	cpath := filepath.Join(libDir, "libnum.c")
	if err := os.WriteFile(cpath, []byte(csrc), 0644); err != nil {
		t.Fatalf("cannot write C source: %v", err)
	}
	libPath := filepath.Join(libDir, "libnum.so")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build num library: %v (%s)", err, string(out))
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

	// load library
	v, err := load.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	handle := v.Str

	// register symbol with return string and one double arg
	sV, err := sym.Builtin([]environment.Value{
		environment.NewString(handle),
		environment.NewString("num_to_str"),
		environment.NewString("string"),
		environment.NewArray([]environment.Value{environment.NewString("double")}),
	})
	if err != nil {
		t.Fatalf("sym failed: %v", err)
	}
	symId := sV.Str

	// call with number
	res, err := call.Builtin([]environment.Value{
		environment.NewString(symId),
		environment.NewNumber(4.2),
	})
	if err != nil {
		t.Fatalf("call returned error: %v", err)
	}
	if res.Type != environment.StringType {
		t.Fatalf("expected string result, got %v", res)
	}
	if res.Str != "num:4.2" {
		t.Fatalf("unexpected string from num_to_str: %q", res.Str)
	}
}
