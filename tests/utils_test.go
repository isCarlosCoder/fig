package tests

import (
	"testing"
)

func TestUtilsOrdChrBasic(t *testing.T) {
	src := `use "utils";
print(utils.ord("A"));
print(utils.chr(65));
print(utils.ord("é"));
print(utils.chr(233));
`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := "65\nA\n233\né"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestUtilsOrdErrors(t *testing.T) {
	src := `use "utils";
_ = utils.ord("");
`
	_, err := runFigSource(t, src)
	if err == nil {
		t.Fatalf("expected error for empty char in ord")
	}
}

func TestUtilsChrErrors(t *testing.T) {
	src := `use "utils";
_ = utils.chr(1.5);
`
	_, err := runFigSource(t, src)
	if err == nil {
		t.Fatalf("expected error for non-integer chr")
	}
}

func TestUtilsExtraHelpers(t *testing.T) {
	src := `use "utils";
print(utils.isWhitespace(" "));
print(utils.isWhitespace(utils.fromCodePoint(160)));
print(utils.isUpper("A"));
print(utils.isLower("a"));
print(utils.isAlphaNum("A"));
print(utils.isAlphaNum("3"));
print(utils.fromCodePoints([72,101,108,108,111]));
print(utils.toCodePoints("Hi"));
print(utils.runeCount("A😀"));
print(utils.byteLength("A😀"));
`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := "true\ntrue\ntrue\ntrue\ntrue\ntrue\nHello\n[72, 105]\n2\n5"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}
