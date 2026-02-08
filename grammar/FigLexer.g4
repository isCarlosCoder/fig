lexer grammar FigLexer;

// Multi-char tokens
LT: '<' ;
GT: '>' ;
LE: '<=' ;
GE: '>=' ;
EQ: '==' ;
NEQ: '!=' ;
AND: '&&' ;
OR: '||' ;

// Single-char tokens
DOT: '.' ;
ASSIGN: '=' ;
PLUS: '+' ;
MINUS: '-' ;
PLUSPLUS: '++' ;
MINUSMINUS: '--' ;
STAR: '*' ;
SLASH: '/' ;
LPAREN: '(' ;
RPAREN: ')' ;
EXCLAM: '!' ;
SEMICOLON: ';' ;
MOD: '%' ;

// Keywords and identifiers
TK_LET: 'let' ;
TK_PRINT: 'print' ;
TK_IF: 'if' ;
TK_ELIF: 'elif' ;
TK_ELSE: 'else' ;
TK_WHILE: 'while' ;
TK_DO: 'do' ;
TK_BREAK: 'break' ;
TK_CONTINUE: 'continue' ;
TK_FOR: 'for' ;
TK_NULL: 'null' ;
TK_FN: 'fn' ;
TK_RETURN: 'return' ;
TK_IMPORT: 'import' ;
TK_USE: 'use' ;
TK_IN: 'in' ;
TK_RANGE: 'range' ;
TK_ENUMERATE: 'enumerate' ;
TK_STRUCT: 'struct' ;
TK_THIS: 'this' ;
COMMA: ',' ;

// Array / object tokens
LBRACKET: '[' ;
RBRACKET: ']' ;
COLON: ':' ;

// Block tokens
LBRACE: '{' ;
RBRACE: '}' ;

// Literals and other tokens
BOOL: 'true' | 'false' ;
ID: (LETTER | '_') (LETTER | DIGIT | '_')* ;
NUMBER: DIGIT+ ('.' DIGIT+)? ;
STRING: '"' (ESC | ~["\\])* '"' | '\'' (ESC | ~['\\])* '\'' ;

// Skip whitespace and comments
WS: [ \t\r\n]+ -> skip ;
COMMENT: '#' ~[\r\n]* -> skip ;

// Fragments for building other tokens
fragment DIGIT: [0-9] ;
fragment LETTER: [a-zA-Z] ;
fragment ESC: '\\' [btnfr"'\\] ;