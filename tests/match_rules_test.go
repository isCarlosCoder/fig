package tests

import "testing"

func TestMatchDuplicatePatternsIsError(t *testing.T) {
	src := `match 1 { 1 => { print("a") } 1 => { print("b") } _ => { print("ok") }}`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for duplicate patterns")
	}
}

func TestMatchMissingWildcardIsError(t *testing.T) {
	src := `match 1 { 1 => { print("a") } 2 => { print("b") } }`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for missing wildcard '_'")
	}
}

func TestMatchWildcardMustBeAlone(t *testing.T) {
	src := `match 1 { _ , 1 => { print("bad") } }`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for wildcard used with other patterns")
	}
}
