package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
)

func TestRunUndefinedVariable(t *testing.T) {
	input := "print(x);"
	var errBuf bytes.Buffer
	err := interpreter.Run(input, "<test>", environment.NewEnv(nil), nil, &errBuf)
	if err == nil {
		t.Fatalf("expected runtime error for undefined variable")
	}
	if !strings.Contains(err.Error(), "not defined") {
		t.Fatalf("expected 'not defined' in error message, got: %v", err)
	}
	if !strings.Contains(errBuf.String(), "variable 'x' not defined") {
		t.Fatalf("expected pretty error printed to errBuf, got: %q", errBuf.String())
	}
}

func TestRunDivisionByZero(t *testing.T) {
	input := "let a = 1 / 0; print(a);"
	var errBuf bytes.Buffer
	err := interpreter.Run(input, "<test>", environment.NewEnv(nil), nil, &errBuf)
	if err == nil {
		t.Fatalf("expected runtime error for division by zero")
	}
	if !strings.Contains(err.Error(), "division by zero") {
		t.Fatalf("expected 'division by zero' in error message, got: %v", err)
	}
	if !strings.Contains(errBuf.String(), "division by zero") {
		t.Fatalf("expected pretty error printed to errBuf, got: %q", errBuf.String())
	}
}
