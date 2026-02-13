// Generated from /home/carlos/projects/golang/FigLang/grammar/FigParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link FigParser}.
 */
public interface FigParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link FigParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(FigParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(FigParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#statements}.
	 * @param ctx the parse tree
	 */
	void enterStatements(FigParser.StatementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#statements}.
	 * @param ctx the parse tree
	 */
	void exitStatements(FigParser.StatementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#exprStmt}.
	 * @param ctx the parse tree
	 */
	void enterExprStmt(FigParser.ExprStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#exprStmt}.
	 * @param ctx the parse tree
	 */
	void exitExprStmt(FigParser.ExprStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#ifStmt}.
	 * @param ctx the parse tree
	 */
	void enterIfStmt(FigParser.IfStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#ifStmt}.
	 * @param ctx the parse tree
	 */
	void exitIfStmt(FigParser.IfStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#whileStmt}.
	 * @param ctx the parse tree
	 */
	void enterWhileStmt(FigParser.WhileStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#whileStmt}.
	 * @param ctx the parse tree
	 */
	void exitWhileStmt(FigParser.WhileStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#doWhileStmt}.
	 * @param ctx the parse tree
	 */
	void enterDoWhileStmt(FigParser.DoWhileStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#doWhileStmt}.
	 * @param ctx the parse tree
	 */
	void exitDoWhileStmt(FigParser.DoWhileStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#breakStmt}.
	 * @param ctx the parse tree
	 */
	void enterBreakStmt(FigParser.BreakStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#breakStmt}.
	 * @param ctx the parse tree
	 */
	void exitBreakStmt(FigParser.BreakStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#continueStmt}.
	 * @param ctx the parse tree
	 */
	void enterContinueStmt(FigParser.ContinueStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#continueStmt}.
	 * @param ctx the parse tree
	 */
	void exitContinueStmt(FigParser.ContinueStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#fnDecl}.
	 * @param ctx the parse tree
	 */
	void enterFnDecl(FigParser.FnDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#fnDecl}.
	 * @param ctx the parse tree
	 */
	void exitFnDecl(FigParser.FnDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#fnParams}.
	 * @param ctx the parse tree
	 */
	void enterFnParams(FigParser.FnParamsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#fnParams}.
	 * @param ctx the parse tree
	 */
	void exitFnParams(FigParser.FnParamsContext ctx);
	/**
	 * Enter a parse tree produced by the {@code paramWithDefaultOrRequired}
	 * labeled alternative in {@link FigParser#paramDecl}.
	 * @param ctx the parse tree
	 */
	void enterParamWithDefaultOrRequired(FigParser.ParamWithDefaultOrRequiredContext ctx);
	/**
	 * Exit a parse tree produced by the {@code paramWithDefaultOrRequired}
	 * labeled alternative in {@link FigParser#paramDecl}.
	 * @param ctx the parse tree
	 */
	void exitParamWithDefaultOrRequired(FigParser.ParamWithDefaultOrRequiredContext ctx);
	/**
	 * Enter a parse tree produced by the {@code paramOptional}
	 * labeled alternative in {@link FigParser#paramDecl}.
	 * @param ctx the parse tree
	 */
	void enterParamOptional(FigParser.ParamOptionalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code paramOptional}
	 * labeled alternative in {@link FigParser#paramDecl}.
	 * @param ctx the parse tree
	 */
	void exitParamOptional(FigParser.ParamOptionalContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#returnStmt}.
	 * @param ctx the parse tree
	 */
	void enterReturnStmt(FigParser.ReturnStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#returnStmt}.
	 * @param ctx the parse tree
	 */
	void exitReturnStmt(FigParser.ReturnStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#importStmt}.
	 * @param ctx the parse tree
	 */
	void enterImportStmt(FigParser.ImportStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#importStmt}.
	 * @param ctx the parse tree
	 */
	void exitImportStmt(FigParser.ImportStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#useStmt}.
	 * @param ctx the parse tree
	 */
	void enterUseStmt(FigParser.UseStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#useStmt}.
	 * @param ctx the parse tree
	 */
	void exitUseStmt(FigParser.UseStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#fnArgs}.
	 * @param ctx the parse tree
	 */
	void enterFnArgs(FigParser.FnArgsContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#fnArgs}.
	 * @param ctx the parse tree
	 */
	void exitFnArgs(FigParser.FnArgsContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#forInit}.
	 * @param ctx the parse tree
	 */
	void enterForInit(FigParser.ForInitContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#forInit}.
	 * @param ctx the parse tree
	 */
	void exitForInit(FigParser.ForInitContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#forStep}.
	 * @param ctx the parse tree
	 */
	void enterForStep(FigParser.ForStepContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#forStep}.
	 * @param ctx the parse tree
	 */
	void exitForStep(FigParser.ForStepContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void enterForStmt(FigParser.ForStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#forStmt}.
	 * @param ctx the parse tree
	 */
	void exitForStmt(FigParser.ForStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code forEnumerate}
	 * labeled alternative in {@link FigParser#forInStmt}.
	 * @param ctx the parse tree
	 */
	void enterForEnumerate(FigParser.ForEnumerateContext ctx);
	/**
	 * Exit a parse tree produced by the {@code forEnumerate}
	 * labeled alternative in {@link FigParser#forInStmt}.
	 * @param ctx the parse tree
	 */
	void exitForEnumerate(FigParser.ForEnumerateContext ctx);
	/**
	 * Enter a parse tree produced by the {@code forRange}
	 * labeled alternative in {@link FigParser#forInStmt}.
	 * @param ctx the parse tree
	 */
	void enterForRange(FigParser.ForRangeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code forRange}
	 * labeled alternative in {@link FigParser#forInStmt}.
	 * @param ctx the parse tree
	 */
	void exitForRange(FigParser.ForRangeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code forIn}
	 * labeled alternative in {@link FigParser#forInStmt}.
	 * @param ctx the parse tree
	 */
	void enterForIn(FigParser.ForInContext ctx);
	/**
	 * Exit a parse tree produced by the {@code forIn}
	 * labeled alternative in {@link FigParser#forInStmt}.
	 * @param ctx the parse tree
	 */
	void exitForIn(FigParser.ForInContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(FigParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(FigParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#structDecl}.
	 * @param ctx the parse tree
	 */
	void enterStructDecl(FigParser.StructDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#structDecl}.
	 * @param ctx the parse tree
	 */
	void exitStructDecl(FigParser.StructDeclContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structField}
	 * labeled alternative in {@link FigParser#structMember}.
	 * @param ctx the parse tree
	 */
	void enterStructField(FigParser.StructFieldContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structField}
	 * labeled alternative in {@link FigParser#structMember}.
	 * @param ctx the parse tree
	 */
	void exitStructField(FigParser.StructFieldContext ctx);
	/**
	 * Enter a parse tree produced by the {@code structMethod}
	 * labeled alternative in {@link FigParser#structMember}.
	 * @param ctx the parse tree
	 */
	void enterStructMethod(FigParser.StructMethodContext ctx);
	/**
	 * Exit a parse tree produced by the {@code structMethod}
	 * labeled alternative in {@link FigParser#structMember}.
	 * @param ctx the parse tree
	 */
	void exitStructMethod(FigParser.StructMethodContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#enumDecl}.
	 * @param ctx the parse tree
	 */
	void enterEnumDecl(FigParser.EnumDeclContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#enumDecl}.
	 * @param ctx the parse tree
	 */
	void exitEnumDecl(FigParser.EnumDeclContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#enumMember}.
	 * @param ctx the parse tree
	 */
	void enterEnumMember(FigParser.EnumMemberContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#enumMember}.
	 * @param ctx the parse tree
	 */
	void exitEnumMember(FigParser.EnumMemberContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclaration(FigParser.VarDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclaration(FigParser.VarDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#varAtribuition}.
	 * @param ctx the parse tree
	 */
	void enterVarAtribuition(FigParser.VarAtribuitionContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#varAtribuition}.
	 * @param ctx the parse tree
	 */
	void exitVarAtribuition(FigParser.VarAtribuitionContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#bindingTarget}.
	 * @param ctx the parse tree
	 */
	void enterBindingTarget(FigParser.BindingTargetContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#bindingTarget}.
	 * @param ctx the parse tree
	 */
	void exitBindingTarget(FigParser.BindingTargetContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#arrayPattern}.
	 * @param ctx the parse tree
	 */
	void enterArrayPattern(FigParser.ArrayPatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#arrayPattern}.
	 * @param ctx the parse tree
	 */
	void exitArrayPattern(FigParser.ArrayPatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#bindingElement}.
	 * @param ctx the parse tree
	 */
	void enterBindingElement(FigParser.BindingElementContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#bindingElement}.
	 * @param ctx the parse tree
	 */
	void exitBindingElement(FigParser.BindingElementContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#objectPattern}.
	 * @param ctx the parse tree
	 */
	void enterObjectPattern(FigParser.ObjectPatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#objectPattern}.
	 * @param ctx the parse tree
	 */
	void exitObjectPattern(FigParser.ObjectPatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#memberAssign}.
	 * @param ctx the parse tree
	 */
	void enterMemberAssign(FigParser.MemberAssignContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#memberAssign}.
	 * @param ctx the parse tree
	 */
	void exitMemberAssign(FigParser.MemberAssignContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#printStmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintStmt(FigParser.PrintStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#printStmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintStmt(FigParser.PrintStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(FigParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(FigParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#logicalOr}.
	 * @param ctx the parse tree
	 */
	void enterLogicalOr(FigParser.LogicalOrContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#logicalOr}.
	 * @param ctx the parse tree
	 */
	void exitLogicalOr(FigParser.LogicalOrContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#logicalAnd}.
	 * @param ctx the parse tree
	 */
	void enterLogicalAnd(FigParser.LogicalAndContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#logicalAnd}.
	 * @param ctx the parse tree
	 */
	void exitLogicalAnd(FigParser.LogicalAndContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#equality}.
	 * @param ctx the parse tree
	 */
	void enterEquality(FigParser.EqualityContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#equality}.
	 * @param ctx the parse tree
	 */
	void exitEquality(FigParser.EqualityContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#comparison}.
	 * @param ctx the parse tree
	 */
	void enterComparison(FigParser.ComparisonContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#comparison}.
	 * @param ctx the parse tree
	 */
	void exitComparison(FigParser.ComparisonContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#term}.
	 * @param ctx the parse tree
	 */
	void enterTerm(FigParser.TermContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#term}.
	 * @param ctx the parse tree
	 */
	void exitTerm(FigParser.TermContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#factor}.
	 * @param ctx the parse tree
	 */
	void enterFactor(FigParser.FactorContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#factor}.
	 * @param ctx the parse tree
	 */
	void exitFactor(FigParser.FactorContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#unary}.
	 * @param ctx the parse tree
	 */
	void enterUnary(FigParser.UnaryContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#unary}.
	 * @param ctx the parse tree
	 */
	void exitUnary(FigParser.UnaryContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#postfix}.
	 * @param ctx the parse tree
	 */
	void enterPostfix(FigParser.PostfixContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#postfix}.
	 * @param ctx the parse tree
	 */
	void exitPostfix(FigParser.PostfixContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#memberName}.
	 * @param ctx the parse tree
	 */
	void enterMemberName(FigParser.MemberNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#memberName}.
	 * @param ctx the parse tree
	 */
	void exitMemberName(FigParser.MemberNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#primary}.
	 * @param ctx the parse tree
	 */
	void enterPrimary(FigParser.PrimaryContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#primary}.
	 * @param ctx the parse tree
	 */
	void exitPrimary(FigParser.PrimaryContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#tryExpr}.
	 * @param ctx the parse tree
	 */
	void enterTryExpr(FigParser.TryExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#tryExpr}.
	 * @param ctx the parse tree
	 */
	void exitTryExpr(FigParser.TryExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#matchExpr}.
	 * @param ctx the parse tree
	 */
	void enterMatchExpr(FigParser.MatchExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#matchExpr}.
	 * @param ctx the parse tree
	 */
	void exitMatchExpr(FigParser.MatchExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code matchArmCase}
	 * labeled alternative in {@link FigParser#matchArm}.
	 * @param ctx the parse tree
	 */
	void enterMatchArmCase(FigParser.MatchArmCaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code matchArmCase}
	 * labeled alternative in {@link FigParser#matchArm}.
	 * @param ctx the parse tree
	 */
	void exitMatchArmCase(FigParser.MatchArmCaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#matchPattern}.
	 * @param ctx the parse tree
	 */
	void enterMatchPattern(FigParser.MatchPatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#matchPattern}.
	 * @param ctx the parse tree
	 */
	void exitMatchPattern(FigParser.MatchPatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void enterArrayLiteral(FigParser.ArrayLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#arrayLiteral}.
	 * @param ctx the parse tree
	 */
	void exitArrayLiteral(FigParser.ArrayLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#objectLiteral}.
	 * @param ctx the parse tree
	 */
	void enterObjectLiteral(FigParser.ObjectLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#objectLiteral}.
	 * @param ctx the parse tree
	 */
	void exitObjectLiteral(FigParser.ObjectLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link FigParser#objectEntry}.
	 * @param ctx the parse tree
	 */
	void enterObjectEntry(FigParser.ObjectEntryContext ctx);
	/**
	 * Exit a parse tree produced by {@link FigParser#objectEntry}.
	 * @param ctx the parse tree
	 */
	void exitObjectEntry(FigParser.ObjectEntryContext ctx);
}