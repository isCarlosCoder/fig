package tests

import "testing"

func useMath(code string) string {
	return "use " + `"` + "strings" + `"` + "; " + code
}

func TestStringsLen(t *testing.T) {
	out, err := runFig(t, useMath("print(strings.len("+`"FigLang"`+"));"))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "7" {
		t.Fatalf("expected '7', got %q", out)
	}
}

func TestStringsUpper(t *testing.T) {
	out, err := runFig(t, useMath("print(strings.upper("+`"hello"`+"));"))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "HELLO" {
		t.Fatalf("expected 'HELLO', got %q", out)
	}
}

func TestStringsLower(t *testing.T) {
	out, err := runFig(t, useMath("print(strings.lower("+`"WORLD"`+"));"))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "world" {
		t.Fatalf("expected 'world', got %q", out)
	}
}

func TestStringsTrim(t *testing.T) {
	out, err := runFig(t, useMath("print(strings.trim("+`"  fig  "`+"));"))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "fig" {
		t.Fatalf("expected 'fig', got %q", out)
	}
}

func TestStringsSplit(t *testing.T) {
	src := useMath(`let p = strings.split("a,b,c", ","); print(p);`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := `["a", "b", "c"]`
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestStringsJoin(t *testing.T) {
	src := useMath(`let p = strings.split("a,b,c", ","); print(strings.join(p, " - "));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "a - b - c" {
		t.Fatalf("expected 'a - b - c', got %q", out)
	}
}

func TestStringsReplace(t *testing.T) {
	src := useMath(`print(strings.replace("hello world", "world", "Fig"));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "hello Fig" {
		t.Fatalf("expected 'hello Fig', got %q", out)
	}
}

func TestStringsContains(t *testing.T) {
	src := useMath(`print(strings.contains("FigLang", "Lang")); print(strings.contains("FigLang", "xyz"));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true\nfalse" {
		t.Fatalf("expected 'true\\nfalse', got %q", out)
	}
}

func TestStringsStartsWith(t *testing.T) {
	src := useMath(`print(strings.startsWith("FigLang", "Fig")); print(strings.startsWith("FigLang", "Lang"));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true\nfalse" {
		t.Fatalf("expected 'true\\nfalse', got %q", out)
	}
}

func TestStringsEndsWith(t *testing.T) {
	src := useMath(`print(strings.endsWith("FigLang", "Lang")); print(strings.endsWith("FigLang", "Fig"));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true\nfalse" {
		t.Fatalf("expected 'true\\nfalse', got %q", out)
	}
}

func TestStringsIndexOf(t *testing.T) {
	src := useMath(`print(strings.indexOf("abcabc", "bc")); print(strings.indexOf("abc", "z"));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "1\n-1" {
		t.Fatalf("expected '1\\n-1', got %q", out)
	}
}

func TestStringsLastIndexOf(t *testing.T) {
	src := useMath(`print(strings.lastIndexOf("abcabc", "bc"));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "4" {
		t.Fatalf("expected '4', got %q", out)
	}
}

func TestStringsSubstring(t *testing.T) {
	src := useMath(`print(strings.substring("FigLang", 0, 3)); print(strings.substring("FigLang", 3, 7));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "Fig\nLang" {
		t.Fatalf("expected 'Fig\\nLang', got %q", out)
	}
}

func TestStringsCharAt(t *testing.T) {
	src := useMath(`print(strings.charAt("Fig", 0)); print(strings.charAt("Fig", 2));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "F\ng" {
		t.Fatalf("expected 'F\\ng', got %q", out)
	}
}

func TestStringsRepeat(t *testing.T) {
	src := useMath(`print(strings.repeat("ha", 3));`)
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "hahaha" {
		t.Fatalf("expected 'hahaha', got %q", out)
	}
}

func TestStringsCharAtOutOfBounds(t *testing.T) {
	src := useMath(`strings.charAt("hi", 5);`)
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for charAt out of bounds")
	}
}

func TestStringsWrongType(t *testing.T) {
	src := useMath("strings.upper(42);")
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for wrong arg type")
	}
}
