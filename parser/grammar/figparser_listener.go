// Code generated from grammar/FigParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FigParser
import "github.com/antlr4-go/antlr/v4"

// FigParserListener is a complete listener for a parse tree produced by FigParser.
type FigParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterExprStmt is called when entering the exprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterWhileStmt is called when entering the whileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterDoWhileStmt is called when entering the doWhileStmt production.
	EnterDoWhileStmt(c *DoWhileStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterFnDecl is called when entering the fnDecl production.
	EnterFnDecl(c *FnDeclContext)

	// EnterFnParams is called when entering the fnParams production.
	EnterFnParams(c *FnParamsContext)

	// EnterReturnStmt is called when entering the returnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterImportStmt is called when entering the importStmt production.
	EnterImportStmt(c *ImportStmtContext)

	// EnterUseStmt is called when entering the useStmt production.
	EnterUseStmt(c *UseStmtContext)

	// EnterFnArgs is called when entering the fnArgs production.
	EnterFnArgs(c *FnArgsContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForStep is called when entering the forStep production.
	EnterForStep(c *ForStepContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterForEnumerate is called when entering the forEnumerate production.
	EnterForEnumerate(c *ForEnumerateContext)

	// EnterForRange is called when entering the forRange production.
	EnterForRange(c *ForRangeContext)

	// EnterForIn is called when entering the forIn production.
	EnterForIn(c *ForInContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterStructField is called when entering the structField production.
	EnterStructField(c *StructFieldContext)

	// EnterStructMethod is called when entering the structMethod production.
	EnterStructMethod(c *StructMethodContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterVarAtribuition is called when entering the varAtribuition production.
	EnterVarAtribuition(c *VarAtribuitionContext)

	// EnterMemberAssign is called when entering the memberAssign production.
	EnterMemberAssign(c *MemberAssignContext)

	// EnterPrintStmt is called when entering the printStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterLogicalOr is called when entering the logicalOr production.
	EnterLogicalOr(c *LogicalOrContext)

	// EnterLogicalAnd is called when entering the logicalAnd production.
	EnterLogicalAnd(c *LogicalAndContext)

	// EnterEquality is called when entering the equality production.
	EnterEquality(c *EqualityContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// EnterPostfix is called when entering the postfix production.
	EnterPostfix(c *PostfixContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterObjectEntry is called when entering the objectEntry production.
	EnterObjectEntry(c *ObjectEntryContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitExprStmt is called when exiting the exprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitWhileStmt is called when exiting the whileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitDoWhileStmt is called when exiting the doWhileStmt production.
	ExitDoWhileStmt(c *DoWhileStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitFnDecl is called when exiting the fnDecl production.
	ExitFnDecl(c *FnDeclContext)

	// ExitFnParams is called when exiting the fnParams production.
	ExitFnParams(c *FnParamsContext)

	// ExitReturnStmt is called when exiting the returnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitImportStmt is called when exiting the importStmt production.
	ExitImportStmt(c *ImportStmtContext)

	// ExitUseStmt is called when exiting the useStmt production.
	ExitUseStmt(c *UseStmtContext)

	// ExitFnArgs is called when exiting the fnArgs production.
	ExitFnArgs(c *FnArgsContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForStep is called when exiting the forStep production.
	ExitForStep(c *ForStepContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitForEnumerate is called when exiting the forEnumerate production.
	ExitForEnumerate(c *ForEnumerateContext)

	// ExitForRange is called when exiting the forRange production.
	ExitForRange(c *ForRangeContext)

	// ExitForIn is called when exiting the forIn production.
	ExitForIn(c *ForInContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitStructField is called when exiting the structField production.
	ExitStructField(c *StructFieldContext)

	// ExitStructMethod is called when exiting the structMethod production.
	ExitStructMethod(c *StructMethodContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitVarAtribuition is called when exiting the varAtribuition production.
	ExitVarAtribuition(c *VarAtribuitionContext)

	// ExitMemberAssign is called when exiting the memberAssign production.
	ExitMemberAssign(c *MemberAssignContext)

	// ExitPrintStmt is called when exiting the printStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitLogicalOr is called when exiting the logicalOr production.
	ExitLogicalOr(c *LogicalOrContext)

	// ExitLogicalAnd is called when exiting the logicalAnd production.
	ExitLogicalAnd(c *LogicalAndContext)

	// ExitEquality is called when exiting the equality production.
	ExitEquality(c *EqualityContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)

	// ExitPostfix is called when exiting the postfix production.
	ExitPostfix(c *PostfixContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitObjectEntry is called when exiting the objectEntry production.
	ExitObjectEntry(c *ObjectEntryContext)
}
