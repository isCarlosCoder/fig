package tests

import "testing"

func TestStringEscapeOctal(t *testing.T) {
	src := `print("\033[H\033[2J")`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := "\x1b[H\x1b[2J"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestStringEscapeHex(t *testing.T) {
	src := `print("\x1b[2J")`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := "\x1b[2J"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}
