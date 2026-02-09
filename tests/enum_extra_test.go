package tests

import (
	"testing"
)

func TestEnumKeysReturnsMemberNames(t *testing.T) {
	src := useObjects(`enum Color { Red Green Blue }
print(objects.keys(Color));`)
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != `["Red", "Green", "Blue"]` {
		t.Fatalf("expected keys, got %q", out)
	}
}

func TestEnumMatchPattern(t *testing.T) {
	src := `enum Color { Red Green Blue }
let c = Color.Green;
let s = match c {
  Color.Red => "red"
  Color.Green => "green"
  Color.Blue => "blue"
  _ => "other"
}
print(s);
`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "green" {
		t.Fatalf("expected 'green', got %q", out)
	}
}

func TestEnumBracketAccess(t *testing.T) {
	src := `enum Color { Red Green Blue }
print(Color["Blue"]);
let x = Color["Blue"];
print(x == Color.Blue);
`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "Color.Blue\ntrue" {
		t.Fatalf("expected \"Color.Blue\\ntrue\", got %q", out)
	}
}
