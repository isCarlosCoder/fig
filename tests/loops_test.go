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

func TestWhileLoopBasic(t *testing.T) {
	input := "let i = 0; while (i < 3) { print(i); i = i + 1; }"
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

func TestBreakAndContinue(t *testing.T) {
	input := "let i = 0; while (i < 6) { i = i + 1; if (i % 2 == 0) { continue; } if (i == 5) { break; } print(i); }"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out := strings.TrimSpace(buf.String())
	// prints odd numbers until break at 5: 1,3
	if out != "1\n3" {
		t.Fatalf("expected 1\n3 got %q", out)
	}
}

func TestDoWhileLoop(t *testing.T) {
	input := "let i = 0; do { print(i); i = i + 1; } while (i < 2);"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	var buf bytes.Buffer
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), &buf)
	tree.Accept(v)
	out := strings.TrimSpace(buf.String())
	if out != "0\n1" {
		t.Fatalf("expected 0\n1 got %q", out)
	}
}

func TestBreakOutsideLoopIsError(t *testing.T) {
	input := "break;"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), nil)
	tree.Accept(v)
	if v.RuntimeErr == nil {
		t.Fatalf("expected runtime error for break outside loop")
	}
	if !strings.Contains(v.RuntimeErr.Error(), "break") {
		t.Fatalf("unexpected runtime error: %v", v.RuntimeErr)
	}
}

func TestContinueOutsideLoopIsError(t *testing.T) {
	input := "continue;"
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	tree := p.Program()
	v := interpreter.NewFigVisitor(environment.NewEnv(nil), nil)
	tree.Accept(v)
	if v.RuntimeErr == nil {
		t.Fatalf("expected runtime error for continue outside loop")
	}
	if !strings.Contains(v.RuntimeErr.Error(), "continue") {
		t.Fatalf("unexpected runtime error: %v", v.RuntimeErr)
	}
}
