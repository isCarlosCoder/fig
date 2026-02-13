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
			// ExitSignal is a controlled exit from system.exit(); propagate as error, don't print.
			if es, ok := r.(environment.ExitSignal); ok {
				err = es
				return
			}
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
	listener.AbortOnError = true // stop parsing on first error for runs
	lex.RemoveErrorListeners()
	lex.AddErrorListener(listener)
	p.AddErrorListener(listener)

	// Parse — abort early on first error from the pretty listener
	var tree antlr.ParseTree
	func() {
		defer func() {
			if r := recover(); r != nil {
				// If the pretty listener caused the panic, surface its error
				if _, ok := r.(error); ok && listener != nil && listener.Occurred {
					// swallow the panic and return via listener.Err
					return
				}
				// otherwise re-panic
				panic(r)
			}
		}()
		tree = p.Program()
	}()
	if listener.Occurred {
		return listener.Err
	}

	// Execute with visitor; collect runtime errors without crashing
	v := NewFigVisitorWithSource(global, out, source)
	v.baseDir = filepath.Dir(filename)
	v.currentFile = filename
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

// RunInEnv parses `source` and executes it using `env` as the *exact* execution
// environment (top-level definitions are stored directly on `env`). This is
// useful for REPL preloading where callers expect loaded symbols to persist.
func RunInEnv(source, filename string, env *environment.Env, out io.Writer, errOut io.Writer) (err error) {
	// Reuse the same parsing + error-listener logic as Run, but attach the
	// visitor to use the provided env directly.
	defer func() {
		if r := recover(); r != nil {
			if es, ok := r.(environment.ExitSignal); ok {
				err = es
				return
			}
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

	p.RemoveErrorListeners()
	listener := NewPrettyErrorListener(source, filename, errOut)
	listener.AbortOnError = true
	lex.RemoveErrorListeners()
	lex.AddErrorListener(listener)
	p.AddErrorListener(listener)

	var tree antlr.ParseTree
	func() {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(error); ok && listener != nil && listener.Occurred {
					return
				}
				panic(r)
			}
		}()
		tree = p.Program()
	}()
	if listener.Occurred {
		return listener.Err
	}

	// Create visitor but force its execution env to the provided env so top-level
	// declarations persist there.
	v := NewFigVisitorWithSource(nil, out, source)
	v.env = env
	v.global = env.Parent() // keep parent linkage if any
	v.baseDir = filepath.Dir(filename)
	v.currentFile = filename
	if projectToml, err := findProjectTomlFrom(v.baseDir); err == nil {
		v.projectRoot = filepath.Dir(projectToml)
	}
	// Instead of calling VisitProgram (which creates a fresh local env),
	// iterate top-level statements and evaluate them directly in the provided
	// `env` so declarations persist there.
	if progCtx, ok := tree.(*parser.ProgramContext); ok {
		for _, st := range progCtx.AllStatements() {
			v.VisitStatements(st.(*parser.StatementsContext))
			if v.RuntimeErr != nil {
				if re, ok := v.RuntimeErr.(*RuntimeError); ok && errOut != nil {
					fmt.Fprint(errOut, re.PrettyError())
				}
				return v.RuntimeErr
			}
		}
	} else {
		tree.Accept(v)
		if v.RuntimeErr != nil {
			if re, ok := v.RuntimeErr.(*RuntimeError); ok && errOut != nil {
				fmt.Fprint(errOut, re.PrettyError())
			}
			return v.RuntimeErr
		}
	}
	return nil
}
