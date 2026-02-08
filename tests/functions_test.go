package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/interpreter"
	"github.com/iscarloscoder/fig/parser"
)

func runFig(t *testing.T, input string) (string, error) {
	t.Helper()
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	return strings.TrimSpace(buf.String()), v.RuntimeErr
}

func TestNullLiteral(t *testing.T) {
	out, err := runFig(t, "let x = null; print(x);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "null" {
		t.Fatalf("expected 'null', got %q", out)
	}
}

func TestNullIsFalsy(t *testing.T) {
	out, err := runFig(t, `let x = null; if (x) { print("truthy"); } else { print("falsy"); }`)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "falsy" {
		t.Fatalf("expected 'falsy', got %q", out)
	}
}

func TestNullEquality(t *testing.T) {
	out, err := runFig(t, "print(null == null); let a = null; print(a == null);")
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "true\ntrue" {
		t.Fatalf("expected 'true\\ntrue', got %q", out)
	}
}

func TestSimpleFunction(t *testing.T) {
	src := `fn hello() { print("Hello"); } hello();`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "Hello" {
		t.Fatalf("expected 'Hello', got %q", out)
	}
}

func TestFunctionWithParams(t *testing.T) {
	src := `fn add(a, b) { return a + b; } print(add(3, 4));`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "7" {
		t.Fatalf("expected '7', got %q", out)
	}
}

func TestFunctionWithLocalVars(t *testing.T) {
	src := `fn square(x) { let y = x * x; return y; } print(square(5));`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "25" {
		t.Fatalf("expected '25', got %q", out)
	}
}

func TestFunctionWithConditional(t *testing.T) {
	src := `fn sign(n) { if (n > 0) { return 1; } elif (n < 0) { return -1; } else { return 0; } } print(sign(-3)); print(sign(5)); print(sign(0));`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "-1\n1\n0" {
		t.Fatalf("expected '-1\\n1\\n0', got %q", out)
	}
}

func TestFunctionWithLoop(t *testing.T) {
	src := `fn sumTo(n) { let s = 0; let i = 1; while (i <= n) { s = s + i; i++; } return s; } print(sumTo(5));`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "15" {
		t.Fatalf("expected '15', got %q", out)
	}
}

func TestRecursiveFunction(t *testing.T) {
	src := `fn fact(n) { if (n <= 1) { return 1; } return n * fact(n - 1); } print(fact(5));`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "120" {
		t.Fatalf("expected '120', got %q", out)
	}
}

func TestEarlyReturn(t *testing.T) {
	src := `fn findEven(n) { let i = 0; while (i <= n) { if (i % 2 == 0) { return i; } i++; } return -1; } print(findEven(7));`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "0" {
		t.Fatalf("expected '0', got %q", out)
	}
}

func TestFunctionNoReturnIsNull(t *testing.T) {
	src := `fn log(x) { print(x); } let result = log("teste"); print(result);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "teste\nnull" {
		t.Fatalf("expected 'teste\\nnull', got %q", out)
	}
}

func TestFunctionArityError(t *testing.T) {
	src := `fn add(a, b) { return a + b; } add(1);`
	_, err := runFig(t, src)
	if err == nil {
		t.Fatal("expected arity error")
	}
}

func TestFunctionScopeIsolation(t *testing.T) {
	src := `let x = 10; fn getX() { return x; } fn shadow() { let x = 99; return x; } print(getX()); print(shadow()); print(x);`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "10\n99\n10" {
		t.Fatalf("expected '10\\n99\\n10', got %q", out)
	}
}

func TestFunctionReturnInForLoop(t *testing.T) {
	src := `fn firstMul3() { for (let i = 1; i < 20; i++) { if (i % 3 == 0) { return i; } } return -1; } print(firstMul3());`
	out, err := runFig(t, src)
	if err != nil {
		t.Fatalf("runtime error: %v", err)
	}
	if out != "3" {
		t.Fatalf("expected '3', got %q", out)
	}
}
