package interpreter

import (
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iscarloscoder/fig/builtins"
	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/parser"
	"github.com/pelletier/go-toml/v2"
)

// FigVisitor evaluates the parse tree. It keeps a global environment and a current (local) environment stack.
type loopSignal int

const (
	loopNone loopSignal = iota
	loopBreak
	loopContinue
)

// returnSignal wraps a return value so it can be propagated up the call stack.
type returnSignal struct {
	value environment.Value
}

type FigVisitor struct {
	parser.BaseFigParserVisitor
	global *environment.Env
	env    *environment.Env // current environment
	out    io.Writer        // output for print

	// RuntimeErr records the first runtime error encountered during visiting.
	RuntimeErr error

	// source lines for creating snippet in runtime errors
	srcLines []string

	// currentFile holds the filename (absolute or relative) of the currently executing source
	currentFile string

	// loopDepth indicates how many nested loops we're currently in (to validate break/continue)
	loopDepth int

	// evalSteps counts expression evaluation steps to detect runaway recursion during debugging
	evalSteps int

	// pendingLoopSignal holds a break/continue signal from a try/onerror block
	// that needs to propagate to the enclosing loop.
	pendingLoopSignal interface{}

	// baseDir is the directory of the currently executing .fig file (for resolving imports)
	baseDir string

	// projectRoot is the directory that contains fig.toml for module resolution
	projectRoot string

	// callDepth tracks nested function call depth to detect runaway recursion during debugging
	callDepth int

	// frames stores a runtime call/import stack for better error diagnostics
	frames []StackFrame

	// importedFiles tracks already-imported absolute paths to prevent circular imports
	importedFiles map[string]bool

	// importedModules caches module objects by entry file path
	importedModules map[string]environment.Value
}

// syncWriter wraps an io.Writer with a mutex for goroutine-safe writes.
type syncWriter struct {
	mu sync.Mutex
	w  io.Writer
}

func (sw *syncWriter) Write(p []byte) (n int, err error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	return sw.w.Write(p)
}

func NewFigVisitor(globalEnv *environment.Env, out io.Writer) *FigVisitor {
	if out == nil {
		out = os.Stdout
	}
	v := &FigVisitor{
		global:          globalEnv,
		out:             out,
		importedFiles:   make(map[string]bool),
		importedModules: make(map[string]environment.Value),
	}
	// start with a local env whose parent is the global env
	v.env = environment.NewEnv(globalEnv)
	return v
}

// NewFigVisitorWithSource creates a visitor and attaches source lines so runtime
// errors can include snippets with carets.
func NewFigVisitorWithSource(globalEnv *environment.Env, out io.Writer, source string) *FigVisitor {
	v := NewFigVisitor(globalEnv, out)
	if source != "" {
		v.srcLines = strings.Split(source, "\n")
	}
	// currentFile will be set by the caller (Run) or by import handling when switching files
	v.currentFile = ""
	return v
}

func (v *FigVisitor) pushEnv() {
	v.env = environment.NewEnv(v.env)
}

func (v *FigVisitor) popEnv() {
	if v.env != nil {
		v.env = v.env.Parent()
	}
}

func (v *FigVisitor) pushFrame(kind, name, file string, line, column int) {
	v.frames = append(v.frames, StackFrame{Kind: kind, Name: name, File: file, Line: line, Column: column})
}

func (v *FigVisitor) popFrame() {
	if len(v.frames) > 0 {
		v.frames = v.frames[:len(v.frames)-1]
	}
}

// makeRuntimeError builds a RuntimeError with snippet information when source is available.
func (v *FigVisitor) makeRuntimeError(line, column int, msg string, length int) *RuntimeError {
	r := &RuntimeError{File: v.currentFile, Line: line, Column: column, Message: msg, ColumnStart: column, Length: length}
	// attach current frames so errors show a stack trace
	if len(v.frames) > 0 {
		r.Frames = append([]StackFrame(nil), v.frames...)
	}
	// prefer snippet from current source if available
	if v.srcLines != nil && line-1 >= 0 && line-1 < len(v.srcLines) {
		ln := v.srcLines[line-1]
		r.Snippet = ln
		if r.ColumnStart < 0 {
			r.ColumnStart = 0
		}
		if r.ColumnStart > len(ln) {
			r.ColumnStart = len(ln)
		}
		if r.ColumnStart+r.Length > len(ln) {
			r.Length = len(ln) - r.ColumnStart
		}
	}

	// If there's no snippet but we have frames, use the most-recent frame to provide context
	if r.Snippet == "" && len(r.Frames) > 0 {
		f := r.Frames[len(r.Frames)-1]
		if f.File != "" && f.Line > 0 {
			if data, err := os.ReadFile(f.File); err == nil {
				lines := strings.Split(string(data), "\n")
				if f.Line-1 >= 0 && f.Line-1 < len(lines) {
					ln := lines[f.Line-1]
					r.Snippet = ln
					// use frame column if available
					r.ColumnStart = f.Column
					if r.ColumnStart < 0 {
						r.ColumnStart = 0
					}
					if r.ColumnStart > len(ln) {
						r.ColumnStart = len(ln)
					}
					if r.ColumnStart+r.Length > len(ln) {
						r.Length = len(ln) - r.ColumnStart
					}
					// reflect frame location as the reported origin
					r.File = f.File
					r.Line = f.Line
					r.Column = f.Column
				}
			}
		}
	}

	return r
}

func (v *FigVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	// run in a fresh local environment that inherits from global
	prev := v.env
	v.env = environment.NewEnv(v.global)
	defer func() { v.env = prev }()

	for _, st := range ctx.AllStatements() {
		v.VisitStatements(st.(*parser.StatementsContext))
		if v.RuntimeErr != nil {
			return nil
		}
	}
	return nil
}

func (v *FigVisitor) VisitStatements(ctx *parser.StatementsContext) interface{} {
	if ctx.VarDeclaration() != nil {
		return v.VisitVarDeclaration(ctx.VarDeclaration().(*parser.VarDeclarationContext))
	}
	if ctx.VarAtribuition() != nil {
		return v.VisitVarAtribuition(ctx.VarAtribuition().(*parser.VarAtribuitionContext))
	}
	if ctx.PrintStmt() != nil {
		return v.VisitPrintStmt(ctx.PrintStmt().(*parser.PrintStmtContext))
	}
	if ctx.IfStmt() != nil {
		return v.VisitIfStmt(ctx.IfStmt().(*parser.IfStmtContext))
	}
	if ctx.WhileStmt() != nil {
		return v.VisitWhileStmt(ctx.WhileStmt().(*parser.WhileStmtContext))
	}
	if ctx.DoWhileStmt() != nil {
		return v.VisitDoWhileStmt(ctx.DoWhileStmt().(*parser.DoWhileStmtContext))
	}
	if ctx.ForStmt() != nil {
		return v.VisitForStmt(ctx.ForStmt().(*parser.ForStmtContext))
	}
	if ctx.ForInStmt() != nil {
		return ctx.ForInStmt().Accept(v)
	}
	if ctx.BreakStmt() != nil {
		return v.VisitBreakStmt(ctx.BreakStmt().(*parser.BreakStmtContext))
	}
	if ctx.ContinueStmt() != nil {
		return v.VisitContinueStmt(ctx.ContinueStmt().(*parser.ContinueStmtContext))
	}
	if ctx.FnDecl() != nil {
		return v.VisitFnDecl(ctx.FnDecl().(*parser.FnDeclContext))
	}
	if ctx.ReturnStmt() != nil {
		return v.VisitReturnStmt(ctx.ReturnStmt().(*parser.ReturnStmtContext))
	}
	if ctx.MemberAssign() != nil {
		return v.VisitMemberAssign(ctx.MemberAssign().(*parser.MemberAssignContext))
	}
	if ctx.ImportStmt() != nil {
		return v.VisitImportStmt(ctx.ImportStmt().(*parser.ImportStmtContext))
	}
	if ctx.UseStmt() != nil {
		return v.VisitUseStmt(ctx.UseStmt().(*parser.UseStmtContext))
	}
	if ctx.StructDecl() != nil {
		return v.VisitStructDecl(ctx.StructDecl().(*parser.StructDeclContext))
	}
	if ctx.EnumDecl() != nil {
		return v.VisitEnumDecl(ctx.EnumDecl().(*parser.EnumDeclContext))
	}
	if ctx.ExprStmt() != nil {
		return v.VisitExprStmt(ctx.ExprStmt().(*parser.ExprStmtContext))
	}
	return nil
}

func (v *FigVisitor) VisitUseStmt(ctx *parser.UseStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}

	raw := ctx.STRING().GetText()
	modName := raw[1 : len(raw)-1]

	mod := builtins.Get(modName)
	if mod == nil {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("unknown builtin module \"%s\"", modName), len(raw))
		return nil
	}

	// Define the module as an object variable with the module name
	v.env.Define(modName, mod.ToObject())

	// If loading the http module, set the FigCaller so HTTP server handlers
	// can invoke Fig functions via the visitor.
	if modName == "http" {
		builtins.FigCaller = func(fn environment.Value, args []environment.Value) error {
			savedErr := v.RuntimeErr
			v.RuntimeErr = nil
			v.callFunction(0, 0, fn, args, "")
			handlerErr := v.RuntimeErr
			v.RuntimeErr = savedErr
			return handlerErr
		}
	}

	// If loading a module that needs to invoke user-defined Fig functions,
	// set FnCaller so builtin helpers can call them (not just builtins).
	if modName == "functional" || modName == "arrays" {
		builtins.FnCaller = func(fn environment.Value, args []environment.Value) (environment.Value, error) {
			savedErr := v.RuntimeErr
			v.RuntimeErr = nil
			result := v.callFunction(0, 0, fn, args, "")
			if v.RuntimeErr != nil {
				err := v.RuntimeErr
				v.RuntimeErr = savedErr
				if re, ok := err.(*RuntimeError); ok {
					return environment.NewNil(), fmt.Errorf("%s", re.Message)
				}
				return environment.NewNil(), err
			}
			v.RuntimeErr = savedErr
			if val, ok := result.(environment.Value); ok {
				return val, nil
			}
			return environment.NewNil(), nil
		}
	}

	// If loading the figtest module, set FigtestCaller so test functions
	// can invoke user-defined Fig callbacks and capture assertion errors.
	if modName == "figtest" {
		builtins.FigtestCaller = func(fn environment.Value, args []environment.Value) (environment.Value, error) {
			savedErr := v.RuntimeErr
			v.RuntimeErr = nil
			result := v.callFunction(0, 0, fn, args, "")
			if v.RuntimeErr != nil {
				err := v.RuntimeErr
				v.RuntimeErr = savedErr
				if re, ok := err.(*RuntimeError); ok {
					return environment.NewNil(), fmt.Errorf("%s", re.Message)
				}
				return environment.NewNil(), err
			}
			v.RuntimeErr = savedErr
			if val, ok := result.(environment.Value); ok {
				return val, nil
			}
			return environment.NewNil(), nil
		}
	}

	// If loading the task module, set TaskSpawner so task.spawn can execute
	// user-defined Fig functions in separate goroutines.
	if modName == "task" {
		// Wrap v.out in a syncWriter so goroutine prints are safe.
		sw := &syncWriter{w: v.out}
		v.out = sw
		builtins.TaskSpawner = func(fn environment.Value, resultCh chan<- builtins.TaskResult) {
			go func() {
				newV := NewFigVisitor(environment.NewEnv(nil), sw)
				result := newV.callFunction(0, 0, fn, nil, "")
				if newV.RuntimeErr != nil {
					if re, ok := newV.RuntimeErr.(*RuntimeError); ok {
						resultCh <- builtins.TaskResult{Err: fmt.Errorf("%s", re.Message)}
					} else {
						resultCh <- builtins.TaskResult{Err: newV.RuntimeErr}
					}
				} else if val, ok := result.(environment.Value); ok {
					resultCh <- builtins.TaskResult{Value: val}
				} else {
					resultCh <- builtins.TaskResult{Value: environment.NewNil()}
				}
			}()
		}
	}

	return nil
}

// ── struct declaration ──
func (v *FigVisitor) VisitStructDecl(ctx *parser.StructDeclContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}

	name := ctx.ID().GetText()
	sd := &environment.StructDef{
		Name:    name,
		Methods: make(map[string]*environment.FuncDef),
	}

	for _, member := range ctx.AllStructMember() {
		switch m := member.(type) {
		case *parser.StructFieldContext:
			fieldName := m.ID().GetText()
			defaultVal := environment.NewNil()
			if m.Expr() != nil {
				defaultVal = v.VisitExpr(m.Expr().(*parser.ExprContext)).(environment.Value)
				if v.RuntimeErr != nil {
					return nil
				}
			}
			sd.Fields = append(sd.Fields, environment.StructField{Name: fieldName, Default: defaultVal})

		case *parser.StructMethodContext:
			methodName := m.ID().GetText()
			var params []environment.Param
			if fp := m.FnParams(); fp != nil {
				for _, pd := range fp.(*parser.FnParamsContext).AllParamDecl() {
					switch pctx := pd.(type) {
					case *parser.ParamWithDefaultOrRequiredContext:
						param := environment.Param{Name: pctx.ID().GetText()}
						if pctx.ASSIGN() != nil {
							param.HasDefault = true
							param.Default = pctx.Expr()
						}
						params = append(params, param)
					case *parser.ParamOptionalContext:
						param := environment.Param{Name: pctx.ID().GetText(), Optional: true}
						params = append(params, param)
					}
				}
			}
			fd := &environment.FuncDef{
				Name:       methodName,
				Params:     params,
				Body:       m.Block(),
				ClosureEnv: v.env,
			}
			sd.Methods[methodName] = fd
		}
	}

	if err := v.env.Define(name, environment.NewStructDef(sd)); err != nil {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(name))
	}
	return nil
}

// enum declaration
func (v *FigVisitor) VisitEnumDecl(ctx *parser.EnumDeclContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	name := ctx.ID().GetText()
	entries := make(map[string]environment.Value)
	var keys []string
	for i, m := range ctx.AllEnumMember() {
		em := m.(*parser.EnumMemberContext)
		mname := em.ID().GetText()
		val := environment.NewEnumMember(name, mname, i)
		entries[mname] = val
		keys = append(keys, mname)
	}
	// Define enum name as an object mapping member names to enum values
	if err := v.env.Define(name, environment.NewObject(entries, keys)); err != nil {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(name))
	}
	return nil
}

func (v *FigVisitor) VisitImportStmt(ctx *parser.ImportStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}

	// Get the string literal and strip quotes
	raw := ctx.STRING().GetText()
	modPath := raw[1 : len(raw)-1]
	alias := ""
	if ctx.ID() != nil {
		alias = ctx.ID().GetText()
	}

	if strings.HasPrefix(modPath, "mod:") {
		if err := v.importModule(modPath, alias, ctx); err != nil {
			v.RuntimeErr = err
		}
		return nil
	}

	// Allow optional alias for local imports; if not provided derive from filename
	var localAlias string
	if ctx.ID() != nil {
		localAlias = ctx.ID().GetText()
	} else {
		// derive base name without extension
		localAlias = strings.TrimSuffix(filepath.Base(modPath), filepath.Ext(modPath))
	}

	// Append .fig extension if not present
	if !strings.HasSuffix(modPath, ".fig") {
		modPath = modPath + ".fig"
	}

	// Build a list of candidate paths (in priority order) and pick the first that exists.
	var candidates []string
	// Absolute path: use directly
	if filepath.IsAbs(modPath) {
		candidates = append(candidates, modPath)
	} else if strings.HasPrefix(modPath, "./") || strings.HasPrefix(modPath, "../") {
		// Explicit relative paths are always relative to the importing file's directory
		candidates = append(candidates, filepath.Join(v.baseDir, modPath))
	} else {
		// 1) relative to the directory of the importing file
		candidates = append(candidates, filepath.Join(v.baseDir, modPath))

		// 2) if projectRoot is known, try projectRoot/src/<modPath> for plain names
		if v.projectRoot != "" {
			if !strings.Contains(modPath, string(os.PathSeparator)) {
				candidates = append(candidates, filepath.Join(v.projectRoot, "src", modPath))
			}
			// 3) try projectRoot/<modPath>
			candidates = append(candidates, filepath.Join(v.projectRoot, modPath))
		}
	}

	// Clean candidates and pick first that exists
	var resolved string
	for _, c := range candidates {
		c = filepath.Clean(c)
		if _, statErr := os.Stat(c); statErr == nil {
			resolved = c
			break
		}
	}

	// If none of the candidates exist, still set resolved to the first candidate so
	// the error message references a sensible path.
	if resolved == "" && len(candidates) > 0 {
		resolved = filepath.Clean(candidates[0])
	}

	// Absolute path for dedup
	absPath, err := filepath.Abs(resolved)
	if err != nil {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("cannot resolve import path: %s", modPath), len(raw))
		return nil
	}

	// Check for circular imports
	if v.importedFiles[absPath] {
		// Already imported — skip silently
		return nil
	}
	v.importedFiles[absPath] = true

	// Read the file
	data, readErr := os.ReadFile(absPath)
	if readErr != nil {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("cannot import \"%s\": %v", modPath, readErr), len(raw))
		return nil
	}

	source := string(data)

	// Parse the imported file
	is := antlr.NewInputStream(source)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)

	// Capture parse/lex errors. Attach listener to both lexer and parser and
	// abort parsing on the first reported error.
	lex.RemoveErrorListeners()
	p.RemoveErrorListeners()
	listener := NewPrettyErrorListener(source, absPath, v.out)
	listener.AbortOnError = true
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
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("parse error in import \"%s\"", modPath), len(raw))
		return nil
	}

	// Execute the imported file in its own module environment and expose it as an object
	// Save and restore visitor state
	prevSrcLines := v.srcLines
	prevBaseDir := v.baseDir
	prevEnv := v.env
	prevCurrentFile := v.currentFile

	moduleEnv := environment.NewEnv(v.env)
	v.env = moduleEnv
	v.srcLines = strings.Split(source, "\n")
	v.baseDir = filepath.Dir(absPath)
	v.currentFile = absPath

	progCtx := tree.(*parser.ProgramContext)
	// push a module frame so runtime errors show the imported file in the stack
	v.pushFrame("module", filepath.Base(absPath), absPath, 1, 0)
	for _, st := range progCtx.AllStatements() {
		v.VisitStatements(st.(*parser.StatementsContext))
		if v.RuntimeErr != nil {
			break
		}
	}
	v.popFrame()

	// Restore state
	v.srcLines = prevSrcLines
	v.baseDir = prevBaseDir
	v.env = prevEnv
	v.currentFile = prevCurrentFile

	if v.RuntimeErr != nil {
		return nil
	}

	entries, keys := moduleEnv.Snapshot()
	moduleObj := environment.NewObject(entries, keys)
	// define using localAlias (derived above)
	if err := v.env.Define(localAlias, moduleObj); err != nil {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(localAlias))
		return nil
	}

	return nil
}

func (v *FigVisitor) importModule(modPath, alias string, ctx *parser.ImportStmtContext) error {
	moduleSpec := strings.TrimPrefix(modPath, "mod:")
	if moduleSpec == "" {
		return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			"import mod: requires a module path", len(modPath))
	}

	projectRoot := v.projectRoot
	if projectRoot == "" {
		projectToml, err := findProjectTomlFrom(v.baseDir)
		if err != nil {
			return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				"cannot locate fig.toml for module import", len(modPath))
		}
		projectRoot = filepath.Dir(projectToml)
		v.projectRoot = projectRoot
	}
	moduleName := filepath.Base(moduleSpec)
	if alias == "" {
		alias = moduleName
	}
	moduleDir := filepath.Join(projectRoot, "_modules", moduleName)
	// If the module directory is not present, try resolving by alias from the project's fig.toml
	if _, statErr := os.Stat(moduleDir); os.IsNotExist(statErr) {
		projectTomlPath := filepath.Join(projectRoot, "fig.toml")
		deps, err := loadProjectDeps(projectTomlPath)
		if err != nil {
			return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				fmt.Sprintf("cannot read project fig.toml: %v", err), len(modPath))
		}
		found := false
		for _, dep := range deps {
			if dep.Alias == moduleSpec || filepath.Base(dep.Location) == moduleSpec {
				if dep.Location != "" {
					moduleName = filepath.Base(dep.Location)
				} else if dep.Source != "" {
					parts := strings.Split(dep.Source, "/")
					moduleName = parts[len(parts)-1]
				}
				moduleDir = filepath.Join(projectRoot, "_modules", moduleName)
				found = true
				break
			}
		}
		if !found {
			return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				fmt.Sprintf("module not installed: %s", moduleSpec), len(modPath))
		}
	}
	moduleTomlPath := filepath.Join(moduleDir, "fig.toml")
	modCfg, err := loadModuleToml(moduleTomlPath)
	if err != nil {
		return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("cannot read module fig.toml: %v", err), len(modPath))
	}

	entry := modCfg.Project.Main
	if entry == "" {
		entry = "src/main.fig"
	}
	entryPath := filepath.Join(moduleDir, entry)
	absPath, err := filepath.Abs(entryPath)
	if err != nil {
		return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("cannot resolve module entry: %s", entryPath), len(modPath))
	}

	if cached, ok := v.importedModules[absPath]; ok {
		if err := v.env.Define(alias, cached); err != nil {
			return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(alias))
		}
		return nil
	}

	if v.importedFiles[absPath] {
		return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("circular module import: %s", moduleSpec), len(modPath))
	}
	v.importedFiles[absPath] = true

	data, readErr := os.ReadFile(absPath)
	if readErr != nil {
		return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("cannot import module \"%s\": %v", moduleSpec, readErr), len(modPath))
	}

	source := string(data)
	is := antlr.NewInputStream(source)
	lex := parser.NewFigLexer(is)
	ts := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewFigParser(ts)

	// attach listener to lexer and parser and abort early on first error
	lex.RemoveErrorListeners()
	p.RemoveErrorListeners()
	listener := NewPrettyErrorListener(source, absPath, v.out)
	listener.AbortOnError = true
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
		return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			fmt.Sprintf("parse error in module \"%s\"", moduleSpec), len(modPath))
	}

	prevSrcLines := v.srcLines
	prevBaseDir := v.baseDir
	prevEnv := v.env
	prevCurrentFile := v.currentFile

	moduleEnv := environment.NewEnv(v.env)
	v.env = moduleEnv
	v.srcLines = strings.Split(source, "\n")
	v.baseDir = filepath.Dir(absPath)
	v.currentFile = absPath

	progCtx := tree.(*parser.ProgramContext)
	// push a module frame so runtime errors show the module file in the stack
	v.pushFrame("module", filepath.Base(absPath), absPath, 1, 0)
	for _, st := range progCtx.AllStatements() {
		v.VisitStatements(st.(*parser.StatementsContext))
		if v.RuntimeErr != nil {
			break
		}
	}
	v.popFrame()

	v.srcLines = prevSrcLines
	v.baseDir = prevBaseDir
	v.env = prevEnv
	v.currentFile = prevCurrentFile

	if v.RuntimeErr != nil {
		return v.RuntimeErr
	}

	entries, keys := moduleEnv.Snapshot()
	moduleObj := environment.NewObject(entries, keys)
	v.importedModules[absPath] = moduleObj
	if err := v.env.Define(alias, moduleObj); err != nil {
		return v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(alias))
	}
	return nil
}

type moduleTomlConfig struct {
	Project moduleProjectConfig `toml:"project"`
}

type moduleProjectConfig struct {
	Name string `toml:"name"`
	Main string `toml:"main"`
}

func loadModuleToml(path string) (moduleTomlConfig, error) {
	var cfg moduleTomlConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func findProjectTomlFrom(startDir string) (string, error) {
	dir := startDir
	for {
		candidate := filepath.Join(dir, "fig.toml")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("fig.toml not found")
		}
		dir = parent
	}
}

// projectDependency represents a single dependency entry in a project's fig.toml.
// This mirrors the structure used by the CLI but is scoped to the interpreter package.
type projectDependency struct {
	Version  string `toml:"version"`
	Source   string `toml:"source"`
	Location string `toml:"location"`
	Alias    string `toml:"alias,omitempty"`
}

func loadProjectDeps(path string) ([]projectDependency, error) {
	var cfg struct {
		Deps map[string]projectDependency `toml:"dependencies"`
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	deps := make([]projectDependency, 0, len(cfg.Deps))
	for name, dep := range cfg.Deps {
		if dep.Alias == "" {
			dep.Alias = name
		}
		deps = append(deps, dep)
	}
	return deps, nil
}

func (v *FigVisitor) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	v.VisitExpr(ctx.Expr().(*parser.ExprContext))
	if v.RuntimeErr != nil {
		return nil
	}
	if v.pendingLoopSignal != nil {
		sig := v.pendingLoopSignal
		v.pendingLoopSignal = nil
		return sig
	}
	return nil
}

func (v *FigVisitor) VisitVarDeclaration(ctx *parser.VarDeclarationContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	name := ctx.ID().GetText()
	if ctx.Expr() != nil {
		val := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return nil
		}
		if v.pendingLoopSignal != nil {
			sig := v.pendingLoopSignal
			v.pendingLoopSignal = nil
			return sig
		}
		if err := v.env.Define(name, val); err != nil {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
			return nil
		}
	} else {
		if err := v.env.Define(name, environment.NewNil()); err != nil {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
			return nil
		}
	}
	return nil
}

func (v *FigVisitor) VisitVarAtribuition(ctx *parser.VarAtribuitionContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	name := ctx.ID().GetText()
	val := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return nil
	}
	if v.pendingLoopSignal != nil {
		sig := v.pendingLoopSignal
		v.pendingLoopSignal = nil
		return sig
	}
	if err := v.env.Assign(name, val); err != nil {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
		return nil
	}
	return nil
}

func (v *FigVisitor) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	exprs := ctx.AllExpr()
	parts := make([]string, 0, len(exprs))
	for _, e := range exprs {
		val := v.VisitExpr(e.(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return nil
		}
		if v.pendingLoopSignal != nil {
			sig := v.pendingLoopSignal
			v.pendingLoopSignal = nil
			return sig
		}
		parts = append(parts, val.String())
	}
	fmt.Fprintln(v.out, strings.Join(parts, " "))
	return nil
}

// Expression evaluation follows the grammar hierarchy.
func (v *FigVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	if ctx.LogicalOr() != nil {
		return v.VisitLogicalOr(ctx.LogicalOr().(*parser.LogicalOrContext))
	}
	// safety fallback — should never happen with a valid grammar
	result := v.VisitChildren(ctx)
	if result == nil {
		return environment.NewNil()
	}
	return result
}

func (v *FigVisitor) VisitLogicalAnd(ctx *parser.LogicalAndContext) interface{} {
	// logicalAnd: equality ( AND equality )* ;
	children := ctx.GetChildren()
	var seq []antlr.Tree
	for _, c := range children {
		switch c.(type) {
		case parser.IEqualityContext, antlr.TerminalNode:
			seq = append(seq, c)
		}
	}
	if len(seq) == 1 {
		return v.VisitEquality(seq[0].(parser.IEqualityContext).(*parser.EqualityContext))
	}
	left := v.VisitEquality(seq[0].(parser.IEqualityContext).(*parser.EqualityContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}
	for i := 1; i < len(seq); i += 2 {
		op := seq[i].(antlr.TerminalNode)
		// short-circuit: if left is falsey, return false without evaluating right
		if !left.IsTruthy() && op.GetSymbol().GetTokenType() == parser.FigParserAND {
			return environment.NewBool(false)
		}
		rightCtx := seq[i+1].(parser.IEqualityContext).(*parser.EqualityContext)
		right := v.VisitEquality(rightCtx).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		left = environment.NewBool(left.IsTruthy() && right.IsTruthy())
	}
	return left
}

func (v *FigVisitor) VisitLogicalOr(ctx *parser.LogicalOrContext) interface{} {
	// logicalOr: logicalAnd ( OR logicalAnd )* ;
	children := ctx.GetChildren()
	var seq []antlr.Tree
	for _, c := range children {
		switch c.(type) {
		case parser.ILogicalAndContext, antlr.TerminalNode:
			seq = append(seq, c)
		}
	}
	if len(seq) == 1 {
		return v.VisitLogicalAnd(seq[0].(parser.ILogicalAndContext).(*parser.LogicalAndContext))
	}
	left := v.VisitLogicalAnd(seq[0].(parser.ILogicalAndContext).(*parser.LogicalAndContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}
	for i := 1; i < len(seq); i += 2 {
		tn := seq[i].(antlr.TerminalNode)
		tok := tn.GetSymbol()
		// short-circuit: if left is truthy, return true without evaluating right
		if left.IsTruthy() && tok.GetTokenType() == parser.FigParserOR {
			return environment.NewBool(true)
		}
		rightCtx := seq[i+1].(parser.ILogicalAndContext).(*parser.LogicalAndContext)
		right := v.VisitLogicalAnd(rightCtx).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		left = environment.NewBool(left.IsTruthy() || right.IsTruthy())
	}
	return left
}

func (v *FigVisitor) VisitEquality(ctx *parser.EqualityContext) interface{} {
	// Build alternating sequence of comparison and operator nodes
	children := ctx.GetChildren()
	var seq []antlr.Tree
	for _, c := range children {
		switch c.(type) {
		case parser.IComparisonContext, antlr.TerminalNode:
			seq = append(seq, c)
		}
	}

	if len(seq) == 1 {
		return v.VisitComparison(seq[0].(parser.IComparisonContext).(*parser.ComparisonContext))
	}

	left := v.VisitComparison(seq[0].(parser.IComparisonContext).(*parser.ComparisonContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}
	for i := 1; i < len(seq); i += 2 {
		tn := seq[i].(antlr.TerminalNode)
		tok := tn.GetSymbol()
		rightCtx := seq[i+1].(parser.IComparisonContext).(*parser.ComparisonContext)
		right := v.VisitComparison(rightCtx).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		switch tok.GetTokenType() {
		case parser.FigParserEQ:
			left = environment.NewBool(valuesEqual(left, right))
		case parser.FigParserNEQ:
			left = environment.NewBool(!valuesEqual(left, right))
		default:
			v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), fmt.Sprintf("unsupported equality operator: %v", tok.GetTokenType()), len(tok.GetText()))
			return environment.NewNil()
		}
	}
	return left
}

func (v *FigVisitor) VisitComparison(ctx *parser.ComparisonContext) interface{} {
	children := ctx.GetChildren()
	var seq []antlr.Tree
	for _, c := range children {
		switch c.(type) {
		case parser.ITermContext, antlr.TerminalNode:
			seq = append(seq, c)
		}
	}
	if len(seq) == 1 {
		return v.VisitTerm(seq[0].(parser.ITermContext).(*parser.TermContext))
	}
	left := v.VisitTerm(seq[0].(parser.ITermContext).(*parser.TermContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}
	for i := 1; i < len(seq); i += 2 {
		tn := seq[i].(antlr.TerminalNode)
		tok := tn.GetSymbol()
		rightCtx := seq[i+1].(parser.ITermContext).(*parser.TermContext)
		right := v.VisitTerm(rightCtx).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		switch tok.GetTokenType() {
		case parser.FigParserGT:
			ln, err := left.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), err.Error(), len(tok.GetText()))
				return environment.NewNil()
			}
			rn, err := right.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), err.Error(), len(tok.GetText()))
				return environment.NewNil()
			}
			left = environment.NewBool(ln > rn)
		case parser.FigParserGE:
			ln, err := left.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), err.Error(), len(tok.GetText()))
				return environment.NewNil()
			}
			rn, err := right.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), err.Error(), len(tok.GetText()))
				return environment.NewNil()
			}
			left = environment.NewBool(ln >= rn)
		case parser.FigParserLT:
			ln, err := left.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), err.Error(), len(tok.GetText()))
				return environment.NewNil()
			}
			rn, err := right.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), err.Error(), len(tok.GetText()))
				return environment.NewNil()
			}
			left = environment.NewBool(ln < rn)
		case parser.FigParserLE:
			ln, err := left.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), err.Error(), len(tok.GetText()))
				return environment.NewNil()
			}
			rn, err := right.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), err.Error(), len(tok.GetText()))
				return environment.NewNil()
			}
			left = environment.NewBool(ln <= rn)
		default:
			v.RuntimeErr = v.makeRuntimeError(tok.GetLine(), tok.GetColumn(), fmt.Sprintf("unsupported comparison operator: %v", tok.GetTokenType()), 1)
			return environment.NewNil()
		}
	}
	return left
}

func (v *FigVisitor) VisitTerm(ctx *parser.TermContext) interface{} {
	children := ctx.GetChildren()
	var seq []antlr.Tree
	for _, c := range children {
		switch c.(type) {
		case parser.IFactorContext, antlr.TerminalNode:
			seq = append(seq, c)
		}
	}
	if len(seq) == 1 {
		return v.VisitFactor(seq[0].(parser.IFactorContext).(*parser.FactorContext))
	}
	left := v.VisitFactor(seq[0].(parser.IFactorContext).(*parser.FactorContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}
	for i := 1; i < len(seq); i += 2 {
		op := seq[i].(antlr.TerminalNode)
		rightCtx := seq[i+1].(parser.IFactorContext).(*parser.FactorContext)
		right := v.VisitFactor(rightCtx).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		switch op := op.GetSymbol().GetTokenType(); op {
		case parser.FigParserPLUS:
			// if either is string, concatenate
			if left.Type == environment.StringType || right.Type == environment.StringType {
				left = environment.NewString(left.String() + right.String())
			} else {
				ln, err := left.AsNumber()
				if err != nil {
					v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
					return environment.NewNil()
				}
				rn, err := right.AsNumber()
				if err != nil {
					v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
					return environment.NewNil()
				}
				left = environment.NewNumber(ln + rn)
			}
		case parser.FigParserMINUS:
			ln, err := left.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
			rn, err := right.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
			left = environment.NewNumber(ln - rn)
		default:
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("unsupported term operator: %v", op), 1)
			return environment.NewNil()
		}
	}
	return left
}

func (v *FigVisitor) VisitFactor(ctx *parser.FactorContext) interface{} {
	children := ctx.GetChildren()
	var seq []antlr.Tree
	for _, c := range children {
		switch c.(type) {
		case parser.IUnaryContext, antlr.TerminalNode:
			seq = append(seq, c)
		}
	}
	if len(seq) == 1 {
		return v.VisitUnary(seq[0].(parser.IUnaryContext).(*parser.UnaryContext))
	}
	left := v.VisitUnary(seq[0].(parser.IUnaryContext).(*parser.UnaryContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}
	for i := 1; i < len(seq); i += 2 {
		op := seq[i].(antlr.TerminalNode)
		rightCtx := seq[i+1].(parser.IUnaryContext).(*parser.UnaryContext)
		right := v.VisitUnary(rightCtx).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		switch op := op.GetSymbol().GetTokenType(); op {
		case parser.FigParserSTAR:
			ln, err := left.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
			rn, err := right.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
			left = environment.NewNumber(ln * rn)
		case parser.FigParserSLASH:
			ln, err := left.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
			rn, err := right.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
			if rn == 0 {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "division by zero", 1)
				return environment.NewNil()
			}
			left = environment.NewNumber(ln / rn)
		case parser.FigParserMOD:
			ln, err := left.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
			rn, err := right.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
			if rn == 0 {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "modulo by zero", 1)
				return environment.NewNil()
			}
			left = environment.NewNumber(math.Mod(ln, rn))
		default:
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("unsupported factor operator: %v", op), 1)
			return environment.NewNil()
		}
	}
	return left
}

func (v *FigVisitor) VisitUnary(ctx *parser.UnaryContext) interface{} {
	// unary: ( MINUS | EXCLAM ) unary | primary ;
	if ctx.GetChildCount() == 2 { // unary operator present
		op := ctx.GetChild(0).(antlr.TerminalNode)
		opText := op.GetSymbol().GetText()
		child := ctx.GetChild(1).(parser.IUnaryContext).(*parser.UnaryContext)

		// prefix ++/-- semantics: must apply to a simple ID (lvalue)
		if opText == "++" || opText == "--" {
			// navigate: unary -> postfix -> primary -> ID
			if child.Postfix() != nil && child.Postfix().Primary() != nil && child.Postfix().Primary().ID() != nil {
				name := child.Postfix().Primary().ID().GetText()
				cur, ok := v.env.Get(name)
				if !ok {
					v.RuntimeErr = v.makeRuntimeError(child.GetStart().GetLine(), child.GetStart().GetColumn(), fmt.Sprintf("variable '%s' not defined", name), len(name))
					return environment.NewNil()
				}
				n, err := cur.AsNumber()
				if err != nil {
					v.RuntimeErr = v.makeRuntimeError(child.GetStart().GetLine(), child.GetStart().GetColumn(), "invalid operand to increment/decrement", 1)
					return environment.NewNil()
				}
				if opText == "++" {
					n = n + 1
				} else {
					n = n - 1
				}
				if err := v.env.Assign(name, environment.NewNumber(n)); err != nil {
					v.RuntimeErr = v.makeRuntimeError(child.GetStart().GetLine(), child.GetStart().GetColumn(), err.Error(), len(name))
					return environment.NewNil()
				}
				return environment.NewNumber(n)
			}
			// otherwise, invalid prefix usage
			v.RuntimeErr = v.makeRuntimeError(op.GetSymbol().GetLine(), op.GetSymbol().GetColumn(), fmt.Sprintf("invalid operand for '%s'", opText), len(opText))
			return environment.NewNil()
		}

		// normal unary ops
		val := v.VisitUnary(child).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		switch op.GetSymbol().GetTokenType() {
		case parser.FigParserMINUS:
			n, err := val.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(op.GetSymbol().GetLine(), op.GetSymbol().GetColumn(), fmt.Sprintf("cannot negate %s", val.TypeName()), 1)
				return environment.NewNil()
			}
			return environment.NewNumber(-n)
		case parser.FigParserEXCLAM:
			return environment.NewBool(!val.IsTruthy())
		default:
			v.RuntimeErr = v.makeRuntimeError(op.GetSymbol().GetLine(), op.GetSymbol().GetColumn(), fmt.Sprintf("unsupported unary operator '%s'", opText), len(opText))
			return environment.NewNil()
		}
	}
	return v.VisitPostfix(ctx.Postfix().(*parser.PostfixContext))
}

func (v *FigVisitor) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	// Debug: count evaluation steps and detect runaway recursion
	if !builtins.IsStepLimitDisabled() {
		v.evalSteps++
		if v.evalSteps > 20000 {
			// Too many evaluation steps: set a runtime error to avoid Go stack overflow
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "maximum evaluation steps exceeded - possible infinite recursion", 1)
			return environment.NewNil()
		}
	}
	if ctx.NUMBER() != nil {
		s := ctx.NUMBER().GetText()
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(s))
			return environment.NewNil()
		}
		return environment.NewNumber(f)
	}
	if ctx.BOOL() != nil {
		b, err := strconv.ParseBool(ctx.BOOL().GetText())
		if err != nil {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(ctx.BOOL().GetText()))
			return environment.NewNil()
		}
		return environment.NewBool(b)
	}
	if ctx.STRING() != nil {
		text := ctx.STRING().GetText()
		unq, err := strconv.Unquote(text)
		if err != nil {
			// try to handle single-quoted strings
			if len(text) >= 2 && text[0] == '\'' && text[len(text)-1] == '\'' {
				unq = text[1 : len(text)-1]
			} else {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), 1)
				return environment.NewNil()
			}
		}
		return environment.NewString(unq)
	}
	if ctx.TK_NULL() != nil {
		return environment.NewNil()
	}
	if ctx.TK_THIS() != nil {
		thisVal, ok := v.env.Get("this")
		if !ok {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				"'this' used outside of a method", 4)
			return environment.NewNil()
		}
		return thisVal
	}
	if ctx.ArrayLiteral() != nil {
		return v.VisitArrayLiteral(ctx.ArrayLiteral().(*parser.ArrayLiteralContext))
	}
	if ctx.ObjectLiteral() != nil {
		return v.VisitObjectLiteral(ctx.ObjectLiteral().(*parser.ObjectLiteralContext))
	}
	// ── try expression: try expr onerror(e) { block } ──
	if ctx.TryExpr() != nil {
		return v.VisitTryExpr(ctx.TryExpr().(*parser.TryExprContext))
	}
	// ── match expression ──
	if ctx.MatchExpr() != nil {
		return v.VisitMatchExpr(ctx.MatchExpr().(*parser.MatchExprContext))
	}
	// ── anonymous function expression: fn(params) block ──
	if ctx.TK_FN() != nil {
		var params []environment.Param
		if ctx.FnParams() != nil {
			fp := ctx.FnParams().(*parser.FnParamsContext)
			for _, pd := range fp.AllParamDecl() {
				switch pctx := pd.(type) {
				case *parser.ParamWithDefaultOrRequiredContext:
					name := pctx.ID().GetText()
					param := environment.Param{Name: name}
					if pctx.ASSIGN() != nil {
						param.HasDefault = true
						param.Default = pctx.Expr()
					}
					params = append(params, param)
				case *parser.ParamOptionalContext:
					name := pctx.ID().GetText()
					param := environment.Param{Name: name, Optional: true}
					params = append(params, param)
				}
			}
		}
		fd := &environment.FuncDef{
			Name:       "<anonymous>",
			Params:     params,
			Body:       ctx.Block().(*parser.BlockContext),
			ClosureEnv: v.env,
		}
		return environment.NewFunction(fd)
	}
	if ctx.ID() != nil {
		name := ctx.ID().GetText()

		// ── function call or struct instantiation: ID LPAREN fnArgs? RPAREN ──
		if ctx.LPAREN() != nil {
			fnVal, ok := v.env.Get(name)
			if !ok {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fmt.Sprintf("'%s' not defined", name), len(name))
				return environment.NewNil()
			}
			args := v.evaluateArgs(ctx.FnArgs())
			if v.RuntimeErr != nil {
				return environment.NewNil()
			}

			// struct instantiation
			if fnVal.Type == environment.StructDefType {
				return v.instantiateStruct(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fnVal.Struct, args)
			}

			return v.callFunction(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), fnVal, args, name).(environment.Value)
		}

		// ── postfix ++/-- ──
		var postfix string
		for _, c := range ctx.GetChildren() {
			switch t := c.(type) {
			case antlr.TerminalNode:
				if t.GetSymbol().GetText() == "++" || t.GetSymbol().GetText() == "--" {
					postfix = t.GetSymbol().GetText()
				}
			}
		}
		val, ok := v.env.Get(name)
		if !ok {
			v.RuntimeErr = v.makeRuntimeError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), fmt.Sprintf("variable '%s' not defined", name), len(name))
			return environment.NewNil()
		}
		if postfix != "" {
			n, err := val.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), "invalid operand to increment/decrement", 1)
				return environment.NewNil()
			}
			old := n
			if postfix == "++" {
				n = n + 1
			} else {
				n = n - 1
			}
			if err := v.env.Assign(name, environment.NewNumber(n)); err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn(), err.Error(), len(name))
				return environment.NewNil()
			}
			return environment.NewNumber(old)
		}
		return val
	}
	if ctx.LPAREN() != nil {
		return v.VisitExpr(ctx.Expr().(*parser.ExprContext))
	}
	return environment.NewNil()
}

// VisitTryExpr handles: try expr onerror(e) { block }
// Evaluates expr; if a RuntimeError occurs, clears it, binds the error message
// to the optional variable, and executes the onerror block.
// A 'return' inside onerror provides the fallback value for the expression.
// A 'break'/'continue' inside onerror sets pendingLoopSignal for the enclosing loop.
func (v *FigVisitor) VisitTryExpr(ctx *parser.TryExprContext) interface{} {
	// The guarded portion can be either an expression or a block (try expr | try { block }).
	// Evaluate accordingly and if no runtime error occurs, return the result.
	if ctx.Expr() != nil {
		// Guarded is an expression
		val := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr == nil {
			return val
		}
	} else {
		// Guarded is a block: it may return a value via 'return' or fall through
		guarded := ctx.Block(0).(*parser.BlockContext)
		res := v.visitBlockRaw(guarded)
		if v.RuntimeErr == nil {
			// No runtime error inside the guarded block — if it returned via 'return', use that
			if ret, ok := res.(returnSignal); ok {
				return ret.value
			}
			if sig, ok := res.(loopSignal); ok {
				v.pendingLoopSignal = sig
				return environment.NewNil()
			}
			return environment.NewNil()
		}
	}

	// At this point v.RuntimeErr != nil: handle the error using the onerror block
	// Extract human-readable message from the error
	var errMsg string
	if re, ok := v.RuntimeErr.(*RuntimeError); ok {
		errMsg = re.Message
	} else {
		errMsg = v.RuntimeErr.Error()
	}

	// Clear the error so execution can continue
	v.RuntimeErr = nil

	// Execute onerror block in a new scope
	v.pushEnv()
	defer v.popEnv()

	// Bind error variable if specified: onerror(e) { ... }
	if ctx.ID() != nil {
		v.env.Define(ctx.ID().GetText(), environment.NewString(errMsg))
	}

	// The onerror block is the last block in the context. If the guarded part was
	// a block, there will be two block nodes (guarded, onerror); otherwise only
	// the onerror block exists.
	var onErrBlock *parser.BlockContext
	blocks := ctx.AllBlock()
	if len(blocks) == 0 {
		// Should not happen per grammar, but be defensive
		return environment.NewNil()
	} else if len(blocks) == 1 {
		onErrBlock = blocks[0].(*parser.BlockContext)
	} else {
		onErrBlock = blocks[1].(*parser.BlockContext)
	}

	result := v.visitBlockRaw(onErrBlock)

	// If block returned via 'return', use its value as the fallback
	if ret, ok := result.(returnSignal); ok {
		return ret.value
	}

	// If block produced break/continue, store for the enclosing loop
	if sig, ok := result.(loopSignal); ok {
		v.pendingLoopSignal = sig
		return environment.NewNil()
	}

	// Block fell through without return — result is null
	return environment.NewNil()
}

// VisitMatchExpr evaluates a match expression: match expr { pattern => body ... }
func (v *FigVisitor) VisitMatchExpr(ctx *parser.MatchExprContext) interface{} {
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}

	// Evaluate and record pattern values up-front to enforce rules:
	// - no duplicate pattern values
	// - there must be a wildcard '_' arm
	arms := ctx.AllMatchArm()
	seenValues := []environment.Value{}
	perArmPatterns := make([][]environment.Value, len(arms))
	wildcardFound := false

	for ai, arm := range arms {
		armCtx := arm.(*parser.MatchArmCaseContext)
		patternCtx := armCtx.MatchPattern().(*parser.MatchPatternContext)
		exprs := patternCtx.AllExpr()

		// detect wildcard usage
		if len(exprs) == 1 && exprs[0].GetText() == "_" {
			perArmPatterns[ai] = []environment.Value{}
			wildcardFound = true
			continue
		}
		// wildcard used with other patterns is invalid
		for _, p := range exprs {
			if p.GetText() == "_" {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "wildcard '_' must be alone in its arm", 1)
				return environment.NewNil()
			}
		}

		// evaluate pattern expressions (one or more)
		var pats []environment.Value
		for _, p := range exprs {
			pv := v.VisitExpr(p.(*parser.ExprContext)).(environment.Value)
			if v.RuntimeErr != nil {
				return environment.NewNil()
			}
			// check duplicate against all previously seen values
			for _, sv := range seenValues {
				if valuesEqual(sv, pv) {
					v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "duplicate pattern value in match", 1)
					return environment.NewNil()
				}
			}
			seenValues = append(seenValues, pv)
			pats = append(pats, pv)
		}
		perArmPatterns[ai] = pats
	}

	if !wildcardFound {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "match expression requires a default '_' arm", 1)
		return environment.NewNil()
	}

	// Evaluate the subject expression now and perform matching using evaluated patterns
	subject := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}

	for ai, arm := range arms {
		armCtx := arm.(*parser.MatchArmCaseContext)
		// wildcard arm
		patternCtx := armCtx.MatchPattern().(*parser.MatchPatternContext)
		exprs := patternCtx.AllExpr()
		if len(exprs) == 1 && exprs[0].GetText() == "_" {
			// execute wildcard body
			if armCtx.Block() != nil {
				v.pushEnv()
				result := v.visitBlockRaw(armCtx.Block().(*parser.BlockContext))
				v.popEnv()

				if ret, ok := result.(returnSignal); ok {
					return ret.value
				}
				if sig, ok := result.(loopSignal); ok {
					v.pendingLoopSignal = sig
					return environment.NewNil()
				}

				return environment.NewNil()
			}
			return v.VisitExpr(armCtx.Expr().(*parser.ExprContext)).(environment.Value)
		}

		// compare against pre-evaluated patterns
		for _, pv := range perArmPatterns[ai] {
			if valuesEqual(subject, pv) {
				if armCtx.Block() != nil {
					v.pushEnv()
					result := v.visitBlockRaw(armCtx.Block().(*parser.BlockContext))
					v.popEnv()

					if ret, ok := result.(returnSignal); ok {
						return ret.value
					}
					if sig, ok := result.(loopSignal); ok {
						v.pendingLoopSignal = sig
						return environment.NewNil()
					}

					return environment.NewNil()
				}
				return v.VisitExpr(armCtx.Expr().(*parser.ExprContext)).(environment.Value)
			}
		}
	}

	return environment.NewNil()
}

// instantiateStruct creates a new instance of a struct, initializing fields and calling init if present.
func (v *FigVisitor) instantiateStruct(line, col int, sd *environment.StructDef, args []environment.Value) environment.Value {
	// Create instance with fields set to defaults
	entries := make(map[string]environment.Value)
	var keys []string
	for _, f := range sd.Fields {
		entries[f.Name] = f.Default
		keys = append(keys, f.Name)
	}
	inst := &environment.Instance{
		Def:    sd,
		Fields: &environment.ObjData{Entries: entries, Keys: keys},
	}
	instVal := environment.NewInstance(inst)

	// Call init if defined
	if initMethod, ok := sd.Methods["init"]; ok {
		// compute required params for init
		required := 0
		for _, p := range initMethod.Params {
			if !p.HasDefault && !p.Optional {
				required++
			}
		}
		if len(args) < required || len(args) > len(initMethod.Params) {
			v.RuntimeErr = v.makeRuntimeError(line, col,
				fmt.Sprintf("'%s' init expects %d argument(s), got %d", sd.Name, len(initMethod.Params), len(args)), len(sd.Name))
			return environment.NewNil()
		}
		v.callMethod(line, col, instVal, environment.NewFunction(initMethod), args)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
	} else if len(args) > 0 {
		v.RuntimeErr = v.makeRuntimeError(line, col,
			fmt.Sprintf("'%s' does not accept constructor arguments (no init method)", sd.Name), len(sd.Name))
		return environment.NewNil()
	}

	return instVal
}

// callMethod calls a method on an instance, binding "this" in the method scope.
func (v *FigVisitor) callMethod(line, col int, receiver environment.Value, fnVal environment.Value, args []environment.Value) environment.Value {
	// Debug: increment call depth and guard against runaway recursion for methods too
	v.callDepth++
	defer func() { v.callDepth-- }()
	if v.callDepth > 100 {
		v.RuntimeErr = v.makeRuntimeError(line, col, "maximum call depth exceeded", 1)
		return environment.NewNil()
	}
	if fnVal.Type != environment.FunctionType || fnVal.Func == nil {
		v.RuntimeErr = v.makeRuntimeError(line, col, "not a method", 1)
		return environment.NewNil()
	}
	fd := fnVal.Func
	name := fd.Name

	// method arity check: missing args allowed if defaults/optionals provided
	required := 0
	for _, p := range fd.Params {
		if !p.HasDefault && !p.Optional {
			required++
		}
	}
	if len(args) < required || len(args) > len(fd.Params) {
		v.RuntimeErr = v.makeRuntimeError(line, col,
			fmt.Sprintf("'%s' expects %d argument(s), got %d", name, len(fd.Params), len(args)), len(name))
		return environment.NewNil()
	}

	prevEnv := v.env
	v.env = environment.NewEnv(fd.ClosureEnv)
	v.env.Define("this", receiver) // bind "this" to the instance
	// bind provided args
	for i, p := range fd.Params {
		if i < len(args) {
			v.env.Define(p.Name, args[i])
		}
	}
	// bind defaults/optionals
	for i, p := range fd.Params {
		if i >= len(args) {
			if p.HasDefault {
				if p.Default != nil {
					val := v.VisitExpr(p.Default.(*parser.ExprContext)).(environment.Value)
					if v.RuntimeErr != nil {
						v.env = prevEnv
						return environment.NewNil()
					}
					v.env.Define(p.Name, val)
				} else {
					v.env.Define(p.Name, environment.NewNil())
				}
			} else if p.Optional {
				v.env.Define(p.Name, environment.NewNil())
			} else {
				v.env.Define(p.Name, environment.NewNil())
			}
		}
	}

	blockCtx := fd.Body.(*parser.BlockContext)
	result := v.visitBlockRaw(blockCtx)
	v.env = prevEnv
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}
	if ret, ok := result.(returnSignal); ok {
		return ret.value
	}
	return environment.NewNil()
}

// callFunction evaluates a function call using a resolved Value and arguments.
func (v *FigVisitor) callFunction(line, col int, fnVal environment.Value, args []environment.Value, callee string) interface{} {
	// Debug: increment call depth and guard against runaway recursion
	v.callDepth++
	defer func() { v.callDepth-- }()
	if v.callDepth > 100 {
		// Too deep: set a runtime error instead of letting Go stack overflow
		v.RuntimeErr = v.makeRuntimeError(line, col, "maximum call depth exceeded", 1)
		return environment.NewNil()
	}

	// Handle builtin functions
	if fnVal.Type == environment.BuiltinFnType {
		if fnVal.Builtin == nil {
			v.RuntimeErr = v.makeRuntimeError(line, col, "builtin function is nil", 1)
			return environment.NewNil()
		}
		result, err := fnVal.Builtin(args)
		if err != nil {
			v.RuntimeErr = v.makeRuntimeError(line, col, err.Error(), len(fnVal.BName))
			return environment.NewNil()
		}
		return result
	}

	name := "<anonymous>"
	if fnVal.Func != nil {
		name = fnVal.Func.Name
	}
	if fnVal.Type != environment.FunctionType || fnVal.Func == nil {
		// Prefer the provided callee description when available (e.g., property name)
		if callee != "" {
			v.RuntimeErr = v.makeRuntimeError(line, col, fmt.Sprintf("'%s' is not a function", callee), len(callee))
			return environment.NewNil()
		}
		// If we have an identifier name use it; otherwise show the value type
		if name != "<anonymous>" {
			v.RuntimeErr = v.makeRuntimeError(line, col, fmt.Sprintf("'%s' is not a function", name), len(name))
			return environment.NewNil()
		}
		v.RuntimeErr = v.makeRuntimeError(line, col, fmt.Sprintf("%s is not a function", fnVal.TypeName()), len(fnVal.TypeName()))
		return environment.NewNil()
	}
	fd := fnVal.Func

	// arity check: allow missing args only when parameters have defaults or are optional
	required := 0
	for _, p := range fd.Params {
		if !p.HasDefault && !p.Optional {
			required++
		}
	}
	if len(args) < required || len(args) > len(fd.Params) {
		v.RuntimeErr = v.makeRuntimeError(line, col,
			fmt.Sprintf("'%s' expects %d argument(s), got %d", name, len(fd.Params), len(args)), len(name))
		return environment.NewNil()
	}

	// push a stack frame for this function call (show where it was defined)
	v.pushFrame("function", name, fd.DefFile, fd.DefLine, 0)
	defer v.popFrame()

	// If the function has a definition file recorded, switch the source context
	// while executing the body so runtime errors show the correct snippet.
	prevSrc := v.srcLines
	prevCurrentFile := v.currentFile
	if fd.DefFile != "" {
		if data, err := os.ReadFile(fd.DefFile); err == nil {
			v.srcLines = strings.Split(string(data), "\n")
			v.currentFile = fd.DefFile
		}
	}
	defer func() {
		v.srcLines = prevSrc
		v.currentFile = prevCurrentFile
	}()

	// create a new scope for the function call (closure over the definition-time env)
	prevEnv := v.env
	v.env = environment.NewEnv(fd.ClosureEnv) // functions close over definition-time scope
	// bind provided args first
	for i, p := range fd.Params {
		if i < len(args) {
			v.env.Define(p.Name, args[i])
		}
	}
	// evaluate & bind defaults/optionals (left-to-right, call-time evaluation)
	for i, p := range fd.Params {
		if i >= len(args) {
			if p.HasDefault {
				if p.Default != nil {
					val := v.VisitExpr(p.Default.(*parser.ExprContext)).(environment.Value)
					if v.RuntimeErr != nil {
						v.env = prevEnv
						return environment.NewNil()
					}
					v.env.Define(p.Name, val)
				} else {
					// optional parameter declared with '?'
					v.env.Define(p.Name, environment.NewNil())
				}
			} else if p.Optional {
				v.env.Define(p.Name, environment.NewNil())
			} else {
				// missing required param — should not happen because of earlier check
				v.env.Define(p.Name, environment.NewNil())
			}
		}
	}

	// execute the body
	blockCtx := fd.Body.(*parser.BlockContext)
	result := v.visitBlockRaw(blockCtx) // raw: without extra pushEnv (we already set up the scope)
	v.env = prevEnv
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}

	// unwrap return signal
	if ret, ok := result.(returnSignal); ok {
		return ret.value
	}
	return environment.NewNil()
}

// evaluateArgs evaluates a FnArgsContext into a slice of Values.
func (v *FigVisitor) evaluateArgs(argsCtx parser.IFnArgsContext) []environment.Value {
	var args []environment.Value
	if argsCtx != nil {
		ac := argsCtx.(*parser.FnArgsContext)
		for _, e := range ac.AllExpr() {
			val := v.VisitExpr(e.(*parser.ExprContext)).(environment.Value)
			if v.RuntimeErr != nil {
				return nil
			}
			args = append(args, val)
		}
	}
	return args
}

// VisitPostfix handles chained member access: primary ( [expr] | .ID | (args) )*
func (v *FigVisitor) VisitPostfix(ctx *parser.PostfixContext) interface{} {
	if v.RuntimeErr != nil {
		return environment.NewNil()
	}
	val := v.VisitPrimary(ctx.Primary().(*parser.PrimaryContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return val
	}

	// receiver tracks the instance when DOT accesses a method, so LPAREN can bind "this"
	var receiver *environment.Value
	// lastKey stores the most recent property/index accessed so we can produce
	// better error messages when a subsequent LPAREN tries to call a non-function
	lastKey := ""

	// iterate through suffix children: LBRACKET/DOT/LPAREN
	children := ctx.GetChildren()
	i := 1 // skip primary (child 0)
	for i < len(children) {
		c := children[i]
		// process terminal nodes only; skip other nodes
		if _, ok := c.(antlr.TerminalNode); !ok {
			i++
			continue
		}
		tn := c.(antlr.TerminalNode)
		tok := tn.GetSymbol()
		switch tok.GetTokenType() {
		case parser.FigParserLBRACKET:
			// index access: [ expr ]
			receiver = nil
			// debug

			// children: '[' expr ']'
			if i+2 < len(children) {
				exprNode := children[i+1]
				idx := v.VisitExpr(exprNode.(*parser.ExprContext)).(environment.Value)
				if v.RuntimeErr != nil {
					return environment.NewNil()
				}
				val = v.indexAccess(tok.GetLine(), tok.GetColumn(), val, idx)
				if v.RuntimeErr != nil {
					return environment.NewNil()
				}
			}
			// advance past '[', expr, ']'
			i += 3
		case parser.FigParserDOT:
			// dot access: .memberName  (memberName = ID | TK_MATCH)

			if i+1 >= len(children) {
				i++
				continue
			}
			// next child may be a MemberNameContext or a TerminalNode
			var key string
			if mn, ok := children[i+1].(*parser.MemberNameContext); ok {
				key = mn.GetText()
				// advance past '.' and memberName
				i += 2
			} else if tn2, ok := children[i+1].(antlr.TerminalNode); ok {
				key = tn2.GetText()
				i += 2
			} else {
				i++
				continue
			}
			// remember lastKey for error messages
			lastKey = key
			// track receiver for method calls on instances
			if val.Type == environment.InstanceType {
				recv := val
				receiver = &recv
			} else {
				receiver = nil
			}
			val = v.dotAccess(tok.GetLine(), tok.GetColumn(), val, key)
			if v.RuntimeErr != nil {
				return environment.NewNil()
			}
		case parser.FigParserLPAREN:
			// function call: ( args? )
			// next child may be IFnArgsContext or ')' terminal
			var args []environment.Value
			if i+1 < len(children) {
				if next, ok := children[i+1].(parser.IFnArgsContext); ok {
					args = v.evaluateArgs(next)
					if v.RuntimeErr != nil {
						return environment.NewNil()
					}
					// advance past '(', fnArgs, ')'
					i += 3
				} else {
					// advance past '(' and ')'
					i += 2
				}
				// handle call/instantiation/method
				if val.Type == environment.StructDefType {
					val = v.instantiateStruct(tok.GetLine(), tok.GetColumn(), val.Struct, args)
					receiver = nil
					if v.RuntimeErr != nil {
						return environment.NewNil()
					}
				} else if receiver != nil && val.Type == environment.FunctionType {
					// method call: inject "this"
					val = v.callMethod(tok.GetLine(), tok.GetColumn(), *receiver, val, args)
					receiver = nil
					if v.RuntimeErr != nil {
						return environment.NewNil()
					}
				} else {
					receiver = nil
					callee := ""
					if lastKey != "" {
						callee = fmt.Sprintf("%s", lastKey)
					}
					val = v.callFunction(tok.GetLine(), tok.GetColumn(), val, args, callee).(environment.Value)
					// reset lastKey after the call
					lastKey = ""
					if v.RuntimeErr != nil {
						return environment.NewNil()
					}
				}
			}
		default:
			// unknown terminal: advance
			i++
		}
	}
	return val
}

// indexAccess returns the element at index of an array or object.
func (v *FigVisitor) indexAccess(line, col int, container, index environment.Value) environment.Value {
	switch container.Type {
	case environment.ArrayType:
		n, err := index.AsNumber()
		if err != nil {
			v.RuntimeErr = v.makeRuntimeError(line, col, "array index must be a number", 1)
			return environment.NewNil()
		}
		idx := int(n)
		arr := *container.Arr
		// Support negative indices (Python-like): -1 refers to last element.
		if idx < 0 {
			idx = len(arr) + idx
		}
		if idx < 0 || idx >= len(arr) {
			v.RuntimeErr = v.makeRuntimeError(line, col, fmt.Sprintf("array index %d out of range (length %d)", idx, len(arr)), 1)
			return environment.NewNil()
		}
		return arr[idx]
	case environment.ObjectType:
		key := index.String()
		if val, ok := container.Obj.Entries[key]; ok {
			return val
		}
		return environment.NewNil()
	default:
		v.RuntimeErr = v.makeRuntimeError(line, col, fmt.Sprintf("cannot index into %s", container.TypeName()), 1)
		return environment.NewNil()
	}
}

// dotAccess returns the value of a property of an object or instance.
func (v *FigVisitor) dotAccess(line, col int, container environment.Value, key string) environment.Value {
	switch container.Type {
	case environment.ObjectType:
		if container.Obj == nil {
			return environment.NewNil()
		}
		if val, ok := container.Obj.Entries[key]; ok {
			return val
		}
		return environment.NewNil()
	case environment.InstanceType:
		inst := container.Inst
		if inst == nil || inst.Fields == nil || inst.Def == nil {
			v.RuntimeErr = v.makeRuntimeError(line, col, "cannot access property on invalid instance", 1)
			return environment.NewNil()
		}
		// first check instance fields
		if val, ok := inst.Fields.Entries[key]; ok {
			return val
		}
		// then check struct methods
		if method, ok := inst.Def.Methods[key]; ok {
			return environment.NewFunction(method)
		}
		return environment.NewNil()
	default:
		v.RuntimeErr = v.makeRuntimeError(line, col, fmt.Sprintf("cannot access property '%s' on %s", key, container.TypeName()), 1)
		return environment.NewNil()
	}
}

// VisitArrayLiteral evaluates [expr, expr, ...]
func (v *FigVisitor) VisitArrayLiteral(ctx *parser.ArrayLiteralContext) interface{} {
	var elems []environment.Value
	for _, e := range ctx.AllExpr() {
		val := v.VisitExpr(e.(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		elems = append(elems, val)
	}
	return environment.NewArray(elems)
}

// VisitObjectLiteral evaluates {key: value, ...}
func (v *FigVisitor) VisitObjectLiteral(ctx *parser.ObjectLiteralContext) interface{} {
	entries := make(map[string]environment.Value)
	var keys []string
	for _, entry := range ctx.AllObjectEntry() {
		ec := entry.(*parser.ObjectEntryContext)
		var key string
		if ec.ID() != nil {
			key = ec.ID().GetText()
		} else if ec.STRING() != nil {
			text := ec.STRING().GetText()
			unq, err := strconv.Unquote(text)
			if err != nil {
				if len(text) >= 2 && text[0] == '\'' && text[len(text)-1] == '\'' {
					unq = text[1 : len(text)-1]
				} else {
					v.RuntimeErr = v.makeRuntimeError(ec.GetStart().GetLine(), ec.GetStart().GetColumn(), err.Error(), 1)
					return environment.NewNil()
				}
			}
			key = unq
		}
		val := v.VisitExpr(ec.Expr().(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return environment.NewNil()
		}
		if _, exists := entries[key]; !exists {
			keys = append(keys, key)
		}
		entries[key] = val
	}
	return environment.NewObject(entries, keys)
}

// VisitMemberAssign handles assignments like a[0] = x; or obj.name = x;
func (v *FigVisitor) VisitMemberAssign(ctx *parser.MemberAssignContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}

	// Parse the access chain from the children.
	// The structure is: expr (access)+ ASSIGN expr SEMICOLON?
	// The first expr gives us the base, the accesses give us the chain,
	// and the last expr is the value to assign.
	allExprs := ctx.AllExpr()
	if len(allExprs) < 2 {
		return nil
	}

	// The value to assign is the expr right after ASSIGN (second-to-last meaningful expr)
	// In the grammar: expr (LBRACKET expr RBRACKET | DOT ID)+ ASSIGN expr SEMICOLON?
	// allExprs includes: base expr, index exprs, and the RHS value expr
	rhs := v.VisitExpr(allExprs[len(allExprs)-1].(*parser.ExprContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return nil
	}

	// Build the access chain from children
	type accessStep struct {
		isIndex bool
		key     string            // for dot access
		idx     environment.Value // for index access
	}

	var chain []accessStep
	base := v.VisitExpr(allExprs[0].(*parser.ExprContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return nil
	}

	// Walk children to build access chain
	children := ctx.GetChildren()
	exprIdx := 1 // skip base expr (index 0), the exprs in brackets are 1..n-2, last is RHS
	for i := 0; i < len(children); i++ {
		tn, ok := children[i].(antlr.TerminalNode)
		if !ok {
			continue
		}
		tok := tn.GetSymbol()
		switch tok.GetTokenType() {
		case parser.FigParserLBRACKET:
			// next expr is the index
			if exprIdx < len(allExprs)-1 {
				idx := v.VisitExpr(allExprs[exprIdx].(*parser.ExprContext)).(environment.Value)
				if v.RuntimeErr != nil {
					return nil
				}
				chain = append(chain, accessStep{isIndex: true, idx: idx})
				exprIdx++
			}
		case parser.FigParserDOT:
			// next node is memberName (rule) or terminal ID
			for j := i + 1; j < len(children); j++ {
				if mn, ok := children[j].(*parser.MemberNameContext); ok {
					chain = append(chain, accessStep{isIndex: false, key: mn.GetText()})
					i = j
					break
				}
				if idTn, ok := children[j].(antlr.TerminalNode); ok && idTn.GetSymbol().GetTokenType() == parser.FigParserID {
					chain = append(chain, accessStep{isIndex: false, key: idTn.GetText()})
					i = j
					break
				}
			}
		}
	}

	if len(chain) == 0 {
		return nil
	}

	// Navigate to the container (all but last step)
	container := base
	for i := 0; i < len(chain)-1; i++ {
		step := chain[i]
		if step.isIndex {
			container = v.indexAccess(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), container, step.idx)
		} else {
			container = v.dotAccess(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), container, step.key)
		}
		if v.RuntimeErr != nil {
			return nil
		}
	}

	// Apply the last step as assignment
	last := chain[len(chain)-1]
	if last.isIndex {
		switch container.Type {
		case environment.ArrayType:
			n, err := last.idx.AsNumber()
			if err != nil {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "array index must be a number", 1)
				return nil
			}
			idx := int(n)
			arr := *container.Arr
			if idx < 0 {
				v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
					fmt.Sprintf("array index %d out of range (length %d)", idx, len(arr)), 1)
				return nil
			}
			if idx >= len(arr) {
				// Auto-grow: fill gaps with null
				for len(*container.Arr) <= idx {
					*container.Arr = append(*container.Arr, environment.NewNil())
				}
			}
			(*container.Arr)[idx] = rhs
		case environment.ObjectType:
			key := last.idx.String()
			if _, exists := container.Obj.Entries[key]; !exists {
				container.Obj.Keys = append(container.Obj.Keys, key)
			}
			container.Obj.Entries[key] = rhs
		default:
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				fmt.Sprintf("cannot index into %s", container.TypeName()), 1)
		}
	} else {
		switch container.Type {
		case environment.ObjectType:
			if _, exists := container.Obj.Entries[last.key]; !exists {
				container.Obj.Keys = append(container.Obj.Keys, last.key)
			}
			container.Obj.Entries[last.key] = rhs
		case environment.InstanceType:
			inst := container.Inst
			if _, exists := inst.Fields.Entries[last.key]; !exists {
				inst.Fields.Keys = append(inst.Fields.Keys, last.key)
			}
			inst.Fields.Entries[last.key] = rhs
		default:
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				fmt.Sprintf("cannot set property '%s' on %s", last.key, container.TypeName()), 1)
			return nil
		}
	}
	return nil
}

// visitBlockRaw executes a block's statements in the current env (no push/pop).
func (v *FigVisitor) visitBlockRaw(ctx *parser.BlockContext) interface{} {
	for _, st := range ctx.AllStatements() {
		res := v.VisitStatements(st.(*parser.StatementsContext))
		if v.RuntimeErr != nil {
			return nil
		}
		if sig, ok := res.(loopSignal); ok {
			return sig
		}
		if ret, ok := res.(returnSignal); ok {
			return ret
		}
	}
	return nil
}

// VisitFnDecl handles function declarations: fn name(params) { body }
func (v *FigVisitor) VisitFnDecl(ctx *parser.FnDeclContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	name := ctx.ID().GetText()

	// collect parameter metadata (name, default, optional)
	var params []environment.Param
	if ctx.FnParams() != nil {
		paramsCtx := ctx.FnParams().(*parser.FnParamsContext)
		for _, pd := range paramsCtx.AllParamDecl() {
			switch pctx := pd.(type) {
			case *parser.ParamWithDefaultOrRequiredContext:
				param := environment.Param{Name: pctx.ID().GetText()}
				if pctx.ASSIGN() != nil {
					param.HasDefault = true
					param.Default = pctx.Expr()
				}
				params = append(params, param)
			case *parser.ParamOptionalContext:
				param := environment.Param{Name: pctx.ID().GetText(), Optional: true}
				params = append(params, param)
			}
		}
	}

	fd := &environment.FuncDef{
		Name:       name,
		Params:     params,
		Body:       ctx.Block(),
		ClosureEnv: v.env, // capture the definition-time environment
		DefFile:    v.currentFile,
		DefLine:    ctx.GetStart().GetLine(),
	}
	if err := v.env.Define(name, environment.NewFunction(fd)); err != nil {
		// allow redefinition of functions (overwrite)
		v.env.Set(name, environment.NewFunction(fd))
	}
	return nil
}

// VisitReturnStmt handles return statements.
func (v *FigVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	if ctx.Expr() != nil {
		val := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return nil
		}
		return returnSignal{value: val}
	}
	return returnSignal{value: environment.NewNil()}
}

func (v *FigVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	// execute block in a fresh local environment
	v.pushEnv()
	defer v.popEnv()
	for _, st := range ctx.AllStatements() {
		res := v.VisitStatements(st.(*parser.StatementsContext))
		if v.RuntimeErr != nil {
			return nil
		}
		if sig, ok := res.(loopSignal); ok {
			// propagate loop control signals up to the enclosing loop
			return sig
		}
		if ret, ok := res.(returnSignal); ok {
			return ret
		}
	}
	return nil
}

func (v *FigVisitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	v.loopDepth++
	defer func() { v.loopDepth-- }()
	for {
		cond := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return nil
		}
		if !cond.IsTruthy() {
			break
		}
		res := v.VisitBlock(ctx.Block().(*parser.BlockContext))
		if v.RuntimeErr != nil {
			return nil
		}
		if ret, ok := res.(returnSignal); ok {
			return ret
		}
		if sig, ok := res.(loopSignal); ok {
			if sig == loopBreak {
				break
			}
			if sig == loopContinue {
				// on continue, execute step semantics if any (for while there is none)
				continue
			}
		}
	}
	return nil
}

// execForStep executes the step expression of a for loop (e.g. n = n + 1).
func (v *FigVisitor) execForStep(ctx *parser.ForStmtContext) {
	fs := ctx.ForStep()
	if fs == nil || v.RuntimeErr != nil {
		return
	}
	if fs.ASSIGN() != nil {
		name := fs.ID().GetText()
		val := v.VisitExpr(fs.Expr().(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return
		}
		if err := v.env.Assign(name, val); err != nil {
			v.RuntimeErr = v.makeRuntimeError(fs.GetStart().GetLine(), fs.GetStart().GetColumn(), err.Error(), len(name))
		}
	} else if fs.Expr() != nil {
		v.VisitExpr(fs.Expr().(*parser.ExprContext))
	}
}

func (v *FigVisitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}

	v.pushEnv() // for-loop introduces its own scope for the init variable
	defer v.popEnv()

	// ── init ──
	if fi := ctx.ForInit(); fi != nil {
		if fi.TK_LET() != nil {
			name := fi.ID().GetText()
			if fi.Expr() != nil {
				val := v.VisitExpr(fi.Expr().(*parser.ExprContext)).(environment.Value)
				if v.RuntimeErr != nil {
					return nil
				}
				if err := v.env.Define(name, val); err != nil {
					v.RuntimeErr = v.makeRuntimeError(fi.GetStart().GetLine(), fi.GetStart().GetColumn(), err.Error(), len(name))
					return nil
				}
			} else {
				if err := v.env.Define(name, environment.NewNil()); err != nil {
					v.RuntimeErr = v.makeRuntimeError(fi.GetStart().GetLine(), fi.GetStart().GetColumn(), err.Error(), len(name))
					return nil
				}
			}
		} else if fi.ASSIGN() != nil {
			name := fi.ID().GetText()
			val := v.VisitExpr(fi.Expr().(*parser.ExprContext)).(environment.Value)
			if v.RuntimeErr != nil {
				return nil
			}
			if err := v.env.Assign(name, val); err != nil {
				v.RuntimeErr = v.makeRuntimeError(fi.GetStart().GetLine(), fi.GetStart().GetColumn(), err.Error(), len(name))
				return nil
			}
		} else if fi.Expr() != nil {
			v.VisitExpr(fi.Expr().(*parser.ExprContext))
			if v.RuntimeErr != nil {
				return nil
			}
		}
	}

	// ── loop ──
	v.loopDepth++
	defer func() { v.loopDepth-- }()

	for {
		// condition (absent = always true)
		if e := ctx.Expr(); e != nil {
			cond := v.VisitExpr(e.(*parser.ExprContext)).(environment.Value)
			if v.RuntimeErr != nil {
				return nil
			}
			if !cond.IsTruthy() {
				break
			}
		}

		// body
		res := v.VisitBlock(ctx.Block().(*parser.BlockContext))
		if v.RuntimeErr != nil {
			return nil
		}

		if ret, ok := res.(returnSignal); ok {
			return ret
		}

		if sig, ok := res.(loopSignal); ok {
			if sig == loopBreak {
				break
			}
			// loopContinue: execute step, then next iteration
			v.execForStep(ctx)
			if v.RuntimeErr != nil {
				return nil
			}
			continue
		}

		// normal end of body: execute step
		v.execForStep(ctx)
		if v.RuntimeErr != nil {
			return nil
		}
	}
	return nil
}

// ── for x in range(start, end) { } ──
// ── for x in range(start, end, step) { } ──
func (v *FigVisitor) VisitForRange(ctx *parser.ForRangeContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}

	varName := ctx.ID().GetText()
	exprs := ctx.AllExpr()

	startVal := v.VisitExpr(exprs[0].(*parser.ExprContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return nil
	}
	endVal := v.VisitExpr(exprs[1].(*parser.ExprContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return nil
	}

	if startVal.Type != environment.NumberType || endVal.Type != environment.NumberType {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			"range() arguments must be numbers", 5)
		return nil
	}

	start := startVal.Num
	end := endVal.Num
	step := 1.0

	if len(exprs) == 3 {
		stepVal := v.VisitExpr(exprs[2].(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return nil
		}
		if stepVal.Type != environment.NumberType {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				"range() step must be a number", 5)
			return nil
		}
		step = stepVal.Num
		if step == 0 {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
				"range() step cannot be zero", 5)
			return nil
		}
	} else if start > end {
		step = -1.0
	}

	v.loopDepth++
	defer func() { v.loopDepth-- }()

	for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
		// create a fresh scope for this iteration so closures capture the
		// current value of the loop variable (capture-by-value semantics)
		v.pushEnv()
		if err := v.env.Define(varName, environment.NewNumber(i)); err != nil {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(varName))
			v.popEnv()
			return nil
		}

		res := v.VisitBlock(ctx.Block().(*parser.BlockContext))
		// always pop the per-iteration env before handling results
		v.popEnv()

		if v.RuntimeErr != nil {
			return nil
		}
		if ret, ok := res.(returnSignal); ok {
			return ret
		}
		if sig, ok := res.(loopSignal); ok {
			if sig == loopBreak {
				break
			}
			// continue
		}
	}
	return nil
}

// ── for idx, val in enumerate(expr) { } ──
func (v *FigVisitor) VisitForEnumerate(ctx *parser.ForEnumerateContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}

	idxName := ctx.ID(0).GetText()
	valName := ctx.ID(1).GetText()

	iterVal := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return nil
	}

	if iterVal.Type != environment.ArrayType {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			"enumerate() argument must be an array", 9)
		return nil
	}

	arr := *iterVal.Arr

	v.loopDepth++
	defer func() { v.loopDepth-- }()

	for i, elem := range arr {
		// per-iteration scope so closures capture current idx/val
		v.pushEnv()
		if err := v.env.Define(idxName, environment.NewNumber(float64(i))); err != nil {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(idxName))
			v.popEnv()
			return nil
		}
		if err := v.env.Define(valName, elem); err != nil {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(valName))
			v.popEnv()
			return nil
		}

		res := v.VisitBlock(ctx.Block().(*parser.BlockContext))
		// pop per-iteration scope before result handling
		v.popEnv()

		if v.RuntimeErr != nil {
			return nil
		}
		if ret, ok := res.(returnSignal); ok {
			return ret
		}
		if sig, ok := res.(loopSignal); ok {
			if sig == loopBreak {
				break
			}
			// continue
		}
	}
	return nil
}

// ── for x in expr { } ── (iterates over array elements)
func (v *FigVisitor) VisitForIn(ctx *parser.ForInContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}

	varName := ctx.ID().GetText()

	iterVal := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
	if v.RuntimeErr != nil {
		return nil
	}

	if iterVal.Type != environment.ArrayType {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(),
			"for..in requires an array", 3)
		return nil
	}

	arr := *iterVal.Arr

	v.loopDepth++
	defer func() { v.loopDepth-- }()

	for _, elem := range arr {
		// per-iteration env so closures capture the current element
		v.pushEnv()
		if err := v.env.Define(varName, elem); err != nil {
			v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), err.Error(), len(varName))
			v.popEnv()
			return nil
		}

		res := v.VisitBlock(ctx.Block().(*parser.BlockContext))
		v.popEnv()

		if v.RuntimeErr != nil {
			return nil
		}
		if ret, ok := res.(returnSignal); ok {
			return ret
		}
		if sig, ok := res.(loopSignal); ok {
			if sig == loopBreak {
				break
			}
			// continue
		}
	}
	return nil
}

func (v *FigVisitor) VisitDoWhileStmt(ctx *parser.DoWhileStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	v.loopDepth++
	defer func() { v.loopDepth-- }()
	for {
		res := v.VisitBlock(ctx.Block().(*parser.BlockContext))
		if v.RuntimeErr != nil {
			return nil
		}
		if ret, ok := res.(returnSignal); ok {
			return ret
		}
		if sig, ok := res.(loopSignal); ok {
			if sig == loopBreak {
				break
			}
			if sig == loopContinue {
				// continue to re-evaluate condition
			}
		}
		cond := v.VisitExpr(ctx.Expr().(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return nil
		}
		if !cond.IsTruthy() {
			break
		}
	}
	return nil
}

func (v *FigVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	if v.loopDepth == 0 {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "'break' não permitido fora de loop", len("break"))
		return nil
	}
	return loopBreak
}

func (v *FigVisitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	if v.loopDepth == 0 {
		v.RuntimeErr = v.makeRuntimeError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), "'continue' não permitido fora de loop", len("continue"))
		return nil
	}
	return loopContinue
}

func (v *FigVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	if v.RuntimeErr != nil {
		return nil
	}
	exprs := ctx.AllExpr()
	blocks := ctx.AllBlock()
	// evaluate if and elif branches in order
	for i := 0; i < len(exprs); i++ {
		cond := v.VisitExpr(exprs[i].(*parser.ExprContext)).(environment.Value)
		if v.RuntimeErr != nil {
			return nil
		}
		if cond.IsTruthy() {
			res := v.VisitBlock(blocks[i].(*parser.BlockContext))
			if sig, ok := res.(loopSignal); ok {
				return sig
			}
			if ret, ok := res.(returnSignal); ok {
				return ret
			}
			return nil
		}
	}
	// optional else
	if len(blocks) > len(exprs) {
		res := v.VisitBlock(blocks[len(blocks)-1].(*parser.BlockContext))
		if sig, ok := res.(loopSignal); ok {
			return sig
		}
		if ret, ok := res.(returnSignal); ok {
			return ret
		}
	}
	return nil
}

// valuesEqual compares two Values for equality by type and content.
// It is cycle-safe: when comparing container types (arrays/objects),
// it tracks pointer pairs already being compared to avoid infinite recursion
// on self-referential structures.
func valuesEqual(a, b environment.Value) bool {
	seen := make(map[[2]uintptr]bool)
	return valuesEqualSeen(a, b, seen)
}

func valuesEqualSeen(a, b environment.Value, seen map[[2]uintptr]bool) bool {
	if a.Type != b.Type {
		return false
	}
	switch a.Type {
	case environment.NumberType:
		aNum, _ := a.AsNumber()
		bNum, _ := b.AsNumber()
		return aNum == bNum
	case environment.StringType:
		return a.Str == b.Str
	case environment.BooleanType:
		return a.Bool == b.Bool
	case environment.NilType:
		return true
	case environment.ArrayType:
		if a.Arr == nil && b.Arr == nil {
			return true
		}
		if a.Arr == nil || b.Arr == nil {
			return false
		}
		// detect cycles using pointer identity of the underlying slice
		pa := uintptr(unsafe.Pointer(a.Arr))
		pb := uintptr(unsafe.Pointer(b.Arr))
		key := [2]uintptr{pa, pb}
		if seen[key] {
			return true
		}
		seen[key] = true
		aa, bb := *a.Arr, *b.Arr
		if len(aa) != len(bb) {
			return false
		}
		for i := range aa {
			if !valuesEqualSeen(aa[i], bb[i], seen) {
				return false
			}
		}
		return true
	case environment.ObjectType:
		if a.Obj == nil && b.Obj == nil {
			return true
		}
		if a.Obj == nil || b.Obj == nil {
			return false
		}
		pa := uintptr(unsafe.Pointer(a.Obj))
		pb := uintptr(unsafe.Pointer(b.Obj))
		key := [2]uintptr{pa, pb}
		if seen[key] {
			return true
		}
		seen[key] = true
		if len(a.Obj.Entries) != len(b.Obj.Entries) {
			return false
		}
		for k, va := range a.Obj.Entries {
			vb, ok := b.Obj.Entries[k]
			if !ok || !valuesEqualSeen(va, vb, seen) {
				return false
			}
		}
		return true
	case environment.EnumMemberType:
		if a.EnumMem == nil || b.EnumMem == nil {
			return false
		}
		return a.EnumMem.EnumName == b.EnumMem.EnumName && a.EnumMem.Name == b.EnumMem.Name
	}
	return false
}
