package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestFfiGenBasic(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"mymath\"\n\n[[functions]]\nname = \"add\"\nsymbol = \"add\"\nreturn = \"int\"\nargs = [\"int\", \"int\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "test.ffi.def.toml")
	outPath := filepath.Join(tmpDir, "bindings.fig")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-output", outPath)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("ffi-gen failed: %v (%s)", err, string(out))
	}
	data, _ := os.ReadFile(outPath)
	code := string(data)

	if !strings.Contains(code, `use "ffi"`) {
		t.Error("missing use ffi")
	}
	if !strings.Contains(code, `ffi.lib_name("mymath")`) {
		t.Error("missing lib_name call")
	}
	if !strings.Contains(code, `ffi.sym(__lib, "add", "int", ["int", "int"])`) {
		t.Error("missing sym call for add with argTypes")
	}
	if !strings.Contains(code, `fn add(a, b)`) {
		t.Error("missing fn add wrapper")
	}
	if !strings.Contains(code, `ffi.call(__sym_add, a, b)`) {
		t.Error("missing call in add wrapper")
	}
}

func TestFfiGenStructs(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"geo\"\n\n[[structs]]\nname = \"Point\"\nfields = [\n  {name=\"x\",type=\"int\"},\n  {name=\"y\",type=\"int\"}\n]\n\n[[structs]]\nname = \"Rect\"\nfields = [\n  {name=\"origin\",type=\"struct:Point\"},\n  {name=\"width\",type=\"double\"},\n  {name=\"height\",type=\"double\"}\n]\n\n[[functions]]\nname = \"area\"\nsymbol = \"rect_area\"\nreturn = \"double\"\nargs = [\"struct:Rect\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "test.ffi.def.toml")
	outPath := filepath.Join(tmpDir, "bindings.fig")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-output", outPath)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("ffi-gen failed: %v (%s)", err, string(out))
	}
	data, _ := os.ReadFile(outPath)
	code := string(data)

	if !strings.Contains(code, `ffi.define_struct("Point"`) {
		t.Error("missing define_struct for Point")
	}
	if !strings.Contains(code, `ffi.define_struct("Rect"`) {
		t.Error("missing define_struct for Rect")
	}
	if !strings.Contains(code, `"struct:Point"`) {
		t.Error("missing nested struct reference in Rect")
	}
	if !strings.Contains(code, `["struct:Rect"]`) {
		t.Error("missing arg_types for area function")
	}
}

func TestFfiGenMixedArgTypes(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"mixed\"\n\n[[functions]]\nname = \"format\"\nreturn = \"string\"\nargs = [\"string\", \"int\", \"double\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "test.ffi.def.toml")
	outPath := filepath.Join(tmpDir, "bindings.fig")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-output", outPath)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("ffi-gen failed: %v (%s)", err, string(out))
	}
	data, _ := os.ReadFile(outPath)
	code := string(data)

	if !strings.Contains(code, `["string", "int", "double"]`) {
		t.Error("missing arg_types for mixed-type function")
	}
}

func TestFfiGenVoidReturn(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"sideeffect\"\n\n[[functions]]\nname = \"init_system\"\nreturn = \"void\"\nargs = []\n\n[[functions]]\nname = \"log_msg\"\nreturn = \"void\"\nargs = [\"string\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "test.ffi.def.toml")
	outPath := filepath.Join(tmpDir, "bindings.fig")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-output", outPath)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("ffi-gen failed: %v (%s)", err, string(out))
	}
	data, _ := os.ReadFile(outPath)
	code := string(data)

	if !strings.Contains(code, `"void"`) {
		t.Error("missing void return type")
	}
	if !strings.Contains(code, `fn init_system()`) {
		t.Error("missing init_system wrapper with no args")
	}
	if !strings.Contains(code, `fn log_msg(a)`) {
		t.Error("missing log_msg wrapper")
	}
}

func TestFfiGenExplicitPath(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"custom\"\npath = \"/opt/libs/libcustom\"\n\n[[functions]]\nname = \"hello\"\nreturn = \"string\"\nargs = []\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "test.ffi.def.toml")
	outPath := filepath.Join(tmpDir, "bindings.fig")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-output", outPath)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("ffi-gen failed: %v (%s)", err, string(out))
	}
	data, _ := os.ReadFile(outPath)
	code := string(data)

	if !strings.Contains(code, `ffi.load("/opt/libs/libcustom" + ffi.lib_ext())`) {
		t.Error("missing explicit path with lib_ext()")
	}
}

func TestFfiGenInit(t *testing.T) {
	root := repoRootForTest(t)
	tmpDir := t.TempDir()
	projDir := filepath.Join(tmpDir, "testplugin")

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-init", projDir)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("ffi-gen -init failed: %v (%s)", err, string(out))
	}

	expectedFiles := []string{"fig.toml", "testplugin.ffi.def.toml", "testplugin.c", "Makefile", "main.fig", "README.md"}
	for _, f := range expectedFiles {
		path := filepath.Join(projDir, f)
		if _, err := os.Stat(path); err != nil {
			t.Errorf("expected file %s not found: %v", f, err)
		}
	}

	data, _ := os.ReadFile(filepath.Join(projDir, "fig.toml"))
	if !strings.Contains(string(data), `enabled = true`) {
		t.Error("fig.toml missing enabled = true")
	}

	outPath := filepath.Join(tmpDir, "bindings.fig")
	defPath := filepath.Join(projDir, "testplugin.ffi.def.toml")
	cmd2 := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-output", outPath)
	cmd2.Dir = root
	if out, err := cmd2.CombinedOutput(); err != nil {
		t.Fatalf("ffi-gen on scaffolded project failed: %v (%s)", err, string(out))
	}
	gen, _ := os.ReadFile(outPath)
	if !strings.Contains(string(gen), `fn add`) {
		t.Error("generated bindings missing add function")
	}
	if !strings.Contains(string(gen), `fn greet`) {
		t.Error("generated bindings missing greet function")
	}
}

func TestFfiGenSymbolDefault(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"test\"\n\n[[functions]]\nname = \"my_func\"\nreturn = \"int\"\nargs = [\"int\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "test.ffi.def.toml")
	outPath := filepath.Join(tmpDir, "bindings.fig")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-output", outPath)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("ffi-gen failed: %v (%s)", err, string(out))
	}
	data, _ := os.ReadFile(outPath)
	code := string(data)

	if !strings.Contains(code, `ffi.sym(__lib, "my_func", "int", ["int"])`) {
		t.Error("expected symbol name to default to function name with argTypes")
	}
}

// --- P2-4: New tests ---

func TestFfiGenInvalidType(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"test\"\n\n[[functions]]\nname = \"bad_func\"\nsymbol = \"bad_func\"\nreturn = \"pointer\"\nargs = [\"int\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "bad.ffi.def.toml")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath)
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected non-zero exit for invalid type, got success")
	}
	outStr := string(out)
	if !strings.Contains(outStr, "unknown return type: pointer") {
		t.Errorf("expected 'unknown return type: pointer' in output, got: %s", outStr)
	}
	if !strings.Contains(outStr, "Supported: int, double, string, void, struct:Name") {
		t.Errorf("expected supported types list in output, got: %s", outStr)
	}
}

func TestFfiGenInvalidArgType(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"test\"\n\n[[functions]]\nname = \"bad_args\"\nreturn = \"int\"\nargs = [\"int\", \"float\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "bad_args.ffi.def.toml")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath)
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("expected non-zero exit for invalid arg type")
	}
	if !strings.Contains(string(out), "unknown type: float") {
		t.Errorf("expected 'unknown type: float' in output, got: %s", out)
	}
}

func TestFfiGenValidateValid(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"mylib\"\n\n[[functions]]\nname = \"add\"\nreturn = \"int\"\nargs = [\"int\", \"int\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "valid.ffi.def.toml")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-validate")
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("--validate should succeed for valid def, got: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "is valid") {
		t.Errorf("expected 'is valid' message, got: %s", out)
	}
}

func TestFfiGenValidateInvalid(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"mylib\"\n\n[[functions]]\nname = \"bad\"\nreturn = \"pointer\"\nargs = [\"uint8\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "invalid.ffi.def.toml")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-validate")
	cmd.Dir = root
	_, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("--validate should fail for invalid def")
	}
}

func TestFfiGenCSignatureComment(t *testing.T) {
	root := repoRootForTest(t)
	defToml := "[library]\nname = \"mymath\"\n\n[[functions]]\nname = \"add\"\nsymbol = \"add\"\nreturn = \"int\"\nargs = [\"int\", \"int\"]\n"
	tmpDir := t.TempDir()
	defPath := filepath.Join(tmpDir, "csig.ffi.def.toml")
	outPath := filepath.Join(tmpDir, "bindings.fig")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-output", outPath)
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generation failed: %v\n%s", err, out)
	}
	data, _ := os.ReadFile(outPath)
	if !strings.Contains(string(data), "# C: int add(int a, int b)") {
		t.Errorf("expected C signature comment, got:\n%s", data)
	}
}

func TestFfiGenCheckHelperMissing(t *testing.T) {
	root := repoRootForTest(t)
	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-check-helper", "/nonexistent/ffi-helper")
	cmd.Dir = root
	_, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("--check-helper should fail for non-existent binary")
	}
}

func TestFfiGenCheckHelperExists(t *testing.T) {
	root := repoRootForTest(t)
	fakeHelper := filepath.Join(t.TempDir(), "fake-helper")
	os.WriteFile(fakeHelper, []byte("#!/bin/sh\n"), 0755)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-check-helper", fakeHelper)
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("--check-helper should succeed, got: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "Helper binary OK") {
		t.Errorf("expected 'Helper binary OK', got: %s", out)
	}
}

func TestFfiGenFigTomlValidation(t *testing.T) {
	root := repoRootForTest(t)
	tmpDir := t.TempDir()

	tomlContent := "[ffi]\nenabled = false\nhelper = \"\"\n"
	tomlPath := filepath.Join(tmpDir, "fig.toml")
	os.WriteFile(tomlPath, []byte(tomlContent), 0644)

	defToml := "[library]\nname = \"test\"\n\n[[functions]]\nname = \"noop\"\nreturn = \"void\"\nargs = []\n"
	defPath := filepath.Join(tmpDir, "test.ffi.def.toml")
	os.WriteFile(defPath, []byte(defToml), 0644)

	cmd := exec.Command("go", "run", "./tools/ffi-gen", "-input", defPath, "-fig-toml", tomlPath)
	cmd.Dir = root
	out, err := cmd.CombinedOutput()
	// fig.toml issues are warnings, generation should still succeed
	if err != nil {
		t.Fatalf("expected success (warnings only), got: %v\n%s", err, out)
	}
	outStr := string(out)
	if !strings.Contains(outStr, "ffi.enabled is not true") {
		t.Errorf("expected warning about ffi.enabled, got: %s", outStr)
	}
}
