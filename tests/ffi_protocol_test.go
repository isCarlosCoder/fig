package tests

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
)

// setupProtocolTest builds the helper binary and test lib, creates a project dir with fig.toml.
func setupProtocolTest(t *testing.T) (proj, bin, libPath string) {
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

// TestProtocolLoadMissingPath verifies that Load("") returns ERR_MISSING_PARAM.
func TestProtocolLoadMissingPath(t *testing.T) {
	proj, bin, _ := setupProtocolTest(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}

	_, err = hc.Load("")
	if err == nil {
		t.Fatal("expected error for empty path, got nil")
	}
	var ffiErr *builtins.FFIError
	if !errors.As(err, &ffiErr) {
		t.Fatalf("expected *FFIError, got %T: %v", err, err)
	}
	if ffiErr.Code != "ERR_MISSING_PARAM" {
		t.Fatalf("expected ERR_MISSING_PARAM, got %s (message: %s)", ffiErr.Code, ffiErr.Message)
	}
}

// TestProtocolLoadNonexistent verifies that Load("/fake/lib.so") returns ERR_DLOPEN_FAILED.
func TestProtocolLoadNonexistent(t *testing.T) {
	proj, bin, _ := setupProtocolTest(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}

	_, err = hc.Load("/fake/lib.so")
	if err == nil {
		t.Fatal("expected error for nonexistent library, got nil")
	}
	var ffiErr *builtins.FFIError
	if !errors.As(err, &ffiErr) {
		t.Fatalf("expected *FFIError, got %T: %v", err, err)
	}
	if ffiErr.Code != "ERR_DLOPEN_FAILED" {
		t.Fatalf("expected ERR_DLOPEN_FAILED, got %s (message: %s)", ffiErr.Code, ffiErr.Message)
	}
}

// TestProtocolSymInvalidHandle verifies that Sym with a bogus handle returns ERR_INVALID_HANDLE.
func TestProtocolSymInvalidHandle(t *testing.T) {
	proj, bin, _ := setupProtocolTest(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}

	_, err = hc.Sym("fake-handle", "func", "void")
	if err == nil {
		t.Fatal("expected error for invalid handle, got nil")
	}
	var ffiErr *builtins.FFIError
	if !errors.As(err, &ffiErr) {
		t.Fatalf("expected *FFIError, got %T: %v", err, err)
	}
	if ffiErr.Code != "ERR_INVALID_HANDLE" {
		t.Fatalf("expected ERR_INVALID_HANDLE, got %s (message: %s)", ffiErr.Code, ffiErr.Message)
	}
}

// TestProtocolSymNonexistent loads a valid library and tries to resolve a symbol that does not exist.
func TestProtocolSymNonexistent(t *testing.T) {
	proj, bin, libPath := setupProtocolTest(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}

	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}

	_, err = hc.Sym(handle, "nonexistent_xyz", "void")
	if err == nil {
		t.Fatal("expected error for nonexistent symbol, got nil")
	}
	var ffiErr *builtins.FFIError
	if !errors.As(err, &ffiErr) {
		t.Fatalf("expected *FFIError, got %T: %v", err, err)
	}
	if ffiErr.Code != "ERR_DLSYM_FAILED" {
		t.Fatalf("expected ERR_DLSYM_FAILED, got %s (message: %s)", ffiErr.Code, ffiErr.Message)
	}
}

// TestProtocolCallInvalidSymbol verifies that CallSymbol with a fake symbol ID returns ERR_INVALID_SYMBOL.
func TestProtocolCallInvalidSymbol(t *testing.T) {
	proj, bin, _ := setupProtocolTest(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}

	_, err = hc.CallSymbol("fake-sym", nil, nil)
	if err == nil {
		t.Fatal("expected error for invalid symbol, got nil")
	}
	var ffiErr *builtins.FFIError
	if !errors.As(err, &ffiErr) {
		t.Fatalf("expected *FFIError, got %T: %v", err, err)
	}
	if ffiErr.Code != "ERR_INVALID_SYMBOL" {
		t.Fatalf("expected ERR_INVALID_SYMBOL, got %s (message: %s)", ffiErr.Code, ffiErr.Message)
	}
}

// TestProtocolFreeInvalidMemID verifies that Free with a bogus memory ID returns ERR_INVALID_MEM_ID.
func TestProtocolFreeInvalidMemID(t *testing.T) {
	proj, bin, _ := setupProtocolTest(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}

	err = hc.Free("fake-mem")
	if err == nil {
		t.Fatal("expected error for invalid mem_id, got nil")
	}
	var ffiErr *builtins.FFIError
	if !errors.As(err, &ffiErr) {
		t.Fatalf("expected *FFIError, got %T: %v", err, err)
	}
	if ffiErr.Code != "ERR_INVALID_MEM_ID" {
		t.Fatalf("expected ERR_INVALID_MEM_ID, got %s (message: %s)", ffiErr.Code, ffiErr.Message)
	}
}

// TestProtocolBurstRequests sends 100 Load requests that all fail with ERR_DLOPEN_FAILED
// and verifies every error is returned correctly without crashing the helper.
func TestProtocolBurstRequests(t *testing.T) {
	proj, bin, _ := setupProtocolTest(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}

	const N = 100
	for i := 0; i < N; i++ {
		fakePath := fmt.Sprintf("/nonexistent/burst_%d.so", i)
		_, err := hc.Load(fakePath)
		if err == nil {
			t.Fatalf("burst request %d: expected error, got nil", i)
		}
		var ffiErr *builtins.FFIError
		if !errors.As(err, &ffiErr) {
			t.Fatalf("burst request %d: expected *FFIError, got %T: %v", i, err, err)
		}
		if ffiErr.Code != "ERR_DLOPEN_FAILED" {
			t.Fatalf("burst request %d: expected ERR_DLOPEN_FAILED, got %s (message: %s)", i, ffiErr.Code, ffiErr.Message)
		}
	}
}

// TestProtocolUnknownCommand is skipped because the client wrapper does not expose
// a way to send arbitrary commands to the helper process.
func TestProtocolUnknownCommand(t *testing.T) {
	t.Skip("cannot send unknown commands via the HelperForTest client wrapper")
}
