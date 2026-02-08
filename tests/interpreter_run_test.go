package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
)

func TestRunSuccess(t *testing.T) {
	input := "let x = 7; print(x);"
	var buf bytes.Buffer
	var errBuf bytes.Buffer
	if err := interpreter.Run(input, "<test>", environment.NewEnv(nil), &buf, &errBuf); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if errBuf.Len() != 0 {
		t.Fatalf("expected no error output, got: %q", errBuf.String())
	}
	out := strings.TrimSpace(buf.String())
	if out != "7" {
		t.Fatalf("expected 7 got %q", out)
	}
}

func TestRunParseError(t *testing.T) {
	input := "let 123 = 1;"
	var buf bytes.Buffer
	var errBuf bytes.Buffer
	err := interpreter.Run(input, "<test>", environment.NewEnv(nil), &buf, &errBuf)
	if err == nil {
		t.Fatalf("expected parse error but got nil")
	}
	out := errBuf.String()
	if !strings.Contains(out, "error:") {
		t.Fatalf("expected pretty error output in error buffer, got: %q", out)
	}
}
