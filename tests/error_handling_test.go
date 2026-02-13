package tests

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripAnsi(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

// --- Unary errors ---

func TestUnaryMinusOnString(t *testing.T) {
	_, err := runFig(t, `let x = -"hello"`)
	if err == nil {
		t.Fatal("expected runtime error for -\"hello\"")
	}
	if !strings.Contains(err.Error(), "cannot negate") {
		t.Fatalf("expected 'cannot negate' in error, got: %v", err)
	}
}

func TestUnaryMinusOnBool(t *testing.T) {
	_, err := runFig(t, `let x = -true`)
	if err == nil {
		t.Fatal("expected runtime error for -true")
	}
	if !strings.Contains(err.Error(), "cannot negate") {
		t.Fatalf("expected 'cannot negate' in error, got: %v", err)
	}
}

func TestUnaryMinusOnNull(t *testing.T) {
	_, err := runFig(t, `let x = -null`)
	if err == nil {
		t.Fatal("expected runtime error for -null")
	}
	if !strings.Contains(err.Error(), "cannot negate") {
		t.Fatalf("expected 'cannot negate' in error, got: %v", err)
	}
}

func TestUnaryNotOnValues(t *testing.T) {
	// !value should always work (returns bool)
	// In Fig: 0 and "" are truthy (only null, false, empty array, empty object are falsy)
	out, err := runFig(t, `print(!true); print(!false); print(!null); print(!0); print(!1); print(!"")`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "false\ntrue\ntrue\nfalse\nfalse\nfalse"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

// --- Arithmetic errors ---

func TestAddNonNumbers(t *testing.T) {
	// string + anything should concatenate (not error)
	out, err := runFig(t, `print("hello" + 42)`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "hello42" {
		t.Fatalf("expected 'hello42', got %q", out)
	}
}

func TestSubtractStrings(t *testing.T) {
	_, err := runFig(t, `let x = "a" - "b"`)
	if err == nil {
		t.Fatal("expected runtime error for string subtraction")
	}
}

func TestMultiplyStrings(t *testing.T) {
	_, err := runFig(t, `let x = "a" * "b"`)
	if err == nil {
		t.Fatal("expected runtime error for string multiplication")
	}
}

func TestDivisionByZero(t *testing.T) {
	_, err := runFig(t, `let x = 10 / 0`)
	if err == nil {
		t.Fatal("expected runtime error for division by zero")
	}
	if !strings.Contains(err.Error(), "division by zero") {
		t.Fatalf("expected 'division by zero' in error, got: %v", err)
	}
}

func TestModuloByZero(t *testing.T) {
	_, err := runFig(t, `let x = 10 % 0`)
	if err == nil {
		t.Fatal("expected runtime error for modulo by zero")
	}
	if !strings.Contains(err.Error(), "modulo by zero") {
		t.Fatalf("expected 'modulo by zero' in error, got: %v", err)
	}
}

// --- Comparison errors ---

func TestCompareStringAndNumber(t *testing.T) {
	_, err := runFig(t, `let x = "hello" > 5`)
	if err == nil {
		t.Fatal("expected runtime error for comparing string > number")
	}
}

// --- Dot access errors ---

func TestDotAccessOnNumber(t *testing.T) {
	_, err := runFig(t, `let x = 42; let y = x.foo`)
	if err == nil {
		t.Fatal("expected runtime error for dot access on number")
	}
	if !strings.Contains(err.Error(), "cannot access property") {
		t.Fatalf("expected 'cannot access property' in error, got: %v", err)
	}
}

func TestDotAccessOnString(t *testing.T) {
	_, err := runFig(t, `let x = "hello"; let y = x.foo`)
	if err == nil {
		t.Fatal("expected runtime error for dot access on string")
	}
	if !strings.Contains(err.Error(), "cannot access property") {
		t.Fatalf("expected 'cannot access property' in error, got: %v", err)
	}
}

func TestDotAccessOnBool(t *testing.T) {
	_, err := runFig(t, `let x = true; let y = x.foo`)
	if err == nil {
		t.Fatal("expected runtime error for dot access on bool")
	}
	if !strings.Contains(err.Error(), "cannot access property") {
		t.Fatalf("expected 'cannot access property' in error, got: %v", err)
	}
}

func TestDotAccessOnNull(t *testing.T) {
	_, err := runFig(t, `let x = null; let y = x.foo`)
	if err == nil {
		t.Fatal("expected runtime error for dot access on null")
	}
	if !strings.Contains(err.Error(), "cannot access property") {
		t.Fatalf("expected 'cannot access property' in error, got: %v", err)
	}
}

// --- Index access errors ---

func TestIndexOnNonIndexable(t *testing.T) {
	_, err := runFig(t, `let x = 42; let y = x[0]`)
	if err == nil {
		t.Fatal("expected runtime error for indexing a number")
	}
	if !strings.Contains(err.Error(), "cannot index") {
		t.Fatalf("expected 'cannot index' in error, got: %v", err)
	}
}

func TestArrayIndexOutOfRange(t *testing.T) {
	_, err := runFig(t, `let arr = [1, 2, 3]; let y = arr[10]`)
	if err == nil {
		t.Fatal("expected runtime error for out-of-range index")
	}
	if !strings.Contains(err.Error(), "out of range") {
		t.Fatalf("expected 'out of range' in error, got: %v", err)
	}
}

func TestArrayNegativeIndex(t *testing.T) {
	out, err := runFig(t, `let arr = [1, 2, 3]; print(arr[-1]); print(arr[-3])`)
	if err != nil {
		t.Fatalf("unexpected runtime error: %v", err)
	}
	expected := "3\n1"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}

	// out-of-range negative index should still error
	_, err = runFig(t, `let arr = [1,2,3]; let y = arr[-4]`)
	if err == nil {
		t.Fatal("expected runtime error for out-of-range negative index")
	}
	if !strings.Contains(err.Error(), "out of range") {
		t.Fatalf("expected 'out of range' in error, got: %v", err)
	}
}

func TestArrayNonNumberIndex(t *testing.T) {
	_, err := runFig(t, `let arr = [1, 2, 3]; let y = arr["foo"]`)
	if err == nil {
		t.Fatal("expected runtime error for non-number index")
	}
	if !strings.Contains(err.Error(), "index must be a number") {
		t.Fatalf("expected 'index must be a number' in error, got: %v", err)
	}
}

// --- Error propagation ---

func TestErrorInNestedExpression(t *testing.T) {
	// error in rhs of addition should propagate
	_, err := runFig(t, `let x = 1 + (2 * "abc")`)
	if err == nil {
		t.Fatal("expected runtime error for nested expression error")
	}
}

func TestErrorInFunctionArg(t *testing.T) {
	// error in argument evaluation should propagate
	_, err := runFig(t, `
		fn add(a, b) { return a + b }
		add(1, -"hello")
	`)
	if err == nil {
		t.Fatal("expected runtime error for error in function arg")
	}
}

func TestErrorInIfCondition(t *testing.T) {
	// error in condition evaluation should propagate
	_, err := runFig(t, `
		if (-"bad") {
			print("yes")
		}
	`)
	if err == nil {
		t.Fatal("expected runtime error in if condition")
	}
}

func TestErrorInWhileCondition(t *testing.T) {
	_, err := runFig(t, `
		while (-"bad") {
			print("loop")
		}
	`)
	if err == nil {
		t.Fatal("expected runtime error in while condition")
	}
}

// --- Error does not crash the process ---

func TestErrorDoesNotPanic(t *testing.T) {
	// These used to panic — now they should produce clean RuntimeErrors
	cases := []string{
		`-"hello"`,
		`-true`,
		`-null`,
		`-[1, 2]`,
	}
	for _, src := range cases {
		_, err := runFig(t, src)
		if err == nil {
			t.Errorf("expected error for %q, got nil", src)
		}
	}
}

// --- RuntimeError formatting ---

func TestRuntimeErrorHasLineInfo(t *testing.T) {
	_, err := runFig(t, `let x = 10 / 0`)
	if err == nil {
		t.Fatal("expected error")
	}
	errStr := err.Error()
	// should contain line:col format
	if !strings.Contains(errStr, ":") {
		t.Fatalf("expected line:col in error, got: %q", errStr)
	}
	if !strings.Contains(errStr, "runtime error") {
		t.Fatalf("expected 'runtime error' in error, got: %q", errStr)
	}
}

func TestRuntimeErrorHasSnippet(t *testing.T) {
	// Use NewFigVisitorWithSource to get snippets
	src := `let x = 10 / 0`
	var buf bytes.Buffer
	var errBuf bytes.Buffer
	err := interpreter.Run(src, "<test>", environment.NewEnv(nil), &buf, &errBuf)
	if err == nil {
		t.Fatal("expected error")
	}
	// errBuf should contain the pretty error with snippet
	errStr := errBuf.String()
	// Strip ANSI escape codes for content checking
	stripped := stripAnsi(errStr)
	if !strings.Contains(stripped, "10 / 0") {
		t.Fatalf("expected snippet with '10 / 0' in error output, got: %q", stripped)
	}
	if !strings.Contains(stripped, "^") {
		t.Fatalf("expected caret marker in error output, got: %q", stripped)
	}
}

func TestPrettyErrorHasColors(t *testing.T) {
	re := &interpreter.RuntimeError{
		Line:        1,
		Column:      8,
		Message:     "division by zero",
		Snippet:     "let x = 10 / 0",
		ColumnStart: 8,
		Length:      1,
	}
	pretty := re.PrettyError()
	// Should contain ANSI red
	if !strings.Contains(pretty, "\x1b[1;31m") {
		t.Fatalf("expected ANSI red in PrettyError(), got: %q", pretty)
	}
	// Should contain reset
	if !strings.Contains(pretty, "\x1b[0m") {
		t.Fatalf("expected ANSI reset in PrettyError(), got: %q", pretty)
	}
}

// --- Recover safety net ---

func TestRecoverFromInternalError(t *testing.T) {
	// interpreter.Run should not crash the process even on unexpected errors.
	// We test the recover by running through Run() which has the safety net.
	var out bytes.Buffer
	var errOut bytes.Buffer
	// This is a valid program so should not error
	err := interpreter.Run(`print("hello")`, "<test>", environment.NewEnv(nil), &out, &errOut)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.TrimSpace(out.String()) != "hello" {
		t.Fatalf("expected 'hello', got %q", out.String())
	}
}

func TestRunReportsRuntimeError(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer
	err := interpreter.Run(`let x = 10 / 0`, "<test>", environment.NewEnv(nil), &out, &errOut)
	if err == nil {
		t.Fatal("expected runtime error from Run()")
	}
	// errOut should contain pretty colored output
	errStr := errOut.String()
	if !strings.Contains(errStr, "division by zero") {
		t.Fatalf("expected 'division by zero' in errOut, got: %q", errStr)
	}
}

// --- Error short-circuit in logical expressions ---

func TestLogicalAndShortCircuitOnError(t *testing.T) {
	// false && error_expr should short-circuit, no error
	out, err := runFig(t, `print(false && -"bad")`)
	if err != nil {
		t.Fatalf("expected short-circuit, got error: %v", err)
	}
	if out != "false" {
		t.Fatalf("expected 'false', got %q", out)
	}
}

func TestLogicalOrShortCircuitOnError(t *testing.T) {
	// true || error_expr should short-circuit, no error
	out, err := runFig(t, `print(true || -"bad")`)
	if err != nil {
		t.Fatalf("expected short-circuit, got error: %v", err)
	}
	if out != "true" {
		t.Fatalf("expected 'true', got %q", out)
	}
}

func TestLogicalAndPropagatesError(t *testing.T) {
	// true && error_expr should propagate the error
	_, err := runFig(t, `let x = true && -"bad"`)
	if err == nil {
		t.Fatal("expected runtime error when && RHS errors")
	}
}

func TestLogicalOrPropagatesError(t *testing.T) {
	// false || error_expr should propagate the error
	_, err := runFig(t, `let x = false || -"bad"`)
	if err == nil {
		t.Fatal("expected runtime error when || RHS errors")
	}
}
