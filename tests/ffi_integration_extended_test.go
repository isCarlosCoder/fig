package tests

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/iscarloscoder/fig/builtins"
)

// setupIntegrationExt builds the ffi-helper binary and the test C library,
// creates a temporary project directory with fig.toml, and returns paths.
func setupIntegrationExt(t *testing.T) (proj, bin, libPath string) {
	t.Helper()
	builtins.StopAllHelpers()
	root := findRepoRoot(t)

	binDir := t.TempDir()
	bin = filepath.Join(binDir, "ffi-helper")
	cmd := exec.Command("go", "build", "-o", bin, "./tools/ffi-helper")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("failed to build helper: %v (%s)", err, string(out))
	}

	libDir := t.TempDir()
	libPath = filepath.Join(libDir, "libtest.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build test lib: %v (%s)", err, string(out))
	}

	proj = t.TempDir()
	fig := filepath.Join(proj, "fig.toml")
	cfg := fmt.Sprintf("[ffi]\nenabled = true\nhelper = %q\ncall_timeout = 10000\n", bin)
	if err := os.WriteFile(fig, []byte(cfg), 0644); err != nil {
		t.Fatalf("write fig.toml: %v", err)
	}
	return proj, bin, libPath
}

// --- mul2 tests ---

func TestIntegrationMul2Zero(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	sym, err := hc.Sym(handle, "mul2", "double")
	if err != nil {
		t.Fatalf("sym: %v", err)
	}
	res, err := hc.CallSymbol(sym, []interface{}{0.0, 100.0}, nil)
	if err != nil {
		t.Fatalf("call mul2(0.0, 100.0): %v", err)
	}
	got, ok := res.(float64)
	if !ok {
		t.Fatalf("expected float64, got %T: %v", res, res)
	}
	if got != 0.0 {
		t.Fatalf("expected 0.0, got %v", got)
	}
}

func TestIntegrationMul2Negative(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	sym, err := hc.Sym(handle, "mul2", "double")
	if err != nil {
		t.Fatalf("sym: %v", err)
	}
	res, err := hc.CallSymbol(sym, []interface{}{-3.0, 4.0}, nil)
	if err != nil {
		t.Fatalf("call mul2(-3.0, 4.0): %v", err)
	}
	got, ok := res.(float64)
	if !ok {
		t.Fatalf("expected float64, got %T: %v", res, res)
	}
	if got != -12.0 {
		t.Fatalf("expected -12.0, got %v", got)
	}
}

// --- sum3 tests ---

func TestIntegrationSum3AllZeros(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	sym, err := hc.Sym(handle, "sum3", "int")
	if err != nil {
		t.Fatalf("sym: %v", err)
	}
	res, err := hc.CallSymbol(sym, []interface{}{0, 0, 0}, []string{"int", "int", "int"})
	if err != nil {
		t.Fatalf("call sum3(0,0,0): %v", err)
	}
	got, ok := res.(float64)
	if !ok {
		t.Fatalf("expected float64, got %T: %v", res, res)
	}
	if int(got) != 0 {
		t.Fatalf("expected 0, got %v", got)
	}
}

func TestIntegrationSum3Large(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	sym, err := hc.Sym(handle, "sum3", "int")
	if err != nil {
		t.Fatalf("sym: %v", err)
	}
	res, err := hc.CallSymbol(sym, []interface{}{1000000, 2000000, 3000000}, []string{"int", "int", "int"})
	if err != nil {
		t.Fatalf("call sum3(1000000,2000000,3000000): %v", err)
	}
	got, ok := res.(float64)
	if !ok {
		t.Fatalf("expected float64, got %T: %v", res, res)
	}
	if int(got) != 6000000 {
		t.Fatalf("expected 6000000, got %v", got)
	}
}

// --- dupstr tests ---

func TestIntegrationDupstrEmpty(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	sym, err := hc.Sym(handle, "dupstr", "string")
	if err != nil {
		t.Fatalf("sym: %v", err)
	}
	res, err := hc.CallSymbol(sym, []interface{}{""}, nil)
	if err != nil {
		t.Fatalf("call dupstr(''): %v", err)
	}
	got, ok := res.(string)
	if !ok {
		t.Fatalf("expected string, got %T: %v", res, res)
	}
	if got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}

func TestIntegrationDupstrUTF8(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	sym, err := hc.Sym(handle, "dupstr", "string")
	if err != nil {
		t.Fatalf("sym: %v", err)
	}
	input := "café ☕ 日本語"
	res, err := hc.CallSymbol(sym, []interface{}{input}, nil)
	if err != nil {
		t.Fatalf("call dupstr(%q): %v", input, err)
	}
	got, ok := res.(string)
	if !ok {
		t.Fatalf("expected string, got %T: %v", res, res)
	}
	if got != input {
		t.Fatalf("expected %q, got %q", input, got)
	}
}

// --- concat tests ---

func TestIntegrationConcatStrings(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	sym, err := hc.Sym(handle, "concat", "string")
	if err != nil {
		t.Fatalf("sym: %v", err)
	}
	res, err := hc.CallSymbol(sym, []interface{}{"hello", " world"}, nil)
	if err != nil {
		t.Fatalf("call concat('hello', ' world'): %v", err)
	}
	got, ok := res.(string)
	if !ok {
		t.Fatalf("expected string, got %T: %v", res, res)
	}
	if got != "hello world" {
		t.Fatalf("expected %q, got %q", "hello world", got)
	}
}

func TestIntegrationConcatEmpty(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	sym, err := hc.Sym(handle, "concat", "string")
	if err != nil {
		t.Fatalf("sym: %v", err)
	}
	res, err := hc.CallSymbol(sym, []interface{}{"", "test"}, nil)
	if err != nil {
		t.Fatalf("call concat('', 'test'): %v", err)
	}
	got, ok := res.(string)
	if !ok {
		t.Fatalf("expected string, got %T: %v", res, res)
	}
	if got != "test" {
		t.Fatalf("expected %q, got %q", "test", got)
	}
}

// --- alloc/write/read test (skipped if Alloc is not exposed) ---

func TestIntegrationAllocWriteRead(t *testing.T) {
	// Alloc, MemWrite, MemRead are not exposed on HelperForTest — skip.
	t.Skip("Alloc/MemWrite/MemRead not exposed on HelperForTest; skipping")
}

// --- multiple symbols from same lib ---

func TestIntegrationMultipleSymsSameLib(t *testing.T) {
	proj, bin, libPath := setupIntegrationExt(t)
	defer builtins.StopAllHelpers()

	hc, err := builtins.GetHelperForTest(proj, bin)
	if err != nil {
		t.Fatalf("get helper: %v", err)
	}
	handle, err := hc.Load(libPath)
	if err != nil {
		t.Fatalf("load: %v", err)
	}

	// resolve mul2
	symMul2, err := hc.Sym(handle, "mul2", "double")
	if err != nil {
		t.Fatalf("sym mul2: %v", err)
	}

	// resolve sum3
	symSum3, err := hc.Sym(handle, "sum3", "int")
	if err != nil {
		t.Fatalf("sym sum3: %v", err)
	}

	// call mul2(2.5, 4.0) = 10.0
	resMul, err := hc.CallSymbol(symMul2, []interface{}{2.5, 4.0}, nil)
	if err != nil {
		t.Fatalf("call mul2(2.5, 4.0): %v", err)
	}
	gotMul, ok := resMul.(float64)
	if !ok {
		t.Fatalf("mul2: expected float64, got %T: %v", resMul, resMul)
	}
	if gotMul != 10.0 {
		t.Fatalf("mul2: expected 10.0, got %v", gotMul)
	}

	// call sum3(10, 20, 30) = 60
	resSum, err := hc.CallSymbol(symSum3, []interface{}{10, 20, 30}, []string{"int", "int", "int"})
	if err != nil {
		t.Fatalf("call sum3(10, 20, 30): %v", err)
	}
	gotSum, ok := resSum.(float64)
	if !ok {
		t.Fatalf("sum3: expected float64, got %T: %v", resSum, resSum)
	}
	if int(gotSum) != 60 {
		t.Fatalf("sum3: expected 60, got %v", gotSum)
	}
}
