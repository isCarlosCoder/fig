package tests

import (
	"strings"
	"testing"
)

func useTypes(code string) string {
	return "use " + `"` + "types" + `"` + "; " + code
}

func TestTypesType(t *testing.T) {
	src := useTypes(`print(types.type(42)); print(types.type("hi")); print(types.type(true)); print(types.type([1])); print(types.type({ a: 1 })); print(types.type(null));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	expected := []string{"number", "string", "boolean", "array", "object", "null"}
	for i, e := range expected {
		if lines[i] != e {
			t.Fatalf("line %d: expected %q, got %q", i, e, lines[i])
		}
	}
}

func TestTypesIsNumber(t *testing.T) {
	src := useTypes(`print(types.isNumber(1)); print(types.isNumber("x"));`)
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

func TestTypesIsString(t *testing.T) {
	src := useTypes(`print(types.isString("ok")); print(types.isString(0));`)
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

func TestTypesIsBool(t *testing.T) {
	src := useTypes(`print(types.isBool(true)); print(types.isBool(1));`)
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

func TestTypesIsArray(t *testing.T) {
	src := useTypes(`print(types.isArray([1, 2])); print(types.isArray("no"));`)
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

func TestTypesIsObject(t *testing.T) {
	src := useTypes(`print(types.isObject({ a: 1 })); print(types.isObject(42));`)
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

func TestTypesIsNil(t *testing.T) {
	src := useTypes(`print(types.isNil(null)); print(types.isNil(0));`)
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

func TestTypesToInt(t *testing.T) {
	src := useTypes(`print(types.toInt(3.7)); print(types.toInt("42")); print(types.toInt(true));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "3" {
		t.Fatalf("expected '3', got %q", lines[0])
	}
	if lines[1] != "42" {
		t.Fatalf("expected '42', got %q", lines[1])
	}
	if lines[2] != "1" {
		t.Fatalf("expected '1', got %q", lines[2])
	}
}

func TestTypesToFloat(t *testing.T) {
	src := useTypes(`print(types.toFloat("3.14")); print(types.toFloat(false));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "3.14" {
		t.Fatalf("expected '3.14', got %q", lines[0])
	}
	if lines[1] != "0" {
		t.Fatalf("expected '0', got %q", lines[1])
	}
}

func TestTypesToString(t *testing.T) {
	src := useTypes(`print(types.toString(123)); print(types.toString(true)); print(types.toString(null));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "123" {
		t.Fatalf("expected '123', got %q", lines[0])
	}
	if lines[1] != "true" {
		t.Fatalf("expected 'true', got %q", lines[1])
	}
	if lines[2] != "null" {
		t.Fatalf("expected 'null', got %q", lines[2])
	}
}

func TestTypesToBool(t *testing.T) {
	src := useTypes(`print(types.toBool(0)); print(types.toBool(1)); print(types.toBool("")); print(types.toBool("fig")); print(types.toBool(null));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	expected := []string{"false", "true", "false", "true", "false"}
	for i, e := range expected {
		if lines[i] != e {
			t.Fatalf("line %d: expected %q, got %q", i, e, lines[i])
		}
	}
}

func TestTypesToIntBadString(t *testing.T) {
	src := useTypes(`types.toInt("abc");`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for toInt with non-numeric string")
	}
}

func TestTypesToFloatBadString(t *testing.T) {
	src := useTypes(`types.toFloat("xyz");`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatalf("expected error for toFloat with non-numeric string")
	}
}

func TestTypesIsFunction(t *testing.T) {
	src := useTypes(`fn soma(a, b) { return a + b } print(types.isFunction(soma)); print(types.type(soma));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if lines[0] != "true" {
		t.Fatalf("expected 'true', got %q", lines[0])
	}
	if lines[1] != "function" {
		t.Fatalf("expected 'function', got %q", lines[1])
	}
}
