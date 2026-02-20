package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
)

// testdataDir returns the absolute path to the testdata directory.
func testdataDir(t *testing.T) string {
	t.Helper()
	d, err := filepath.Abs("testdata")
	if err != nil {
		t.Fatal(err)
	}
	return d
}

// runFigSource runs Fig source code with baseDir set to the testdata directory,
// so import statements resolve correctly.
func runFigSource(t *testing.T, src string) (string, error) {
	t.Helper()
	dir := testdataDir(t)
	fakeFile := filepath.Join(dir, "_test_.fig")
	var buf bytes.Buffer
	err := interpreter.Run(src, fakeFile, environment.NewEnv(nil), &buf, &buf)
	return strings.TrimSpace(buf.String()), err
}

// runFigFile runs a .fig file from the testdata directory.
func runFigFile(t *testing.T, name string) (string, error) {
	t.Helper()
	dir := testdataDir(t)
	path := filepath.Join(dir, name)
	data, readErr := os.ReadFile(path)
	if readErr != nil {
		t.Fatalf("cannot read %s: %v", path, readErr)
	}
	var buf bytes.Buffer
	err := interpreter.Run(string(data), path, environment.NewEnv(nil), &buf, &buf)
	return strings.TrimSpace(buf.String()), err
}

func TestImportBasicFunction(t *testing.T) {
	src := `import "math_utils"
print(math_utils.soma(2, 3));`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "5" {
		t.Fatalf("expected '5', got %q", out)
	}
}

func TestImportVariable(t *testing.T) {
	src := `import "math_utils"
print(math_utils.PI);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3.14159" {
		t.Fatalf("expected '3.14159', got %q", out)
	}
}

func TestImportMultipleModules(t *testing.T) {
	src := `import "math_utils"
import "greet"
print(math_utils.soma(10, 20));
print(greet.saudar("Fig"));`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	expected := "30\nOlá, Fig!"
	if out != expected {
		t.Fatalf("expected %q, got %q", expected, out)
	}
}

func TestImportTransitive(t *testing.T) {
	src := `import "transitive"
print(transitive.dobro(7));`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "14" {
		t.Fatalf("expected '14', got %q", out)
	}
}

func TestImportRuntimeFileInModule(t *testing.T) {
	// create a small module file inside testdata/foo/bar
	dir := testdataDir(t)
	modDir := filepath.Join(dir, "foo", "bar")
	if err := os.MkdirAll(modDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	defer os.RemoveAll(filepath.Join(dir, "foo")) // cleanup
	libContent := `use "runtime"`
	libContent += "\n" + `let x = runtime.file(); let y = runtime.dir()`
	if err := os.WriteFile(filepath.Join(modDir, "lib.fig"), []byte(libContent), 0o644); err != nil {
		t.Fatalf("write lib.fig: %v", err)
	}

	src := `import "foo/bar/lib.fig"
print(lib.x);
print(lib.y)`
	// y will store runtime.dir()
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 2 {
		t.Fatalf("expected two lines of output, got %d", len(lines))
	}
	expected := filepath.Join(dir, "foo", "bar", "lib.fig")
	if lines[0] != expected {
		t.Fatalf("expected file %q, got %q", expected, lines[0])
	}
	expectedDir := filepath.Dir(expected)
	if lines[1] != expectedDir {
		t.Fatalf("expected dir %q, got %q", expectedDir, lines[1])
	}
}

func TestImportTransitiveFunctionAvailable(t *testing.T) {
	src := `import "math_utils"
print(math_utils.mult(3, 4));`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "12" {
		t.Fatalf("expected '12', got %q", out)
	}
}

func TestImportWithExtension(t *testing.T) {
	src := `import "with_extension.fig"
print(with_extension.imported);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true" {
		t.Fatalf("expected 'true', got %q", out)
	}
}

func TestImportCircularDoesNotLoop(t *testing.T) {
	src := `import "circular_a"
print(circular_a.fromA);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "A" {
		t.Fatalf("expected 'A', got %q", out)
	}
}

func TestImportNonexistentFile(t *testing.T) {
	src := `import "nao_existe"`
	_, err := runFigSource(t, src)
	if err == nil {
		t.Fatal("expected error for nonexistent import")
	}
}

func TestImportIdempotent(t *testing.T) {
	src := `import "math_utils"
import "math_utils"
print(math_utils.soma(1, 1));`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "2" {
		t.Fatalf("expected '2', got %q", out)
	}
}

func TestImportLocalAlias(t *testing.T) {
	src := `import "math_utils" utils
print(utils.soma(2, 3));`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "5" {
		t.Fatalf("expected '5', got %q", out)
	}
}

func TestImportLocalAliasWithSubdirFromRoot(t *testing.T) {
	out, err := runFigFile(t, "main_using_src.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "5" {
		t.Fatalf("expected '5', got %q", out)
	}
}

func TestImportLocalAliasWithSubdirFromInside(t *testing.T) {
	out, err := runFigFile(t, "src/main_inside.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "9" {
		t.Fatalf("expected '9', got %q", out)
	}
}

func TestImportWithDotSlashFromInside(t *testing.T) {
	out, err := runFigFile(t, "src/main_inside_rel.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "9" {
		t.Fatalf("expected '9', got %q", out)
	}
}

func TestImportWithParentDirFromDeep(t *testing.T) {
	out, err := runFigFile(t, "src/subdir/deep_import.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3" {
		t.Fatalf("expected '3', got %q", out)
	}
}

func TestImportPlainNameResolvesRelativeToFile(t *testing.T) {
	// create a temporary file in src that imports "utils.fig" without path
	out, err := runFigFile(t, "src/main_inside.fig")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "9" {
		t.Fatalf("expected '9', got %q", out)
	}
}

func TestImportUsesmath(t *testing.T) {
	src := `import "uses_math"
print(uses_math.quadrado(5));`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "25" {
		t.Fatalf("expected '25', got %q", out)
	}
}

func TestImportModuleAlias(t *testing.T) {
	src := `import "mod:isCarlosCoder/myfigtestdependency" lib
print(lib.hello("Fig"));`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "hello, Fig" {
		t.Fatalf("expected 'hello, Fig', got %q", out)
	}
}

func TestImportModuleDefaultName(t *testing.T) {
	src := `import "mod:isCarlosCoder/myfigtestdependency"
print(myfigtestdependency.magic);`
	out, err := runFigSource(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "42" {
		t.Fatalf("expected '42', got %q", out)
	}
}
