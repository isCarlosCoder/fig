package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// hasAsanSupport checks if gcc supports -fsanitize=address.
func hasAsanSupport(t *testing.T) bool {
	t.Helper()
	probe := exec.Command("gcc", "-fsanitize=address", "-shared", "-fPIC", "-x", "c", "-o", "/dev/null", "-")
	probe.Stdin = strings.NewReader("void f(){}\n")
	return probe.Run() == nil
}

// TestAsanCleanLibrary builds the ASAN-clean test library and a standalone
// ASAN-instrumented driver, runs the driver, and verifies zero ASAN errors.
func TestAsanCleanLibrary(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("ASAN test only supported on Linux")
	}
	if !hasAsanSupport(t) {
		t.Skip("gcc does not support -fsanitize=address on this system")
	}

	root := repoRootForTest(t)
	tmpDir := t.TempDir()

	// 1. Build the C library with ASAN instrumentation
	libPath := filepath.Join(tmpDir, "libasan_clean.so")
	libSrc := filepath.Join(root, "tests", "ffi_integration", "lib_asan_clean.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-fsanitize=address", "-g", "-O1", "-o", libPath, libSrc)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build ASAN library: %v (%s)", err, string(out))
	}

	// 2. Build the ASAN driver (also instrumented)
	driverPath := filepath.Join(tmpDir, "asan_driver")
	driverSrc := filepath.Join(root, "tests", "ffi_integration", "asan_driver.c")
	gcc2 := exec.Command("gcc", "-fsanitize=address", "-g", "-O1", "-o", driverPath, driverSrc, "-ldl")
	if out, err := gcc2.CombinedOutput(); err != nil {
		t.Fatalf("failed to build ASAN driver: %v (%s)", err, string(out))
	}

	// 3. Run the driver with the ASAN library
	driver := exec.Command(driverPath, libPath)
	driver.Env = append(os.Environ(), "ASAN_OPTIONS=detect_leaks=1:halt_on_error=1:abort_on_error=0")
	out, err := driver.CombinedOutput()
	output := string(out)
	t.Logf("ASAN driver output:\n%s", output)

	if err != nil {
		t.Fatalf("ASAN driver failed (possible memory error): %v\nOutput:\n%s", err, output)
	}

	// 4. Verify no ASAN errors in output
	if strings.Contains(output, "ERROR: AddressSanitizer") {
		t.Fatalf("AddressSanitizer error detected:\n%s", output)
	}
	if strings.Contains(output, "ERROR: LeakSanitizer") {
		t.Fatalf("LeakSanitizer error detected:\n%s", output)
	}
	if !strings.Contains(output, "PASSED") {
		t.Errorf("driver did not report PASSED")
	}
}

// TestAsanWithMainLib runs the ASAN driver against the main lib.c test library.
func TestAsanWithMainLib(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("ASAN test only supported on Linux")
	}
	if !hasAsanSupport(t) {
		t.Skip("gcc does not support -fsanitize=address on this system")
	}

	root := repoRootForTest(t)
	tmpDir := t.TempDir()

	// Build main lib.c with ASAN
	libPath := filepath.Join(tmpDir, "libtest.so")
	libSrc := filepath.Join(root, "tests", "ffi_integration", "lib.c")
	gcc := exec.Command("gcc", "-shared", "-fPIC", "-fsanitize=address", "-g", "-O1", "-o", libPath, libSrc)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("failed to build ASAN lib: %v (%s)", err, string(out))
	}

	// Build driver
	driverPath := filepath.Join(tmpDir, "asan_driver")
	driverSrc := filepath.Join(root, "tests", "ffi_integration", "asan_driver.c")
	gcc2 := exec.Command("gcc", "-fsanitize=address", "-g", "-O1", "-o", driverPath, driverSrc, "-ldl")
	if out, err := gcc2.CombinedOutput(); err != nil {
		t.Fatalf("failed to build ASAN driver: %v (%s)", err, string(out))
	}

	driver := exec.Command(driverPath, libPath)
	driver.Env = append(os.Environ(), "ASAN_OPTIONS=detect_leaks=1:halt_on_error=1:abort_on_error=0")
	out, err := driver.CombinedOutput()
	output := string(out)
	t.Logf("ASAN driver output:\n%s", output)

	if err != nil {
		t.Fatalf("ASAN driver failed: %v\nOutput:\n%s", err, output)
	}
	if strings.Contains(output, "ERROR: AddressSanitizer") {
		t.Fatalf("AddressSanitizer error detected:\n%s", output)
	}
}

// TestAsanScriptExists verifies that ASAN/Valgrind check scripts exist and are executable.
func TestAsanScriptExists(t *testing.T) {
	root := repoRootForTest(t)
	scripts := []string{"tools/asan-check.sh", "tools/valgrind-check.sh"}
	for _, s := range scripts {
		p := filepath.Join(root, s)
		info, err := os.Stat(p)
		if err != nil {
			t.Errorf("script %s not found: %v", s, err)
			continue
		}
		if info.Mode()&0111 == 0 {
			t.Errorf("script %s is not executable", s)
		}
	}
}

// TestValgrindAvailability checks if valgrind is installed.
func TestValgrindAvailability(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Valgrind test only on Linux")
	}
	if _, err := exec.LookPath("valgrind"); err != nil {
		t.Skip("valgrind not installed")
	}
	cmd := exec.Command("valgrind", "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Skipf("valgrind not functional: %v", err)
	}
	t.Logf("valgrind version: %s", strings.TrimSpace(string(out)))
}

// TestAsanBuildIntegration verifies ASAN-instrumented builds contain expected symbols.
func TestAsanBuildIntegration(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("ASAN build test only on Linux")
	}
	if !hasAsanSupport(t) {
		t.Skip("gcc does not support -fsanitize=address")
	}

	root := repoRootForTest(t)
	tmpDir := t.TempDir()
	libPath := filepath.Join(tmpDir, "libasan_clean.so")
	cpath := filepath.Join(root, "tests", "ffi_integration", "lib_asan_clean.c")

	gcc := exec.Command("gcc", "-shared", "-fPIC", "-fsanitize=address", "-g", "-O1", "-o", libPath, cpath)
	if out, err := gcc.CombinedOutput(); err != nil {
		t.Fatalf("ASAN build failed: %v (%s)", err, string(out))
	}

	nm := exec.Command("nm", "-D", libPath)
	out, err := nm.CombinedOutput()
	if err != nil {
		t.Fatalf("nm failed: %v (%s)", err, string(out))
	}
	symbols := string(out)
	expected := []string{"asan_safe_add", "asan_safe_concat", "asan_safe_strlen"}
	for _, s := range expected {
		if !strings.Contains(symbols, s) {
			t.Errorf("missing symbol %s in ASAN-built library", s)
		}
	}
}
