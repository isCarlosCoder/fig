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
