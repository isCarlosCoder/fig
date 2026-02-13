package tests

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
)

// Ensure importing a local file with '*' injects its top-level symbols into the importer scope.
func TestImportLocalStar(t *testing.T) {
	// use existing testdata/src/utils.fig which defines `fn soma(a, b)`
	dir := testdataDir(t)
	fakeFile := filepath.Join(dir, "_test_star_.fig")
	src := `import "./src/utils.fig" *
print(soma(2, 3));`
	var buf bytes.Buffer
	err := interpreter.Run(src, fakeFile, environment.NewEnv(nil), &buf, &buf)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	out := strings.TrimSpace(buf.String())
	if out != "5" {
		t.Fatalf("expected '5', got %q", out)
	}
}
