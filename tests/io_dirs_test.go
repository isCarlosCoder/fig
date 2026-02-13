package tests

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/interpreter"
)

func runSrcToString(src, fname string) (string, error) {
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer
	err := interpreter.Run(src, fname, nil, &outBuf, &errBuf)
	if err != nil {
		return "", err
	}
	if errBuf.Len() > 0 {
		return "", fmt.Errorf("stderr: %s", errBuf.String())
	}
	return strings.TrimSpace(outBuf.String()), nil
}

func TestIoMkdirAndIsDir(t *testing.T) {
	proj := t.TempDir()
	path := filepath.Join(proj, "dirA")
	src := fmt.Sprintf("use \"io\"\nio.mkdir(%q)\nprint(io.isDir(%q))", path, path)
	out, err := runSrcToString(src, "_test_.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true" {
		t.Fatalf("expected 'true', got %q", out)
	}
}

func TestIoMkdirAllAndReadDir(t *testing.T) {
	proj := t.TempDir()
	nested := filepath.Join(proj, "a/b/c")
	src := fmt.Sprintf("use \"io\"\nio.mkdirAll(%q)\nprint(io.exists(%q))", nested, nested)
	out, err := runSrcToString(src, "_test_.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true" {
		t.Fatalf("expected 'true', got %q", out)
	}

	// create files to list
	f1 := filepath.Join(proj, "a", "x.txt")
	_ = os.WriteFile(f1, []byte("x"), 0644)
	f2 := filepath.Join(proj, "a", "b", "y.txt")
	_ = os.WriteFile(f2, []byte("y"), 0644)

	srcList := fmt.Sprintf("use \"io\"\nuse \"arrays\"\nlet arr = io.readDir(%q)\nfor i in range(0, arrays.len(arr)) { print(arr[i]) }", filepath.Join(proj, "a"))
	out2, err := runSrcToString(srcList, "_test_.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	// should contain x.txt and b (order not guaranteed)
	if !(containsLine(out2, "x.txt") && containsLine(out2, "b")) {
		t.Fatalf("readDir missing entries: %q", out2)
	}
}

func TestIoRmdirAndRmdirAll(t *testing.T) {
	proj := t.TempDir()
	path := filepath.Join(proj, "tmpdir")
	_ = os.Mkdir(path, 0755)
	// rmdir on empty dir
	src := fmt.Sprintf("use \"io\"\nio.rmdir(%q)\nprint(io.exists(%q))", path, path)
	out, err := runSrcToString(src, "_test_.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "false" {
		t.Fatalf("expected 'false', got %q", out)
	}

	// rmdirAll on nested tree
	nested := filepath.Join(proj, "r1/r2/r3")
	_ = os.MkdirAll(nested, 0755)
	file := filepath.Join(nested, "f.txt")
	_ = os.WriteFile(file, []byte("ok"), 0644)
	src2 := fmt.Sprintf("use \"io\"\nio.rmdirAll(%q)\nprint(io.exists(%q))", filepath.Join(proj, "r1"), filepath.Join(proj, "r1"))
	out2, err := runSrcToString(src2, "_test_.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out2 != "false" {
		t.Fatalf("expected 'false', got %q", out2)
	}
}

// helper: check if output contains exact line
func containsLine(s, sub string) bool {
	for _, ln := range strings.Split(s, "\n") {
		if ln == sub {
			return true
		}
	}
	return false
}
