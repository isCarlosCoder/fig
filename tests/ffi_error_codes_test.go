package tests

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
)

// setupHelperForErrors builds the helper binary and test lib, creates a project dir.
func setupHelperForErrors(t *testing.T) (proj, bin, libPath string) {
	t.Helper()
	builtins.StopAllHelpers()
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin = filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	libDir := t.TempDir()
	libPath = filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	proj = t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\ncall_timeout = 10000\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("write fig.toml: %v", err)
	}
	return proj, bin, libPath
}

// assertFFIErrorCode checks that err is an *FFIError with the expected code.
func assertFFIErrorCode(t *testing.T, err error, expectedCode string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error with code %s, got nil", expectedCode)
	}
	var ffiErr *builtins.FFIError
	if !errors.As(err, &ffiErr) {
		t.Fatalf("expected *FFIError, got %T: %v", err, err)
	}
	if ffiErr.Code != expectedCode {
		t.Fatalf("expected error code %s, got %s (message: %s)", expectedCode, ffiErr.Code, ffiErr.Message)
	}
}

// TestFfiErrorCodeDlopenFailed verifies ERR_DLOPEN_FAILED when loading a non-existent library.
func TestFfiErrorCodeDlopenFailed(t *testing.T) {
	proj, bin, _ := setupHelperForErrors(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	_, err = hc.Load("/nonexistent/path/libfake.so")
	if err == nil {
		t.Fatal("expected error for dlopen failure, got nil")
	}
	assertFFIErrorCode(t, err, "ERR_DLOPEN_FAILED")
}

// TestFfiErrorCodeInvalidHandle verifies ERR_INVALID_HANDLE when resolving
// a symbol on a bogus handle.
func TestFfiErrorCodeInvalidHandle(t *testing.T) {
	proj, bin, _ := setupHelperForErrors(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	_, err = hc.Sym("bogus-handle-999", "sum2", "double")
	if err == nil {
		t.Fatal("expected error for invalid handle, got nil")
	}
	assertFFIErrorCode(t, err, "ERR_INVALID_HANDLE")
}

// TestFfiErrorCodeDlsymFailed verifies ERR_DLSYM_FAILED for a non-existent symbol.
func TestFfiErrorCodeDlsymFailed(t *testing.T) {
	proj, bin, libPath := setupHelperForErrors(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	_, err = hc.Sym(handle, "nonexistent_function_xyz", "void")
	if err == nil {
		t.Fatal("expected error for dlsym failure, got nil")
	}
	assertFFIErrorCode(t, err, "ERR_DLSYM_FAILED")
}

// TestFfiErrorCodeInvalidMemID verifies ERR_INVALID_MEM_ID when freeing a bogus id.
func TestFfiErrorCodeInvalidMemID(t *testing.T) {
	proj, bin, _ := setupHelperForErrors(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	err = hc.Free("nonexistent-mem-999")
	if err == nil {
		t.Fatal("expected error for invalid mem_id, got nil")
	}
	assertFFIErrorCode(t, err, "ERR_INVALID_MEM_ID")
}

// TestFfiErrorStructuredFormat verifies the FFIError string format: [CODE] message.
func TestFfiErrorStructuredFormat(t *testing.T) {
	proj, bin, _ := setupHelperForErrors(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}

	_, err = hc.Load("/nonexistent/fake.so")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	var ffiErr *builtins.FFIError
	if !errors.As(err, &ffiErr) {
		t.Fatalf("expected *FFIError, got %T: %v", err, err)
	}
	if ffiErr.Code != "ERR_DLOPEN_FAILED" {
		t.Fatalf("expected ERR_DLOPEN_FAILED, got %s", ffiErr.Code)
	}
	// Verify Error() format: [ERR_DLOPEN_FAILED] ...
	s := ffiErr.Error()
	if !strings.Contains(s, "[ERR_DLOPEN_FAILED]") {
		t.Fatalf("Error() should contain [ERR_DLOPEN_FAILED], got: %s", s)
	}
	if ffiErr.Message == "" {
		t.Fatal("expected non-empty error message")
	}
}
