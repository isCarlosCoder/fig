package tests

import (
	"os"
	"strings"
	"testing"
)

func useIO(code string) string {
	return `use "io"; ` + code
}

func TestIOWriteAndReadFile(t *testing.T) {
	tmp := t.TempDir() + "/test.txt"
	code := `io.writeFile("` + tmp + `", "hello fig"); let c = io.readFile("` + tmp + `"); print(c)`
	out, err := runFig(t, useIO(code))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "hello fig" {
		t.Errorf("expected 'hello fig', got %q", out)
	}
}

func TestIOAppendFile(t *testing.T) {
	tmp := t.TempDir() + "/test.txt"
	code := `io.writeFile("` + tmp + `", "hello"); io.appendFile("` + tmp + `", " world"); let c = io.readFile("` + tmp + `"); print(c)`
	out, err := runFig(t, useIO(code))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "hello world" {
		t.Errorf("expected 'hello world', got %q", out)
	}
}

func TestIOExistsTrue(t *testing.T) {
	tmp := t.TempDir() + "/test.txt"
	os.WriteFile(tmp, []byte("data"), 0644)
	code := `let e = io.exists("` + tmp + `"); print(e)`
	out, err := runFig(t, useIO(code))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "true" {
		t.Errorf("expected true, got %q", out)
	}
}

func TestIOExistsFalse(t *testing.T) {
	code := `let e = io.exists("/tmp/fig_nonexistent_file_xyz.txt"); print(e)`
	out, err := runFig(t, useIO(code))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "false" {
		t.Errorf("expected false, got %q", out)
	}
}

func TestIODeleteFile(t *testing.T) {
	tmp := t.TempDir() + "/test.txt"
	os.WriteFile(tmp, []byte("data"), 0644)
	code := `io.deleteFile("` + tmp + `"); let e = io.exists("` + tmp + `"); print(e)`
	out, err := runFig(t, useIO(code))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "false" {
		t.Errorf("expected false, got %q", out)
	}
}

func TestIOReadFileNotFound(t *testing.T) {
	code := `io.readFile("/tmp/fig_nonexistent_xyz.txt")`
	_, err := runFig(t, useIO(code))
	if err == nil {
		t.Fatal("expected error for readFile on missing file")
	}
}

func TestIODeleteFileNotFound(t *testing.T) {
	code := `io.deleteFile("/tmp/fig_nonexistent_xyz.txt")`
	_, err := runFig(t, useIO(code))
	if err == nil {
		t.Fatal("expected error for deleteFile on missing file")
	}
}

func TestIOWriteFileArgError(t *testing.T) {
	_, err := runFig(t, useIO(`io.writeFile("path")`))
	if err == nil {
		t.Fatal("expected error for writeFile with 1 arg")
	}
}

func TestIOReadFileArgError(t *testing.T) {
	_, err := runFig(t, useIO(`io.readFile()`))
	if err == nil {
		t.Fatal("expected error for readFile with 0 args")
	}
}

func TestIOExistsArgError(t *testing.T) {
	_, err := runFig(t, useIO(`io.exists()`))
	if err == nil {
		t.Fatal("expected error for exists with 0 args")
	}
}

func TestIOWriteReadMultiline(t *testing.T) {
	tmp := t.TempDir() + "/multi.txt"
	code := `io.writeFile("` + tmp + `", "line1\nline2\nline3"); let c = io.readFile("` + tmp + `"); print(c)`
	out, err := runFig(t, useIO(code))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "line1") {
		t.Errorf("expected multiline content, got %q", out)
	}
}

func TestIOAppendCreatesFile(t *testing.T) {
	tmp := t.TempDir() + "/new.txt"
	code := `io.appendFile("` + tmp + `", "created"); let c = io.readFile("` + tmp + `"); print(c)`
	out, err := runFig(t, useIO(code))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "created" {
		t.Errorf("expected 'created', got %q", out)
	}
}
