package interpreter

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/iscarloscoder/fig/environment"
	"github.com/iscarloscoder/fig/parser"
)

// compileNativeFunc validates a FuncDef annotated with @native and, on
// success, fills FuncDef.NativeImpl with a fast-path evaluator that works on
// float64 arguments. On validation failure returns an error (definition
// should be rejected). Currently supports a small, safe subset of expressions
// (numeric literals, parameters, unary -, binary + - * / %, parentheses and
// calls to math.<func> for a whitelist of functions).
func compileNativeFunc(fd *environment.FuncDef) error {
	// only simple numeric parameters (no defaults/optionals)
	for _, p := range fd.Params {
		if p.HasDefault || p.Optional {
			return fmt.Errorf("native functions do not support default or optional parameters")
		}
	}

	// body must be a single return statement with an expression
	blk, ok := fd.Body.(*parser.BlockContext)
	if !ok {
		return fmt.Errorf("native function body not a block")
	}
	stmts := blk.AllStatements()
	if len(stmts) != 1 {
		return fmt.Errorf("native function must contain exactly one statement (a return)")
	}
	if stmts[0].(*parser.StatementsContext).ReturnStmt() == nil {
		return fmt.Errorf("native function must be a single return expression")
	}
	rctx := stmts[0].(*parser.StatementsContext).ReturnStmt().Expr().(*parser.ExprContext)

	// build parameter name -> index map
	paramNames := make([]string, len(fd.Params))
	paramIndex := make(map[string]int)
	for i, p := range fd.Params {
		paramNames[i] = p.Name
		paramIndex[p.Name] = i
	}

	// validator + evaluator closures that walk the parser expression nodes.
	var evalExpr func(ctx *parser.ExprContext, args []float64) (float64, error)
	var evalLogicalOr func(ctx *parser.LogicalOrContext, args []float64) (float64, error)
	var evalLogicalAnd func(ctx *parser.LogicalAndContext, args []float64) (float64, error)
	var evalEquality func(ctx *parser.EqualityContext, args []float64) (float64, error)
	var evalComparison func(ctx *parser.ComparisonContext, args []float64) (float64, error)
	var evalTerm func(ctx *parser.TermContext, args []float64) (float64, error)
	var evalFactor func(ctx *parser.FactorContext, args []float64) (float64, error)
	var evalUnary func(ctx *parser.UnaryContext, args []float64) (float64, error)
	var evalPostfix func(ctx *parser.PostfixContext, args []float64) (float64, error)
	var evalPrimary func(ctx *parser.PrimaryContext, args []float64) (float64, error)

	// Disallow logical/equality/comparison at top-level for native funcs — keep
	// semantics simple (only arithmetic + math.* calls). However the grammar
	// wraps arithmetic under Comparison/Equality/Logical nodes; allow only the
	// degenerate case (single child) for those nodes.
	evalExpr = func(ctx *parser.ExprContext, args []float64) (float64, error) {
		if lo := ctx.Conditional().LogicalOr(); lo != nil {
			return evalLogicalOr(lo.(*parser.LogicalOrContext), args)
		}
		return 0, fmt.Errorf("unsupported expression node")
	}

	evalLogicalOr = func(ctx *parser.LogicalOrContext, args []float64) (float64, error) {
		// only allow single child (no || chains)
		children := ctx.GetChildren()
		count := 0
		for _, c := range children {
			switch c.(type) {
			case parser.ILogicalAndContext, antlr.TerminalNode:
				count++
			}
		}
		if count != 1 {
			return 0, fmt.Errorf("logical operators not permitted in native functions")
		}
		for _, c := range ctx.GetChildren() {
			if la, ok := c.(parser.ILogicalAndContext); ok {
				return evalLogicalAnd(la.(*parser.LogicalAndContext), args)
			}
		}
		return 0, fmt.Errorf("logical operators not permitted in native functions")
	}

	evalLogicalAnd = func(ctx *parser.LogicalAndContext, args []float64) (float64, error) {
		// allow only single child
		children := ctx.GetChildren()
		count := 0
		for _, c := range children {
			switch c.(type) {
			case parser.IEqualityContext, antlr.TerminalNode:
				count++
			}
		}
		if count != 1 {
			return 0, fmt.Errorf("logical operators not permitted in native functions")
		}
		for _, c := range ctx.GetChildren() {
			if eq, ok := c.(parser.IEqualityContext); ok {
				return evalEquality(eq.(*parser.EqualityContext), args)
			}
		}
		return 0, fmt.Errorf("equality operators not permitted in native functions")
	}

	evalEquality = func(ctx *parser.EqualityContext, args []float64) (float64, error) {
		// allow only single comparison (no ==/!=)
		children := ctx.GetChildren()
		count := 0
		for _, c := range children {
			switch c.(type) {
			case parser.IComparisonContext, antlr.TerminalNode:
				count++
			}
		}
		if count != 1 {
			return 0, fmt.Errorf("equality operators not permitted in native functions")
		}
		for _, c := range ctx.GetChildren() {
			if cmp, ok := c.(parser.IComparisonContext); ok {
				return evalComparison(cmp.(*parser.ComparisonContext), args)
			}
		}
		return 0, fmt.Errorf("comparison operators not permitted in native functions")
	}

	evalComparison = func(ctx *parser.ComparisonContext, args []float64) (float64, error) {
		// allow only single term (no <,>,<=,>=)
		children := ctx.GetChildren()
		count := 0
		for _, c := range children {
			switch c.(type) {
			case parser.ITermContext, antlr.TerminalNode:
				count++
			}
		}
		if count != 1 {
			return 0, fmt.Errorf("comparison operators not permitted in native functions")
		}
		for _, c := range ctx.GetChildren() {
			if t, ok := c.(parser.ITermContext); ok {
				return evalTerm(t.(*parser.TermContext), args)
			}
		}
		return 0, fmt.Errorf("comparison operators not permitted in native functions")
	}

	evalTerm = func(ctx *parser.TermContext, args []float64) (float64, error) {
		children := ctx.GetChildren()
		var seq []antlr.Tree
		for _, c := range children {
			switch c.(type) {
			case parser.IFactorContext, antlr.TerminalNode:
				seq = append(seq, c)
			}
		}
		if len(seq) == 1 {
			return evalFactor(seq[0].(parser.IFactorContext).(*parser.FactorContext), args)
		}
		left, err := evalFactor(seq[0].(parser.IFactorContext).(*parser.FactorContext), args)
		if err != nil {
			return 0, err
		}
		for i := 1; i < len(seq); i += 2 {
			op := seq[i].(antlr.TerminalNode).GetSymbol().GetTokenType()
			right, err := evalFactor(seq[i+1].(parser.IFactorContext).(*parser.FactorContext), args)
			if err != nil {
				return 0, err
			}
			switch op {
			case parser.FigParserPLUS:
				left = left + right
			case parser.FigParserMINUS:
				left = left - right
			default:
				return 0, fmt.Errorf("unsupported term operator in native function")
			}
		}
		return left, nil
	}

	evalFactor = func(ctx *parser.FactorContext, args []float64) (float64, error) {
		children := ctx.GetChildren()
		var seq []antlr.Tree
		for _, c := range children {
			switch c.(type) {
			case parser.IUnaryContext, antlr.TerminalNode:
				seq = append(seq, c)
			}
		}
		if len(seq) == 1 {
			return evalUnary(seq[0].(parser.IUnaryContext).(*parser.UnaryContext), args)
		}
		left, err := evalUnary(seq[0].(parser.IUnaryContext).(*parser.UnaryContext), args)
		if err != nil {
			return 0, err
		}
		for i := 1; i < len(seq); i += 2 {
			op := seq[i].(antlr.TerminalNode).GetSymbol().GetTokenType()
			right, err := evalUnary(seq[i+1].(parser.IUnaryContext).(*parser.UnaryContext), args)
			if err != nil {
				return 0, err
			}
			switch op {
			case parser.FigParserSTAR:
				left = left * right
			case parser.FigParserSLASH:
				if right == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				left = left / right
			case parser.FigParserMOD:
				if right == 0 {
					return 0, fmt.Errorf("modulo by zero")
				}
				left = math.Mod(left, right)
			default:
				return 0, fmt.Errorf("unsupported factor operator in native function")
			}
		}
		return left, nil
	}

	evalUnary = func(ctx *parser.UnaryContext, args []float64) (float64, error) {
		if ctx.GetChildCount() == 2 {
			op := ctx.GetChild(0).(antlr.TerminalNode)
			child := ctx.GetChild(1).(parser.IUnaryContext).(*parser.UnaryContext)
			val, err := evalUnary(child, args)
			if err != nil {
				return 0, err
			}
			switch op.GetSymbol().GetTokenType() {
			case parser.FigParserMINUS:
				return -val, nil
			default:
				return 0, fmt.Errorf("unsupported unary operator in native function")
			}
		}
		return evalPostfix(ctx.Postfix().(*parser.PostfixContext), args)
	}

	evalPostfix = func(ctx *parser.PostfixContext, args []float64) (float64, error) {
		// base primary
		prim := ctx.Primary().(*parser.PrimaryContext)
		// If primary is an ID followed by DOT <member> and LPAREN => math.<fn>(...)
		// Otherwise the primary itself should be a number or parameter.
		// If there are no suffixes, evaluate primary directly.
		children := ctx.GetChildren()
		if len(children) == 1 {
			return evalPrimary(prim, args)
		}

		// We'll support only math.<fn>(...) style calls on identifiers
		// Find pattern: ID DOT memberName LPAREN fnArgs? RPAREN
		// Validate the chain corresponds to that.
		// The simplest approach: get textual representation and check prefix 'math.'
		text := ctx.GetText()
		// expect something like math.exp(...)
		if strings.HasPrefix(text, "math.") {
			// parse function name between '.' and '('
			dot := strings.Index(text, ".")
			paren := strings.Index(text, "(")
			if dot < 0 || paren < 0 || paren <= dot+1 {
				return 0, fmt.Errorf("invalid math.* call in native function")
			}
			fname := text[dot+1 : paren]
			// evaluate arguments inside parentheses by using parser-provided fnArgs
			// find FnArgsContext if present in children
			var fnArgsCtx parser.IFnArgsContext
			for _, ch := range children {
				if a, ok := ch.(parser.IFnArgsContext); ok {
					fnArgsCtx = a
					break
				}
			}
			var argVals []float64
			if fnArgsCtx != nil {
				ac := fnArgsCtx.(*parser.FnArgsContext)
				for _, e := range ac.AllExpr() {
					v, err := evalExpr(e.(*parser.ExprContext), args)
					if err != nil {
						return 0, err
					}
					argVals = append(argVals, v)
				}
			}

			switch fname {
			case "exp":
				if len(argVals) != 1 {
					return 0, fmt.Errorf("math.exp expects 1 arg")
				}
				return math.Exp(argVals[0]), nil
			case "sin":
				if len(argVals) != 1 {
					return 0, fmt.Errorf("math.sin expects 1 arg")
				}
				return math.Sin(argVals[0]), nil
			case "cos":
				if len(argVals) != 1 {
					return 0, fmt.Errorf("math.cos expects 1 arg")
				}
				return math.Cos(argVals[0]), nil
			case "sqrt":
				if len(argVals) != 1 {
					return 0, fmt.Errorf("math.sqrt expects 1 arg")
				}
				return math.Sqrt(argVals[0]), nil
			case "log":
				if len(argVals) != 1 {
					return 0, fmt.Errorf("math.log expects 1 arg")
				}
				return math.Log(argVals[0]), nil
			case "abs":
				if len(argVals) != 1 {
					return 0, fmt.Errorf("math.abs expects 1 arg")
				}
				return math.Abs(argVals[0]), nil
			case "pow":
				if len(argVals) != 2 {
					return 0, fmt.Errorf("math.pow expects 2 args")
				}
				return math.Pow(argVals[0], argVals[1]), nil
			default:
				return 0, fmt.Errorf("unsupported math function '%s' in native function", fname)
			}
		}

		return 0, fmt.Errorf("unsupported postfix expression in native function: %s", text)
	}

	evalPrimary = func(ctx *parser.PrimaryContext, args []float64) (float64, error) {
		if ctx.NUMBER() != nil {
			s := ctx.NUMBER().GetText()
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return 0, err
			}
			return f, nil
		}
		if ctx.ID() != nil {
			name := ctx.ID().GetText()
			// parameter reference
			if idx, ok := paramIndex[name]; ok {
				return args[idx], nil
			}
			return 0, fmt.Errorf("unknown identifier '%s' in native function", name)
		}
		if ctx.LPAREN() != nil {
			return evalExpr(ctx.Expr().(*parser.ExprContext), args)
		}
		return 0, fmt.Errorf("unsupported primary in native function")
	}

	// Quick validation pass: try evaluating with zero values for params to
	// ensure structure is valid (errors indicate unsupported constructs).
	zeroArgs := make([]float64, len(paramNames))
	if _, err := evalExpr(rctx, zeroArgs); err != nil {
		return fmt.Errorf("native function validation failed: %v", err)
	}

	// Build native implementation closure
	fd.NativeImpl = func(vargs []environment.Value) (environment.Value, error) {
		if len(vargs) != len(paramNames) {
			return environment.NewNil(), fmt.Errorf("native: wrong number of args")
		}
		vals := make([]float64, len(vargs))
		for i, av := range vargs {
			if av.Type != environment.NumberType {
				return environment.NewNil(), fmt.Errorf("native: argument %d is not a number", i)
			}
			vals[i] = av.Num
		}
		res, err := evalExpr(rctx, vals)
		if err != nil {
			return environment.NewNil(), err
		}
		return environment.NewNumber(res), nil
	}

	fd.IsNative = true
	return nil
}
