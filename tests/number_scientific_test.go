package tests

import (
	"strings"
	"testing"
)

func TestScientificNotationInteger(t *testing.T) {
	out, err := runFig(t, `
		let x = 1e9
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "1000000000" {
		t.Fatalf("expected '1000000000', got %q", out)
	}
}

func TestScientificNotationFloat(t *testing.T) {
	out, err := runFig(t, `
		let x = 1.2e-3
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "0.0012" {
		t.Fatalf("expected '0.0012', got %q", out)
	}
}

func TestScientificNotationArithmetic(t *testing.T) {
	out, err := runFig(t, `
		let a = 1e3
		let b = 1e3
		print(a * b)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "1000000" {
		t.Fatalf("expected '1000000', got %q", out)
	}
}

func TestScientificNotationUppercaseE(t *testing.T) {
	out, err := runFig(t, `
		let x = 2E2
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "200" {
		t.Fatalf("expected '200', got %q", out)
	}
}

func TestScientificNotationWithPlusSign(t *testing.T) {
	out, err := runFig(t, `
		let x = 1e+3
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "1000" {
		t.Fatalf("expected '1000', got %q", out)
	}
}

func TestUnaryMinusScientific(t *testing.T) {
	out, err := runFig(t, `
		let x = -1e3
		print(x)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "-1000" {
		t.Fatalf("expected '-1000', got %q", out)
	}
}

func TestScientificInArrayAndObject(t *testing.T) {
	out, err := runFig(t, `
		let arr = [1e2, 2e2]
		let o = { v: 1e2 }
		print(arr)
		print(o.v)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) < 2 {
		t.Fatalf("unexpected output: %q", out)
	}
	if lines[0] != "[100, 200]" {
		t.Fatalf("expected '[100, 200]', got %q", lines[0])
	}
	if lines[1] != "100" {
		t.Fatalf("expected '100', got %q", lines[1])
	}
}

func TestScientificNotationComparisonAndAdd(t *testing.T) {
	out, err := runFig(t, `
		if (1e3 == 1000) { print("ok") }
		print(1e3 + 500)
	`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) < 2 {
		t.Fatalf("unexpected output: %q", out)
	}
	if lines[0] != "ok" {
		t.Fatalf("expected 'ok' in first line, got %q", lines[0])
	}
	if lines[1] != "1500" {
		t.Fatalf("expected '1500', got %q", lines[1])
	}
}
