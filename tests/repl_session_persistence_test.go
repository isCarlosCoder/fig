package tests

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// TestReplSessionPersistence verifies that the REPL preserves definitions in the
// session: a `let` followed by a bare reference should print the stored value.
func TestReplSessionPersistence(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("stdin pipe: %v", err)
	}

	// ensure `go run .` executes from the repository root
	if wd, err := os.Getwd(); err == nil {
		cmd.Dir = filepath.Join(wd, "..")
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("start cli: %v (stderr=%s)", err, stderr.String())
	}

	// define a variable, then reference it as a bare expression
	io.WriteString(stdin, "let a = 1\n")
	io.WriteString(stdin, "a\n")
	io.WriteString(stdin, "exit\n")
	stdin.Close()

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("process error: %v (stderr=%s)", err, stderr.String())
		}
	case <-time.After(5 * time.Second):
		cmd.Process.Kill()
		t.Fatal("timeout waiting for fig process")
	}

	out := stdout.String()
	if !bytes.Contains(stdout.Bytes(), []byte("1")) {
		t.Fatalf("expected '1' in stdout, got:\n%s\nstderr:\n%s", out, stderr.String())
	}
}
