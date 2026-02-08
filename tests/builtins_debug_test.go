package tests

import (
"strings"
"testing"
)

func useDebug(code string) string {
return "use " + `"` + "debug" + `"` + "; " + code
}

func TestDebugInspectNumber(t *testing.T) {
out, err := runFig(t, useDebug(`print(debug.inspect(42));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "<number: 42>" { t.Fatalf("expected '<number: 42>', got %q", out) }
}

func TestDebugInspectString(t *testing.T) {
out, err := runFig(t, useDebug(`print(debug.inspect("hi"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "<string: hi>" { t.Fatalf("expected '<string: hi>', got %q", out) }
}

func TestDebugInspectNull(t *testing.T) {
out, err := runFig(t, useDebug(`print(debug.inspect(null));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "<null: null>" { t.Fatalf("expected '<null: null>', got %q", out) }
}

func TestDebugDumpNumber(t *testing.T) {
out, err := runFig(t, useDebug(`print(debug.dump(42));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "number(42)" { t.Fatalf("expected 'number(42)', got %q", out) }
}

func TestDebugDumpString(t *testing.T) {
out, err := runFig(t, useDebug(`print(debug.dump("hi"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != `string("hi")` { t.Fatalf("expected 'string(\"hi\")', got %q", out) }
}

func TestDebugDumpArray(t *testing.T) {
out, err := runFig(t, useDebug(`print(debug.dump([1, 2]));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if !strings.Contains(out, "array(2)") { t.Fatalf("expected array dump, got %q", out) }
if !strings.Contains(out, "[0] number(1)") { t.Fatalf("expected element 0, got %q", out) }
}

func TestDebugDumpObject(t *testing.T) {
out, err := runFig(t, useDebug(`print(debug.dump({ a: 1 }));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if !strings.Contains(out, "object(1)") { t.Fatalf("expected object dump, got %q", out) }
if !strings.Contains(out, "a: number(1)") { t.Fatalf("expected key dump, got %q", out) }
}

func TestDebugType(t *testing.T) {
out, err := runFig(t, useDebug(`print(debug.type("x"));`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "string" { t.Fatalf("expected 'string', got %q", out) }
}

func TestDebugAssertPass(t *testing.T) {
out, err := runFig(t, useDebug(`debug.assert(1 == 1, "ok"); print("pass");`))
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "pass" { t.Fatalf("expected 'pass', got %q", out) }
}

func TestDebugAssertFail(t *testing.T) {
_, err := runFig(t, useDebug(`debug.assert(1 == 2, "math broken");`))
if err == nil { t.Fatalf("expected error for failed assert") }
}

func TestDebugPanic(t *testing.T) {
_, err := runFig(t, useDebug(`debug.panic("boom");`))
if err == nil { t.Fatalf("expected error from panic") }
}
