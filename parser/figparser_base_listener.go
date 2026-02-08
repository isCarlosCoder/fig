// Code generated from /home/carlos/projects/golang/FigLang/grammar/FigParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FigParser
import "github.com/antlr4-go/antlr/v4"

// BaseFigParserListener is a complete listener for a parse tree produced by FigParser.
type BaseFigParserListener struct{}

var _ FigParserListener = &BaseFigParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFigParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFigParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFigParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFigParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseFigParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseFigParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseFigParserListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseFigParserListener) ExitStatements(ctx *StatementsContext) {}

// EnterExprStmt is called when production exprStmt is entered.
func (s *BaseFigParserListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production exprStmt is exited.
func (s *BaseFigParserListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseFigParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseFigParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseFigParserListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseFigParserListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterDoWhileStmt is called when production doWhileStmt is entered.
func (s *BaseFigParserListener) EnterDoWhileStmt(ctx *DoWhileStmtContext) {}

// ExitDoWhileStmt is called when production doWhileStmt is exited.
func (s *BaseFigParserListener) ExitDoWhileStmt(ctx *DoWhileStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseFigParserListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseFigParserListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseFigParserListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseFigParserListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterFnDecl is called when production fnDecl is entered.
func (s *BaseFigParserListener) EnterFnDecl(ctx *FnDeclContext) {}

// ExitFnDecl is called when production fnDecl is exited.
func (s *BaseFigParserListener) ExitFnDecl(ctx *FnDeclContext) {}

// EnterFnParams is called when production fnParams is entered.
func (s *BaseFigParserListener) EnterFnParams(ctx *FnParamsContext) {}

// ExitFnParams is called when production fnParams is exited.
func (s *BaseFigParserListener) ExitFnParams(ctx *FnParamsContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseFigParserListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseFigParserListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterImportStmt is called when production importStmt is entered.
func (s *BaseFigParserListener) EnterImportStmt(ctx *ImportStmtContext) {}

// ExitImportStmt is called when production importStmt is exited.
func (s *BaseFigParserListener) ExitImportStmt(ctx *ImportStmtContext) {}

// EnterUseStmt is called when production useStmt is entered.
func (s *BaseFigParserListener) EnterUseStmt(ctx *UseStmtContext) {}

// ExitUseStmt is called when production useStmt is exited.
func (s *BaseFigParserListener) ExitUseStmt(ctx *UseStmtContext) {}

// EnterFnArgs is called when production fnArgs is entered.
func (s *BaseFigParserListener) EnterFnArgs(ctx *FnArgsContext) {}

// ExitFnArgs is called when production fnArgs is exited.
func (s *BaseFigParserListener) ExitFnArgs(ctx *FnArgsContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseFigParserListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseFigParserListener) ExitForInit(ctx *ForInitContext) {}

// EnterForStep is called when production forStep is entered.
func (s *BaseFigParserListener) EnterForStep(ctx *ForStepContext) {}

// ExitForStep is called when production forStep is exited.
func (s *BaseFigParserListener) ExitForStep(ctx *ForStepContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseFigParserListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseFigParserListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterForEnumerate is called when production forEnumerate is entered.
func (s *BaseFigParserListener) EnterForEnumerate(ctx *ForEnumerateContext) {}

// ExitForEnumerate is called when production forEnumerate is exited.
func (s *BaseFigParserListener) ExitForEnumerate(ctx *ForEnumerateContext) {}

// EnterForRange is called when production forRange is entered.
func (s *BaseFigParserListener) EnterForRange(ctx *ForRangeContext) {}

// ExitForRange is called when production forRange is exited.
func (s *BaseFigParserListener) ExitForRange(ctx *ForRangeContext) {}

// EnterForIn is called when production forIn is entered.
func (s *BaseFigParserListener) EnterForIn(ctx *ForInContext) {}

// ExitForIn is called when production forIn is exited.
func (s *BaseFigParserListener) ExitForIn(ctx *ForInContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseFigParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseFigParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseFigParserListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseFigParserListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterStructField is called when production structField is entered.
func (s *BaseFigParserListener) EnterStructField(ctx *StructFieldContext) {}

// ExitStructField is called when production structField is exited.
func (s *BaseFigParserListener) ExitStructField(ctx *StructFieldContext) {}

// EnterStructMethod is called when production structMethod is entered.
func (s *BaseFigParserListener) EnterStructMethod(ctx *StructMethodContext) {}

// ExitStructMethod is called when production structMethod is exited.
func (s *BaseFigParserListener) ExitStructMethod(ctx *StructMethodContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseFigParserListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseFigParserListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterVarAtribuition is called when production varAtribuition is entered.
func (s *BaseFigParserListener) EnterVarAtribuition(ctx *VarAtribuitionContext) {}

// ExitVarAtribuition is called when production varAtribuition is exited.
func (s *BaseFigParserListener) ExitVarAtribuition(ctx *VarAtribuitionContext) {}

// EnterMemberAssign is called when production memberAssign is entered.
func (s *BaseFigParserListener) EnterMemberAssign(ctx *MemberAssignContext) {}

// ExitMemberAssign is called when production memberAssign is exited.
func (s *BaseFigParserListener) ExitMemberAssign(ctx *MemberAssignContext) {}

// EnterPrintStmt is called when production printStmt is entered.
func (s *BaseFigParserListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production printStmt is exited.
func (s *BaseFigParserListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseFigParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseFigParserListener) ExitExpr(ctx *ExprContext) {}

// EnterLogicalOr is called when production logicalOr is entered.
func (s *BaseFigParserListener) EnterLogicalOr(ctx *LogicalOrContext) {}

// ExitLogicalOr is called when production logicalOr is exited.
func (s *BaseFigParserListener) ExitLogicalOr(ctx *LogicalOrContext) {}

// EnterLogicalAnd is called when production logicalAnd is entered.
func (s *BaseFigParserListener) EnterLogicalAnd(ctx *LogicalAndContext) {}

// ExitLogicalAnd is called when production logicalAnd is exited.
func (s *BaseFigParserListener) ExitLogicalAnd(ctx *LogicalAndContext) {}

// EnterEquality is called when production equality is entered.
func (s *BaseFigParserListener) EnterEquality(ctx *EqualityContext) {}

// ExitEquality is called when production equality is exited.
func (s *BaseFigParserListener) ExitEquality(ctx *EqualityContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseFigParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseFigParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseFigParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseFigParserListener) ExitTerm(ctx *TermContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseFigParserListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseFigParserListener) ExitFactor(ctx *FactorContext) {}

// EnterUnary is called when production unary is entered.
func (s *BaseFigParserListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BaseFigParserListener) ExitUnary(ctx *UnaryContext) {}

// EnterPostfix is called when production postfix is entered.
func (s *BaseFigParserListener) EnterPostfix(ctx *PostfixContext) {}

// ExitPostfix is called when production postfix is exited.
func (s *BaseFigParserListener) ExitPostfix(ctx *PostfixContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseFigParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseFigParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterTryExpr is called when production tryExpr is entered.
func (s *BaseFigParserListener) EnterTryExpr(ctx *TryExprContext) {}

// ExitTryExpr is called when production tryExpr is exited.
func (s *BaseFigParserListener) ExitTryExpr(ctx *TryExprContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseFigParserListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseFigParserListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseFigParserListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseFigParserListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterObjectEntry is called when production objectEntry is entered.
func (s *BaseFigParserListener) EnterObjectEntry(ctx *ObjectEntryContext) {}

// ExitObjectEntry is called when production objectEntry is exited.
func (s *BaseFigParserListener) ExitObjectEntry(ctx *ObjectEntryContext) {}
