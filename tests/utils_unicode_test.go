package tests

import (
	"testing"
)

func TestUtilsUnicodeHelpers(t *testing.T) {
	src := `use "utils";
print(utils.codePointAt("A", 0));
print(utils.codePointAt("😀", 0));
print(utils.fromCodePoint(128512));
print(utils.codePoints("A😀"));
let comb = "e" + utils.fromCodePoint(769);
print(utils.normalize("NFC", comb));
print(utils.isLetter("é"));
print(utils.isDigit("3"));
`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := "65\n128512\n😀\n[65, 128512]\né\ntrue\ntrue"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestUtilsUnicodeErrors(t *testing.T) {
	// out of range index
	src := `use "utils";
_ = utils.codePointAt("A", 1);
`
	_, err := runFigSource(t, src)
	if err == nil {
		t.Fatalf("expected error for index out of range")
	}
}
