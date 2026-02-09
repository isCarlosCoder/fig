package tests

import (
	"strings"
	"testing"
)

func TestImportRuntimeErrorShowsFileAndSnippet(t *testing.T) {
	out, err := runFigFile(t, "error_bad_number.fig")
	if err == nil {
		t.Fatalf("expected runtime error, got output: %q", out)
	}
	// error should contain filename and line number and helpful message
	errStr := err.Error()
	if !strings.Contains(errStr, "error_bad_number.fig:") {
		t.Fatalf("expected error to mention filename, got: %v", errStr)
	}
	if !strings.Contains(errStr, "not a number") {
		t.Fatalf("expected error message about 'not a number', got: %v", errStr)
	}
	// ensure snippet and caret are present
	if !strings.Contains(errStr, "let x = s - 1") || !strings.Contains(errStr, "^") {
		t.Fatalf("expected snippet and caret in error, got: %v", errStr)
	}
}
