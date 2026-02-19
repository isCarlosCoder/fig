// Code generated from /home/carlos/projects/golang/FigLang/grammar/FigParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

	// Visit a parse tree produced by FigParser#paramWithDefaultOrRequired.
	VisitParamWithDefaultOrRequired(ctx *ParamWithDefaultOrRequiredContext) interface{}

	// Visit a parse tree produced by FigParser#paramOptional.
	VisitParamOptional(ctx *ParamOptionalContext) interface{}

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

	// Visit a parse tree produced by FigParser#enumDecl.
	VisitEnumDecl(ctx *EnumDeclContext) interface{}

	// Visit a parse tree produced by FigParser#enumMember.
	VisitEnumMember(ctx *EnumMemberContext) interface{}

	// Visit a parse tree produced by FigParser#varDeclaration.
	VisitVarDeclaration(ctx *VarDeclarationContext) interface{}

	// Visit a parse tree produced by FigParser#varAtribuition.
	VisitVarAtribuition(ctx *VarAtribuitionContext) interface{}

	// Visit a parse tree produced by FigParser#bindingTarget.
	VisitBindingTarget(ctx *BindingTargetContext) interface{}

	// Visit a parse tree produced by FigParser#arrayPattern.
	VisitArrayPattern(ctx *ArrayPatternContext) interface{}

	// Visit a parse tree produced by FigParser#bindingElement.
	VisitBindingElement(ctx *BindingElementContext) interface{}

	// Visit a parse tree produced by FigParser#objectPattern.
	VisitObjectPattern(ctx *ObjectPatternContext) interface{}

	// Visit a parse tree produced by FigParser#memberAssign.
	VisitMemberAssign(ctx *MemberAssignContext) interface{}

	// Visit a parse tree produced by FigParser#printStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by FigParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by FigParser#conditional.
	VisitConditional(ctx *ConditionalContext) interface{}

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

	// Visit a parse tree produced by FigParser#memberName.
	VisitMemberName(ctx *MemberNameContext) interface{}

	// Visit a parse tree produced by FigParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by FigParser#tryExpr.
	VisitTryExpr(ctx *TryExprContext) interface{}

	// Visit a parse tree produced by FigParser#matchExpr.
	VisitMatchExpr(ctx *MatchExprContext) interface{}

	// Visit a parse tree produced by FigParser#matchArmCase.
	VisitMatchArmCase(ctx *MatchArmCaseContext) interface{}

	// Visit a parse tree produced by FigParser#matchPattern.
	VisitMatchPattern(ctx *MatchPatternContext) interface{}

	// Visit a parse tree produced by FigParser#arrayLiteralSimple.
	VisitArrayLiteralSimple(ctx *ArrayLiteralSimpleContext) interface{}

	// Visit a parse tree produced by FigParser#arrayCompForIn.
	VisitArrayCompForIn(ctx *ArrayCompForInContext) interface{}

	// Visit a parse tree produced by FigParser#arrayCompForRange.
	VisitArrayCompForRange(ctx *ArrayCompForRangeContext) interface{}

	// Visit a parse tree produced by FigParser#arrayCompForEnumerate.
	VisitArrayCompForEnumerate(ctx *ArrayCompForEnumerateContext) interface{}

	// Visit a parse tree produced by FigParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by FigParser#objectEntry.
	VisitObjectEntry(ctx *ObjectEntryContext) interface{}
}
