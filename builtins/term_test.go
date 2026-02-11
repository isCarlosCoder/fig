package builtins

import (
	"bytes"
	"testing"

	"github.com/iscarloscoder/fig/environment"
)

func TestAltScreenEnterExit(t *testing.T) {
	old := stdout
	var buf bytes.Buffer
	stdout = &buf
	defer func() { stdout = old }()

	mod := Get("term")
	if mod == nil {
		t.Fatal("term module not found in registry")
	}

	enter := mod.Entries["enterAltScreen"]
	exit := mod.Entries["exitAltScreen"]
	if enter.Type != environment.BuiltinFnType || exit.Type != environment.BuiltinFnType {
		t.Fatal("expected builtin functions")
	}

	// call enter
	_, err := enter.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("enterAltScreen failed: %v", err)
	}
	if buf.String() != "\x1b[?1049h" {
		t.Fatalf("unexpected enter sequence: %q", buf.String())
	}
	buf.Reset()

	// call exit
	_, err = exit.Builtin([]environment.Value{})
	if err != nil {
		t.Fatalf("exitAltScreen failed: %v", err)
	}
	if buf.String() != "\x1b[?1049l" {
		t.Fatalf("unexpected exit sequence: %q", buf.String())
	}
}
