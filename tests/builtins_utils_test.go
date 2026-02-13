package tests

import (
	"testing"
)

func TestUtilsZipBasic(t *testing.T) {
	src := `use "utils";
let a = [1,2,3]; let b = ["x","y","z"]; let z = utils.zip(a,b); print(z[0][0]); print(z[0][1]); print(z[2][0]); print(z[2][1]);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "1\nx\n3\nz" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestUtilsZipDifferentLengths(t *testing.T) {
	src := `use "utils"; use "arrays";
let a = [1,2]; let b = [10,20,30]; let z = utils.zip(a,b);
print(arrays.len(z));
print(z[0][0]); print(z[0][1]); print(z[1][0]); print(z[1][1]);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	// zip cuts to the shortest array; z length == 2 and elements match
	if out != "2\n1\n10\n2\n20" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestUtilsZipArgError(t *testing.T) {
	src := `print(utils.zip(1, [2,3]))`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for non-array arg")
	}
}
