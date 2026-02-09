package tests

import (
	"strings"
	"testing"
)

func useObjects(code string) string {
	return "use " + `"` + "objects" + `"` + "; " + code
}

func TestObjectsKeys(t *testing.T) {
	src := useObjects(`let o = { a: 1, b: 2 }; print(objects.keys(o));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != `["a", "b"]` {
		t.Fatalf("expected keys, got %q", out)
	}
}

func TestObjectsValues(t *testing.T) {
	src := useObjects(`let o = { x: 10, y: 20 }; print(objects.values(o));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "[10, 20]" {
		t.Fatalf("expected values, got %q", out)
	}
}

func TestObjectsEntries(t *testing.T) {
	src := useObjects(`let o = { a: 1 }; print(objects.entries(o));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := `[["a", 1]]`
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestObjectsHasKey(t *testing.T) {
	src := useObjects(`let o = { nome: "Fig" }; print(objects.hasKey(o, "nome")); print(objects.hasKey(o, "xyz"));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "true" {
		t.Fatalf("expected 'true', got %q", lines[0])
	}
	if lines[1] != "false" {
		t.Fatalf("expected 'false', got %q", lines[1])
	}
}

func TestObjectsDeleteKey(t *testing.T) {
	src := useObjects(`let o = { a: 1, b: 2 }; let v = objects.deleteKey(o, "a"); print(v); print(objects.size(o));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "1" {
		t.Fatalf("expected '1', got %q", lines[0])
	}
	if lines[1] != "1" {
		t.Fatalf("expected '1', got %q", lines[1])
	}
}

func TestObjectsDeleteKeyMissing(t *testing.T) {
	src := useObjects(`let o = { a: 1 }; let v = objects.deleteKey(o, "z"); print(v);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "null" {
		t.Fatalf("expected 'null', got %q", out)
	}
}

func TestObjectsMerge(t *testing.T) {
	src := useObjects(`let a = { x: 1 }; let b = { x: 99, y: 2 }; let m = objects.merge(a, b); print(m);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "{x: 99, y: 2}" {
		t.Fatalf("expected merged object, got %q", out)
	}
}

func TestObjectsClone(t *testing.T) {
	src := useObjects(`let o = { a: 1 }; let c = objects.clone(o); c.a = 999; print(o); print(c);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "{a: 1}" {
		t.Fatalf("expected original unchanged, got %q", lines[0])
	}
	if lines[1] != "{a: 999}" {
		t.Fatalf("expected clone changed, got %q", lines[1])
	}
}

func TestObjectsSize(t *testing.T) {
	src := useObjects(`let o = { a: 1, b: 2, c: 3 }; print(objects.size(o));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3" {
		t.Fatalf("expected '3', got %q", out)
	}
}

func TestObjectsClear(t *testing.T) {
	src := useObjects(`let o = { a: 1, b: 2 }; objects.clear(o); print(objects.size(o)); print(o);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "0" {
		t.Fatalf("expected '0', got %q", lines[0])
	}
	if lines[1] != "{}" {
		t.Fatalf("expected '{}', got %q", lines[1])
	}
}

func TestObjectsWrongType(t *testing.T) {
	src := useObjects(`objects.keys("not_object");`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for non-object argument")
	}
}

func TestObjectsEntriesMultiple(t *testing.T) {
	src := useObjects(`let o = { x: 1, y: 2 }; let e = objects.entries(o); print(e);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := `[["x", 1], ["y", 2]]`
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}
