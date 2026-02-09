package tests

import (
	"strings"
	"testing"
)

func useArrays(code string) string {
	return "use " + `"` + "arrays" + `"` + "; " + code
}

func TestArraysPush(t *testing.T) {
	src := useArrays(`let a = [1, 2, 3]; arrays.push(a, 4); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[1, 2, 3, 4]" {
		t.Fatalf("expected '[1, 2, 3, 4]', got %q", out)
	}
}

func TestArraysPop(t *testing.T) {
	src := useArrays(`let a = [1, 2, 3]; let v = arrays.pop(a); print(v); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	if lines[0] != "3" {
		t.Fatalf("expected '3', got %q", lines[0])
	}
	if lines[1] != "[1, 2]" {
		t.Fatalf("expected '[1, 2]', got %q", lines[1])
	}
}

func TestArraysShift(t *testing.T) {
	src := useArrays(`let a = [1, 2, 3]; let v = arrays.shift(a); print(v); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "1" {
		t.Fatalf("expected '1', got %q", lines[0])
	}
	if lines[1] != "[2, 3]" {
		t.Fatalf("expected '[2, 3]', got %q", lines[1])
	}
}

func TestArraysUnshift(t *testing.T) {
	src := useArrays(`let a = [2, 3]; arrays.unshift(a, 1); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[1, 2, 3]" {
		t.Fatalf("expected '[1, 2, 3]', got %q", out)
	}
}

func TestArraysInsert(t *testing.T) {
	src := useArrays(`let a = [10, 30]; arrays.insert(a, 1, 20); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[10, 20, 30]" {
		t.Fatalf("expected '[10, 20, 30]', got %q", out)
	}
}

func TestArraysRemove(t *testing.T) {
	src := useArrays(`let a = [10, 20, 30]; let v = arrays.remove(a, 1); print(v); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "20" {
		t.Fatalf("expected '20', got %q", lines[0])
	}
	if lines[1] != "[10, 30]" {
		t.Fatalf("expected '[10, 30]', got %q", lines[1])
	}
}

func TestArraysSlice(t *testing.T) {
	src := useArrays(`let a = [1, 2, 3, 4, 5]; print(arrays.slice(a, 1, 4));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[2, 3, 4]" {
		t.Fatalf("expected '[2, 3, 4]', got %q", out)
	}
}

func TestArraysConcat(t *testing.T) {
	src := useArrays(`print(arrays.concat([1, 2], [3, 4]));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[1, 2, 3, 4]" {
		t.Fatalf("expected '[1, 2, 3, 4]', got %q", out)
	}
}

func TestArraysReverse(t *testing.T) {
	src := useArrays(`let a = [1, 2, 3]; arrays.reverse(a); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[3, 2, 1]" {
		t.Fatalf("expected '[3, 2, 1]', got %q", out)
	}
}

func TestArraysSort(t *testing.T) {
	src := useArrays(`let a = [3, 1, 4, 1, 5]; arrays.sort(a); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[1, 1, 3, 4, 5]" {
		t.Fatalf("expected '[1, 1, 3, 4, 5]', got %q", out)
	}
}

func TestArraysIndex(t *testing.T) {
	src := useArrays(`let a = ["a", "b", "c"]; print(arrays.index(a, "b")); print(arrays.index(a, "z"));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "1" {
		t.Fatalf("expected '1', got %q", lines[0])
	}
	if lines[1] != "-1" {
		t.Fatalf("expected '-1', got %q", lines[1])
	}
}

func TestArraysContains(t *testing.T) {
	src := useArrays(`let a = [1, 2, 3]; print(arrays.contains(a, 2)); print(arrays.contains(a, 9));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "true" {
		t.Fatalf("expected 'true', got %q", lines[0])
	}
	if lines[1] != "false" {
		t.Fatalf("expected 'false', got %q", lines[1])
	}
}

func TestArraysUnique(t *testing.T) {
	src := useArrays(`print(arrays.unique([1, 2, 2, 3, 3, 3]));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[1, 2, 3]" {
		t.Fatalf("expected '[1, 2, 3]', got %q", out)
	}
}

func TestArraysShuffle(t *testing.T) {
	src := useArrays(`let a = [1, 2, 3, 4, 5]; arrays.shuffle(a); print(arrays.len(a));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "5" {
		t.Fatalf("expected '5', got %q", out)
	}
}

func TestArraysLen(t *testing.T) {
	src := useArrays(`print(arrays.len([10, 20, 30]));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3" {
		t.Fatalf("expected '3', got %q", out)
	}
}

func TestArraysPopEmpty(t *testing.T) {
	src := useArrays(`let a = []; arrays.pop(a);`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for pop on empty array")
	}
}

func TestArraysWrongType(t *testing.T) {
	src := useArrays(`arrays.push("not_array", 1);`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for non-array argument")
	}
}

func TestArraysSortStrings(t *testing.T) {
	src := useArrays(`let a = ["banana", "apple", "cherry"]; arrays.sort(a); print(a);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := `["apple", "banana", "cherry"]`
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}
