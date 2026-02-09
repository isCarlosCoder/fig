package tests

import (
	"testing"
)

func TestEnumBasicPrintingAndEquality(t *testing.T) {
	src := `enum Color {
  Red
  Green
  Blue
}
print(Color.Red);
let c = Color.Red;
print(c == Color.Red);
print(c == Color.Green);
`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := "Color.Red\ntrue\nfalse"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}
