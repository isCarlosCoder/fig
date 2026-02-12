package tests

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// TestFfiLogLevelWarnQuiet verifies that at log level "warn" (default),
// normal requests do NOT produce any output on stderr.
func TestFfiLogLevelWarnQuiet(t *testing.T) {
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	build := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	build.Dir = root
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	// Start helper with --log-level warn
	cmd := exec.Command(bin, "--server", "--log-level", "warn")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start helper: %v", err)
	}
	defer func() {
		stdin.Close()
		cmd.Wait()
	}()

	// Collect stderr in background
	stderrCh := make(chan string, 1)
	go func() {
		sc := bufio.NewScanner(stderrPipe)
		var lines []string
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		stderrCh <- strings.Join(lines, "\n")
	}()

	enc := json.NewEncoder(stdin)
	dec := json.NewDecoder(stdout)

	// Send a few normal requests: ping, call (no symbol — fallback sum)
	for i := 0; i < 5; i++ {
		enc.Encode(map[string]interface{}{"cmd": "ping", "id": float64(i + 1)})
		var resp map[string]interface{}
		if err := dec.Decode(&resp); err != nil {
			t.Fatalf("decode error on ping %d: %v", i, err)
		}
	}

	// Close stdin to signal EOF
	stdin.Close()
	cmd.Wait()

	// Collect stderr
	var stderrOut string
	select {
	case stderrOut = <-stderrCh:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for stderr collection")
	}

	// At warn level, "server started" (info) and "req ping" (debug) should NOT appear
	if strings.Contains(stderrOut, "[DEBUG]") {
		t.Errorf("expected no DEBUG output at warn level, got:\n%s", stderrOut)
	}
	if strings.Contains(stderrOut, "[INFO]") {
		t.Errorf("expected no INFO output at warn level, got:\n%s", stderrOut)
	}
}

// TestFfiLogLevelDebugVerbose verifies that at log level "debug",
// request details appear on stderr.
func TestFfiLogLevelDebugVerbose(t *testing.T) {
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	build := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	build.Dir = root
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	// Build a test library
	libDir := t.TempDir()
	libPath := filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	// Start helper with --log-level debug
	cmd := exec.Command(bin, "--server", "--log-level", "debug")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start helper: %v", err)
	}
	defer func() {
		stdin.Close()
		cmd.Wait()
	}()

	// Collect stderr in background
	stderrCh := make(chan string, 1)
	go func() {
		sc := bufio.NewScanner(stderrPipe)
		var lines []string
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		stderrCh <- strings.Join(lines, "\n")
	}()

	enc := json.NewEncoder(stdin)
	dec := json.NewDecoder(stdout)

	// Send a ping
	enc.Encode(map[string]interface{}{"cmd": "ping", "id": float64(1)})
	var resp map[string]interface{}
	dec.Decode(&resp)

	// Load a library
	enc.Encode(map[string]interface{}{"cmd": "load", "path": libPath, "id": float64(2)})
	dec.Decode(&resp)
	handle, _ := resp["result"].(map[string]interface{})["handle"].(string)

	// Resolve a symbol
	enc.Encode(map[string]interface{}{"cmd": "sym", "handle": handle, "name": "sum3", "rtype": "int", "id": float64(3)})
	dec.Decode(&resp)

	// Close
	stdin.Close()
	cmd.Wait()

	// Collect stderr
	var stderrOut string
	select {
	case stderrOut = <-stderrCh:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for stderr collection")
	}

	// At debug level, we should see [INFO] "server started" and [DEBUG] "req ... id=" lines
	if !strings.Contains(stderrOut, "[INFO]") {
		t.Errorf("expected INFO output at debug level, got:\n%s", stderrOut)
	}
	if !strings.Contains(stderrOut, "[DEBUG]") {
		t.Errorf("expected DEBUG output at debug level, got:\n%s", stderrOut)
	}
	// Verify specific request debug-logging
	if !strings.Contains(stderrOut, "req") || !strings.Contains(stderrOut, "ping") {
		t.Errorf("expected request debug log for 'ping', got:\n%s", stderrOut)
	}

	// Also verify we see "server started" info-level message
	if !strings.Contains(stderrOut, "server started") {
		t.Errorf("expected 'server started' info log at debug level, got:\n%s", stderrOut)
	}
}

// TestFfiLogLevelEnvVar verifies that FFI_LOG_LEVEL env var is respected.
func TestFfiLogLevelEnvVar(t *testing.T) {
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin := filepath.Join(binDir, "ffi-helper")
	build := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	build.Dir = root
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	// Start helper with env var FFI_LOG_LEVEL=debug (no --log-level flag)
	cmd := exec.Command(bin, "--server")
	cmd.Env = append(os.Environ(), "FFI_LOG_LEVEL=debug")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start helper: %v", err)
	}
	defer func() {
		stdin.Close()
		cmd.Wait()
	}()

	stderrCh := make(chan string, 1)
	go func() {
		sc := bufio.NewScanner(stderrPipe)
		var lines []string
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		stderrCh <- strings.Join(lines, "\n")
	}()

	enc := json.NewEncoder(stdin)
	dec := json.NewDecoder(stdout)

	// Send a ping
	enc.Encode(map[string]interface{}{"cmd": "ping", "id": float64(1)})
	var resp map[string]interface{}
	dec.Decode(&resp)

	stdin.Close()
	cmd.Wait()

	var stderrOut string
	select {
	case stderrOut = <-stderrCh:
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for stderr collection")
	}

	// env var should enable debug output
	if !strings.Contains(stderrOut, "[DEBUG]") {
		t.Errorf("expected DEBUG output via FFI_LOG_LEVEL=debug env, got:\n%s", stderrOut)
	}

	_ = fmt.Sprintf("") // keep fmt import
}
