package tests

import (
	"strings"
	"testing"
)

func TestInfiniteRecursionIsHandled(t *testing.T) {
	src := `fn f() { return f() } f()`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected runtime error for infinite recursion, got nil")
	}
	if err.Error() != "maximum call depth exceeded" {
		// Allow the exact formatting from runtime error wrapper
		if !strings.Contains(err.Error(), "maximum call depth") {
			t.Fatalf("unexpected error message: %v", err)
		}
	}
}
