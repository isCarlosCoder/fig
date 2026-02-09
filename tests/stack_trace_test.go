package tests

import (
	"strings"
	"testing"
)

func TestRuntimeErrorShowsStackTraceForImportedFunction(t *testing.T) {
	out, err := runFigFile(t, "stack_main.fig")
	if err == nil {
		t.Fatalf("expected runtime error, got output: %q", out)
	}
	errStr := err.Error()
	if !strings.Contains(errStr, "stack_module.fig") {
		t.Fatalf("expected stack trace to mention module filename, got: %v", errStr)
	}
	if !strings.Contains(errStr, "bad") {
		t.Fatalf("expected stack trace to mention function name 'bad', got: %v", errStr)
	}
}

func TestRuntimeErrorShowsStackTraceForLocalFunction(t *testing.T) {
	out, err := runFigFile(t, "stack_local.fig")
	if err == nil {
		t.Fatalf("expected runtime error, got output: %q", out)
	}
	errStr := err.Error()
	if !strings.Contains(errStr, "stack_local.fig") {
		t.Fatalf("expected stack trace to mention filename, got: %v", errStr)
	}
	if !strings.Contains(errStr, "bad") {
		t.Fatalf("expected stack trace to mention function name 'bad', got: %v", errStr)
	}
}
