package tests

import (
	"strings"
	"testing"
)

func TestMathAbs(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.abs(-5)); print(math.abs(3));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "5\n3" {
		t.Fatalf("expected '5\\n3', got %q", out)
	}
}

func TestMathPow(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.pow(2, 10));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "1024" {
		t.Fatalf("expected '1024', got %q", out)
	}
}

func TestMathSqrt(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.sqrt(144));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "12" {
		t.Fatalf("expected '12', got %q", out)
	}
}

func TestMathCbrt(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.cbrt(27));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3" {
		t.Fatalf("expected '3', got %q", out)
	}
}

func TestMathFloor(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.floor(3.7));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3" {
		t.Fatalf("expected '3', got %q", out)
	}
}

func TestMathCeil(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.ceil(3.2));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "4" {
		t.Fatalf("expected '4', got %q", out)
	}
}

func TestMathRound(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.round(3.5)); print(math.round(3.4));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "4\n3" {
		t.Fatalf("expected '4\\n3', got %q", out)
	}
}

func TestMathMinMax(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.min(3, 7)); print(math.max(3, 7));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3\n7" {
		t.Fatalf("expected '3\\n7', got %q", out)
	}
}

func TestMathClamp(t *testing.T) {
	src := "use " + `"` + "math" + `"` + "; print(math.clamp(15, 0, 10)); print(math.clamp(-5, 0, 10)); print(math.clamp(5, 0, 10));"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10\n0\n5" {
		t.Fatalf("expected '10\\n0\\n5', got %q", out)
	}
}

func TestMathTrig(t *testing.T) {
	src := "use " + `"` + "math" + `"` + "; print(math.sin(0)); print(math.cos(0)); print(math.tan(0));"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "0\n1\n0" {
		t.Fatalf("expected '0\\n1\\n0', got %q", out)
	}
}

func TestMathLog(t *testing.T) {
	src := "use " + `"` + "math" + `"` + "; print(math.log(1)); print(math.log10(100));"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "0\n2" {
		t.Fatalf("expected '0\\n2', got %q", out)
	}
}

func TestMathExp(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.exp(0));")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "1" {
		t.Fatalf("expected '1', got %q", out)
	}
}

func TestMathConstants(t *testing.T) {
	out, err := runFig(t, "use "+`"`+"math"+`"`+"; print(math.PI); print(math.E);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %q", len(lines), out)
	}
	if !strings.HasPrefix(lines[0], "3.14159") {
		t.Fatalf("expected PI '3.14159...', got %q", lines[0])
	}
	if !strings.HasPrefix(lines[1], "2.71828") {
		t.Fatalf("expected E '2.71828...', got %q", lines[1])
	}
}

func TestMathRand(t *testing.T) {
	src := "use " + `"` + "math" + `"` + "; let r = math.rand(); print(r >= 0 && r < 1);"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true" {
		t.Fatalf("expected 'true', got %q", out)
	}
}

func TestMathRandInt(t *testing.T) {
	src := "use " + `"` + "math" + `"` + "; let r = math.randInt(1, 10); print(r >= 1 && r < 10);"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true" {
		t.Fatalf("expected 'true', got %q", out)
	}
}

func TestMathInExpression(t *testing.T) {
	src := "use " + `"` + "math" + `"` + "; let area = math.PI * math.pow(5, 2); print(math.round(area));"
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "79" {
		t.Fatalf("expected '79', got %q", out)
	}
}

func TestUseUnknownModule(t *testing.T) {
	_, err := runFig(t, "use "+`"`+"inexistente"+`"`)
	if err == nil {
		t.Fatal("expected error for unknown module")
	}
}

func TestMathWrongArgType(t *testing.T) {
	src := "use " + `"` + "math" + `"` + `; math.abs("text");`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for wrong arg type")
	}
}

func TestMathWrongArity(t *testing.T) {
	src := "use " + `"` + "math" + `"` + "; math.abs(1, 2);"
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected error for wrong arity")
	}
}
