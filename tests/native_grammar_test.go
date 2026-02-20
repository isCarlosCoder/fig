package tests

import "testing"

func TestNativeGrammarLevelDeclaration(t *testing.T) {
	out, err := runFig(t, `@native fn triple(x) { return x * 3 } print(triple(4));`)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "12" {
		t.Fatalf("expected '12', got %q", out)
	}
}

func TestNativeGrammarWithAttributeString(t *testing.T) {
	out, err := runFig(t, `@native("fallback=true") fn maybe(x) { return 2 * x } print(maybe(5));`)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10" {
		t.Fatalf("expected '10', got %q", out)
	}
}
