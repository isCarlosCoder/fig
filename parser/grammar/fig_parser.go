// Code generated from grammar/FigParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // FigParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FigParser struct {
	*antlr.BaseParser
}

var FigParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func figparserParserInit() {
	staticData := &FigParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'&&'", "'||'", "'.'",
		"'='", "'+'", "'-'", "'++'", "'--'", "'*'", "'/'", "'('", "')'", "'!'",
		"';'", "'%'", "'let'", "'print'", "'if'", "'elif'", "'else'", "'while'",
		"'do'", "'break'", "'continue'", "'for'", "'null'", "'fn'", "'return'",
		"'import'", "'use'", "'in'", "'range'", "'enumerate'", "'struct'", "'this'",
		"'try'", "'onerror'", "','", "'['", "']'", "':'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "LT", "GT", "LE", "GE", "EQ", "NEQ", "AND", "OR", "DOT", "ASSIGN",
		"PLUS", "MINUS", "PLUSPLUS", "MINUSMINUS", "STAR", "SLASH", "LPAREN",
		"RPAREN", "EXCLAM", "SEMICOLON", "MOD", "TK_LET", "TK_PRINT", "TK_IF",
		"TK_ELIF", "TK_ELSE", "TK_WHILE", "TK_DO", "TK_BREAK", "TK_CONTINUE",
		"TK_FOR", "TK_NULL", "TK_FN", "TK_RETURN", "TK_IMPORT", "TK_USE", "TK_IN",
		"TK_RANGE", "TK_ENUMERATE", "TK_STRUCT", "TK_THIS", "TK_TRY", "TK_ONERROR",
		"COMMA", "LBRACKET", "RBRACKET", "COLON", "LBRACE", "RBRACE", "BOOL",
		"ID", "NUMBER", "STRING", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statements", "exprStmt", "ifStmt", "whileStmt", "doWhileStmt",
		"breakStmt", "continueStmt", "fnDecl", "fnParams", "returnStmt", "importStmt",
		"useStmt", "fnArgs", "forInit", "forStep", "forStmt", "forInStmt", "block",
		"structDecl", "structMember", "varDeclaration", "varAtribuition", "memberAssign",
		"printStmt", "expr", "logicalOr", "logicalAnd", "equality", "comparison",
		"term", "factor", "unary", "postfix", "primary", "tryExpr", "arrayLiteral",
		"objectLiteral", "objectEntry",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 483, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 1, 0, 4, 0, 80, 8, 0, 11, 0, 12, 0, 81, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 103, 8, 1, 1, 2, 1, 2, 3,
		2, 107, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 120, 8, 3, 10, 3, 12, 3, 123, 9, 3, 1, 3, 1, 3, 3, 3, 127,
		8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 142, 8, 5, 1, 6, 1, 6, 3, 6, 146, 8, 6, 1, 7, 1, 7, 3,
		7, 150, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 156, 8, 8, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 5, 9, 164, 8, 9, 10, 9, 12, 9, 167, 9, 9, 1, 10, 1, 10,
		3, 10, 171, 8, 10, 1, 10, 3, 10, 174, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11,
		179, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 184, 8, 12, 1, 13, 1, 13, 1, 13,
		5, 13, 189, 8, 13, 10, 13, 12, 13, 192, 9, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 198, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 204, 8, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 3, 15, 210, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16,
		215, 8, 16, 1, 16, 1, 16, 3, 16, 219, 8, 16, 1, 16, 1, 16, 3, 16, 223,
		8, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 249, 8, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 260, 8, 17, 1, 18, 1, 18,
		5, 18, 264, 8, 18, 10, 18, 12, 18, 267, 9, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 5, 19, 275, 8, 19, 10, 19, 12, 19, 278, 9, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 3, 20, 285, 8, 20, 1, 20, 3, 20, 288, 8, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 294, 8, 20, 1, 20, 1, 20, 3, 20, 298,
		8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 304, 8, 21, 1, 21, 3, 21, 307,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 313, 8, 22, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 4, 23, 322, 8, 23, 11, 23, 12, 23, 323,
		1, 23, 1, 23, 1, 23, 3, 23, 329, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 3, 24, 336, 8, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 5, 26, 343, 8,
		26, 10, 26, 12, 26, 346, 9, 26, 1, 27, 1, 27, 1, 27, 5, 27, 351, 8, 27,
		10, 27, 12, 27, 354, 9, 27, 1, 28, 1, 28, 1, 28, 5, 28, 359, 8, 28, 10,
		28, 12, 28, 362, 9, 28, 1, 29, 1, 29, 1, 29, 5, 29, 367, 8, 29, 10, 29,
		12, 29, 370, 9, 29, 1, 30, 1, 30, 1, 30, 5, 30, 375, 8, 30, 10, 30, 12,
		30, 378, 9, 30, 1, 31, 1, 31, 1, 31, 5, 31, 383, 8, 31, 10, 31, 12, 31,
		386, 9, 31, 1, 32, 1, 32, 1, 32, 3, 32, 391, 8, 32, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 402, 8, 33, 1, 33, 5,
		33, 405, 8, 33, 10, 33, 12, 33, 408, 9, 33, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 421, 8, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 428, 8, 34, 1, 34, 1, 34, 1, 34,
		3, 34, 433, 8, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 439, 8, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 446, 8, 35, 1, 35, 3, 35, 449, 8, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 457, 8, 36, 10, 36, 12,
		36, 460, 9, 36, 3, 36, 462, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1,
		37, 5, 37, 470, 8, 37, 10, 37, 12, 37, 473, 9, 37, 3, 37, 475, 8, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 0, 0, 39, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 0, 7, 1,
		0, 5, 6, 1, 0, 1, 4, 1, 0, 11, 12, 2, 0, 15, 16, 21, 21, 2, 0, 12, 14,
		19, 19, 1, 0, 13, 14, 2, 0, 51, 51, 53, 53, 527, 0, 79, 1, 0, 0, 0, 2,
		102, 1, 0, 0, 0, 4, 104, 1, 0, 0, 0, 6, 108, 1, 0, 0, 0, 8, 128, 1, 0,
		0, 0, 10, 134, 1, 0, 0, 0, 12, 143, 1, 0, 0, 0, 14, 147, 1, 0, 0, 0, 16,
		151, 1, 0, 0, 0, 18, 160, 1, 0, 0, 0, 20, 168, 1, 0, 0, 0, 22, 175, 1,
		0, 0, 0, 24, 180, 1, 0, 0, 0, 26, 185, 1, 0, 0, 0, 28, 203, 1, 0, 0, 0,
		30, 209, 1, 0, 0, 0, 32, 211, 1, 0, 0, 0, 34, 259, 1, 0, 0, 0, 36, 261,
		1, 0, 0, 0, 38, 270, 1, 0, 0, 0, 40, 297, 1, 0, 0, 0, 42, 299, 1, 0, 0,
		0, 44, 308, 1, 0, 0, 0, 46, 314, 1, 0, 0, 0, 48, 330, 1, 0, 0, 0, 50, 337,
		1, 0, 0, 0, 52, 339, 1, 0, 0, 0, 54, 347, 1, 0, 0, 0, 56, 355, 1, 0, 0,
		0, 58, 363, 1, 0, 0, 0, 60, 371, 1, 0, 0, 0, 62, 379, 1, 0, 0, 0, 64, 390,
		1, 0, 0, 0, 66, 392, 1, 0, 0, 0, 68, 438, 1, 0, 0, 0, 70, 440, 1, 0, 0,
		0, 72, 452, 1, 0, 0, 0, 74, 465, 1, 0, 0, 0, 76, 478, 1, 0, 0, 0, 78, 80,
		3, 2, 1, 0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 5, 0, 0, 1, 84, 1, 1, 0,
		0, 0, 85, 103, 3, 42, 21, 0, 86, 103, 3, 44, 22, 0, 87, 103, 3, 46, 23,
		0, 88, 103, 3, 48, 24, 0, 89, 103, 3, 6, 3, 0, 90, 103, 3, 8, 4, 0, 91,
		103, 3, 10, 5, 0, 92, 103, 3, 32, 16, 0, 93, 103, 3, 34, 17, 0, 94, 103,
		3, 12, 6, 0, 95, 103, 3, 14, 7, 0, 96, 103, 3, 16, 8, 0, 97, 103, 3, 20,
		10, 0, 98, 103, 3, 22, 11, 0, 99, 103, 3, 24, 12, 0, 100, 103, 3, 38, 19,
		0, 101, 103, 3, 4, 2, 0, 102, 85, 1, 0, 0, 0, 102, 86, 1, 0, 0, 0, 102,
		87, 1, 0, 0, 0, 102, 88, 1, 0, 0, 0, 102, 89, 1, 0, 0, 0, 102, 90, 1, 0,
		0, 0, 102, 91, 1, 0, 0, 0, 102, 92, 1, 0, 0, 0, 102, 93, 1, 0, 0, 0, 102,
		94, 1, 0, 0, 0, 102, 95, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0, 102, 97, 1, 0,
		0, 0, 102, 98, 1, 0, 0, 0, 102, 99, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102,
		101, 1, 0, 0, 0, 103, 3, 1, 0, 0, 0, 104, 106, 3, 50, 25, 0, 105, 107,
		5, 20, 0, 0, 106, 105, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 5, 1, 0,
		0, 0, 108, 109, 5, 24, 0, 0, 109, 110, 5, 17, 0, 0, 110, 111, 3, 50, 25,
		0, 111, 112, 5, 18, 0, 0, 112, 121, 3, 36, 18, 0, 113, 114, 5, 25, 0, 0,
		114, 115, 5, 17, 0, 0, 115, 116, 3, 50, 25, 0, 116, 117, 5, 18, 0, 0, 117,
		118, 3, 36, 18, 0, 118, 120, 1, 0, 0, 0, 119, 113, 1, 0, 0, 0, 120, 123,
		1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 126, 1, 0,
		0, 0, 123, 121, 1, 0, 0, 0, 124, 125, 5, 26, 0, 0, 125, 127, 3, 36, 18,
		0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 7, 1, 0, 0, 0, 128,
		129, 5, 27, 0, 0, 129, 130, 5, 17, 0, 0, 130, 131, 3, 50, 25, 0, 131, 132,
		5, 18, 0, 0, 132, 133, 3, 36, 18, 0, 133, 9, 1, 0, 0, 0, 134, 135, 5, 28,
		0, 0, 135, 136, 3, 36, 18, 0, 136, 137, 5, 27, 0, 0, 137, 138, 5, 17, 0,
		0, 138, 139, 3, 50, 25, 0, 139, 141, 5, 18, 0, 0, 140, 142, 5, 20, 0, 0,
		141, 140, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 11, 1, 0, 0, 0, 143, 145,
		5, 29, 0, 0, 144, 146, 5, 20, 0, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1,
		0, 0, 0, 146, 13, 1, 0, 0, 0, 147, 149, 5, 30, 0, 0, 148, 150, 5, 20, 0,
		0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 15, 1, 0, 0, 0, 151,
		152, 5, 33, 0, 0, 152, 153, 5, 51, 0, 0, 153, 155, 5, 17, 0, 0, 154, 156,
		3, 18, 9, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0,
		0, 0, 157, 158, 5, 18, 0, 0, 158, 159, 3, 36, 18, 0, 159, 17, 1, 0, 0,
		0, 160, 165, 5, 51, 0, 0, 161, 162, 5, 44, 0, 0, 162, 164, 5, 51, 0, 0,
		163, 161, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165, 163, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 19, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 168, 170, 5,
		34, 0, 0, 169, 171, 3, 50, 25, 0, 170, 169, 1, 0, 0, 0, 170, 171, 1, 0,
		0, 0, 171, 173, 1, 0, 0, 0, 172, 174, 5, 20, 0, 0, 173, 172, 1, 0, 0, 0,
		173, 174, 1, 0, 0, 0, 174, 21, 1, 0, 0, 0, 175, 176, 5, 35, 0, 0, 176,
		178, 5, 53, 0, 0, 177, 179, 5, 20, 0, 0, 178, 177, 1, 0, 0, 0, 178, 179,
		1, 0, 0, 0, 179, 23, 1, 0, 0, 0, 180, 181, 5, 36, 0, 0, 181, 183, 5, 53,
		0, 0, 182, 184, 5, 20, 0, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0,
		184, 25, 1, 0, 0, 0, 185, 190, 3, 50, 25, 0, 186, 187, 5, 44, 0, 0, 187,
		189, 3, 50, 25, 0, 188, 186, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190, 188,
		1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 27, 1, 0, 0, 0, 192, 190, 1, 0,
		0, 0, 193, 194, 5, 22, 0, 0, 194, 197, 5, 51, 0, 0, 195, 196, 5, 10, 0,
		0, 196, 198, 3, 50, 25, 0, 197, 195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0,
		198, 204, 1, 0, 0, 0, 199, 200, 5, 51, 0, 0, 200, 201, 5, 10, 0, 0, 201,
		204, 3, 50, 25, 0, 202, 204, 3, 50, 25, 0, 203, 193, 1, 0, 0, 0, 203, 199,
		1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 29, 1, 0, 0, 0, 205, 206, 5, 51,
		0, 0, 206, 207, 5, 10, 0, 0, 207, 210, 3, 50, 25, 0, 208, 210, 3, 50, 25,
		0, 209, 205, 1, 0, 0, 0, 209, 208, 1, 0, 0, 0, 210, 31, 1, 0, 0, 0, 211,
		212, 5, 31, 0, 0, 212, 214, 5, 17, 0, 0, 213, 215, 3, 28, 14, 0, 214, 213,
		1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 218, 5, 20,
		0, 0, 217, 219, 3, 50, 25, 0, 218, 217, 1, 0, 0, 0, 218, 219, 1, 0, 0,
		0, 219, 220, 1, 0, 0, 0, 220, 222, 5, 20, 0, 0, 221, 223, 3, 30, 15, 0,
		222, 221, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224,
		225, 5, 18, 0, 0, 225, 226, 3, 36, 18, 0, 226, 33, 1, 0, 0, 0, 227, 228,
		5, 31, 0, 0, 228, 229, 5, 51, 0, 0, 229, 230, 5, 44, 0, 0, 230, 231, 5,
		51, 0, 0, 231, 232, 5, 37, 0, 0, 232, 233, 5, 39, 0, 0, 233, 234, 5, 17,
		0, 0, 234, 235, 3, 50, 25, 0, 235, 236, 5, 18, 0, 0, 236, 237, 3, 36, 18,
		0, 237, 260, 1, 0, 0, 0, 238, 239, 5, 31, 0, 0, 239, 240, 5, 51, 0, 0,
		240, 241, 5, 37, 0, 0, 241, 242, 5, 38, 0, 0, 242, 243, 5, 17, 0, 0, 243,
		244, 3, 50, 25, 0, 244, 245, 5, 44, 0, 0, 245, 248, 3, 50, 25, 0, 246,
		247, 5, 44, 0, 0, 247, 249, 3, 50, 25, 0, 248, 246, 1, 0, 0, 0, 248, 249,
		1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 5, 18, 0, 0, 251, 252, 3, 36,
		18, 0, 252, 260, 1, 0, 0, 0, 253, 254, 5, 31, 0, 0, 254, 255, 5, 51, 0,
		0, 255, 256, 5, 37, 0, 0, 256, 257, 3, 50, 25, 0, 257, 258, 3, 36, 18,
		0, 258, 260, 1, 0, 0, 0, 259, 227, 1, 0, 0, 0, 259, 238, 1, 0, 0, 0, 259,
		253, 1, 0, 0, 0, 260, 35, 1, 0, 0, 0, 261, 265, 5, 48, 0, 0, 262, 264,
		3, 2, 1, 0, 263, 262, 1, 0, 0, 0, 264, 267, 1, 0, 0, 0, 265, 263, 1, 0,
		0, 0, 265, 266, 1, 0, 0, 0, 266, 268, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0,
		268, 269, 5, 49, 0, 0, 269, 37, 1, 0, 0, 0, 270, 271, 5, 40, 0, 0, 271,
		272, 5, 51, 0, 0, 272, 276, 5, 48, 0, 0, 273, 275, 3, 40, 20, 0, 274, 273,
		1, 0, 0, 0, 275, 278, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 276, 277, 1, 0,
		0, 0, 277, 279, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 279, 280, 5, 49, 0, 0,
		280, 39, 1, 0, 0, 0, 281, 284, 5, 51, 0, 0, 282, 283, 5, 10, 0, 0, 283,
		285, 3, 50, 25, 0, 284, 282, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 287,
		1, 0, 0, 0, 286, 288, 5, 20, 0, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0,
		0, 0, 288, 298, 1, 0, 0, 0, 289, 290, 5, 33, 0, 0, 290, 291, 5, 51, 0,
		0, 291, 293, 5, 17, 0, 0, 292, 294, 3, 18, 9, 0, 293, 292, 1, 0, 0, 0,
		293, 294, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 5, 18, 0, 0, 296,
		298, 3, 36, 18, 0, 297, 281, 1, 0, 0, 0, 297, 289, 1, 0, 0, 0, 298, 41,
		1, 0, 0, 0, 299, 300, 5, 22, 0, 0, 300, 303, 5, 51, 0, 0, 301, 302, 5,
		10, 0, 0, 302, 304, 3, 50, 25, 0, 303, 301, 1, 0, 0, 0, 303, 304, 1, 0,
		0, 0, 304, 306, 1, 0, 0, 0, 305, 307, 5, 20, 0, 0, 306, 305, 1, 0, 0, 0,
		306, 307, 1, 0, 0, 0, 307, 43, 1, 0, 0, 0, 308, 309, 5, 51, 0, 0, 309,
		310, 5, 10, 0, 0, 310, 312, 3, 50, 25, 0, 311, 313, 5, 20, 0, 0, 312, 311,
		1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 45, 1, 0, 0, 0, 314, 321, 3, 50,
		25, 0, 315, 316, 5, 45, 0, 0, 316, 317, 3, 50, 25, 0, 317, 318, 5, 46,
		0, 0, 318, 322, 1, 0, 0, 0, 319, 320, 5, 9, 0, 0, 320, 322, 5, 51, 0, 0,
		321, 315, 1, 0, 0, 0, 321, 319, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323,
		321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326,
		5, 10, 0, 0, 326, 328, 3, 50, 25, 0, 327, 329, 5, 20, 0, 0, 328, 327, 1,
		0, 0, 0, 328, 329, 1, 0, 0, 0, 329, 47, 1, 0, 0, 0, 330, 331, 5, 23, 0,
		0, 331, 332, 5, 17, 0, 0, 332, 333, 3, 50, 25, 0, 333, 335, 5, 18, 0, 0,
		334, 336, 5, 20, 0, 0, 335, 334, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336,
		49, 1, 0, 0, 0, 337, 338, 3, 52, 26, 0, 338, 51, 1, 0, 0, 0, 339, 344,
		3, 54, 27, 0, 340, 341, 5, 8, 0, 0, 341, 343, 3, 54, 27, 0, 342, 340, 1,
		0, 0, 0, 343, 346, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344, 345, 1, 0, 0,
		0, 345, 53, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 347, 352, 3, 56, 28, 0, 348,
		349, 5, 7, 0, 0, 349, 351, 3, 56, 28, 0, 350, 348, 1, 0, 0, 0, 351, 354,
		1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 55, 1, 0,
		0, 0, 354, 352, 1, 0, 0, 0, 355, 360, 3, 58, 29, 0, 356, 357, 7, 0, 0,
		0, 357, 359, 3, 58, 29, 0, 358, 356, 1, 0, 0, 0, 359, 362, 1, 0, 0, 0,
		360, 358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 57, 1, 0, 0, 0, 362, 360,
		1, 0, 0, 0, 363, 368, 3, 60, 30, 0, 364, 365, 7, 1, 0, 0, 365, 367, 3,
		60, 30, 0, 366, 364, 1, 0, 0, 0, 367, 370, 1, 0, 0, 0, 368, 366, 1, 0,
		0, 0, 368, 369, 1, 0, 0, 0, 369, 59, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0,
		371, 376, 3, 62, 31, 0, 372, 373, 7, 2, 0, 0, 373, 375, 3, 62, 31, 0, 374,
		372, 1, 0, 0, 0, 375, 378, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 377,
		1, 0, 0, 0, 377, 61, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 379, 384, 3, 64,
		32, 0, 380, 381, 7, 3, 0, 0, 381, 383, 3, 64, 32, 0, 382, 380, 1, 0, 0,
		0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385,
		63, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 388, 7, 4, 0, 0, 388, 391, 3,
		64, 32, 0, 389, 391, 3, 66, 33, 0, 390, 387, 1, 0, 0, 0, 390, 389, 1, 0,
		0, 0, 391, 65, 1, 0, 0, 0, 392, 406, 3, 68, 34, 0, 393, 394, 5, 45, 0,
		0, 394, 395, 3, 50, 25, 0, 395, 396, 5, 46, 0, 0, 396, 405, 1, 0, 0, 0,
		397, 398, 5, 9, 0, 0, 398, 405, 5, 51, 0, 0, 399, 401, 5, 17, 0, 0, 400,
		402, 3, 26, 13, 0, 401, 400, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402, 403,
		1, 0, 0, 0, 403, 405, 5, 18, 0, 0, 404, 393, 1, 0, 0, 0, 404, 397, 1, 0,
		0, 0, 404, 399, 1, 0, 0, 0, 405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0,
		406, 407, 1, 0, 0, 0, 407, 67, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 439,
		5, 52, 0, 0, 410, 439, 5, 50, 0, 0, 411, 439, 5, 53, 0, 0, 412, 439, 5,
		32, 0, 0, 413, 439, 5, 41, 0, 0, 414, 439, 3, 72, 36, 0, 415, 439, 3, 74,
		37, 0, 416, 439, 3, 70, 35, 0, 417, 418, 5, 33, 0, 0, 418, 420, 5, 17,
		0, 0, 419, 421, 3, 18, 9, 0, 420, 419, 1, 0, 0, 0, 420, 421, 1, 0, 0, 0,
		421, 422, 1, 0, 0, 0, 422, 423, 5, 18, 0, 0, 423, 439, 3, 36, 18, 0, 424,
		425, 5, 51, 0, 0, 425, 427, 5, 17, 0, 0, 426, 428, 3, 26, 13, 0, 427, 426,
		1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 439, 5, 18,
		0, 0, 430, 432, 5, 51, 0, 0, 431, 433, 7, 5, 0, 0, 432, 431, 1, 0, 0, 0,
		432, 433, 1, 0, 0, 0, 433, 439, 1, 0, 0, 0, 434, 435, 5, 17, 0, 0, 435,
		436, 3, 50, 25, 0, 436, 437, 5, 18, 0, 0, 437, 439, 1, 0, 0, 0, 438, 409,
		1, 0, 0, 0, 438, 410, 1, 0, 0, 0, 438, 411, 1, 0, 0, 0, 438, 412, 1, 0,
		0, 0, 438, 413, 1, 0, 0, 0, 438, 414, 1, 0, 0, 0, 438, 415, 1, 0, 0, 0,
		438, 416, 1, 0, 0, 0, 438, 417, 1, 0, 0, 0, 438, 424, 1, 0, 0, 0, 438,
		430, 1, 0, 0, 0, 438, 434, 1, 0, 0, 0, 439, 69, 1, 0, 0, 0, 440, 441, 5,
		42, 0, 0, 441, 442, 3, 50, 25, 0, 442, 448, 5, 43, 0, 0, 443, 445, 5, 17,
		0, 0, 444, 446, 5, 51, 0, 0, 445, 444, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0,
		446, 447, 1, 0, 0, 0, 447, 449, 5, 18, 0, 0, 448, 443, 1, 0, 0, 0, 448,
		449, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 451, 3, 36, 18, 0, 451, 71,
		1, 0, 0, 0, 452, 461, 5, 45, 0, 0, 453, 458, 3, 50, 25, 0, 454, 455, 5,
		44, 0, 0, 455, 457, 3, 50, 25, 0, 456, 454, 1, 0, 0, 0, 457, 460, 1, 0,
		0, 0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 462, 1, 0, 0, 0,
		460, 458, 1, 0, 0, 0, 461, 453, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0, 462,
		463, 1, 0, 0, 0, 463, 464, 5, 46, 0, 0, 464, 73, 1, 0, 0, 0, 465, 474,
		5, 48, 0, 0, 466, 471, 3, 76, 38, 0, 467, 468, 5, 44, 0, 0, 468, 470, 3,
		76, 38, 0, 469, 467, 1, 0, 0, 0, 470, 473, 1, 0, 0, 0, 471, 469, 1, 0,
		0, 0, 471, 472, 1, 0, 0, 0, 472, 475, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0,
		474, 466, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476,
		477, 5, 49, 0, 0, 477, 75, 1, 0, 0, 0, 478, 479, 7, 6, 0, 0, 479, 480,
		5, 47, 0, 0, 480, 481, 3, 50, 25, 0, 481, 77, 1, 0, 0, 0, 56, 81, 102,
		106, 121, 126, 141, 145, 149, 155, 165, 170, 173, 178, 183, 190, 197, 203,
		209, 214, 218, 222, 248, 259, 265, 276, 284, 287, 293, 297, 303, 306, 312,
		321, 323, 328, 335, 344, 352, 360, 368, 376, 384, 390, 401, 404, 406, 420,
		427, 432, 438, 445, 448, 458, 461, 471, 474,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FigParserInit initializes any static state used to implement FigParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFigParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FigParserInit() {
	staticData := &FigParserParserStaticData
	staticData.once.Do(figparserParserInit)
}

// NewFigParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFigParser(input antlr.TokenStream) *FigParser {
	FigParserInit()
	this := new(FigParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FigParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "FigParser.g4"

	return this
}

// FigParser tokens.
const (
	FigParserEOF          = antlr.TokenEOF
	FigParserLT           = 1
	FigParserGT           = 2
	FigParserLE           = 3
	FigParserGE           = 4
	FigParserEQ           = 5
	FigParserNEQ          = 6
	FigParserAND          = 7
	FigParserOR           = 8
	FigParserDOT          = 9
	FigParserASSIGN       = 10
	FigParserPLUS         = 11
	FigParserMINUS        = 12
	FigParserPLUSPLUS     = 13
	FigParserMINUSMINUS   = 14
	FigParserSTAR         = 15
	FigParserSLASH        = 16
	FigParserLPAREN       = 17
	FigParserRPAREN       = 18
	FigParserEXCLAM       = 19
	FigParserSEMICOLON    = 20
	FigParserMOD          = 21
	FigParserTK_LET       = 22
	FigParserTK_PRINT     = 23
	FigParserTK_IF        = 24
	FigParserTK_ELIF      = 25
	FigParserTK_ELSE      = 26
	FigParserTK_WHILE     = 27
	FigParserTK_DO        = 28
	FigParserTK_BREAK     = 29
	FigParserTK_CONTINUE  = 30
	FigParserTK_FOR       = 31
	FigParserTK_NULL      = 32
	FigParserTK_FN        = 33
	FigParserTK_RETURN    = 34
	FigParserTK_IMPORT    = 35
	FigParserTK_USE       = 36
	FigParserTK_IN        = 37
	FigParserTK_RANGE     = 38
	FigParserTK_ENUMERATE = 39
	FigParserTK_STRUCT    = 40
	FigParserTK_THIS      = 41
	FigParserTK_TRY       = 42
	FigParserTK_ONERROR   = 43
	FigParserCOMMA        = 44
	FigParserLBRACKET     = 45
	FigParserRBRACKET     = 46
	FigParserCOLON        = 47
	FigParserLBRACE       = 48
	FigParserRBRACE       = 49
	FigParserBOOL         = 50
	FigParserID           = 51
	FigParserNUMBER       = 52
	FigParserSTRING       = 53
	FigParserWS           = 54
	FigParserCOMMENT      = 55
)

// FigParser rules.
const (
	FigParserRULE_program        = 0
	FigParserRULE_statements     = 1
	FigParserRULE_exprStmt       = 2
	FigParserRULE_ifStmt         = 3
	FigParserRULE_whileStmt      = 4
	FigParserRULE_doWhileStmt    = 5
	FigParserRULE_breakStmt      = 6
	FigParserRULE_continueStmt   = 7
	FigParserRULE_fnDecl         = 8
	FigParserRULE_fnParams       = 9
	FigParserRULE_returnStmt     = 10
	FigParserRULE_importStmt     = 11
	FigParserRULE_useStmt        = 12
	FigParserRULE_fnArgs         = 13
	FigParserRULE_forInit        = 14
	FigParserRULE_forStep        = 15
	FigParserRULE_forStmt        = 16
	FigParserRULE_forInStmt      = 17
	FigParserRULE_block          = 18
	FigParserRULE_structDecl     = 19
	FigParserRULE_structMember   = 20
	FigParserRULE_varDeclaration = 21
	FigParserRULE_varAtribuition = 22
	FigParserRULE_memberAssign   = 23
	FigParserRULE_printStmt      = 24
	FigParserRULE_expr           = 25
	FigParserRULE_logicalOr      = 26
	FigParserRULE_logicalAnd     = 27
	FigParserRULE_equality       = 28
	FigParserRULE_comparison     = 29
	FigParserRULE_term           = 30
	FigParserRULE_factor         = 31
	FigParserRULE_unary          = 32
	FigParserRULE_postfix        = 33
	FigParserRULE_primary        = 34
	FigParserRULE_tryExpr        = 35
	FigParserRULE_arrayLiteral   = 36
	FigParserRULE_objectLiteral  = 37
	FigParserRULE_objectEntry    = 38
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatements() []IStatementsContext
	Statements(i int) IStatementsContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(FigParserEOF, 0)
}

func (s *ProgramContext) AllStatements() []IStatementsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementsContext); ok {
			len++
		}
	}

	tst := make([]IStatementsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementsContext); ok {
			tst[i] = t.(IStatementsContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statements(i int) IStatementsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FigParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17212991867613184) != 0) {
		{
			p.SetState(78)
			p.Statements()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(83)
		p.Match(FigParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDeclaration() IVarDeclarationContext
	VarAtribuition() IVarAtribuitionContext
	MemberAssign() IMemberAssignContext
	PrintStmt() IPrintStmtContext
	IfStmt() IIfStmtContext
	WhileStmt() IWhileStmtContext
	DoWhileStmt() IDoWhileStmtContext
	ForStmt() IForStmtContext
	ForInStmt() IForInStmtContext
	BreakStmt() IBreakStmtContext
	ContinueStmt() IContinueStmtContext
	FnDecl() IFnDeclContext
	ReturnStmt() IReturnStmtContext
	ImportStmt() IImportStmtContext
	UseStmt() IUseStmtContext
	StructDecl() IStructDeclContext
	ExprStmt() IExprStmtContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) VarDeclaration() IVarDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *StatementsContext) VarAtribuition() IVarAtribuitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAtribuitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarAtribuitionContext)
}

func (s *StatementsContext) MemberAssign() IMemberAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemberAssignContext)
}

func (s *StatementsContext) PrintStmt() IPrintStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStmtContext)
}

func (s *StatementsContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementsContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StatementsContext) DoWhileStmt() IDoWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoWhileStmtContext)
}

func (s *StatementsContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementsContext) ForInStmt() IForInStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInStmtContext)
}

func (s *StatementsContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementsContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementsContext) FnDecl() IFnDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnDeclContext)
}

func (s *StatementsContext) ReturnStmt() IReturnStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementsContext) ImportStmt() IImportStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStmtContext)
}

func (s *StatementsContext) UseStmt() IUseStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseStmtContext)
}

func (s *StatementsContext) StructDecl() IStructDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *StatementsContext) ExprStmt() IExprStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStmtContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FigParserRULE_statements)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.VarDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.VarAtribuition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.MemberAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
			p.PrintStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(89)
			p.IfStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(90)
			p.WhileStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(91)
			p.DoWhileStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(92)
			p.ForStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(93)
			p.ForInStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(94)
			p.BreakStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(95)
			p.ContinueStmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(96)
			p.FnDecl()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(97)
			p.ReturnStmt()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(98)
			p.ImportStmt()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(99)
			p.UseStmt()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(100)
			p.StructDecl()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(101)
			p.ExprStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprStmtContext is an interface to support dynamic dispatch.
type IExprStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsExprStmtContext differentiates from other interfaces.
	IsExprStmtContext()
}

type ExprStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStmtContext() *ExprStmtContext {
	var p = new(ExprStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_exprStmt
	return p
}

func InitEmptyExprStmtContext(p *ExprStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_exprStmt
}

func (*ExprStmtContext) IsExprStmtContext() {}

func NewExprStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStmtContext {
	var p = new(ExprStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_exprStmt

	return p
}

func (s *ExprStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *ExprStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterExprStmt(s)
	}
}

func (s *ExprStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitExprStmt(s)
	}
}

func (s *ExprStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitExprStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ExprStmt() (localctx IExprStmtContext) {
	localctx = NewExprStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FigParserRULE_exprStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Expr()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(105)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_IF() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllTK_ELIF() []antlr.TerminalNode
	TK_ELIF(i int) antlr.TerminalNode
	TK_ELSE() antlr.TerminalNode

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) TK_IF() antlr.TerminalNode {
	return s.GetToken(FigParserTK_IF, 0)
}

func (s *IfStmtContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(FigParserLPAREN)
}

func (s *IfStmtContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, i)
}

func (s *IfStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(FigParserRPAREN)
}

func (s *IfStmtContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, i)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) AllTK_ELIF() []antlr.TerminalNode {
	return s.GetTokens(FigParserTK_ELIF)
}

func (s *IfStmtContext) TK_ELIF(i int) antlr.TerminalNode {
	return s.GetToken(FigParserTK_ELIF, i)
}

func (s *IfStmtContext) TK_ELSE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_ELSE, 0)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FigParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(FigParserTK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Expr()
	}
	{
		p.SetState(111)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Block()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_ELIF {
		{
			p.SetState(113)
			p.Match(FigParserTK_ELIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Expr()
		}
		{
			p.SetState(116)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Block()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserTK_ELSE {
		{
			p.SetState(124)
			p.Match(FigParserTK_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_WHILE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_whileStmt
	return p
}

func InitEmptyWhileStmtContext(p *WhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_whileStmt
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) TK_WHILE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_WHILE, 0)
}

func (s *WhileStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *WhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) WhileStmt() (localctx IWhileStmtContext) {
	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FigParserRULE_whileStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(FigParserTK_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Expr()
	}
	{
		p.SetState(131)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDoWhileStmtContext is an interface to support dynamic dispatch.
type IDoWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_DO() antlr.TerminalNode
	Block() IBlockContext
	TK_WHILE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsDoWhileStmtContext differentiates from other interfaces.
	IsDoWhileStmtContext()
}

type DoWhileStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoWhileStmtContext() *DoWhileStmtContext {
	var p = new(DoWhileStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_doWhileStmt
	return p
}

func InitEmptyDoWhileStmtContext(p *DoWhileStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_doWhileStmt
}

func (*DoWhileStmtContext) IsDoWhileStmtContext() {}

func NewDoWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoWhileStmtContext {
	var p = new(DoWhileStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_doWhileStmt

	return p
}

func (s *DoWhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DoWhileStmtContext) TK_DO() antlr.TerminalNode {
	return s.GetToken(FigParserTK_DO, 0)
}

func (s *DoWhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DoWhileStmtContext) TK_WHILE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_WHILE, 0)
}

func (s *DoWhileStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *DoWhileStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DoWhileStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *DoWhileStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *DoWhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoWhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoWhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterDoWhileStmt(s)
	}
}

func (s *DoWhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitDoWhileStmt(s)
	}
}

func (s *DoWhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitDoWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) DoWhileStmt() (localctx IDoWhileStmtContext) {
	localctx = NewDoWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FigParserRULE_doWhileStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(FigParserTK_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Block()
	}
	{
		p.SetState(136)
		p.Match(FigParserTK_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Expr()
	}
	{
		p.SetState(139)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(140)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_BREAK() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) TK_BREAK() antlr.TerminalNode {
	return s.GetToken(FigParserTK_BREAK, 0)
}

func (s *BreakStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FigParserRULE_breakStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(FigParserTK_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(144)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_CONTINUE() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) TK_CONTINUE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_CONTINUE, 0)
}

func (s *ContinueStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FigParserRULE_continueStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(FigParserTK_CONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(148)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnDeclContext is an interface to support dynamic dispatch.
type IFnDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_FN() antlr.TerminalNode
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	FnParams() IFnParamsContext

	// IsFnDeclContext differentiates from other interfaces.
	IsFnDeclContext()
}

type FnDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnDeclContext() *FnDeclContext {
	var p = new(FnDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_fnDecl
	return p
}

func InitEmptyFnDeclContext(p *FnDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_fnDecl
}

func (*FnDeclContext) IsFnDeclContext() {}

func NewFnDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnDeclContext {
	var p = new(FnDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_fnDecl

	return p
}

func (s *FnDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FnDeclContext) TK_FN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FN, 0)
}

func (s *FnDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *FnDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *FnDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *FnDeclContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FnDeclContext) FnParams() IFnParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnParamsContext)
}

func (s *FnDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterFnDecl(s)
	}
}

func (s *FnDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitFnDecl(s)
	}
}

func (s *FnDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitFnDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) FnDecl() (localctx IFnDeclContext) {
	localctx = NewFnDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FigParserRULE_fnDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(FigParserTK_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID {
		{
			p.SetState(154)
			p.FnParams()
		}

	}
	{
		p.SetState(157)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnParamsContext is an interface to support dynamic dispatch.
type IFnParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFnParamsContext differentiates from other interfaces.
	IsFnParamsContext()
}

type FnParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnParamsContext() *FnParamsContext {
	var p = new(FnParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_fnParams
	return p
}

func InitEmptyFnParamsContext(p *FnParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_fnParams
}

func (*FnParamsContext) IsFnParamsContext() {}

func NewFnParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnParamsContext {
	var p = new(FnParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_fnParams

	return p
}

func (s *FnParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FnParamsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(FigParserID)
}

func (s *FnParamsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(FigParserID, i)
}

func (s *FnParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *FnParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *FnParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterFnParams(s)
	}
}

func (s *FnParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitFnParams(s)
	}
}

func (s *FnParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitFnParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) FnParams() (localctx IFnParamsContext) {
	localctx = NewFnParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FigParserRULE_fnParams)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(161)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_RETURN() antlr.TerminalNode
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_returnStmt
	return p
}

func InitEmptyReturnStmtContext(p *ReturnStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_returnStmt
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) TK_RETURN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_RETURN, 0)
}

func (s *ReturnStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FigParserRULE_returnStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(FigParserTK_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(169)
			p.Expr()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(172)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_IMPORT() antlr.TerminalNode
	STRING() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_importStmt
	return p
}

func InitEmptyImportStmtContext(p *ImportStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_importStmt
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_importStmt

	return p
}

func (s *ImportStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStmtContext) TK_IMPORT() antlr.TerminalNode {
	return s.GetToken(FigParserTK_IMPORT, 0)
}

func (s *ImportStmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(FigParserSTRING, 0)
}

func (s *ImportStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterImportStmt(s)
	}
}

func (s *ImportStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitImportStmt(s)
	}
}

func (s *ImportStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitImportStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ImportStmt() (localctx IImportStmtContext) {
	localctx = NewImportStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FigParserRULE_importStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(FigParserTK_IMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(FigParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(177)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUseStmtContext is an interface to support dynamic dispatch.
type IUseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_USE() antlr.TerminalNode
	STRING() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsUseStmtContext differentiates from other interfaces.
	IsUseStmtContext()
}

type UseStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseStmtContext() *UseStmtContext {
	var p = new(UseStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_useStmt
	return p
}

func InitEmptyUseStmtContext(p *UseStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_useStmt
}

func (*UseStmtContext) IsUseStmtContext() {}

func NewUseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseStmtContext {
	var p = new(UseStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_useStmt

	return p
}

func (s *UseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UseStmtContext) TK_USE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_USE, 0)
}

func (s *UseStmtContext) STRING() antlr.TerminalNode {
	return s.GetToken(FigParserSTRING, 0)
}

func (s *UseStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *UseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterUseStmt(s)
	}
}

func (s *UseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitUseStmt(s)
	}
}

func (s *UseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitUseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) UseStmt() (localctx IUseStmtContext) {
	localctx = NewUseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FigParserRULE_useStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(FigParserTK_USE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(FigParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(182)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFnArgsContext is an interface to support dynamic dispatch.
type IFnArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFnArgsContext differentiates from other interfaces.
	IsFnArgsContext()
}

type FnArgsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnArgsContext() *FnArgsContext {
	var p = new(FnArgsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_fnArgs
	return p
}

func InitEmptyFnArgsContext(p *FnArgsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_fnArgs
}

func (*FnArgsContext) IsFnArgsContext() {}

func NewFnArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnArgsContext {
	var p = new(FnArgsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_fnArgs

	return p
}

func (s *FnArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FnArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *FnArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FnArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *FnArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *FnArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterFnArgs(s)
	}
}

func (s *FnArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitFnArgs(s)
	}
}

func (s *FnArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitFnArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) FnArgs() (localctx IFnArgsContext) {
	localctx = NewFnArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FigParserRULE_fnArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Expr()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(186)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Expr()
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_LET() antlr.TerminalNode
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_forInit
	return p
}

func InitEmptyForInitContext(p *ForInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_forInit
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) TK_LET() antlr.TerminalNode {
	return s.GetToken(FigParserTK_LET, 0)
}

func (s *ForInitContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ForInitContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FigParserASSIGN, 0)
}

func (s *ForInitContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FigParserRULE_forInit)
	var _la int

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Match(FigParserTK_LET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(195)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(196)
				p.Expr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(202)
			p.Expr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStepContext is an interface to support dynamic dispatch.
type IForStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext

	// IsForStepContext differentiates from other interfaces.
	IsForStepContext()
}

type ForStepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStepContext() *ForStepContext {
	var p = new(ForStepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_forStep
	return p
}

func InitEmptyForStepContext(p *ForStepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_forStep
}

func (*ForStepContext) IsForStepContext() {}

func NewForStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStepContext {
	var p = new(ForStepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_forStep

	return p
}

func (s *ForStepContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStepContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ForStepContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FigParserASSIGN, 0)
}

func (s *ForStepContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterForStep(s)
	}
}

func (s *ForStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitForStep(s)
	}
}

func (s *ForStepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitForStep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ForStep() (localctx IForStepContext) {
	localctx = NewForStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FigParserRULE_forStep)
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Expr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_FOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	ForInit() IForInitContext
	Expr() IExprContext
	ForStep() IForStepContext

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_forStmt
	return p
}

func InitEmptyForStmtContext(p *ForStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_forStmt
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) TK_FOR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FOR, 0)
}

func (s *ForStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *ForStmtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(FigParserSEMICOLON)
}

func (s *ForStmtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, i)
}

func (s *ForStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *ForStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStmtContext) ForInit() IForInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForStmtContext) ForStep() IForStepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStepContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FigParserRULE_forStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(FigParserTK_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17211767910985728) != 0 {
		{
			p.SetState(213)
			p.ForInit()
		}

	}
	{
		p.SetState(216)
		p.Match(FigParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17211767906791424) != 0 {
		{
			p.SetState(217)
			p.Expr()
		}

	}
	{
		p.SetState(220)
		p.Match(FigParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17211767906791424) != 0 {
		{
			p.SetState(221)
			p.ForStep()
		}

	}
	{
		p.SetState(224)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForInStmtContext is an interface to support dynamic dispatch.
type IForInStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForInStmtContext differentiates from other interfaces.
	IsForInStmtContext()
}

type ForInStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInStmtContext() *ForInStmtContext {
	var p = new(ForInStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_forInStmt
	return p
}

func InitEmptyForInStmtContext(p *ForInStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_forInStmt
}

func (*ForInStmtContext) IsForInStmtContext() {}

func NewForInStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInStmtContext {
	var p = new(ForInStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_forInStmt

	return p
}

func (s *ForInStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInStmtContext) CopyAll(ctx *ForInStmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForInStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForRangeContext struct {
	ForInStmtContext
}

func NewForRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeContext {
	var p = new(ForRangeContext)

	InitEmptyForInStmtContext(&p.ForInStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForInStmtContext))

	return p
}

func (s *ForRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeContext) TK_FOR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FOR, 0)
}

func (s *ForRangeContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ForRangeContext) TK_IN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_IN, 0)
}

func (s *ForRangeContext) TK_RANGE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_RANGE, 0)
}

func (s *ForRangeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *ForRangeContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForRangeContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForRangeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ForRangeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *ForRangeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *ForRangeContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterForRange(s)
	}
}

func (s *ForRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitForRange(s)
	}
}

func (s *ForRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitForRange(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForEnumerateContext struct {
	ForInStmtContext
}

func NewForEnumerateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForEnumerateContext {
	var p = new(ForEnumerateContext)

	InitEmptyForInStmtContext(&p.ForInStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForInStmtContext))

	return p
}

func (s *ForEnumerateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForEnumerateContext) TK_FOR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FOR, 0)
}

func (s *ForEnumerateContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(FigParserID)
}

func (s *ForEnumerateContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(FigParserID, i)
}

func (s *ForEnumerateContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, 0)
}

func (s *ForEnumerateContext) TK_IN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_IN, 0)
}

func (s *ForEnumerateContext) TK_ENUMERATE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_ENUMERATE, 0)
}

func (s *ForEnumerateContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *ForEnumerateContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForEnumerateContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *ForEnumerateContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForEnumerateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterForEnumerate(s)
	}
}

func (s *ForEnumerateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitForEnumerate(s)
	}
}

func (s *ForEnumerateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitForEnumerate(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForInContext struct {
	ForInStmtContext
}

func NewForInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInContext {
	var p = new(ForInContext)

	InitEmptyForInStmtContext(&p.ForInStmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForInStmtContext))

	return p
}

func (s *ForInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInContext) TK_FOR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FOR, 0)
}

func (s *ForInContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ForInContext) TK_IN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_IN, 0)
}

func (s *ForInContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForInContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterForIn(s)
	}
}

func (s *ForInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitForIn(s)
	}
}

func (s *ForInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitForIn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ForInStmt() (localctx IForInStmtContext) {
	localctx = NewForInStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FigParserRULE_forInStmt)
	var _la int

	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(FigParserTK_ENUMERATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Expr()
		}
		{
			p.SetState(235)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Block()
		}

	case 2:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(FigParserTK_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Expr()
		}
		{
			p.SetState(244)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Expr()
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserCOMMA {
			{
				p.SetState(246)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(247)
				p.Expr()
			}

		}
		{
			p.SetState(250)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.Block()
		}

	case 3:
		localctx = NewForInContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(253)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Expr()
		}
		{
			p.SetState(257)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatements() []IStatementsContext
	Statements(i int) IStatementsContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACE, 0)
}

func (s *BlockContext) AllStatements() []IStatementsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementsContext); ok {
			len++
		}
	}

	tst := make([]IStatementsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementsContext); ok {
			tst[i] = t.(IStatementsContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statements(i int) IStatementsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FigParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17212991867613184) != 0 {
		{
			p.SetState(262)
			p.Statements()
		}

		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(268)
		p.Match(FigParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStructMember() []IStructMemberContext
	StructMember(i int) IStructMemberContext

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_structDecl
	return p
}

func InitEmptyStructDeclContext(p *StructDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_structDecl
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) TK_STRUCT() antlr.TerminalNode {
	return s.GetToken(FigParserTK_STRUCT, 0)
}

func (s *StructDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *StructDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACE, 0)
}

func (s *StructDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACE, 0)
}

func (s *StructDeclContext) AllStructMember() []IStructMemberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructMemberContext); ok {
			len++
		}
	}

	tst := make([]IStructMemberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructMemberContext); ok {
			tst[i] = t.(IStructMemberContext)
			i++
		}
	}

	return tst
}

func (s *StructDeclContext) StructMember(i int) IStructMemberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemberContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructMemberContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) StructDecl() (localctx IStructDeclContext) {
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FigParserRULE_structDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(FigParserTK_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_FN || _la == FigParserID {
		{
			p.SetState(273)
			p.StructMember()
		}

		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(279)
		p.Match(FigParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructMemberContext is an interface to support dynamic dispatch.
type IStructMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStructMemberContext differentiates from other interfaces.
	IsStructMemberContext()
}

type StructMemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemberContext() *StructMemberContext {
	var p = new(StructMemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_structMember
	return p
}

func InitEmptyStructMemberContext(p *StructMemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_structMember
}

func (*StructMemberContext) IsStructMemberContext() {}

func NewStructMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemberContext {
	var p = new(StructMemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_structMember

	return p
}

func (s *StructMemberContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemberContext) CopyAll(ctx *StructMemberContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StructMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructFieldContext struct {
	StructMemberContext
}

func NewStructFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructFieldContext {
	var p = new(StructFieldContext)

	InitEmptyStructMemberContext(&p.StructMemberContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructMemberContext))

	return p
}

func (s *StructFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *StructFieldContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FigParserASSIGN, 0)
}

func (s *StructFieldContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StructFieldContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *StructFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterStructField(s)
	}
}

func (s *StructFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitStructField(s)
	}
}

func (s *StructFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitStructField(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructMethodContext struct {
	StructMemberContext
}

func NewStructMethodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructMethodContext {
	var p = new(StructMethodContext)

	InitEmptyStructMemberContext(&p.StructMemberContext)
	p.parser = parser
	p.CopyAll(ctx.(*StructMemberContext))

	return p
}

func (s *StructMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMethodContext) TK_FN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FN, 0)
}

func (s *StructMethodContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *StructMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *StructMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *StructMethodContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StructMethodContext) FnParams() IFnParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnParamsContext)
}

func (s *StructMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterStructMethod(s)
	}
}

func (s *StructMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitStructMethod(s)
	}
}

func (s *StructMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitStructMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) StructMember() (localctx IStructMemberContext) {
	localctx = NewStructMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FigParserRULE_structMember)
	var _la int

	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		localctx = NewStructFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(282)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(283)
				p.Expr()
			}

		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserSEMICOLON {
			{
				p.SetState(286)
				p.Match(FigParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case FigParserTK_FN:
		localctx = NewStructMethodContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(289)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(292)
				p.FnParams()
			}

		}
		{
			p.SetState(295)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.Block()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_LET() antlr.TerminalNode
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_varDeclaration
	return p
}

func InitEmptyVarDeclarationContext(p *VarDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_varDeclaration
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) TK_LET() antlr.TerminalNode {
	return s.GetToken(FigParserTK_LET, 0)
}

func (s *VarDeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *VarDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FigParserASSIGN, 0)
}

func (s *VarDeclarationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitVarDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) VarDeclaration() (localctx IVarDeclarationContext) {
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FigParserRULE_varDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(FigParserTK_LET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserASSIGN {
		{
			p.SetState(301)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Expr()
		}

	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(305)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarAtribuitionContext is an interface to support dynamic dispatch.
type IVarAtribuitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expr() IExprContext
	SEMICOLON() antlr.TerminalNode

	// IsVarAtribuitionContext differentiates from other interfaces.
	IsVarAtribuitionContext()
}

type VarAtribuitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAtribuitionContext() *VarAtribuitionContext {
	var p = new(VarAtribuitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_varAtribuition
	return p
}

func InitEmptyVarAtribuitionContext(p *VarAtribuitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_varAtribuition
}

func (*VarAtribuitionContext) IsVarAtribuitionContext() {}

func NewVarAtribuitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAtribuitionContext {
	var p = new(VarAtribuitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_varAtribuition

	return p
}

func (s *VarAtribuitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAtribuitionContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *VarAtribuitionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FigParserASSIGN, 0)
}

func (s *VarAtribuitionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarAtribuitionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *VarAtribuitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAtribuitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAtribuitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterVarAtribuition(s)
	}
}

func (s *VarAtribuitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitVarAtribuition(s)
	}
}

func (s *VarAtribuitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitVarAtribuition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) VarAtribuition() (localctx IVarAtribuitionContext) {
	localctx = NewVarAtribuitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FigParserRULE_varAtribuition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Expr()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(311)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMemberAssignContext is an interface to support dynamic dispatch.
type IMemberAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	ASSIGN() antlr.TerminalNode
	AllLBRACKET() []antlr.TerminalNode
	LBRACKET(i int) antlr.TerminalNode
	AllRBRACKET() []antlr.TerminalNode
	RBRACKET(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsMemberAssignContext differentiates from other interfaces.
	IsMemberAssignContext()
}

type MemberAssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberAssignContext() *MemberAssignContext {
	var p = new(MemberAssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_memberAssign
	return p
}

func InitEmptyMemberAssignContext(p *MemberAssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_memberAssign
}

func (*MemberAssignContext) IsMemberAssignContext() {}

func NewMemberAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberAssignContext {
	var p = new(MemberAssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_memberAssign

	return p
}

func (s *MemberAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberAssignContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MemberAssignContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MemberAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FigParserASSIGN, 0)
}

func (s *MemberAssignContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(FigParserLBRACKET)
}

func (s *MemberAssignContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(FigParserLBRACKET, i)
}

func (s *MemberAssignContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(FigParserRBRACKET)
}

func (s *MemberAssignContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, i)
}

func (s *MemberAssignContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(FigParserDOT)
}

func (s *MemberAssignContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(FigParserDOT, i)
}

func (s *MemberAssignContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(FigParserID)
}

func (s *MemberAssignContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(FigParserID, i)
}

func (s *MemberAssignContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *MemberAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterMemberAssign(s)
	}
}

func (s *MemberAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitMemberAssign(s)
	}
}

func (s *MemberAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitMemberAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) MemberAssign() (localctx IMemberAssignContext) {
	localctx = NewMemberAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FigParserRULE_memberAssign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Expr()
	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FigParserDOT || _la == FigParserLBRACKET {
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case FigParserLBRACKET:
			{
				p.SetState(315)
				p.Match(FigParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(316)
				p.Expr()
			}
			{
				p.SetState(317)
				p.Match(FigParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case FigParserDOT:
			{
				p.SetState(319)
				p.Match(FigParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(320)
				p.Match(FigParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(325)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(326)
		p.Expr()
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(327)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStmtContext is an interface to support dynamic dispatch.
type IPrintStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_PRINT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expr() IExprContext
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsPrintStmtContext differentiates from other interfaces.
	IsPrintStmtContext()
}

type PrintStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStmtContext() *PrintStmtContext {
	var p = new(PrintStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_printStmt
	return p
}

func InitEmptyPrintStmtContext(p *PrintStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_printStmt
}

func (*PrintStmtContext) IsPrintStmtContext() {}

func NewPrintStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStmtContext {
	var p = new(PrintStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_printStmt

	return p
}

func (s *PrintStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStmtContext) TK_PRINT() antlr.TerminalNode {
	return s.GetToken(FigParserTK_PRINT, 0)
}

func (s *PrintStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *PrintStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *PrintStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterPrintStmt(s)
	}
}

func (s *PrintStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitPrintStmt(s)
	}
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) PrintStmt() (localctx IPrintStmtContext) {
	localctx = NewPrintStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FigParserRULE_printStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(FigParserTK_PRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(332)
		p.Expr()
	}
	{
		p.SetState(333)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(334)
			p.Match(FigParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOr() ILogicalOrContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) LogicalOr() ILogicalOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FigParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.LogicalOr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalOrContext is an interface to support dynamic dispatch.
type ILogicalOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAnd() []ILogicalAndContext
	LogicalAnd(i int) ILogicalAndContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLogicalOrContext differentiates from other interfaces.
	IsLogicalOrContext()
}

type LogicalOrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrContext() *LogicalOrContext {
	var p = new(LogicalOrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_logicalOr
	return p
}

func InitEmptyLogicalOrContext(p *LogicalOrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_logicalOr
}

func (*LogicalOrContext) IsLogicalOrContext() {}

func NewLogicalOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrContext {
	var p = new(LogicalOrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_logicalOr

	return p
}

func (s *LogicalOrContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrContext) AllLogicalAnd() []ILogicalAndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndContext); ok {
			tst[i] = t.(ILogicalAndContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrContext) LogicalAnd(i int) ILogicalAndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndContext)
}

func (s *LogicalOrContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(FigParserOR)
}

func (s *LogicalOrContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(FigParserOR, i)
}

func (s *LogicalOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterLogicalOr(s)
	}
}

func (s *LogicalOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitLogicalOr(s)
	}
}

func (s *LogicalOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitLogicalOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) LogicalOr() (localctx ILogicalOrContext) {
	localctx = NewLogicalOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FigParserRULE_logicalOr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.LogicalAnd()
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserOR {
		{
			p.SetState(340)
			p.Match(FigParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.LogicalAnd()
		}

		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalAndContext is an interface to support dynamic dispatch.
type ILogicalAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEquality() []IEqualityContext
	Equality(i int) IEqualityContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsLogicalAndContext differentiates from other interfaces.
	IsLogicalAndContext()
}

type LogicalAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndContext() *LogicalAndContext {
	var p = new(LogicalAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_logicalAnd
	return p
}

func InitEmptyLogicalAndContext(p *LogicalAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_logicalAnd
}

func (*LogicalAndContext) IsLogicalAndContext() {}

func NewLogicalAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndContext {
	var p = new(LogicalAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_logicalAnd

	return p
}

func (s *LogicalAndContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndContext) AllEquality() []IEqualityContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEqualityContext); ok {
			len++
		}
	}

	tst := make([]IEqualityContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEqualityContext); ok {
			tst[i] = t.(IEqualityContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndContext) Equality(i int) IEqualityContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityContext)
}

func (s *LogicalAndContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(FigParserAND)
}

func (s *LogicalAndContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(FigParserAND, i)
}

func (s *LogicalAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterLogicalAnd(s)
	}
}

func (s *LogicalAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitLogicalAnd(s)
	}
}

func (s *LogicalAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitLogicalAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) LogicalAnd() (localctx ILogicalAndContext) {
	localctx = NewLogicalAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FigParserRULE_logicalAnd)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Equality()
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserAND {
		{
			p.SetState(348)
			p.Match(FigParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(349)
			p.Equality()
		}

		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityContext is an interface to support dynamic dispatch.
type IEqualityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparison() []IComparisonContext
	Comparison(i int) IComparisonContext
	AllEQ() []antlr.TerminalNode
	EQ(i int) antlr.TerminalNode
	AllNEQ() []antlr.TerminalNode
	NEQ(i int) antlr.TerminalNode

	// IsEqualityContext differentiates from other interfaces.
	IsEqualityContext()
}

type EqualityContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityContext() *EqualityContext {
	var p = new(EqualityContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_equality
	return p
}

func InitEmptyEqualityContext(p *EqualityContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_equality
}

func (*EqualityContext) IsEqualityContext() {}

func NewEqualityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityContext {
	var p = new(EqualityContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_equality

	return p
}

func (s *EqualityContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityContext) AllComparison() []IComparisonContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonContext); ok {
			len++
		}
	}

	tst := make([]IComparisonContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonContext); ok {
			tst[i] = t.(IComparisonContext)
			i++
		}
	}

	return tst
}

func (s *EqualityContext) Comparison(i int) IComparisonContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *EqualityContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(FigParserEQ)
}

func (s *EqualityContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(FigParserEQ, i)
}

func (s *EqualityContext) AllNEQ() []antlr.TerminalNode {
	return s.GetTokens(FigParserNEQ)
}

func (s *EqualityContext) NEQ(i int) antlr.TerminalNode {
	return s.GetToken(FigParserNEQ, i)
}

func (s *EqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterEquality(s)
	}
}

func (s *EqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitEquality(s)
	}
}

func (s *EqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Equality() (localctx IEqualityContext) {
	localctx = NewEqualityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FigParserRULE_equality)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Comparison()
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserEQ || _la == FigParserNEQ {
		{
			p.SetState(356)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FigParserEQ || _la == FigParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(357)
			p.Comparison()
		}

		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllGT() []antlr.TerminalNode
	GT(i int) antlr.TerminalNode
	AllGE() []antlr.TerminalNode
	GE(i int) antlr.TerminalNode
	AllLT() []antlr.TerminalNode
	LT(i int) antlr.TerminalNode
	AllLE() []antlr.TerminalNode
	LE(i int) antlr.TerminalNode

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_comparison
	return p
}

func InitEmptyComparisonContext(p *ComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_comparison
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ComparisonContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(FigParserGT)
}

func (s *ComparisonContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(FigParserGT, i)
}

func (s *ComparisonContext) AllGE() []antlr.TerminalNode {
	return s.GetTokens(FigParserGE)
}

func (s *ComparisonContext) GE(i int) antlr.TerminalNode {
	return s.GetToken(FigParserGE, i)
}

func (s *ComparisonContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(FigParserLT)
}

func (s *ComparisonContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(FigParserLT, i)
}

func (s *ComparisonContext) AllLE() []antlr.TerminalNode {
	return s.GetTokens(FigParserLE)
}

func (s *ComparisonContext) LE(i int) antlr.TerminalNode {
	return s.GetToken(FigParserLE, i)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FigParserRULE_comparison)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Term()
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0 {
		{
			p.SetState(364)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(365)
			p.Term()
		}

		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(FigParserPLUS)
}

func (s *TermContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(FigParserPLUS, i)
}

func (s *TermContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(FigParserMINUS)
}

func (s *TermContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(FigParserMINUS, i)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FigParserRULE_term)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(371)
		p.Factor()
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(372)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FigParserPLUS || _la == FigParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(373)
				p.Factor()
			}

		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnary() []IUnaryContext
	Unary(i int) IUnaryContext
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode
	AllMOD() []antlr.TerminalNode
	MOD(i int) antlr.TerminalNode

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) AllUnary() []IUnaryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryContext); ok {
			len++
		}
	}

	tst := make([]IUnaryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryContext); ok {
			tst[i] = t.(IUnaryContext)
			i++
		}
	}

	return tst
}

func (s *FactorContext) Unary(i int) IUnaryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *FactorContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(FigParserSTAR)
}

func (s *FactorContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(FigParserSTAR, i)
}

func (s *FactorContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(FigParserSLASH)
}

func (s *FactorContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(FigParserSLASH, i)
}

func (s *FactorContext) AllMOD() []antlr.TerminalNode {
	return s.GetTokens(FigParserMOD)
}

func (s *FactorContext) MOD(i int) antlr.TerminalNode {
	return s.GetToken(FigParserMOD, i)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FigParserRULE_factor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(379)
		p.Unary()
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2195456) != 0 {
		{
			p.SetState(380)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2195456) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(381)
			p.Unary()
		}

		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Unary() IUnaryContext
	MINUS() antlr.TerminalNode
	EXCLAM() antlr.TerminalNode
	PLUSPLUS() antlr.TerminalNode
	MINUSMINUS() antlr.TerminalNode
	Postfix() IPostfixContext

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *UnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FigParserMINUS, 0)
}

func (s *UnaryContext) EXCLAM() antlr.TerminalNode {
	return s.GetToken(FigParserEXCLAM, 0)
}

func (s *UnaryContext) PLUSPLUS() antlr.TerminalNode {
	return s.GetToken(FigParserPLUSPLUS, 0)
}

func (s *UnaryContext) MINUSMINUS() antlr.TerminalNode {
	return s.GetToken(FigParserMINUSMINUS, 0)
}

func (s *UnaryContext) Postfix() IPostfixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPostfixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPostfixContext)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitUnary(s)
	}
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FigParserRULE_unary)
	var _la int

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserMINUS, FigParserPLUSPLUS, FigParserMINUSMINUS, FigParserEXCLAM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&552960) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(388)
			p.Unary()
		}

	case FigParserLPAREN, FigParserTK_NULL, FigParserTK_FN, FigParserTK_THIS, FigParserTK_TRY, FigParserLBRACKET, FigParserLBRACE, FigParserBOOL, FigParserID, FigParserNUMBER, FigParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
			p.Postfix()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPostfixContext is an interface to support dynamic dispatch.
type IPostfixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	AllLBRACKET() []antlr.TerminalNode
	LBRACKET(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRBRACKET() []antlr.TerminalNode
	RBRACKET(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	AllFnArgs() []IFnArgsContext
	FnArgs(i int) IFnArgsContext

	// IsPostfixContext differentiates from other interfaces.
	IsPostfixContext()
}

type PostfixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixContext() *PostfixContext {
	var p = new(PostfixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_postfix
	return p
}

func InitEmptyPostfixContext(p *PostfixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_postfix
}

func (*PostfixContext) IsPostfixContext() {}

func NewPostfixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixContext {
	var p = new(PostfixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_postfix

	return p
}

func (s *PostfixContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PostfixContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(FigParserLBRACKET)
}

func (s *PostfixContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(FigParserLBRACKET, i)
}

func (s *PostfixContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PostfixContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PostfixContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(FigParserRBRACKET)
}

func (s *PostfixContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, i)
}

func (s *PostfixContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(FigParserDOT)
}

func (s *PostfixContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(FigParserDOT, i)
}

func (s *PostfixContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(FigParserID)
}

func (s *PostfixContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(FigParserID, i)
}

func (s *PostfixContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(FigParserLPAREN)
}

func (s *PostfixContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, i)
}

func (s *PostfixContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(FigParserRPAREN)
}

func (s *PostfixContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, i)
}

func (s *PostfixContext) AllFnArgs() []IFnArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFnArgsContext); ok {
			len++
		}
	}

	tst := make([]IFnArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFnArgsContext); ok {
			tst[i] = t.(IFnArgsContext)
			i++
		}
	}

	return tst
}

func (s *PostfixContext) FnArgs(i int) IFnArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnArgsContext)
}

func (s *PostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterPostfix(s)
	}
}

func (s *PostfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitPostfix(s)
	}
}

func (s *PostfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitPostfix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Postfix() (localctx IPostfixContext) {
	localctx = NewPostfixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FigParserRULE_postfix)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(392)
		p.Primary()
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(404)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case FigParserLBRACKET:
				{
					p.SetState(393)
					p.Match(FigParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(394)
					p.Expr()
				}
				{
					p.SetState(395)
					p.Match(FigParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case FigParserDOT:
				{
					p.SetState(397)
					p.Match(FigParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(398)
					p.Match(FigParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case FigParserLPAREN:
				{
					p.SetState(399)
					p.Match(FigParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(401)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17211767906791424) != 0 {
					{
						p.SetState(400)
						p.FnArgs()
					}

				}
				{
					p.SetState(403)
					p.Match(FigParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TK_NULL() antlr.TerminalNode
	TK_THIS() antlr.TerminalNode
	ArrayLiteral() IArrayLiteralContext
	ObjectLiteral() IObjectLiteralContext
	TryExpr() ITryExprContext
	TK_FN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	FnParams() IFnParamsContext
	ID() antlr.TerminalNode
	FnArgs() IFnArgsContext
	PLUSPLUS() antlr.TerminalNode
	MINUSMINUS() antlr.TerminalNode
	Expr() IExprContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FigParserNUMBER, 0)
}

func (s *PrimaryContext) BOOL() antlr.TerminalNode {
	return s.GetToken(FigParserBOOL, 0)
}

func (s *PrimaryContext) STRING() antlr.TerminalNode {
	return s.GetToken(FigParserSTRING, 0)
}

func (s *PrimaryContext) TK_NULL() antlr.TerminalNode {
	return s.GetToken(FigParserTK_NULL, 0)
}

func (s *PrimaryContext) TK_THIS() antlr.TerminalNode {
	return s.GetToken(FigParserTK_THIS, 0)
}

func (s *PrimaryContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *PrimaryContext) ObjectLiteral() IObjectLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *PrimaryContext) TryExpr() ITryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITryExprContext)
}

func (s *PrimaryContext) TK_FN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FN, 0)
}

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *PrimaryContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PrimaryContext) FnParams() IFnParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnParamsContext)
}

func (s *PrimaryContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *PrimaryContext) FnArgs() IFnArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnArgsContext)
}

func (s *PrimaryContext) PLUSPLUS() antlr.TerminalNode {
	return s.GetToken(FigParserPLUSPLUS, 0)
}

func (s *PrimaryContext) MINUSMINUS() antlr.TerminalNode {
	return s.GetToken(FigParserMINUSMINUS, 0)
}

func (s *PrimaryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FigParserRULE_primary)
	var _la int

	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.Match(FigParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.Match(FigParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
			p.Match(FigParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(412)
			p.Match(FigParserTK_NULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(413)
			p.Match(FigParserTK_THIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(414)
			p.ArrayLiteral()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(415)
			p.ObjectLiteral()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(416)
			p.TryExpr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(417)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(418)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(419)
				p.FnParams()
			}

		}
		{
			p.SetState(422)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Block()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(424)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17211767906791424) != 0 {
			{
				p.SetState(426)
				p.FnArgs()
			}

		}
		{
			p.SetState(429)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(430)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(431)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FigParserPLUSPLUS || _la == FigParserMINUSMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(434)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Expr()
		}
		{
			p.SetState(436)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITryExprContext is an interface to support dynamic dispatch.
type ITryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_TRY() antlr.TerminalNode
	Expr() IExprContext
	TK_ONERROR() antlr.TerminalNode
	Block() IBlockContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsTryExprContext differentiates from other interfaces.
	IsTryExprContext()
}

type TryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTryExprContext() *TryExprContext {
	var p = new(TryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_tryExpr
	return p
}

func InitEmptyTryExprContext(p *TryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_tryExpr
}

func (*TryExprContext) IsTryExprContext() {}

func NewTryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryExprContext {
	var p = new(TryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_tryExpr

	return p
}

func (s *TryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TryExprContext) TK_TRY() antlr.TerminalNode {
	return s.GetToken(FigParserTK_TRY, 0)
}

func (s *TryExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TryExprContext) TK_ONERROR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_ONERROR, 0)
}

func (s *TryExprContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TryExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *TryExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *TryExprContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *TryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterTryExpr(s)
	}
}

func (s *TryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitTryExpr(s)
	}
}

func (s *TryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitTryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) TryExpr() (localctx ITryExprContext) {
	localctx = NewTryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, FigParserRULE_tryExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.Match(FigParserTK_TRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(441)
		p.Expr()
	}
	{
		p.SetState(442)
		p.Match(FigParserTK_ONERROR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserLPAREN {
		{
			p.SetState(443)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(444)
				p.Match(FigParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(447)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(450)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACKET, 0)
}

func (s *ArrayLiteralContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, 0)
}

func (s *ArrayLiteralContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ArrayLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FigParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(FigParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17211767906791424) != 0 {
		{
			p.SetState(453)
			p.Expr()
		}
		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(454)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(455)
				p.Expr()
			}

			p.SetState(460)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(463)
		p.Match(FigParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllObjectEntry() []IObjectEntryContext
	ObjectEntry(i int) IObjectEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_objectLiteral
	return p
}

func InitEmptyObjectLiteralContext(p *ObjectLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_objectLiteral
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACE, 0)
}

func (s *ObjectLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACE, 0)
}

func (s *ObjectLiteralContext) AllObjectEntry() []IObjectEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectEntryContext); ok {
			len++
		}
	}

	tst := make([]IObjectEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectEntryContext); ok {
			tst[i] = t.(IObjectEntryContext)
			i++
		}
	}

	return tst
}

func (s *ObjectLiteralContext) ObjectEntry(i int) IObjectEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectEntryContext)
}

func (s *ObjectLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ObjectLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitObjectLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FigParserRULE_objectLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(465)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID || _la == FigParserSTRING {
		{
			p.SetState(466)
			p.ObjectEntry()
		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(467)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(468)
				p.ObjectEntry()
			}

			p.SetState(473)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(476)
		p.Match(FigParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectEntryContext is an interface to support dynamic dispatch.
type IObjectEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	Expr() IExprContext
	ID() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsObjectEntryContext differentiates from other interfaces.
	IsObjectEntryContext()
}

type ObjectEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectEntryContext() *ObjectEntryContext {
	var p = new(ObjectEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_objectEntry
	return p
}

func InitEmptyObjectEntryContext(p *ObjectEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_objectEntry
}

func (*ObjectEntryContext) IsObjectEntryContext() {}

func NewObjectEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectEntryContext {
	var p = new(ObjectEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_objectEntry

	return p
}

func (s *ObjectEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(FigParserCOLON, 0)
}

func (s *ObjectEntryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ObjectEntryContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ObjectEntryContext) STRING() antlr.TerminalNode {
	return s.GetToken(FigParserSTRING, 0)
}

func (s *ObjectEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterObjectEntry(s)
	}
}

func (s *ObjectEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitObjectEntry(s)
	}
}

func (s *ObjectEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitObjectEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ObjectEntry() (localctx IObjectEntryContext) {
	localctx = NewObjectEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FigParserRULE_objectEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(478)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FigParserID || _la == FigParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(479)
		p.Match(FigParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(480)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
