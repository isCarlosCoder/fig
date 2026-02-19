package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
)

func TestNativeSimpleAndBracketSyntax(t *testing.T) {
	src := `@native fn add(x, y) { return x + y }
@[native()] fn dbl(x) { return x * 2 }
print(add(2,3))
print(dbl(5))`
	var buf bytes.Buffer
	if err := interpreter.Run(src, "<test>", environment.NewEnv(nil), &buf, nil); err != nil {
		t.Fatalf("Run error: %v", err)
	}
	out := strings.TrimSpace(buf.String())
	if out != "5\n10" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestNativeSigmoid(t *testing.T) {
	src := `@native fn sigmoid(x) { return 1 / (1 + math.exp(-x)) }
print(sigmoid(0))`
	var buf bytes.Buffer
	if err := interpreter.Run(src, "<test>", environment.NewEnv(nil), &buf, nil); err != nil {
		t.Fatalf("Run error: %v", err)
	}
	out := strings.TrimSpace(buf.String())
	if out != "0.5" {
		t.Fatalf("expected 0.5 got %q", out)
	}
}

func TestNativeValidationRejectsUnsupportedBody(t *testing.T) {
	src := `@native fn bad(x) { let y = x; return y }`
	var buf bytes.Buffer
	err := interpreter.Run(src, "<test>", environment.NewEnv(nil), &buf, nil)
	if err == nil {
		t.Fatalf("expected error defining native function with unsupported body")
	}
	if !strings.Contains(err.Error(), "native function must contain exactly one statement") {
		t.Fatalf("unexpected error: %v", err)
	}
}
