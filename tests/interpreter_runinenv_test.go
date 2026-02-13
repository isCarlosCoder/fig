package tests

import (
	"bytes"
	"testing"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
)

func TestRunInEnvPersistsTopLevelDefs(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer
	env := environment.NewEnv(nil)
	src := `let x = 42
fn f() { return x }
print("loaded")`
	if err := interpreter.RunInEnv(src, "<test>", env, &out, &errOut); err != nil {
		t.Fatalf("RunInEnv error: %v, stderr=%s", err, errOut.String())
	}
	// now evaluate expressions in a new Run using the same env — should see x and f
	out.Reset()
	if err := interpreter.Run("print(x); print(f());", "<test>", env, &out, &errOut); err != nil {
		t.Fatalf("Run after preload error: %v, stderr=%s", err, errOut.String())
	}
	if out.String() == "" || !(bytes.Contains(out.Bytes(), []byte("42")) && bytes.Contains(out.Bytes(), []byte("42"))) {
		t.Fatalf("expected '42' from env, got: %q", out.String())
	}
}
