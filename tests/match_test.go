package tests

import "testing"

// ── Basic match statement ──

func TestMatchBasicNumber(t *testing.T) {
	src := `
let x = 2
match x {
	1 => { print("one") }
	2 => { print("two") }
	3 => { print("three") }
	_ => { }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "two" {
		t.Errorf("expected 'two', got %q", out)
	}
}

func TestMatchBasicString(t *testing.T) {
	src := `
let lang = "fig"
match lang {
	"go" => { print("Go") }
	"fig" => { print("Fig!") }
	_ => { print("other") }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "Fig!" {
		t.Errorf("expected 'Fig!', got %q", out)
	}
}

// ── Wildcard / default arm ──

func TestMatchWildcard(t *testing.T) {
	src := `
let x = 99
match x {
	1 => { print("one") }
	2 => { print("two") }
	_ => { print("default") }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "default" {
		t.Errorf("expected 'default', got %q", out)
	}
}

// ── No arm matches, no wildcard → nil ──

func TestMatchNoMatch(t *testing.T) {
	src := `
let x = 42
match x {
	1 => { print("one") }
	2 => { print("two") }
	_ => { }
}
print("done")
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "done" {
		t.Errorf("expected 'done', got %q", out)
	}
}

// ── Match as expression (inline value) ──

func TestMatchExpressionInline(t *testing.T) {
	src := `
let x = 1
let result = match x {
	1 => "um"
	2 => "dois"
	_ => "outro"
}
print(result)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "um" {
		t.Errorf("expected 'um', got %q", out)
	}
}

func TestMatchExpressionInlineDefault(t *testing.T) {
	src := `
let x = 100
let result = match x {
	1 => "um"
	2 => "dois"
	_ => "outro"
}
print(result)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "outro" {
		t.Errorf("expected 'outro', got %q", out)
	}
}

// ── Match with multiple values per arm ──

func TestMatchMultipleValues(t *testing.T) {
	src := `
let x = 3
match x {
	1, 3, 5 => { print("odd") }
	2, 4, 6 => { print("even") }
	_ => { print("other") }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "odd" {
		t.Errorf("expected 'odd', got %q", out)
	}
}

func TestMatchMultipleValuesSecond(t *testing.T) {
	src := `
let x = 4
match x {
	1, 3, 5 => { print("odd") }
	2, 4, 6 => { print("even") }
	_ => { print("other") }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "even" {
		t.Errorf("expected 'even', got %q", out)
	}
}

// ── Match with boolean values ──

func TestMatchBoolean(t *testing.T) {
	src := `
let flag = true
let result = match flag {
	true => "yes"
	false => "no"
	_ => "unknown"
}
print(result)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "yes" {
		t.Errorf("expected 'yes', got %q", out)
	}
}

// ── Match with expression as subject ──

func TestMatchExpressionSubject(t *testing.T) {
	src := `
let a = 3
let b = 2
match a + b {
	4 => { print("four") }
	5 => { print("five") }
	6 => { print("six") }
	_ => { }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "five" {
		t.Errorf("expected 'five', got %q", out)
	}
}

// ── Match with block body containing multiple statements ──

func TestMatchBlockMultipleStatements(t *testing.T) {
	src := `
let x = 2
match x {
	1 => {
		let msg = "one"
		print(msg)
	}
	2 => {
		let a = "tw"
		let b = "o"
		print(a + b)
	}
	_ => { }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "two" {
		t.Errorf("expected 'two', got %q", out)
	}
}

// ── Match with return inside a function ──

func TestMatchReturnInsideFunction(t *testing.T) {
	src := `
fn classify(n) {
	return match n {
		1 => "one"
		2 => "two"
		_ => "many"
	}
}
print(classify(1))
print(classify(2))
print(classify(99))
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "one\ntwo\nmany" {
		t.Errorf("expected 'one\\ntwo\\nmany', got %q", out)
	}
}

// ── Match with null ──

func TestMatchNull(t *testing.T) {
	src := `
let x = null
let result = match x {
	null => "nulo"
	_ => "valor"
}
print(result)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "nulo" {
		t.Errorf("expected 'nulo', got %q", out)
	}
}

// ── Match with type checking (using expression) ──

func TestMatchTypeCheck(t *testing.T) {
	src := `
use "types"
let x = 42
match types.type(x) {
	"number" => { print("is number") }
	"string" => { print("is string") }
	_ => { print("other type") }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "is number" {
		t.Errorf("expected 'is number', got %q", out)
	}
}

// ── Match with expressions in patterns ──

func TestMatchExpressionPatterns(t *testing.T) {
	src := `
let x = 10
let target = 5 * 2
match x {
	5 * 2 => { print("ten") }
	_ => { print("other") }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "ten" {
		t.Errorf("expected 'ten', got %q", out)
	}
}

// ── Match first arm wins ──

func TestMatchFirstArmWins(t *testing.T) {
	src := `
let x = 1
match x {
	1 => { print("first") }
	2 => { print("second") }
	_ => { }
}
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "first" {
		t.Errorf("expected 'first', got %q", out)
	}
}

// ── Match no match returns nil for expression ──

func TestMatchNoMatchReturnsNull(t *testing.T) {
	src := `
let x = 99
let result = match x {
	1 => "one"
	2 => "two"
	_ => null
}
print(result)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "null" {
		t.Errorf("expected 'null', got %q", out)
	}
}

// ── Match with multiple values inline expression ──

func TestMatchMultipleValuesInline(t *testing.T) {
	src := `
let x = 5
let result = match x {
	1, 3, 5 => "odd"
	2, 4, 6 => "even"
	_ => "big"
}
print(result)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "odd" {
		t.Errorf("expected 'odd', got %q", out)
	}
}

// ── Match with string concatenation as subject ──

func TestMatchStringConcat(t *testing.T) {
	src := `
let greeting = "hello" + " " + "world"
let result = match greeting {
	"hello world" => "matched!"
	_ => "nope"
}
print(result)
`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "matched!" {
		t.Errorf("expected 'matched!', got %q", out)
	}
}
