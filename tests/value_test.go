package tests

import (
	"testing"

	"github.com/iscarloscoder/fig/environment"
)

func TestValueHelpers(t *testing.T) {
	n := environment.NewNumber(3.14)
	if n.Type != environment.NumberType {
		t.Fatalf("expected number type")
	}
	if v, err := n.AsNumber(); err != nil || v != 3.14 {
		t.Fatalf("AsNumber failed: %v %v", v, err)
	}
	if !n.IsTruthy() {
		t.Fatalf("number should be truthy")
	}

	s := environment.NewString("hi")
	if _, err := s.AsNumber(); err == nil {
		t.Fatalf("expected AsNumber error for string")
	}
	if str, err := s.AsString(); err != nil || str != "hi" {
		t.Fatalf("AsString failed: %v", err)
	}

	b := environment.NewBool(false)
	if b.IsTruthy() {
		t.Fatalf("false should be falsey")
	}
	if _, err := b.AsBool(); err != nil {
		t.Fatalf("AsBool failed: %v", err)
	}

	nilv := environment.NewNil()
	if nilv.IsTruthy() {
		t.Fatalf("nil should be falsey")
	}
}
