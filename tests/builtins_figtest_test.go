package tests

import (
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
)

func useFigtest(code string) string {
	return "use \"figtest\"\n" + code
}

func resetFigtest() {
	builtins.ResetFigtest()
}

func TestFigtestTestPass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("soma", fn() { figtest.assert(2 + 2 == 4) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestTestFail(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("falha", fn() { figtest.assert(1 == 2, "deveria falhar") })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestDescribe(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.describe(\"math\", fn() {\n" +
		"  figtest.test(\"soma\", fn() { figtest.assertEq(2 + 2, 4) })\n" +
		"  figtest.test(\"sub\", fn() { figtest.assertEq(5 - 3, 2) })\n" +
		"})\n")
	_, err := runFig(t, code)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 2 {
		t.Fatalf("expected 2 passed, got %d", s.Passed())
	}
	found := false
	for _, line := range s.Output() {
		if strings.Contains(line, "math") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected describe group name in output, got %v", s.Output())
	}
}

func TestFigtestSkip(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.skip("wip", fn() { figtest.assert(false) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Skipped() != 1 {
		t.Fatalf("expected 1 skipped, got %d", s.Skipped())
	}
	if s.Passed() != 0 {
		t.Fatalf("expected 0 passed, got %d", s.Passed())
	}
}

func TestFigtestOnly(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.only("isolado", fn() { figtest.assert(true) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertTrue(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("ok", fn() { figtest.assert(true) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertFalse(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("fail", fn() { figtest.assert(false) })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertEqPass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("eq", fn() { figtest.assertEq(42, 42) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertEqFail(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("neq", fn() { figtest.assertEq(1, 2) })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertNeqPass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("neq", fn() { figtest.assertNeq(1, 2) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertNeqFail(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("eq", fn() { figtest.assertNeq(5, 5) })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertErrorPass(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.test(\"err\", fn() {\n" +
		"  figtest.assertError(fn() { figtest.assert(false, \"boom\") })\n" +
		"})\n")
	_, err := runFig(t, code)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertErrorFail(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.test(\"no-err\", fn() {\n" +
		"  figtest.assertError(fn() { let x = 1 })\n" +
		"})\n")
	_, _ = runFig(t, code)
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertNoErrorPass(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.test(\"ok\", fn() {\n" +
		"  figtest.assertNoError(fn() { let x = 1 })\n" +
		"})\n")
	_, err := runFig(t, code)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertNoErrorFail(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.test(\"err\", fn() {\n" +
		"  figtest.assertNoError(fn() { figtest.assert(false, \"boom\") })\n" +
		"})\n")
	_, _ = runFig(t, code)
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertNearPass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("near", fn() { figtest.assertNear(3.14, 3.15, 0.02) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertNearFail(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("far", fn() { figtest.assertNear(1.0, 2.0, 0.1) })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertContainsStringPass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("has", fn() { figtest.assertContains("hello world", "world") })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertContainsStringFail(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("nope", fn() { figtest.assertContains("hello", "xyz") })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertContainsArrayPass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("arr", fn() { figtest.assertContains([1, 2, 3], 2) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertContainsArrayFail(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("no", fn() { figtest.assertContains([1, 2, 3], 9) })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertTypePass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("type", fn() { figtest.assertType(42, "number") })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertTypeFail(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("type", fn() { figtest.assertType("hi", "number") })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestAssertLengthStringPass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("len", fn() { figtest.assertLength("abc", 3) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertLengthArrayPass(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("len", fn() { figtest.assertLength([1, 2], 2) })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertLengthFail(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("len", fn() { figtest.assertLength("ab", 5) })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
}

func TestFigtestBeforeEach(t *testing.T) {
	resetFigtest()
	code := useFigtest("let counter = 0\n" +
		"figtest.describe(\"hooks\", fn() {\n" +
		"  figtest.beforeEach(fn() { counter = counter + 1 })\n" +
		"  figtest.test(\"first\", fn() { figtest.assert(counter == 1) })\n" +
		"  figtest.test(\"second\", fn() { figtest.assert(counter == 2) })\n" +
		"})\n")
	_, err := runFig(t, code)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 2 {
		t.Fatalf("expected 2 passed, got %d (failed: %d)", s.Passed(), s.Failed())
	}
}

func TestFigtestAfterEach(t *testing.T) {
	resetFigtest()
	code := useFigtest("let counter = 0\n" +
		"figtest.describe(\"hooks\", fn() {\n" +
		"  figtest.afterEach(fn() { counter = counter + 10 })\n" +
		"  figtest.test(\"first\", fn() { figtest.assertEq(counter, 0) })\n" +
		"  figtest.test(\"second\", fn() { figtest.assertEq(counter, 10) })\n" +
		"})\n")
	_, err := runFig(t, code)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 2 {
		t.Fatalf("expected 2 passed, got %d (failed: %d)", s.Passed(), s.Failed())
	}
}

func TestFigtestBeforeAll(t *testing.T) {
	resetFigtest()
	code := useFigtest("let setup = false\n" +
		"figtest.describe(\"hooks\", fn() {\n" +
		"  figtest.beforeAll(fn() { setup = true })\n" +
		"  figtest.test(\"check\", fn() { figtest.assert(setup) })\n" +
		"})\n")
	_, err := runFig(t, code)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestReset(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.test(\"a\", fn() { figtest.assert(true) })\n" +
		"figtest.reset()\n" +
		"figtest.test(\"b\", fn() { figtest.assert(true) })\n")
	_, err := runFig(t, code)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed after reset, got %d", s.Passed())
	}
}

func TestFigtestCounters(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.test(\"ok\", fn() { figtest.assert(true) })\n" +
		"figtest.test(\"fail\", fn() { figtest.assert(false) })\n" +
		"figtest.skip(\"skip\", fn() {})\n" +
		"print(figtest.count())\n" +
		"print(figtest.passed())\n" +
		"print(figtest.failed())\n" +
		"print(figtest.skipped())\n")
	out, _ := runFig(t, code)
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) < 4 {
		t.Fatalf("expected 4 output lines, got %d: %q", len(lines), out)
	}
	if lines[0] != "3" {
		t.Fatalf("expected count=3, got %q", lines[0])
	}
	if lines[1] != "1" {
		t.Fatalf("expected passed=1, got %q", lines[1])
	}
	if lines[2] != "1" {
		t.Fatalf("expected failed=1, got %q", lines[2])
	}
	if lines[3] != "1" {
		t.Fatalf("expected skipped=1, got %q", lines[3])
	}
}

func TestFigtestSummary(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.test(\"pass\", fn() { figtest.assert(true) })\n" +
		"print(figtest.summary())\n")
	out, _ := runFig(t, code)
	if !strings.Contains(out, "1 passed") {
		t.Fatalf("expected summary to contain '1 passed', got %q", out)
	}
}

func TestFigtestMultipleDescribeGroups(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.describe(\"group A\", fn() {\n" +
		"  figtest.test(\"a1\", fn() { figtest.assert(true) })\n" +
		"})\n" +
		"figtest.describe(\"group B\", fn() {\n" +
		"  figtest.test(\"b1\", fn() { figtest.assert(true) })\n" +
		"  figtest.test(\"b2\", fn() { figtest.assert(true) })\n" +
		"})\n")
	_, err := runFig(t, code)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 3 {
		t.Fatalf("expected 3 passed, got %d", s.Passed())
	}
}

func TestFigtestOutputFormat(t *testing.T) {
	resetFigtest()
	code := useFigtest("figtest.test(\"pass test\", fn() { figtest.assert(true) })\n" +
		"figtest.test(\"fail test\", fn() { figtest.assert(false) })\n" +
		"figtest.skip(\"skip test\", fn() {})\n")
	_, _ = runFig(t, code)
	s := builtins.GetFigtestState()
	output := s.Output()
	foundPass := false
	foundFail := false
	foundSkip := false
	for _, line := range output {
		if strings.Contains(line, "\u2713") && strings.Contains(line, "pass test") {
			foundPass = true
		}
		if strings.Contains(line, "\u2717") && strings.Contains(line, "fail test") {
			foundFail = true
		}
		if strings.Contains(line, "\u25CB") && strings.Contains(line, "skip test") {
			foundSkip = true
		}
	}
	if !foundPass {
		t.Fatalf("expected pass marker in output: %v", output)
	}
	if !foundFail {
		t.Fatalf("expected fail marker in output: %v", output)
	}
	if !foundSkip {
		t.Fatalf("expected skip marker in output: %v", output)
	}
}

func TestFigtestAssertEqCustomMsg(t *testing.T) {
	resetFigtest()
	_, _ = runFig(t, useFigtest(`figtest.test("msg", fn() { figtest.assertEq(1, 2, "custom msg") })`))
	s := builtins.GetFigtestState()
	if s.Failed() != 1 {
		t.Fatalf("expected 1 failed, got %d", s.Failed())
	}
	found := false
	for _, line := range s.Output() {
		if strings.Contains(line, "custom msg") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected custom message in output, got %v", s.Output())
	}
}

func TestFigtestAssertTypeString(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("str", fn() { figtest.assertType("hello", "string") })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertTypeBool(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("bool", fn() { figtest.assertType(true, "boolean") })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}

func TestFigtestAssertTypeArray(t *testing.T) {
	resetFigtest()
	_, err := runFig(t, useFigtest(`figtest.test("arr", fn() { figtest.assertType([1, 2], "array") })`))
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	s := builtins.GetFigtestState()
	if s.Passed() != 1 {
		t.Fatalf("expected 1 passed, got %d", s.Passed())
	}
}
