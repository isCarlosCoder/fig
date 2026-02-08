// Code generated from grammar/FigParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FigParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FigParser.
type FigParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FigParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by FigParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by FigParser#exprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by FigParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by FigParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by FigParser#doWhileStmt.
	VisitDoWhileStmt(ctx *DoWhileStmtContext) interface{}

	// Visit a parse tree produced by FigParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by FigParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by FigParser#fnDecl.
	VisitFnDecl(ctx *FnDeclContext) interface{}

	// Visit a parse tree produced by FigParser#fnParams.
	VisitFnParams(ctx *FnParamsContext) interface{}

	// Visit a parse tree produced by FigParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by FigParser#importStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by FigParser#useStmt.
	VisitUseStmt(ctx *UseStmtContext) interface{}

	// Visit a parse tree produced by FigParser#fnArgs.
	VisitFnArgs(ctx *FnArgsContext) interface{}

	// Visit a parse tree produced by FigParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by FigParser#forStep.
	VisitForStep(ctx *ForStepContext) interface{}

	// Visit a parse tree produced by FigParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by FigParser#forEnumerate.
	VisitForEnumerate(ctx *ForEnumerateContext) interface{}

	// Visit a parse tree produced by FigParser#forRange.
	VisitForRange(ctx *ForRangeContext) interface{}

	// Visit a parse tree produced by FigParser#forIn.
	VisitForIn(ctx *ForInContext) interface{}

	// Visit a parse tree produced by FigParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by FigParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by FigParser#structField.
	VisitStructField(ctx *StructFieldContext) interface{}

	// Visit a parse tree produced by FigParser#structMethod.
	VisitStructMethod(ctx *StructMethodContext) interface{}

	// Visit a parse tree produced by FigParser#varDeclaration.
	VisitVarDeclaration(ctx *VarDeclarationContext) interface{}

	// Visit a parse tree produced by FigParser#varAtribuition.
	VisitVarAtribuition(ctx *VarAtribuitionContext) interface{}

	// Visit a parse tree produced by FigParser#memberAssign.
	VisitMemberAssign(ctx *MemberAssignContext) interface{}

	// Visit a parse tree produced by FigParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by FigParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by FigParser#logicalOr.
	VisitLogicalOr(ctx *LogicalOrContext) interface{}

	// Visit a parse tree produced by FigParser#logicalAnd.
	VisitLogicalAnd(ctx *LogicalAndContext) interface{}

	// Visit a parse tree produced by FigParser#equality.
	VisitEquality(ctx *EqualityContext) interface{}

	// Visit a parse tree produced by FigParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by FigParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by FigParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by FigParser#unary.
	VisitUnary(ctx *UnaryContext) interface{}

	// Visit a parse tree produced by FigParser#postfix.
	VisitPostfix(ctx *PostfixContext) interface{}

	// Visit a parse tree produced by FigParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by FigParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}

	// Visit a parse tree produced by FigParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by FigParser#objectEntry.
	VisitObjectEntry(ctx *ObjectEntryContext) interface{}
}
