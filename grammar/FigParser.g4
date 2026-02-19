parser grammar FigParser;

options {
    tokenVocab=FigLexer;
}

program
    : statements* EOF
    ;

statements
    : varDeclaration
    | varAtribuition
    | memberAssign
    | printStmt
    | ifStmt
    | whileStmt
    | doWhileStmt
    | forStmt
    | forInStmt
    | breakStmt
    | continueStmt
    | fnDecl
    | returnStmt
    | importStmt
    | useStmt
    | structDecl
    | enumDecl
    | exprStmt
    ;

exprStmt
    : expr SEMICOLON?
    ;

ifStmt
    : TK_IF LPAREN expr RPAREN block (TK_ELIF LPAREN expr RPAREN block)* (TK_ELSE block)?
    ;

whileStmt
    : TK_WHILE LPAREN expr RPAREN block
    ;

doWhileStmt
    : TK_DO block TK_WHILE LPAREN expr RPAREN SEMICOLON?
    ;

breakStmt
    : TK_BREAK SEMICOLON?
    ;

continueStmt
    : TK_CONTINUE SEMICOLON?
    ;

fnDecl
    : TK_FN ID LPAREN fnParams? RPAREN block
    ;

fnParams
    : paramDecl (COMMA paramDecl)*
    ;

paramDecl
    : ID (ASSIGN expr)?        # paramWithDefaultOrRequired
    | ID QUESTION              # paramOptional
    ;

returnStmt
    : TK_RETURN expr? SEMICOLON?
    ;

importStmt
    : TK_IMPORT STRING (ID | STAR)? SEMICOLON?
    ;

useStmt
    : TK_USE STRING SEMICOLON?
    ;

fnArgs
    : expr (COMMA expr)*
    ;

forInit
    : TK_LET ID (ASSIGN expr)?
    | ID ASSIGN expr
    | expr
    ;

forStep
    : ID ASSIGN expr
    | expr
    ;

forStmt
    : TK_FOR LPAREN forInit? SEMICOLON expr? SEMICOLON forStep? RPAREN block
    ;

forInStmt
    : TK_FOR ID COMMA ID TK_IN TK_ENUMERATE LPAREN expr RPAREN block    # forEnumerate
    | TK_FOR ID TK_IN TK_RANGE LPAREN expr (COMMA expr (COMMA expr)?)? RPAREN block  # forRange
    | TK_FOR ID TK_IN expr block                                          # forIn
    ;

block
    : LBRACE statements* RBRACE
    ;

structDecl
    : TK_STRUCT ID LBRACE structMember* RBRACE
    ;

structMember
    : ID (ASSIGN expr)? SEMICOLON?       # structField
    | TK_FN ID LPAREN fnParams? RPAREN block  # structMethod
    ;

enumDecl
    : TK_ENUM ID LBRACE enumMember* RBRACE
    ;

enumMember
    : ID SEMICOLON?
    ;

varDeclaration  
    : TK_LET bindingTarget (ASSIGN expr)? SEMICOLON?
    ;

varAtribuition
    : bindingTarget (ASSIGN | PLUSEQ | MINUSEQ | STAREQ | SLASHEQ | MODEQ) expr SEMICOLON?
    ;

// Binding targets for destructuring assignments / declarations
bindingTarget
    : ID
    | arrayPattern
    | objectPattern
    ;

arrayPattern
    : LBRACKET (bindingElement (COMMA bindingElement)*)? RBRACKET
    ;

bindingElement
    : ID
    | arrayPattern
    | objectPattern
    ;

objectPattern
    : LBRACE (ID (COMMA ID)*)? RBRACE
    ;

memberAssign
    : expr (LBRACKET expr RBRACKET | DOT memberName)+ (ASSIGN | PLUSEQ | MINUSEQ | STAREQ | SLASHEQ | MODEQ) expr SEMICOLON?
    ;

printStmt
    : TK_PRINT LPAREN (expr (COMMA expr)*)? RPAREN SEMICOLON?
    ;

expr: conditional ;
conditional: logicalOr ( QUESTION expr COLON conditional )? ;
logicalOr: logicalAnd ( OR logicalAnd )* ;
logicalAnd: equality ( AND equality )* ;
equality: comparison ( ( EQ | NEQ ) comparison )* ;
comparison: term ( ( GT | GE | LT | LE ) term )* ;
term: factor ( ( PLUS | MINUS ) factor )* ;
factor: unary ( ( STAR | SLASH | MOD ) unary )* ;
unary: ( MINUS | EXCLAM | PLUSPLUS | MINUSMINUS ) unary | postfix ;
postfix: primary ( LBRACKET expr RBRACKET | DOT memberName | LPAREN fnArgs? RPAREN )* ;
memberName: ID | TK_MATCH ;
primary
    : NUMBER
    | BOOL
    | STRING
    | TK_NULL
    | TK_THIS
    | arrayLiteral
    | objectLiteral
    | tryExpr
    | matchExpr
    | TK_FN LPAREN fnParams? RPAREN block
    | ID LPAREN fnArgs? RPAREN
    | ID ( PLUSPLUS | MINUSMINUS )?
    | LPAREN expr RPAREN
    ;

tryExpr
    : TK_TRY (expr | block) TK_ONERROR (LPAREN ID? RPAREN)? block
    ;

matchExpr
    : TK_MATCH expr LBRACE matchArm+ RBRACE
    ;

matchArm
    : matchPattern ARROW (block | expr)       # matchArmCase
    ;  

matchPattern
    : expr (COMMA expr)*
    ;

arrayLiteral
    : LBRACKET (expr (COMMA expr)*)? RBRACKET                                 # arrayLiteralSimple
    | LBRACKET expr TK_FOR ID TK_IN expr RBRACKET                            # arrayCompForIn
    | LBRACKET expr TK_FOR ID TK_IN TK_RANGE LPAREN expr (COMMA expr (COMMA expr)?)? RPAREN RBRACKET  # arrayCompForRange
    | LBRACKET expr TK_FOR ID COMMA ID TK_IN TK_ENUMERATE LPAREN expr RPAREN RBRACKET  # arrayCompForEnumerate
    ;

objectLiteral
    : LBRACE (objectEntry (COMMA objectEntry)*)? RBRACE
    ;

objectEntry
    : (ID | STRING) COLON expr
    ;