package tests

import "testing"

func TestImportNativeFunction(t *testing.T) {
	out, err := runFigSource(t, `import "native_mod"
print(native_mod.double(5));`)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10" {
		t.Fatalf("expected '10', got %q", out)
	}
}

func TestImportNativeVariable(t *testing.T) {
	out, err := runFigSource(t, `import "native_mod"
print(native_mod.exported_value);`)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "7" {
		t.Fatalf("expected '7', got %q", out)
	}
}
