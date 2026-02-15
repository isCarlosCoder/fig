package tests

import (
	"strings"
	"testing"
)

// ── range ──

func TestForRangeBasic(t *testing.T) {
	src := `
let r = ""
for i in range(0, 5) {
  r = r + i
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "01234" {
		t.Errorf("expected '01234', got %q", out)
	}
}

func TestForRangeSingleArg(t *testing.T) {
	src := `
let r = ""
for i in range(5) {
  r = r + i
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "01234" {
		t.Errorf("expected '01234' for single-arg range, got %q", out)
	}
}

func TestForRangeStep(t *testing.T) {
	src := `
let r = ""
for i in range(0, 10, 2) {
  r = r + i
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "02468" {
		t.Errorf("expected '02468', got %q", out)
	}
}

func TestForRangeAutoNegative(t *testing.T) {
	src := `
let r = ""
for i in range(5, 0) {
  r = r + i
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "54321" {
		t.Errorf("expected '54321', got %q", out)
	}
}

func TestForRangeNegativeStep(t *testing.T) {
	src := `
let r = ""
for i in range(10, 0, -3) {
  r = r + i
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "10741" {
		t.Errorf("expected '10741', got %q", out)
	}
}

func TestForRangeEmpty(t *testing.T) {
	src := `
let r = "empty"
for i in range(5, 5) {
  r = "not empty"
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "empty" {
		t.Errorf("expected 'empty', got %q", out)
	}
}

func TestForRangeBreak(t *testing.T) {
	src := `
let r = ""
for i in range(0, 100) {
  if (i == 3) { break }
  r = r + i
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "012" {
		t.Errorf("expected '012', got %q", out)
	}
}

func TestForRangeContinue(t *testing.T) {
	src := `
let r = ""
for i in range(0, 5) {
  if (i == 2) { continue }
  r = r + i
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "0134" {
		t.Errorf("expected '0134', got %q", out)
	}
}

func TestForRangeStepZeroError(t *testing.T) {
	src := `for i in range(0, 10, 0) { print(i) }`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for step=0")
	}
}

func TestForRangeNonIntegerError(t *testing.T) {
	src := `for i in range(2.5) { print(i) }`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for non-integer range argument")
	}
}

// ── enumerate ──

func TestForEnumerateBasic(t *testing.T) {
	src := `
let arr = ["a", "b", "c"]
let r = ""
for i, v in enumerate(arr) {
  r = r + i + v
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "0a1b2c" {
		t.Errorf("expected '0a1b2c', got %q", out)
	}
}

func TestForEnumerateEmpty(t *testing.T) {
	src := `
let arr = []
let r = "empty"
for i, v in enumerate(arr) {
  r = "not empty"
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "empty" {
		t.Errorf("expected 'empty', got %q", out)
	}
}

func TestForEnumerateBreak(t *testing.T) {
	src := `
let arr = [10, 20, 30, 40, 50]
let r = ""
for i, v in enumerate(arr) {
  if (i == 2) { break }
  r = r + v
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "1020" {
		t.Errorf("expected '1020', got %q", out)
	}
}

func TestForEnumerateNotArray(t *testing.T) {
	src := `for i, v in enumerate("hello") { print(i) }`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for non-array enumerate")
	}
}

// ── for..in (plain array iteration) ──

func TestForInArray(t *testing.T) {
	src := `
let arr = ["maca", "banana", "uva"]
let r = ""
for f in arr {
  r = r + f
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "macabananauva" {
		t.Errorf("expected 'macabananauva', got %q", out)
	}
}

func TestForInBreak(t *testing.T) {
	src := `
let arr = [1, 2, 3, 4, 5]
let r = ""
for x in arr {
  if (x == 3) { break }
  r = r + x
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "12" {
		t.Errorf("expected '12', got %q", out)
	}
}

func TestForInContinue(t *testing.T) {
	src := `
let arr = [1, 2, 3, 4, 5]
let r = ""
for x in arr {
  if (x == 3) { continue }
  r = r + x
}
print(r)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "1245" {
		t.Errorf("expected '1245', got %q", out)
	}
}

func TestForInNotArray(t *testing.T) {
	src := `for x in 42 { print(x) }`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for non-array for..in")
	}
}

func TestForRangeInFunction(t *testing.T) {
	src := `
fn soma(n) {
  let s = 0
  for i in range(0, n) {
    s = s + i
  }
  return s
}
print(soma(5))
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "10" {
		t.Errorf("expected '10', got %q", out)
	}
}

func TestForEnumerateReturn(t *testing.T) {
	src := `
fn findIndex(arr, target) {
  for i, v in enumerate(arr) {
    if (v == target) { return i }
  }
  return -1
}
print(findIndex(["a", "b", "c", "d"], "c"))
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out) != "2" {
		t.Errorf("expected '2', got %q", out)
	}
}
