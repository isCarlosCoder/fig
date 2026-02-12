package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
)

// buildHelperAndProject builds the ffi-helper binary, creates a temp project with fig.toml,
// and returns (helperBin, projectDir, libTempDir).
func buildHelperAndProject(t *testing.T) (string, string, string) {
	t.Helper()
	root := repoRootForTest(t)
	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
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
	libDir := t.TempDir()
	return bin, proj, libDir
}

// compileTestLib compiles a C source from tests/ffi_integration/<name>.c into libDir
// with the given extension and returns the full path to the built library.
func compileTestLib(t *testing.T, libDir, name, ext string) string {
	t.Helper()
	root := repoRootForTest(t)
	libPath := filepath.Join(libDir, "libtest"+ext)
	cpath := filepath.Join(root, "tests", "ffi_integration", name+".c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}
	return libPath
}

// setupTestProject chdir to project dir and return old cwd
func setupTestProject(t *testing.T, proj string) (string, error) {
	t.Helper()
	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("cannot get cwd: %v", err)
	}
	os.Chdir(proj)
	return old, nil
}

// restoreDir restores the working directory
func restoreDir(old string) {
	os.Chdir(old)
}

// TestFfiLibExt verifies that lib_ext() returns the correct extension for the current OS.
func TestFfiLibExt(t *testing.T) {
	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	libExtFn := mod.Entries["lib_ext"]
	if libExtFn.Type != environment.BuiltinFnType {
		t.Fatalf("expected builtin lib_ext, got %v", libExtFn.Type)
	}
	v, err := libExtFn.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("lib_ext() returned error: %v", err)
	}
	if v.Type != environment.StringType {
		t.Fatalf("expected string, got %v", v.Type)
	}
	var expected string
	switch runtime.GOOS {
	case "darwin":
		expected = ".dylib"
	case "windows":
		expected = ".dll"
	default:
		expected = ".so"
	}
	if v.Str != expected {
		t.Fatalf("lib_ext() = %q, want %q (GOOS=%s)", v.Str, expected, runtime.GOOS)
	}
}

// TestFfiLibName verifies that lib_name(base) returns the correct filename.
func TestFfiLibName(t *testing.T) {
	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	libNameFn := mod.Entries["lib_name"]
	if libNameFn.Type != environment.BuiltinFnType {
		t.Fatalf("expected builtin lib_name, got %v", libNameFn.Type)
	}
	v, err := libNameFn.Builtin([]environment.Value{environment.NewString("mymath")})
	if err != nil {
		t.Fatalf("lib_name() returned error: %v", err)
	}
	if v.Type != environment.StringType {
		t.Fatalf("expected string, got %v", v.Type)
	}
	var expected string
	switch runtime.GOOS {
	case "windows":
		expected = "mymath.dll"
	case "darwin":
		expected = "libmymath.dylib"
	default:
		expected = "libmymath.so"
	}
	if v.Str != expected {
		t.Fatalf("lib_name('mymath') = %q, want %q (GOOS=%s)", v.Str, expected, runtime.GOOS)
	}
}

// TestFfiLibNameEmpty verifies lib_name errors on no args.
func TestFfiLibNameEmpty(t *testing.T) {
	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	libNameFn := mod.Entries["lib_name"]
	_, err := libNameFn.Builtin([]environment.Value{})
	if err == nil {
		t.Fatal("expected error for lib_name() with no args")
	}
}

// TestFfiLibExtGoFunc verifies the exported Go functions directly.
func TestFfiLibExtGoFunc(t *testing.T) {
	ext := builtins.LibExt()
	switch runtime.GOOS {
	case "darwin":
		if ext != ".dylib" {
			t.Fatalf("LibExt() = %q, want .dylib", ext)
		}
	case "windows":
		if ext != ".dll" {
			t.Fatalf("LibExt() = %q, want .dll", ext)
		}
	default:
		if ext != ".so" {
			t.Fatalf("LibExt() = %q, want .so", ext)
		}
	}
}

// TestFfiLibNameGoFunc verifies the exported Go functions directly.
func TestFfiLibNameGoFunc(t *testing.T) {
	name := builtins.LibName("crypto")
	switch runtime.GOOS {
	case "windows":
		if name != "crypto.dll" {
			t.Fatalf("LibName('crypto') = %q, want crypto.dll", name)
		}
	case "darwin":
		if name != "libcrypto.dylib" {
			t.Fatalf("LibName('crypto') = %q, want libcrypto.dylib", name)
		}
	default:
		if name != "libcrypto.so" {
			t.Fatalf("LibName('crypto') = %q, want libcrypto.so", name)
		}
	}
}

// TestFfiLoadWithLibExt verifies that the load function can use lib_ext-based paths.
// This is a cross-platform integration test that compiles a C lib with the correct extension.
func TestFfiLoadWithLibExt(t *testing.T) {
	builtins.StopAllHelpers()

	bin, proj, libDir := buildHelperAndProject(t)
	_ = bin

	// compile the test lib with the correct platform extension
	ext := builtins.LibExt()
	libPath := compileTestLib(t, libDir, "lib", ext)

	old, _ := setupTestProject(t, proj)
	defer restoreDir(old)

	mod := builtins.Get("ffi")
	loadFn := mod.Entries["load"]
	v, err := loadFn.Builtin([]environment.Value{environment.NewString(libPath)})
	if err != nil {
		t.Fatalf("load(%s) failed: %v", libPath, err)
	}
	if v.Type != environment.StringType {
		t.Fatalf("expected string handle, got %v", v.Type)
	}
}
