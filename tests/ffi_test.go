package tests

import (
	"testing"

	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
)

func TestFfiEnabledDefaultFalse(t *testing.T) {
	mod := builtins.Get("ffi")
	if mod == nil {
		t.Fatal("ffi module not found")
	}
	en := mod.Entries["enabled"]
	if en.Type != environment.BuiltinFnType {
		t.Fatalf("expected builtin, got %v", en.Type)
	}
	v, err := en.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("enabled() returned error: %v", err)
	}
	if v.Type != environment.BooleanType || v.Bool != false {
		t.Fatalf("expected enabled=false by default, got %v", v)
	}
}
