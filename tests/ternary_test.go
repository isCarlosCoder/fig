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

func TestTernaryOperator(t *testing.T) {
	// simple ternary
	input := "print((1 > 0) ? 3 : 4)"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out := strings.TrimSpace(buf.String())
	if out != "3" {
		t.Fatalf("expected 3 got %q", out)
	}

	// nested / right-associative
	input = "print(false ? 1 : true ? 2 : 3)"
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

	// precedence: ternary has lower precedence than logical-or
	input = "print((1 > 0 || false) ? 9 : 8)"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out = strings.TrimSpace(buf.String())
	if out != "9" {
		t.Fatalf("expected 9 got %q", out)
	}
}
