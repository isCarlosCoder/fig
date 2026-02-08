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

func TestIfElifElse(t *testing.T) {
	// simple if true
	input := "if (true) { print(1); }"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out := strings.TrimSpace(buf.String())
	if out != "1" {
		t.Fatalf("expected 1 got %q", out)
	}

	// if false with else
	input = "if (false) { print(1); } else { print(2); }"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out = strings.TrimSpace(buf.String())
	if out != "2" {
		t.Fatalf("expected 2 got %q", out)
	}

	// elif chain
	input = "if (false) { print(1); } elif (true) { print(2); } else { print(3); }"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out = strings.TrimSpace(buf.String())
	if out != "2" {
		t.Fatalf("expected 2 got %q", out)
	}

	// variable defined in block should not leak to outer scope
	input = "if (true) { let x = 5; } print(x);"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	if v.RuntimeErr == nil {
		t.Fatalf("expected runtime error for undefined variable x")
	}
	if !strings.Contains(v.RuntimeErr.Error(), "variable 'x' not defined") {
		t.Fatalf("unexpected runtime error: %v", v.RuntimeErr)
	}
}
