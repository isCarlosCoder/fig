// Code generated from grammar/FigParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FigParser
import "github.com/antlr4-go/antlr/v4"

type BaseFigParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFigParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitDoWhileStmt(ctx *DoWhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitFnDecl(ctx *FnDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitFnParams(ctx *FnParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitImportStmt(ctx *ImportStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitUseStmt(ctx *UseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitFnArgs(ctx *FnArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitForInit(ctx *ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitForStep(ctx *ForStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitForEnumerate(ctx *ForEnumerateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitForRange(ctx *ForRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitForIn(ctx *ForInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitStructField(ctx *StructFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitStructMethod(ctx *StructMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitVarDeclaration(ctx *VarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitVarAtribuition(ctx *VarAtribuitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitMemberAssign(ctx *MemberAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitLogicalOr(ctx *LogicalOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitLogicalAnd(ctx *LogicalAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitEquality(ctx *EqualityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitUnary(ctx *UnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitPostfix(ctx *PostfixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitTryExpr(ctx *TryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFigParserVisitor) VisitObjectEntry(ctx *ObjectEntryContext) interface{} {
	return v.VisitChildren(ctx)
}
