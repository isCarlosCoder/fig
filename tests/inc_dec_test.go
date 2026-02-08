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

func runProg(t *testing.T, input string) string {
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	if v.RuntimeErr != nil {
		t.Fatalf("runtime error: %v", v.RuntimeErr)
	}
	return strings.TrimSpace(buf.String())
}

func TestPostfixIncrement(t *testing.T) {
	out := runProg(t, "let n = 0; n++; print(n);")
	if out != "1" {
		t.Fatalf("expected 1, got %q", out)
	}
}

func TestPostfixReturnsOld(t *testing.T) {
	out := runProg(t, "let n = 0; print(n++); print(n);")
	if out != "0\n1" {
		t.Fatalf("expected '0\\n1', got %q", out)
	}
}

func TestPrefixReturnsNew(t *testing.T) {
	out := runProg(t, "let n = 0; print(++n); print(n);")
	if out != "1\n1" {
		t.Fatalf("expected '1\\n1', got %q", out)
	}
}

func TestForWithPostfixStep(t *testing.T) {
	out := runProg(t, "for (let i = 0; i < 3; i++) { print(i); } print(\"done\");")
	if out != "0\n1\n2\ndone" {
		t.Fatalf("expected '0\\n1\\n2\\ndone', got %q", out)
	}
}

func TestDecrement(t *testing.T) {
	out := runProg(t, "let x = 3; print(x--); print(--x); print(x);")
	if out != "3\n1\n1" {
		t.Fatalf("expected '3\\n1\\n1', got %q", out)
	}
}
