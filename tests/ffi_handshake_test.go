package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
)

// TestFfiHandshakeSuccess verifies that handshake succeeds with matching version.
func TestFfiHandshakeSuccess(t *testing.T) {
	builtins.StopAllHelpers()
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\ncall_timeout = 10000\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("write fig.toml: %v", err)
	}

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("expected handshake to succeed, got: %v", err)
	}

	// Verify the helper is functional after handshake
	_, err = hc.Load("/nonexistent.so")
	if err == nil {
		t.Fatal("expected dlopen error")
	}
	if strings.Contains(err.Error(), "handshake") || strings.Contains(err.Error(), "mismatch") {
		t.Fatalf("unexpected handshake error: %v", err)
	}

	builtins.StopAllHelpers()
}

// TestFfiHandshakeFullFlow verifies the helper works correctly after handshake.
func TestFfiHandshakeFullFlow(t *testing.T) {
	builtins.StopAllHelpers()
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	proj := t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\ncall_timeout = 10000\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("write fig.toml: %v", err)
	}

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("handshake failed: %v", err)
	}

	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load after handshake failed: %v", err)
	}
	if handle == "" {
		t.Fatal("expected non-empty handle")
	}
	sym, err := hc.Sym(handle, "mul2", "double")
	if err != nil {
		t.Fatalf("sym after handshake failed: %v", err)
	}
	result, err := hc.CallSymbol(sym, []interface{}{5.0, 6.0}, nil)
	if err != nil {
		t.Fatalf("call after handshake failed: %v", err)
	}
	if num, ok := result.(float64); !ok || num != 30.0 {
		t.Fatalf("expected 30.0, got %v", result)
	}

	builtins.StopAllHelpers()
}
