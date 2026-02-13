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

// TestCLIReplWithIFlag verifies that `fig -i <file>` preloads a .fig file into the REPL
// environment so variables/functions defined in the file are available interactively.
func TestCLIReplWithIFlag(t *testing.T) {
	td := t.TempDir()
	path := filepath.Join(td, "preload.fig")
	src := `let answer = 123
fn say() { print("hello") }
`
	if err := os.WriteFile(path, []byte(src), 0644); err != nil {
		t.Fatalf("write preload file: %v", err)
	}

	// Run the CLI with -i and interact with the REPL: print(answer) then exit
	cmd := exec.Command("go", "run", ".", "-i", path)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("stdin pipe: %v", err)
	}

	// ensure `go run .` executes from the repository root (one level up from tests/)
	if wd, err := os.Getwd(); err == nil {
		cmd.Dir = filepath.Join(wd, "..")
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("start cli: %v (stderr=%s)", err, stderr.String())
	}

	// send commands to REPL: bare-expression echo + function call
	io.WriteString(stdin, "answer\n")
	io.WriteString(stdin, "say()\n")
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
	if !bytes.Contains(stdout.Bytes(), []byte("123")) {
		t.Fatalf("expected '123' in stdout, got:\n%s\nstderr:\n%s", out, stderr.String())
	}
	if !bytes.Contains(stdout.Bytes(), []byte("hello")) {
		t.Fatalf("expected 'hello' in stdout, got:\n%s\nstderr:\n%s", out, stderr.String())
	}
}
