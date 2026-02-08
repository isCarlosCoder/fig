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

func TestInterpreterLetAndPrint(t *testing.T) {
	input := "let x = 10; let y = x + 5; print(y);"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out := strings.TrimSpace(buf.String())
	if out != "15" {
		t.Fatalf("expected 15 got %q", out)
	}
}

func TestInterpreterAssignAndPrint(t *testing.T) {
	input := "let a = 2; a = a * 3; print(a);"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out := strings.TrimSpace(buf.String())
	if out != "6" {
		t.Fatalf("expected 6 got %q", out)
	}
}

func TestLogicalOperatorsAndShortCircuit(t *testing.T) {
	// basic && and ||
	input := "let a = true && false; print(a);"
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

	input = "let b = true || notdefined; print(b);"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	// should not panic due to short-circuit
	tree.Accept(v)
	out = strings.TrimSpace(buf.String())
	if out != "true" {
		t.Fatalf("expected true got %q", out)
	}

	input = "let c = false && notdefined; print(c);"
	is = antlr.NewInputStream(input)
	lex = parser.NewFigLexer(is)
	ts = antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p = parser.NewFigParser(ts)
	tree = p.Program()
	buf.Reset()
	v = interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	// should not panic due to short-circuit
	tree.Accept(v)
	out = strings.TrimSpace(buf.String())
	if out != "false" {
		t.Fatalf("expected false got %q", out)
	}
}

func TestModuloOperator(t *testing.T) {
	input := "let m = 10 % 3; print(m);"
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
}
