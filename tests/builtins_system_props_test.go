package tests

import (
	"testing"

	"github.com/iscarloscoder/fig/builtins"
)

func TestSystemArgvAndCwdProperties(t *testing.T) {
	// set and restore runtime state via builtins globals
	prevArgs := builtins.ScriptArgs
	prevCwd := builtins.ScriptCwd
	defer func() {
		builtins.ScriptArgs = prevArgs
		builtins.ScriptCwd = prevCwd
	}()

	// set values via the runtime state
	builtins.ScriptArgs = []string{"x", "y"}
	builtins.ScriptCwd = "/tmp"

	out, err := runFig(t, `use "system"; print(system.argv()); print(system.cwd());`)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[\"x\", \"y\"]\n/tmp" {
		t.Fatalf("unexpected output: %q", out)
	}
}
