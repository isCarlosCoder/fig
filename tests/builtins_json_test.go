package tests

import (
"strings"
"testing"
)

func useJson(code string) string {
return "use " + `"` + "json" + `"` + "; " + code
}

func TestJsonStringify(t *testing.T) {
src := useJson(`print(json.stringify({ a: 1 }));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if !strings.Contains(out, `"a"`) || !strings.Contains(out, "1") {
t.Fatalf("expected JSON object, got %q", out)
}
}

func TestJsonParse(t *testing.T) {
src := useJson(`let o = json.parse("{\"x\":42}"); print(o.x);`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "42" { t.Fatalf("expected '42', got %q", out) }
}

func TestJsonParseArray(t *testing.T) {
src := useJson(`let a = json.parse("[1,2,3]"); print(a);`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "[1, 2, 3]" { t.Fatalf("expected '[1, 2, 3]', got %q", out) }
}

func TestJsonParseNull(t *testing.T) {
src := useJson(`print(json.parse("null"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "null" { t.Fatalf("expected 'null', got %q", out) }
}

func TestJsonParseBool(t *testing.T) {
src := useJson(`print(json.parse("true")); print(json.parse("false"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
lines := strings.Split(out, "\n")
if lines[0] != "true" { t.Fatalf("expected 'true', got %q", lines[0]) }
if lines[1] != "false" { t.Fatalf("expected 'false', got %q", lines[1]) }
}

func TestJsonStringifyArray(t *testing.T) {
src := useJson(`print(json.stringify([1, 2, 3]));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "[1,2,3]" { t.Fatalf("expected '[1,2,3]', got %q", out) }
}

func TestJsonSerialize(t *testing.T) {
src := useJson(`let s = json.serialize([1]); print(s);`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if !strings.Contains(out, "1") { t.Fatalf("expected pretty JSON, got %q", out) }
}

func TestJsonDeserialize(t *testing.T) {
src := useJson(`let o = json.deserialize("{\"k\":99}"); print(o.k);`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "99" { t.Fatalf("expected '99', got %q", out) }
}

func TestJsonRoundTrip(t *testing.T) {
src := useJson(`let o = { a: 1, b: "two" }; let s = json.stringify(o); let p = json.parse(s); print(p.a); print(p.b);`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
lines := strings.Split(out, "\n")
if lines[0] != "1" { t.Fatalf("expected '1', got %q", lines[0]) }
if lines[1] != "two" { t.Fatalf("expected 'two', got %q", lines[1]) }
}

func TestJsonParseBadInput(t *testing.T) {
_, err := runFig(t, useJson(`json.parse("not json!!!");`))
if err == nil { t.Fatalf("expected error for invalid JSON") }
}

func TestJsonStringifyNested(t *testing.T) {
src := useJson(`print(json.stringify({ a: [1, 2] }));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if !strings.Contains(out, `"a"`) && !strings.Contains(out, "[1,2]") {
t.Fatalf("expected nested JSON, got %q", out)
}
}
