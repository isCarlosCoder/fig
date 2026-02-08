package tests

import (
	"testing"

	"github.com/iscarloscoder/fig/environment"
)

func TestEnvDefineAndAssign(t *testing.T) {
	root := environment.NewEnv(nil)
	if err := root.Define("x", environment.NewNumber(1)); err != nil {
		t.Fatalf("Define failed: %v", err)
	}
	if err := root.Define("x", environment.NewNumber(2)); err == nil {
		t.Fatalf("expected error on redefine in same scope")
	}

	child := environment.NewEnv(root)
	if err := child.Assign("x", environment.NewNumber(42)); err != nil {
		t.Fatalf("Assign to parent var failed: %v", err)
	}
	v, ok := root.Get("x")
	if !ok {
		t.Fatalf("root x missing")
	}
	if num, _ := v.AsNumber(); num != 42 {
		t.Fatalf("expected 42 got %v", num)
	}

	if err := child.Assign("y", environment.NewNumber(3)); err == nil {
		t.Fatalf("expected error assigning non-existent var")
	}

	// shadowing: define in child with same name as parent should succeed
	if err := child.Define("x", environment.NewNumber(5)); err != nil {
		t.Fatalf("child Define shadow failed: %v", err)
	}
	vchild, ok := child.Get("x")
	if !ok {
		t.Fatalf("child x missing")
	}
	if num, _ := vchild.AsNumber(); num != 5 {
		t.Fatalf("expected child x 5 got %v", num)
	}
	// parent remains 42
	vparent, _ := root.Get("x")
	if num, _ := vparent.AsNumber(); num != 42 {
		t.Fatalf("expected parent x 42 got %v", num)
	}
}

func TestEnvHas(t *testing.T) {
	root := environment.NewEnv(nil)
	if root.Has("a") {
		t.Fatalf("expected Has false for undefined variable")
	}
	if err := root.Define("a", environment.NewNumber(10)); err != nil {
		t.Fatalf("Define failed: %v", err)
	}
	if !root.Has("a") {
		t.Fatalf("expected Has true for defined variable")
	}
	child := environment.NewEnv(root)
	if !child.Has("a") {
		t.Fatalf("child should see parent's variable")
	}
	// shadowing in child
	if err := child.Define("a", environment.NewNumber(20)); err != nil {
		t.Fatalf("child Define failed: %v", err)
	}
	if !child.Has("a") {
		t.Fatalf("child should still have its variable after shadowing")
	}
	if !root.Has("a") {
		t.Fatalf("root should still have its variable")
	}
	if child.Has("b") {
		t.Fatalf("Has should be false for missing variable")
	}
}
