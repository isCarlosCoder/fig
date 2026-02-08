package tests

import (
	"testing"
)

func TestArrayLiteralEmpty(t *testing.T) {
	out, err := runFig(t, "let a = []; print(a);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[]" {
		t.Fatalf("expected '[]', got %q", out)
	}
}

func TestArrayLiteralValues(t *testing.T) {
	out, err := runFig(t, "let a = [1, 2, 3]; print(a);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[1, 2, 3]" {
		t.Fatalf("expected '[1, 2, 3]', got %q", out)
	}
}

func TestArrayMixedTypes(t *testing.T) {
	src := "let a = [1, " + `"hello"` + ", true, null]; print(a);"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := "[1, " + `"hello"` + ", true, null]"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestArrayIndexAccess(t *testing.T) {
	out, err := runFig(t, "let a = [10, 20, 30]; print(a[0]); print(a[1]); print(a[2]);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10\n20\n30" {
		t.Fatalf("expected '10\\n20\\n30', got %q", out)
	}
}

func TestArrayIndexAssign(t *testing.T) {
	out, err := runFig(t, "let a = [1, 2, 3]; a[1] = 99; print(a);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[1, 99, 3]" {
		t.Fatalf("expected '[1, 99, 3]', got %q", out)
	}
}

func TestArrayOutOfBounds(t *testing.T) {
	_, err := runFig(t, "let a = [1, 2]; print(a[5]);")
	if err == nil {
		t.Fatal("expected out of bounds error")
	}
}

func TestArrayIsTruthy(t *testing.T) {
	src := `let a = [1]; if (a) { print("yes"); } else { print("no"); }`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "yes" {
		t.Fatalf("expected 'yes', got %q", out)
	}
}

func TestEmptyArrayIsFalsy(t *testing.T) {
	src := `let a = []; if (a) { print("yes"); } else { print("no"); }`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "no" {
		t.Fatalf("expected 'no', got %q", out)
	}
}

func TestArrayEquality(t *testing.T) {
	out, err := runFig(t, "print([1, 2] == [1, 2]); print([1, 2] == [1, 3]);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true\nfalse" {
		t.Fatalf("expected 'true\\nfalse', got %q", out)
	}
}

func TestNestedArrays(t *testing.T) {
	out, err := runFig(t, "let a = [[1, 2], [3, 4]]; print(a[0][1]); print(a[1][0]);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "2\n3" {
		t.Fatalf("expected '2\\n3', got %q", out)
	}
}

func TestObjectLiteralEmpty(t *testing.T) {
	out, err := runFig(t, "let o = {}; print(o);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "{}" {
		t.Fatalf("expected '{}', got %q", out)
	}
}

func TestObjectLiteral(t *testing.T) {
	src := `let o = {name: "Fig", version: 1}; print(o);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := `{name: "Fig", version: 1}`
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestObjectDotAccess(t *testing.T) {
	src := `let o = {name: "Fig", year: 2026}; print(o.name); print(o.year);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "Fig\n2026" {
		t.Fatalf("expected 'Fig\\n2026', got %q", out)
	}
}

func TestObjectBracketAccess(t *testing.T) {
	src := `let o = {x: 10}; print(o["x"]);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10" {
		t.Fatalf("expected '10', got %q", out)
	}
}

func TestObjectDotAssign(t *testing.T) {
	out, err := runFig(t, "let o = {x: 1}; o.x = 99; print(o.x);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "99" {
		t.Fatalf("expected '99', got %q", out)
	}
}

func TestObjectBracketAssign(t *testing.T) {
	src := `let o = {x: 1}; o["x"] = 42; print(o.x);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "42" {
		t.Fatalf("expected '42', got %q", out)
	}
}

func TestObjectAddNewKey(t *testing.T) {
	src := `let o = {}; o.name = "Fig"; print(o.name);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "Fig" {
		t.Fatalf("expected 'Fig', got %q", out)
	}
}

func TestObjectMissingKeyReturnsNull(t *testing.T) {
	out, err := runFig(t, "let o = {a: 1}; print(o.b);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "null" {
		t.Fatalf("expected 'null', got %q", out)
	}
}

func TestObjectIsTruthy(t *testing.T) {
	src := `let o = {a: 1}; if (o) { print("yes"); } else { print("no"); }`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "yes" {
		t.Fatalf("expected 'yes', got %q", out)
	}
}

func TestEmptyObjectIsFalsy(t *testing.T) {
	src := `let o = {}; if (o) { print("yes"); } else { print("no"); }`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "no" {
		t.Fatalf("expected 'no', got %q", out)
	}
}

func TestObjectEquality(t *testing.T) {
	out, err := runFig(t, "print({a: 1} == {a: 1}); print({a: 1} == {a: 2});")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true\nfalse" {
		t.Fatalf("expected 'true\\nfalse', got %q", out)
	}
}

func TestNestedObjects(t *testing.T) {
	out, err := runFig(t, "let o = {inner: {x: 42}}; print(o.inner.x);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "42" {
		t.Fatalf("expected '42', got %q", out)
	}
}

func TestArrayOfObjects(t *testing.T) {
	src := `let a = [{name: "a"}, {name: "b"}]; print(a[0].name); print(a[1].name);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "a\nb" {
		t.Fatalf("expected 'a\\nb', got %q", out)
	}
}

func TestObjectWithArray(t *testing.T) {
	out, err := runFig(t, "let o = {items: [10, 20, 30]}; print(o.items[1]);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "20" {
		t.Fatalf("expected '20', got %q", out)
	}
}

func TestArrayInLoop(t *testing.T) {
	src := "let a = [10, 20, 30]; let s = 0; for (let i = 0; i < 3; i++) { s = s + a[i]; } print(s);"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "60" {
		t.Fatalf("expected '60', got %q", out)
	}
}

func TestFunctionReturnsArray(t *testing.T) {
	src := "fn makeArr(n) { let a = []; let i = 0; while (i < n) { a[i] = i; i++; } return a; } print(makeArr(3));"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[0, 1, 2]" {
		t.Fatalf("expected '[0, 1, 2]', got %q", out)
	}
}

func TestFunctionReturnsObject(t *testing.T) {
	src := `fn make(n, val) { let o = {}; o.name = n; o.val = val; return o; } let r = make("x", 5); print(r.name); print(r.val);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "x\n5" {
		t.Fatalf("expected 'x\\n5', got %q", out)
	}
}
