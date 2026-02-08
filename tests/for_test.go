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

func TestForLoopBasic(t *testing.T) {
	input := "for (let i = 0; i < 3; i = i + 1) { print(i); }"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out := strings.TrimSpace(buf.String())
	if out != "0\n1\n2" {
		t.Fatalf("expected 0\n1\n2 got %q", out)
	}
}

func TestForScopeIsolation(t *testing.T) {
	input := "for (let x = 1; x < 2; x = x + 1) { let y = 5; } print(y);"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), nil)
	tree.Accept(v)
	if v.RuntimeErr == nil {
		t.Fatalf("expected runtime error for undefined variable y")
	}
	if !strings.Contains(v.RuntimeErr.Error(), "variable 'y' not defined") {
		t.Fatalf("unexpected runtime error: %v", v.RuntimeErr)
	}
}
