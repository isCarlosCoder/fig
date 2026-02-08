package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iscarloscoder/fig/interpreter"
	"github.com/iscarloscoder/fig/parser"
)

func TestPrettyErrorListenerShowsContext(t *testing.T) {
	input := "let 123 = 1;" // invalid: expected identifier after 'let'
	is := antlr.NewInputStream(input)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)
	p.RemoveErrorListeners()
	var buf bytes.Buffer
	listener := interpreter.NewPrettyErrorListener(input, "<test>", &buf)
	p.AddErrorListener(listener)

	p.Program() // should trigger SyntaxError

	out := buf.String()
	if !strings.Contains(out, "<test>") || !strings.Contains(out, "error:") {
		t.Fatalf("expected pretty error output, got: %q", out)
	}
	if !strings.Contains(out, "123") {
		t.Fatalf("expected snippet with offending token '123' in output: %q", out)
	}
	if !strings.Contains(out, "^") {
		t.Fatalf("expected caret marker in output: %q", out)
	}
}
