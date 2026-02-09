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
	// ensure snippet and caret are present for the originating frame
	if !strings.Contains(errStr, "let x = s - 1") || !strings.Contains(errStr, "^") {
		t.Fatalf("expected snippet and caret in error, got: %v", errStr)
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
	// ensure snippet and caret are present for the originating frame
	if !strings.Contains(errStr, "let x = s - 1") || !strings.Contains(errStr, "^") {
		t.Fatalf("expected snippet and caret in error, got: %v", errStr)
	}
}
