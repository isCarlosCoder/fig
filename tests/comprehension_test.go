package tests

import "testing"

func TestArrayComprehensionForIn(t *testing.T) {
	src := `let arr = [1,2,3]; let res = [x * 2 for x in arr]; print(res);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[2, 4, 6]" {
		t.Fatalf("expected '[2, 4, 6]', got %q", out)
	}
}

func TestArrayComprehensionRange(t *testing.T) {
	src := `let res = [i for i in range(0, 5)]; print(res);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[0, 1, 2, 3, 4]" {
		t.Fatalf("expected '[0, 1, 2, 3, 4]', got %q", out)
	}
}

func TestArrayComprehensionEnumerate(t *testing.T) {
	src := `let arr = ["a","b"]; let pairs = [idx for idx, v in enumerate(arr)]; print(pairs);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[0, 1]" {
		t.Fatalf("expected '[0, 1]', got %q", out)
	}
}

func TestArrayComprehensionNestedUse(t *testing.T) {
	src := `let square = fn(n) { return n * n }; let res = [square(i) for i in range(0,4)]; print(res);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[0, 1, 4, 9]" {
		t.Fatalf("expected '[0, 1, 4, 9]', got %q", out)
	}
}
