package tests

import "testing"

func useRegex(code string) string {
return "use " + `"` + "regex" + `"` + "; " + code
}

func TestRegexMatch(t *testing.T) {
src := useRegex(`print(regex.match("hello123", "[0-9]+"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "true" { t.Fatalf("expected 'true', got %q", out) }
}

func TestRegexMatchFalse(t *testing.T) {
src := useRegex(`print(regex.match("hello", "[0-9]+"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "false" { t.Fatalf("expected 'false', got %q", out) }
}

func TestRegexFindAll(t *testing.T) {
src := useRegex(`print(regex.findAll("abc123def456", "[0-9]+"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
expected := `["123", "456"]`
if out != expected { t.Fatalf("expected %q, got %q", expected, out) }
}

func TestRegexFindAllNoMatch(t *testing.T) {
src := useRegex(`print(regex.findAll("hello", "[0-9]+"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "[]" { t.Fatalf("expected '[]', got %q", out) }
}

func TestRegexReplaceRegex(t *testing.T) {
src := useRegex(`print(regex.replaceRegex("foo123bar456", "[0-9]+", "#"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "foo#bar#" { t.Fatalf("expected 'foo#bar#', got %q", out) }
}

func TestRegexSplitRegex(t *testing.T) {
src := useRegex(`print(regex.splitRegex("one:two::three", ":+"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
expected := `["one", "two", "three"]`
if out != expected { t.Fatalf("expected %q, got %q", expected, out) }
}

func TestRegexBadPattern(t *testing.T) {
_, err := runFig(t, useRegex(`regex.match("hi", "[invalid");`))
if err == nil { t.Fatalf("expected error for invalid regex pattern") }
}

func TestRegexWrongType(t *testing.T) {
_, err := runFig(t, useRegex(`regex.match(123, "abc");`))
if err == nil { t.Fatalf("expected error for non-string argument") }
}

func TestRegexReplaceAll(t *testing.T) {
src := useRegex(`print(regex.replaceRegex("a1b2c3", "[a-z]", "X"));`)
out, err := runFig(t, src)
if err != nil { t.Fatalf("runtime error: %v", err) }
if out != "X1X2X3" { t.Fatalf("expected 'X1X2X3', got %q", out) }
}
