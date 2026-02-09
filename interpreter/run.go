package interpreter

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/parser"
)

// Run parses and executes `source`.
// `out` is used for program output (print), `errOut` is used for parse/runtime error printing (pretty listener).
// Returns an error if parsing failed or a runtime error occurred.
func Run(source, filename string, global *environment.Env, out io.Writer, errOut io.Writer) (err error) {
	// safety net: recover from unexpected panics and produce a clean error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("internal error: %v", r)
			if errOut != nil {
				fmt.Fprintf(errOut, "\x1b[1;31minternal error:\x1b[0m %v\n", r)
			}
		}
	}()

	is := antlr.NewInputStream(source)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)

	// Attach pretty error listener to errOut (avoid mixing errors with program stdout)
	p.RemoveErrorListeners()
	listener := NewPrettyErrorListener(source, filename, errOut)
	p.AddErrorListener(listener)

	// Parse
	tree := p.Program()
	if listener.Occurred {
		return listener.Err
	}

	// Execute with visitor; collect runtime errors without crashing
	v := NewFigVisitorWithSource(global, out, source)
	v.baseDir = filepath.Dir(filename)
	if projectToml, err := findProjectTomlFrom(v.baseDir); err == nil {
		v.projectRoot = filepath.Dir(projectToml)
	}
	tree.Accept(v)
	if v.RuntimeErr != nil {
		// print a pretty runtime snippet to errOut as well
		if re, ok := v.RuntimeErr.(*RuntimeError); ok && errOut != nil {
			fmt.Fprint(errOut, re.PrettyError())
		}
		return v.RuntimeErr
	}
	return nil
}
