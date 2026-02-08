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

func TestUnaryNotOperator(t *testing.T) {
	// !true -> false
	input := "let a = !true; print(a);"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out := strings.TrimSpace(buf.String())
	if out != "false" {
		t.Fatalf("expected false got %q", out)
	}

	// !!false -> false (double negation)
	input = "let b = !!false; print(b);"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out = strings.TrimSpace(buf.String())
	if out != "false" {
		t.Fatalf("expected false got %q", out)
	}

	// ! on number (truthiness): numbers are truthy by default -> !0 == false
	input = "let c = !0; print(c);"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out = strings.TrimSpace(buf.String())
	if out != "false" {
		t.Fatalf("expected false got %q", out)
	}

	// use in if condition
	input = "if (!false) { print(1); }"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out = strings.TrimSpace(buf.String())
	if out != "1" {
		t.Fatalf("expected 1 got %q", out)
	}
}
