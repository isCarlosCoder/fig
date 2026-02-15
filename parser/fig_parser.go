// Code generated from /home/carlos/projects/golang/FigLang/grammar/FigParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"'='", "'+'", "'-'", "'++'", "'--'", "'*'", "'/'", "'('", "')'", "'?'",
		"'!'", "';'", "'%'", "'let'", "'print'", "'if'", "'elif'", "'else'",
		"'while'", "'do'", "'break'", "'continue'", "'for'", "'null'", "'fn'",
		"'return'", "'import'", "'use'", "'in'", "'range'", "'enumerate'", "'struct'",
		"'enum'", "'this'", "'try'", "'onerror'", "'match'", "','", "'=>'",
		"'['", "']'", "':'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "LT", "GT", "LE", "GE", "EQ", "NEQ", "AND", "OR", "DOT", "ASSIGN",
		"PLUS", "MINUS", "PLUSPLUS", "MINUSMINUS", "STAR", "SLASH", "LPAREN",
		"RPAREN", "QUESTION", "EXCLAM", "SEMICOLON", "MOD", "TK_LET", "TK_PRINT",
		"TK_IF", "TK_ELIF", "TK_ELSE", "TK_WHILE", "TK_DO", "TK_BREAK", "TK_CONTINUE",
		"TK_FOR", "TK_NULL", "TK_FN", "TK_RETURN", "TK_IMPORT", "TK_USE", "TK_IN",
		"TK_RANGE", "TK_ENUMERATE", "TK_STRUCT", "TK_ENUM", "TK_THIS", "TK_TRY",
		"TK_ONERROR", "TK_MATCH", "COMMA", "ARROW", "LBRACKET", "RBRACKET",
		"COLON", "LBRACE", "RBRACE", "BOOL", "ID", "NUMBER", "STRING", "WS",
		"COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statements", "exprStmt", "ifStmt", "whileStmt", "doWhileStmt",
		"breakStmt", "continueStmt", "fnDecl", "fnParams", "paramDecl", "returnStmt",
		"importStmt", "useStmt", "fnArgs", "forInit", "forStep", "forStmt",
		"forInStmt", "block", "structDecl", "structMember", "enumDecl", "enumMember",
		"varDeclaration", "varAtribuition", "bindingTarget", "arrayPattern",
		"bindingElement", "objectPattern", "memberAssign", "printStmt", "expr",
		"logicalOr", "logicalAnd", "equality", "comparison", "term", "factor",
		"unary", "postfix", "memberName", "primary", "tryExpr", "matchExpr",
		"matchArm", "matchPattern", "arrayLiteral", "objectLiteral", "objectEntry",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 652, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 1, 0, 5, 0, 102, 8, 0, 10, 0, 12, 0,
		105, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 127,
		8, 1, 1, 2, 1, 2, 3, 2, 131, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 144, 8, 3, 10, 3, 12, 3, 147, 9, 3,
		1, 3, 1, 3, 3, 3, 151, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 166, 8, 5, 1, 6, 1, 6, 3, 6,
		170, 8, 6, 1, 7, 1, 7, 3, 7, 174, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 180,
		8, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 188, 8, 9, 10, 9, 12, 9,
		191, 9, 9, 1, 10, 1, 10, 1, 10, 3, 10, 196, 8, 10, 1, 10, 1, 10, 3, 10,
		200, 8, 10, 1, 11, 1, 11, 3, 11, 204, 8, 11, 1, 11, 3, 11, 207, 8, 11,
		1, 12, 1, 12, 1, 12, 3, 12, 212, 8, 12, 1, 12, 3, 12, 215, 8, 12, 1, 13,
		1, 13, 1, 13, 3, 13, 220, 8, 13, 1, 14, 1, 14, 1, 14, 5, 14, 225, 8, 14,
		10, 14, 12, 14, 228, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 234, 8,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 240, 8, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 3, 16, 246, 8, 16, 1, 17, 1, 17, 1, 17, 3, 17, 251, 8, 17, 1, 17,
		1, 17, 3, 17, 255, 8, 17, 1, 17, 1, 17, 3, 17, 259, 8, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 3, 18, 285, 8, 18, 3, 18, 287, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 298, 8, 18, 1, 19, 1, 19, 5,
		19, 302, 8, 19, 10, 19, 12, 19, 305, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 5, 20, 313, 8, 20, 10, 20, 12, 20, 316, 9, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 3, 21, 323, 8, 21, 1, 21, 3, 21, 326, 8, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 3, 21, 332, 8, 21, 1, 21, 1, 21, 3, 21, 336, 8,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 342, 8, 22, 10, 22, 12, 22, 345,
		9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 351, 8, 23, 1, 24, 1, 24, 1,
		24, 1, 24, 3, 24, 357, 8, 24, 1, 24, 3, 24, 360, 8, 24, 1, 25, 1, 25, 1,
		25, 1, 25, 3, 25, 366, 8, 25, 1, 26, 1, 26, 1, 26, 3, 26, 371, 8, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 5, 27, 377, 8, 27, 10, 27, 12, 27, 380, 9, 27,
		3, 27, 382, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 3, 28, 389, 8, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 395, 8, 29, 10, 29, 12, 29, 398, 9,
		29, 3, 29, 400, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 4, 30, 411, 8, 30, 11, 30, 12, 30, 412, 1, 30, 1, 30, 1,
		30, 3, 30, 418, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 425, 8,
		31, 10, 31, 12, 31, 428, 9, 31, 3, 31, 430, 8, 31, 1, 31, 1, 31, 3, 31,
		434, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 5, 33, 441, 8, 33, 10, 33,
		12, 33, 444, 9, 33, 1, 34, 1, 34, 1, 34, 5, 34, 449, 8, 34, 10, 34, 12,
		34, 452, 9, 34, 1, 35, 1, 35, 1, 35, 5, 35, 457, 8, 35, 10, 35, 12, 35,
		460, 9, 35, 1, 36, 1, 36, 1, 36, 5, 36, 465, 8, 36, 10, 36, 12, 36, 468,
		9, 36, 1, 37, 1, 37, 1, 37, 5, 37, 473, 8, 37, 10, 37, 12, 37, 476, 9,
		37, 1, 38, 1, 38, 1, 38, 5, 38, 481, 8, 38, 10, 38, 12, 38, 484, 9, 38,
		1, 39, 1, 39, 1, 39, 3, 39, 489, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 500, 8, 40, 1, 40, 5, 40, 503, 8,
		40, 10, 40, 12, 40, 506, 9, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 522, 8,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 529, 8, 42, 1, 42, 1, 42,
		1, 42, 3, 42, 534, 8, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 540, 8, 42,
		1, 43, 1, 43, 1, 43, 3, 43, 545, 8, 43, 1, 43, 1, 43, 1, 43, 3, 43, 550,
		8, 43, 1, 43, 3, 43, 553, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1,
		44, 4, 44, 561, 8, 44, 11, 44, 12, 44, 562, 1, 44, 1, 44, 1, 45, 1, 45,
		1, 45, 1, 45, 3, 45, 571, 8, 45, 1, 46, 1, 46, 1, 46, 5, 46, 576, 8, 46,
		10, 46, 12, 46, 579, 9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 585, 8,
		47, 10, 47, 12, 47, 588, 9, 47, 3, 47, 590, 8, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 613, 8, 47,
		3, 47, 615, 8, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 633,
		8, 47, 1, 48, 1, 48, 1, 48, 1, 48, 5, 48, 639, 8, 48, 10, 48, 12, 48, 642,
		9, 48, 3, 48, 644, 8, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 0, 0, 50, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 0, 9, 2,
		0, 15, 15, 55, 55, 1, 0, 5, 6, 1, 0, 1, 4, 1, 0, 11, 12, 2, 0, 15, 16,
		22, 22, 2, 0, 12, 14, 20, 20, 2, 0, 46, 46, 55, 55, 1, 0, 13, 14, 2, 0,
		55, 55, 57, 57, 712, 0, 103, 1, 0, 0, 0, 2, 126, 1, 0, 0, 0, 4, 128, 1,
		0, 0, 0, 6, 132, 1, 0, 0, 0, 8, 152, 1, 0, 0, 0, 10, 158, 1, 0, 0, 0, 12,
		167, 1, 0, 0, 0, 14, 171, 1, 0, 0, 0, 16, 175, 1, 0, 0, 0, 18, 184, 1,
		0, 0, 0, 20, 199, 1, 0, 0, 0, 22, 201, 1, 0, 0, 0, 24, 208, 1, 0, 0, 0,
		26, 216, 1, 0, 0, 0, 28, 221, 1, 0, 0, 0, 30, 239, 1, 0, 0, 0, 32, 245,
		1, 0, 0, 0, 34, 247, 1, 0, 0, 0, 36, 297, 1, 0, 0, 0, 38, 299, 1, 0, 0,
		0, 40, 308, 1, 0, 0, 0, 42, 335, 1, 0, 0, 0, 44, 337, 1, 0, 0, 0, 46, 348,
		1, 0, 0, 0, 48, 352, 1, 0, 0, 0, 50, 361, 1, 0, 0, 0, 52, 370, 1, 0, 0,
		0, 54, 372, 1, 0, 0, 0, 56, 388, 1, 0, 0, 0, 58, 390, 1, 0, 0, 0, 60, 403,
		1, 0, 0, 0, 62, 419, 1, 0, 0, 0, 64, 435, 1, 0, 0, 0, 66, 437, 1, 0, 0,
		0, 68, 445, 1, 0, 0, 0, 70, 453, 1, 0, 0, 0, 72, 461, 1, 0, 0, 0, 74, 469,
		1, 0, 0, 0, 76, 477, 1, 0, 0, 0, 78, 488, 1, 0, 0, 0, 80, 490, 1, 0, 0,
		0, 82, 507, 1, 0, 0, 0, 84, 539, 1, 0, 0, 0, 86, 541, 1, 0, 0, 0, 88, 556,
		1, 0, 0, 0, 90, 566, 1, 0, 0, 0, 92, 572, 1, 0, 0, 0, 94, 632, 1, 0, 0,
		0, 96, 634, 1, 0, 0, 0, 98, 647, 1, 0, 0, 0, 100, 102, 3, 2, 1, 0, 101,
		100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104,
		1, 0, 0, 0, 104, 106, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107, 5, 0,
		0, 1, 107, 1, 1, 0, 0, 0, 108, 127, 3, 48, 24, 0, 109, 127, 3, 50, 25,
		0, 110, 127, 3, 60, 30, 0, 111, 127, 3, 62, 31, 0, 112, 127, 3, 6, 3, 0,
		113, 127, 3, 8, 4, 0, 114, 127, 3, 10, 5, 0, 115, 127, 3, 34, 17, 0, 116,
		127, 3, 36, 18, 0, 117, 127, 3, 12, 6, 0, 118, 127, 3, 14, 7, 0, 119, 127,
		3, 16, 8, 0, 120, 127, 3, 22, 11, 0, 121, 127, 3, 24, 12, 0, 122, 127,
		3, 26, 13, 0, 123, 127, 3, 40, 20, 0, 124, 127, 3, 44, 22, 0, 125, 127,
		3, 4, 2, 0, 126, 108, 1, 0, 0, 0, 126, 109, 1, 0, 0, 0, 126, 110, 1, 0,
		0, 0, 126, 111, 1, 0, 0, 0, 126, 112, 1, 0, 0, 0, 126, 113, 1, 0, 0, 0,
		126, 114, 1, 0, 0, 0, 126, 115, 1, 0, 0, 0, 126, 116, 1, 0, 0, 0, 126,
		117, 1, 0, 0, 0, 126, 118, 1, 0, 0, 0, 126, 119, 1, 0, 0, 0, 126, 120,
		1, 0, 0, 0, 126, 121, 1, 0, 0, 0, 126, 122, 1, 0, 0, 0, 126, 123, 1, 0,
		0, 0, 126, 124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 3, 1, 0, 0, 0, 128,
		130, 3, 64, 32, 0, 129, 131, 5, 21, 0, 0, 130, 129, 1, 0, 0, 0, 130, 131,
		1, 0, 0, 0, 131, 5, 1, 0, 0, 0, 132, 133, 5, 25, 0, 0, 133, 134, 5, 17,
		0, 0, 134, 135, 3, 64, 32, 0, 135, 136, 5, 18, 0, 0, 136, 145, 3, 38, 19,
		0, 137, 138, 5, 26, 0, 0, 138, 139, 5, 17, 0, 0, 139, 140, 3, 64, 32, 0,
		140, 141, 5, 18, 0, 0, 141, 142, 3, 38, 19, 0, 142, 144, 1, 0, 0, 0, 143,
		137, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146,
		1, 0, 0, 0, 146, 150, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 27,
		0, 0, 149, 151, 3, 38, 19, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0,
		0, 151, 7, 1, 0, 0, 0, 152, 153, 5, 28, 0, 0, 153, 154, 5, 17, 0, 0, 154,
		155, 3, 64, 32, 0, 155, 156, 5, 18, 0, 0, 156, 157, 3, 38, 19, 0, 157,
		9, 1, 0, 0, 0, 158, 159, 5, 29, 0, 0, 159, 160, 3, 38, 19, 0, 160, 161,
		5, 28, 0, 0, 161, 162, 5, 17, 0, 0, 162, 163, 3, 64, 32, 0, 163, 165, 5,
		18, 0, 0, 164, 166, 5, 21, 0, 0, 165, 164, 1, 0, 0, 0, 165, 166, 1, 0,
		0, 0, 166, 11, 1, 0, 0, 0, 167, 169, 5, 30, 0, 0, 168, 170, 5, 21, 0, 0,
		169, 168, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 13, 1, 0, 0, 0, 171, 173,
		5, 31, 0, 0, 172, 174, 5, 21, 0, 0, 173, 172, 1, 0, 0, 0, 173, 174, 1,
		0, 0, 0, 174, 15, 1, 0, 0, 0, 175, 176, 5, 34, 0, 0, 176, 177, 5, 55, 0,
		0, 177, 179, 5, 17, 0, 0, 178, 180, 3, 18, 9, 0, 179, 178, 1, 0, 0, 0,
		179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 5, 18, 0, 0, 182,
		183, 3, 38, 19, 0, 183, 17, 1, 0, 0, 0, 184, 189, 3, 20, 10, 0, 185, 186,
		5, 47, 0, 0, 186, 188, 3, 20, 10, 0, 187, 185, 1, 0, 0, 0, 188, 191, 1,
		0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 19, 1, 0, 0,
		0, 191, 189, 1, 0, 0, 0, 192, 195, 5, 55, 0, 0, 193, 194, 5, 10, 0, 0,
		194, 196, 3, 64, 32, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196,
		200, 1, 0, 0, 0, 197, 198, 5, 55, 0, 0, 198, 200, 5, 19, 0, 0, 199, 192,
		1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 21, 1, 0, 0, 0, 201, 203, 5, 35,
		0, 0, 202, 204, 3, 64, 32, 0, 203, 202, 1, 0, 0, 0, 203, 204, 1, 0, 0,
		0, 204, 206, 1, 0, 0, 0, 205, 207, 5, 21, 0, 0, 206, 205, 1, 0, 0, 0, 206,
		207, 1, 0, 0, 0, 207, 23, 1, 0, 0, 0, 208, 209, 5, 36, 0, 0, 209, 211,
		5, 57, 0, 0, 210, 212, 7, 0, 0, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0,
		0, 0, 212, 214, 1, 0, 0, 0, 213, 215, 5, 21, 0, 0, 214, 213, 1, 0, 0, 0,
		214, 215, 1, 0, 0, 0, 215, 25, 1, 0, 0, 0, 216, 217, 5, 37, 0, 0, 217,
		219, 5, 57, 0, 0, 218, 220, 5, 21, 0, 0, 219, 218, 1, 0, 0, 0, 219, 220,
		1, 0, 0, 0, 220, 27, 1, 0, 0, 0, 221, 226, 3, 64, 32, 0, 222, 223, 5, 47,
		0, 0, 223, 225, 3, 64, 32, 0, 224, 222, 1, 0, 0, 0, 225, 228, 1, 0, 0,
		0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 29, 1, 0, 0, 0, 228,
		226, 1, 0, 0, 0, 229, 230, 5, 23, 0, 0, 230, 233, 5, 55, 0, 0, 231, 232,
		5, 10, 0, 0, 232, 234, 3, 64, 32, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1,
		0, 0, 0, 234, 240, 1, 0, 0, 0, 235, 236, 5, 55, 0, 0, 236, 237, 5, 10,
		0, 0, 237, 240, 3, 64, 32, 0, 238, 240, 3, 64, 32, 0, 239, 229, 1, 0, 0,
		0, 239, 235, 1, 0, 0, 0, 239, 238, 1, 0, 0, 0, 240, 31, 1, 0, 0, 0, 241,
		242, 5, 55, 0, 0, 242, 243, 5, 10, 0, 0, 243, 246, 3, 64, 32, 0, 244, 246,
		3, 64, 32, 0, 245, 241, 1, 0, 0, 0, 245, 244, 1, 0, 0, 0, 246, 33, 1, 0,
		0, 0, 247, 248, 5, 32, 0, 0, 248, 250, 5, 17, 0, 0, 249, 251, 3, 30, 15,
		0, 250, 249, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252,
		254, 5, 21, 0, 0, 253, 255, 3, 64, 32, 0, 254, 253, 1, 0, 0, 0, 254, 255,
		1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 258, 5, 21, 0, 0, 257, 259, 3, 32,
		16, 0, 258, 257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0,
		260, 261, 5, 18, 0, 0, 261, 262, 3, 38, 19, 0, 262, 35, 1, 0, 0, 0, 263,
		264, 5, 32, 0, 0, 264, 265, 5, 55, 0, 0, 265, 266, 5, 47, 0, 0, 266, 267,
		5, 55, 0, 0, 267, 268, 5, 38, 0, 0, 268, 269, 5, 40, 0, 0, 269, 270, 5,
		17, 0, 0, 270, 271, 3, 64, 32, 0, 271, 272, 5, 18, 0, 0, 272, 273, 3, 38,
		19, 0, 273, 298, 1, 0, 0, 0, 274, 275, 5, 32, 0, 0, 275, 276, 5, 55, 0,
		0, 276, 277, 5, 38, 0, 0, 277, 278, 5, 39, 0, 0, 278, 279, 5, 17, 0, 0,
		279, 286, 3, 64, 32, 0, 280, 281, 5, 47, 0, 0, 281, 284, 3, 64, 32, 0,
		282, 283, 5, 47, 0, 0, 283, 285, 3, 64, 32, 0, 284, 282, 1, 0, 0, 0, 284,
		285, 1, 0, 0, 0, 285, 287, 1, 0, 0, 0, 286, 280, 1, 0, 0, 0, 286, 287,
		1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 5, 18, 0, 0, 289, 290, 3, 38,
		19, 0, 290, 298, 1, 0, 0, 0, 291, 292, 5, 32, 0, 0, 292, 293, 5, 55, 0,
		0, 293, 294, 5, 38, 0, 0, 294, 295, 3, 64, 32, 0, 295, 296, 3, 38, 19,
		0, 296, 298, 1, 0, 0, 0, 297, 263, 1, 0, 0, 0, 297, 274, 1, 0, 0, 0, 297,
		291, 1, 0, 0, 0, 298, 37, 1, 0, 0, 0, 299, 303, 5, 52, 0, 0, 300, 302,
		3, 2, 1, 0, 301, 300, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0,
		0, 0, 303, 304, 1, 0, 0, 0, 304, 306, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0,
		306, 307, 5, 53, 0, 0, 307, 39, 1, 0, 0, 0, 308, 309, 5, 41, 0, 0, 309,
		310, 5, 55, 0, 0, 310, 314, 5, 52, 0, 0, 311, 313, 3, 42, 21, 0, 312, 311,
		1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315, 1, 0,
		0, 0, 315, 317, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 318, 5, 53, 0, 0,
		318, 41, 1, 0, 0, 0, 319, 322, 5, 55, 0, 0, 320, 321, 5, 10, 0, 0, 321,
		323, 3, 64, 32, 0, 322, 320, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 325,
		1, 0, 0, 0, 324, 326, 5, 21, 0, 0, 325, 324, 1, 0, 0, 0, 325, 326, 1, 0,
		0, 0, 326, 336, 1, 0, 0, 0, 327, 328, 5, 34, 0, 0, 328, 329, 5, 55, 0,
		0, 329, 331, 5, 17, 0, 0, 330, 332, 3, 18, 9, 0, 331, 330, 1, 0, 0, 0,
		331, 332, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 334, 5, 18, 0, 0, 334,
		336, 3, 38, 19, 0, 335, 319, 1, 0, 0, 0, 335, 327, 1, 0, 0, 0, 336, 43,
		1, 0, 0, 0, 337, 338, 5, 42, 0, 0, 338, 339, 5, 55, 0, 0, 339, 343, 5,
		52, 0, 0, 340, 342, 3, 46, 23, 0, 341, 340, 1, 0, 0, 0, 342, 345, 1, 0,
		0, 0, 343, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 346, 1, 0, 0, 0,
		345, 343, 1, 0, 0, 0, 346, 347, 5, 53, 0, 0, 347, 45, 1, 0, 0, 0, 348,
		350, 5, 55, 0, 0, 349, 351, 5, 21, 0, 0, 350, 349, 1, 0, 0, 0, 350, 351,
		1, 0, 0, 0, 351, 47, 1, 0, 0, 0, 352, 353, 5, 23, 0, 0, 353, 356, 3, 52,
		26, 0, 354, 355, 5, 10, 0, 0, 355, 357, 3, 64, 32, 0, 356, 354, 1, 0, 0,
		0, 356, 357, 1, 0, 0, 0, 357, 359, 1, 0, 0, 0, 358, 360, 5, 21, 0, 0, 359,
		358, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 49, 1, 0, 0, 0, 361, 362, 3,
		52, 26, 0, 362, 363, 5, 10, 0, 0, 363, 365, 3, 64, 32, 0, 364, 366, 5,
		21, 0, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 51, 1, 0, 0,
		0, 367, 371, 5, 55, 0, 0, 368, 371, 3, 54, 27, 0, 369, 371, 3, 58, 29,
		0, 370, 367, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 369, 1, 0, 0, 0, 371,
		53, 1, 0, 0, 0, 372, 381, 5, 49, 0, 0, 373, 378, 3, 56, 28, 0, 374, 375,
		5, 47, 0, 0, 375, 377, 3, 56, 28, 0, 376, 374, 1, 0, 0, 0, 377, 380, 1,
		0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 382, 1, 0, 0,
		0, 380, 378, 1, 0, 0, 0, 381, 373, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382,
		383, 1, 0, 0, 0, 383, 384, 5, 50, 0, 0, 384, 55, 1, 0, 0, 0, 385, 389,
		5, 55, 0, 0, 386, 389, 3, 54, 27, 0, 387, 389, 3, 58, 29, 0, 388, 385,
		1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 388, 387, 1, 0, 0, 0, 389, 57, 1, 0,
		0, 0, 390, 399, 5, 52, 0, 0, 391, 396, 5, 55, 0, 0, 392, 393, 5, 47, 0,
		0, 393, 395, 5, 55, 0, 0, 394, 392, 1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396,
		394, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 400, 1, 0, 0, 0, 398, 396,
		1, 0, 0, 0, 399, 391, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 401, 1, 0,
		0, 0, 401, 402, 5, 53, 0, 0, 402, 59, 1, 0, 0, 0, 403, 410, 3, 64, 32,
		0, 404, 405, 5, 49, 0, 0, 405, 406, 3, 64, 32, 0, 406, 407, 5, 50, 0, 0,
		407, 411, 1, 0, 0, 0, 408, 409, 5, 9, 0, 0, 409, 411, 3, 82, 41, 0, 410,
		404, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 410,
		1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 415, 5, 10,
		0, 0, 415, 417, 3, 64, 32, 0, 416, 418, 5, 21, 0, 0, 417, 416, 1, 0, 0,
		0, 417, 418, 1, 0, 0, 0, 418, 61, 1, 0, 0, 0, 419, 420, 5, 24, 0, 0, 420,
		429, 5, 17, 0, 0, 421, 426, 3, 64, 32, 0, 422, 423, 5, 47, 0, 0, 423, 425,
		3, 64, 32, 0, 424, 422, 1, 0, 0, 0, 425, 428, 1, 0, 0, 0, 426, 424, 1,
		0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428, 426, 1, 0, 0,
		0, 429, 421, 1, 0, 0, 0, 429, 430, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431,
		433, 5, 18, 0, 0, 432, 434, 5, 21, 0, 0, 433, 432, 1, 0, 0, 0, 433, 434,
		1, 0, 0, 0, 434, 63, 1, 0, 0, 0, 435, 436, 3, 66, 33, 0, 436, 65, 1, 0,
		0, 0, 437, 442, 3, 68, 34, 0, 438, 439, 5, 8, 0, 0, 439, 441, 3, 68, 34,
		0, 440, 438, 1, 0, 0, 0, 441, 444, 1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 442,
		443, 1, 0, 0, 0, 443, 67, 1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 445, 450, 3,
		70, 35, 0, 446, 447, 5, 7, 0, 0, 447, 449, 3, 70, 35, 0, 448, 446, 1, 0,
		0, 0, 449, 452, 1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0,
		451, 69, 1, 0, 0, 0, 452, 450, 1, 0, 0, 0, 453, 458, 3, 72, 36, 0, 454,
		455, 7, 1, 0, 0, 455, 457, 3, 72, 36, 0, 456, 454, 1, 0, 0, 0, 457, 460,
		1, 0, 0, 0, 458, 456, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459, 71, 1, 0,
		0, 0, 460, 458, 1, 0, 0, 0, 461, 466, 3, 74, 37, 0, 462, 463, 7, 2, 0,
		0, 463, 465, 3, 74, 37, 0, 464, 462, 1, 0, 0, 0, 465, 468, 1, 0, 0, 0,
		466, 464, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 73, 1, 0, 0, 0, 468, 466,
		1, 0, 0, 0, 469, 474, 3, 76, 38, 0, 470, 471, 7, 3, 0, 0, 471, 473, 3,
		76, 38, 0, 472, 470, 1, 0, 0, 0, 473, 476, 1, 0, 0, 0, 474, 472, 1, 0,
		0, 0, 474, 475, 1, 0, 0, 0, 475, 75, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0,
		477, 482, 3, 78, 39, 0, 478, 479, 7, 4, 0, 0, 479, 481, 3, 78, 39, 0, 480,
		478, 1, 0, 0, 0, 481, 484, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 482, 483,
		1, 0, 0, 0, 483, 77, 1, 0, 0, 0, 484, 482, 1, 0, 0, 0, 485, 486, 7, 5,
		0, 0, 486, 489, 3, 78, 39, 0, 487, 489, 3, 80, 40, 0, 488, 485, 1, 0, 0,
		0, 488, 487, 1, 0, 0, 0, 489, 79, 1, 0, 0, 0, 490, 504, 3, 84, 42, 0, 491,
		492, 5, 49, 0, 0, 492, 493, 3, 64, 32, 0, 493, 494, 5, 50, 0, 0, 494, 503,
		1, 0, 0, 0, 495, 496, 5, 9, 0, 0, 496, 503, 3, 82, 41, 0, 497, 499, 5,
		17, 0, 0, 498, 500, 3, 28, 14, 0, 499, 498, 1, 0, 0, 0, 499, 500, 1, 0,
		0, 0, 500, 501, 1, 0, 0, 0, 501, 503, 5, 18, 0, 0, 502, 491, 1, 0, 0, 0,
		502, 495, 1, 0, 0, 0, 502, 497, 1, 0, 0, 0, 503, 506, 1, 0, 0, 0, 504,
		502, 1, 0, 0, 0, 504, 505, 1, 0, 0, 0, 505, 81, 1, 0, 0, 0, 506, 504, 1,
		0, 0, 0, 507, 508, 7, 6, 0, 0, 508, 83, 1, 0, 0, 0, 509, 540, 5, 56, 0,
		0, 510, 540, 5, 54, 0, 0, 511, 540, 5, 57, 0, 0, 512, 540, 5, 33, 0, 0,
		513, 540, 5, 43, 0, 0, 514, 540, 3, 94, 47, 0, 515, 540, 3, 96, 48, 0,
		516, 540, 3, 86, 43, 0, 517, 540, 3, 88, 44, 0, 518, 519, 5, 34, 0, 0,
		519, 521, 5, 17, 0, 0, 520, 522, 3, 18, 9, 0, 521, 520, 1, 0, 0, 0, 521,
		522, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523, 524, 5, 18, 0, 0, 524, 540,
		3, 38, 19, 0, 525, 526, 5, 55, 0, 0, 526, 528, 5, 17, 0, 0, 527, 529, 3,
		28, 14, 0, 528, 527, 1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 530, 1, 0,
		0, 0, 530, 540, 5, 18, 0, 0, 531, 533, 5, 55, 0, 0, 532, 534, 7, 7, 0,
		0, 533, 532, 1, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534, 540, 1, 0, 0, 0, 535,
		536, 5, 17, 0, 0, 536, 537, 3, 64, 32, 0, 537, 538, 5, 18, 0, 0, 538, 540,
		1, 0, 0, 0, 539, 509, 1, 0, 0, 0, 539, 510, 1, 0, 0, 0, 539, 511, 1, 0,
		0, 0, 539, 512, 1, 0, 0, 0, 539, 513, 1, 0, 0, 0, 539, 514, 1, 0, 0, 0,
		539, 515, 1, 0, 0, 0, 539, 516, 1, 0, 0, 0, 539, 517, 1, 0, 0, 0, 539,
		518, 1, 0, 0, 0, 539, 525, 1, 0, 0, 0, 539, 531, 1, 0, 0, 0, 539, 535,
		1, 0, 0, 0, 540, 85, 1, 0, 0, 0, 541, 544, 5, 44, 0, 0, 542, 545, 3, 64,
		32, 0, 543, 545, 3, 38, 19, 0, 544, 542, 1, 0, 0, 0, 544, 543, 1, 0, 0,
		0, 545, 546, 1, 0, 0, 0, 546, 552, 5, 45, 0, 0, 547, 549, 5, 17, 0, 0,
		548, 550, 5, 55, 0, 0, 549, 548, 1, 0, 0, 0, 549, 550, 1, 0, 0, 0, 550,
		551, 1, 0, 0, 0, 551, 553, 5, 18, 0, 0, 552, 547, 1, 0, 0, 0, 552, 553,
		1, 0, 0, 0, 553, 554, 1, 0, 0, 0, 554, 555, 3, 38, 19, 0, 555, 87, 1, 0,
		0, 0, 556, 557, 5, 46, 0, 0, 557, 558, 3, 64, 32, 0, 558, 560, 5, 52, 0,
		0, 559, 561, 3, 90, 45, 0, 560, 559, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0,
		562, 560, 1, 0, 0, 0, 562, 563, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564,
		565, 5, 53, 0, 0, 565, 89, 1, 0, 0, 0, 566, 567, 3, 92, 46, 0, 567, 570,
		5, 48, 0, 0, 568, 571, 3, 38, 19, 0, 569, 571, 3, 64, 32, 0, 570, 568,
		1, 0, 0, 0, 570, 569, 1, 0, 0, 0, 571, 91, 1, 0, 0, 0, 572, 577, 3, 64,
		32, 0, 573, 574, 5, 47, 0, 0, 574, 576, 3, 64, 32, 0, 575, 573, 1, 0, 0,
		0, 576, 579, 1, 0, 0, 0, 577, 575, 1, 0, 0, 0, 577, 578, 1, 0, 0, 0, 578,
		93, 1, 0, 0, 0, 579, 577, 1, 0, 0, 0, 580, 589, 5, 49, 0, 0, 581, 586,
		3, 64, 32, 0, 582, 583, 5, 47, 0, 0, 583, 585, 3, 64, 32, 0, 584, 582,
		1, 0, 0, 0, 585, 588, 1, 0, 0, 0, 586, 584, 1, 0, 0, 0, 586, 587, 1, 0,
		0, 0, 587, 590, 1, 0, 0, 0, 588, 586, 1, 0, 0, 0, 589, 581, 1, 0, 0, 0,
		589, 590, 1, 0, 0, 0, 590, 591, 1, 0, 0, 0, 591, 633, 5, 50, 0, 0, 592,
		593, 5, 49, 0, 0, 593, 594, 3, 64, 32, 0, 594, 595, 5, 32, 0, 0, 595, 596,
		5, 55, 0, 0, 596, 597, 5, 38, 0, 0, 597, 598, 3, 64, 32, 0, 598, 599, 5,
		50, 0, 0, 599, 633, 1, 0, 0, 0, 600, 601, 5, 49, 0, 0, 601, 602, 3, 64,
		32, 0, 602, 603, 5, 32, 0, 0, 603, 604, 5, 55, 0, 0, 604, 605, 5, 38, 0,
		0, 605, 606, 5, 39, 0, 0, 606, 607, 5, 17, 0, 0, 607, 614, 3, 64, 32, 0,
		608, 609, 5, 47, 0, 0, 609, 612, 3, 64, 32, 0, 610, 611, 5, 47, 0, 0, 611,
		613, 3, 64, 32, 0, 612, 610, 1, 0, 0, 0, 612, 613, 1, 0, 0, 0, 613, 615,
		1, 0, 0, 0, 614, 608, 1, 0, 0, 0, 614, 615, 1, 0, 0, 0, 615, 616, 1, 0,
		0, 0, 616, 617, 5, 18, 0, 0, 617, 618, 5, 50, 0, 0, 618, 633, 1, 0, 0,
		0, 619, 620, 5, 49, 0, 0, 620, 621, 3, 64, 32, 0, 621, 622, 5, 32, 0, 0,
		622, 623, 5, 55, 0, 0, 623, 624, 5, 47, 0, 0, 624, 625, 5, 55, 0, 0, 625,
		626, 5, 38, 0, 0, 626, 627, 5, 40, 0, 0, 627, 628, 5, 17, 0, 0, 628, 629,
		3, 64, 32, 0, 629, 630, 5, 18, 0, 0, 630, 631, 5, 50, 0, 0, 631, 633, 1,
		0, 0, 0, 632, 580, 1, 0, 0, 0, 632, 592, 1, 0, 0, 0, 632, 600, 1, 0, 0,
		0, 632, 619, 1, 0, 0, 0, 633, 95, 1, 0, 0, 0, 634, 643, 5, 52, 0, 0, 635,
		640, 3, 98, 49, 0, 636, 637, 5, 47, 0, 0, 637, 639, 3, 98, 49, 0, 638,
		636, 1, 0, 0, 0, 639, 642, 1, 0, 0, 0, 640, 638, 1, 0, 0, 0, 640, 641,
		1, 0, 0, 0, 641, 644, 1, 0, 0, 0, 642, 640, 1, 0, 0, 0, 643, 635, 1, 0,
		0, 0, 643, 644, 1, 0, 0, 0, 644, 645, 1, 0, 0, 0, 645, 646, 5, 53, 0, 0,
		646, 97, 1, 0, 0, 0, 647, 648, 7, 8, 0, 0, 648, 649, 5, 51, 0, 0, 649,
		650, 3, 64, 32, 0, 650, 99, 1, 0, 0, 0, 77, 103, 126, 130, 145, 150, 165,
		169, 173, 179, 189, 195, 199, 203, 206, 211, 214, 219, 226, 233, 239, 245,
		250, 254, 258, 284, 286, 297, 303, 314, 322, 325, 331, 335, 343, 350, 356,
		359, 365, 370, 378, 381, 388, 396, 399, 410, 412, 417, 426, 429, 433, 442,
		450, 458, 466, 474, 482, 488, 499, 502, 504, 521, 528, 533, 539, 544, 549,
		552, 562, 570, 577, 586, 589, 612, 614, 632, 640, 643,
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
	FigParserQUESTION     = 19
	FigParserEXCLAM       = 20
	FigParserSEMICOLON    = 21
	FigParserMOD          = 22
	FigParserTK_LET       = 23
	FigParserTK_PRINT     = 24
	FigParserTK_IF        = 25
	FigParserTK_ELIF      = 26
	FigParserTK_ELSE      = 27
	FigParserTK_WHILE     = 28
	FigParserTK_DO        = 29
	FigParserTK_BREAK     = 30
	FigParserTK_CONTINUE  = 31
	FigParserTK_FOR       = 32
	FigParserTK_NULL      = 33
	FigParserTK_FN        = 34
	FigParserTK_RETURN    = 35
	FigParserTK_IMPORT    = 36
	FigParserTK_USE       = 37
	FigParserTK_IN        = 38
	FigParserTK_RANGE     = 39
	FigParserTK_ENUMERATE = 40
	FigParserTK_STRUCT    = 41
	FigParserTK_ENUM      = 42
	FigParserTK_THIS      = 43
	FigParserTK_TRY       = 44
	FigParserTK_ONERROR   = 45
	FigParserTK_MATCH     = 46
	FigParserCOMMA        = 47
	FigParserARROW        = 48
	FigParserLBRACKET     = 49
	FigParserRBRACKET     = 50
	FigParserCOLON        = 51
	FigParserLBRACE       = 52
	FigParserRBRACE       = 53
	FigParserBOOL         = 54
	FigParserID           = 55
	FigParserNUMBER       = 56
	FigParserSTRING       = 57
	FigParserWS           = 58
	FigParserCOMMENT      = 59
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
	FigParserRULE_paramDecl      = 10
	FigParserRULE_returnStmt     = 11
	FigParserRULE_importStmt     = 12
	FigParserRULE_useStmt        = 13
	FigParserRULE_fnArgs         = 14
	FigParserRULE_forInit        = 15
	FigParserRULE_forStep        = 16
	FigParserRULE_forStmt        = 17
	FigParserRULE_forInStmt      = 18
	FigParserRULE_block          = 19
	FigParserRULE_structDecl     = 20
	FigParserRULE_structMember   = 21
	FigParserRULE_enumDecl       = 22
	FigParserRULE_enumMember     = 23
	FigParserRULE_varDeclaration = 24
	FigParserRULE_varAtribuition = 25
	FigParserRULE_bindingTarget  = 26
	FigParserRULE_arrayPattern   = 27
	FigParserRULE_bindingElement = 28
	FigParserRULE_objectPattern  = 29
	FigParserRULE_memberAssign   = 30
	FigParserRULE_printStmt      = 31
	FigParserRULE_expr           = 32
	FigParserRULE_logicalOr      = 33
	FigParserRULE_logicalAnd     = 34
	FigParserRULE_equality       = 35
	FigParserRULE_comparison     = 36
	FigParserRULE_term           = 37
	FigParserRULE_factor         = 38
	FigParserRULE_unary          = 39
	FigParserRULE_postfix        = 40
	FigParserRULE_memberName     = 41
	FigParserRULE_primary        = 42
	FigParserRULE_tryExpr        = 43
	FigParserRULE_matchExpr      = 44
	FigParserRULE_matchArm       = 45
	FigParserRULE_matchPattern   = 46
	FigParserRULE_arrayLiteral   = 47
	FigParserRULE_objectLiteral  = 48
	FigParserRULE_objectEntry    = 49
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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275386155985432576) != 0 {
		{
			p.SetState(100)
			p.Statements()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(106)
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
	EnumDecl() IEnumDeclContext
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

func (s *StatementsContext) EnumDecl() IEnumDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDeclContext)
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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.VarDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.VarAtribuition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.MemberAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.PrintStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.IfStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)
			p.WhileStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(114)
			p.DoWhileStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(115)
			p.ForStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(116)
			p.ForInStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(117)
			p.BreakStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(118)
			p.ContinueStmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(119)
			p.FnDecl()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(120)
			p.ReturnStmt()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(121)
			p.ImportStmt()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(122)
			p.UseStmt()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(123)
			p.StructDecl()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(124)
			p.EnumDecl()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(125)
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
		p.SetState(128)
		p.Expr()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(129)
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
		p.SetState(132)
		p.Match(FigParserTK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Expr()
	}
	{
		p.SetState(135)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Block()
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_ELIF {
		{
			p.SetState(137)
			p.Match(FigParserTK_ELIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Expr()
		}
		{
			p.SetState(140)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Block()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserTK_ELSE {
		{
			p.SetState(148)
			p.Match(FigParserTK_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
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
		p.SetState(152)
		p.Match(FigParserTK_WHILE)
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
	{
		p.SetState(154)
		p.Expr()
	}
	{
		p.SetState(155)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
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
		p.SetState(158)
		p.Match(FigParserTK_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Block()
	}
	{
		p.SetState(160)
		p.Match(FigParserTK_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Expr()
	}
	{
		p.SetState(163)
		p.Match(FigParserRPAREN)
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

	if _la == FigParserSEMICOLON {
		{
			p.SetState(164)
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
		p.SetState(167)
		p.Match(FigParserTK_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(168)
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
		p.SetState(171)
		p.Match(FigParserTK_CONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.SetState(175)
		p.Match(FigParserTK_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID {
		{
			p.SetState(178)
			p.FnParams()
		}

	}
	{
		p.SetState(181)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
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
	AllParamDecl() []IParamDeclContext
	ParamDecl(i int) IParamDeclContext
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

func (s *FnParamsContext) AllParamDecl() []IParamDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamDeclContext); ok {
			len++
		}
	}

	tst := make([]IParamDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamDeclContext); ok {
			tst[i] = t.(IParamDeclContext)
			i++
		}
	}

	return tst
}

func (s *FnParamsContext) ParamDecl(i int) IParamDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamDeclContext); ok {
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

	return t.(IParamDeclContext)
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
		p.SetState(184)
		p.ParamDecl()
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(185)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.ParamDecl()
		}

		p.SetState(191)
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

// IParamDeclContext is an interface to support dynamic dispatch.
type IParamDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParamDeclContext differentiates from other interfaces.
	IsParamDeclContext()
}

type ParamDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamDeclContext() *ParamDeclContext {
	var p = new(ParamDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_paramDecl
	return p
}

func InitEmptyParamDeclContext(p *ParamDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_paramDecl
}

func (*ParamDeclContext) IsParamDeclContext() {}

func NewParamDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamDeclContext {
	var p = new(ParamDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_paramDecl

	return p
}

func (s *ParamDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamDeclContext) CopyAll(ctx *ParamDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParamDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParamWithDefaultOrRequiredContext struct {
	ParamDeclContext
}

func NewParamWithDefaultOrRequiredContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamWithDefaultOrRequiredContext {
	var p = new(ParamWithDefaultOrRequiredContext)

	InitEmptyParamDeclContext(&p.ParamDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamDeclContext))

	return p
}

func (s *ParamWithDefaultOrRequiredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamWithDefaultOrRequiredContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ParamWithDefaultOrRequiredContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FigParserASSIGN, 0)
}

func (s *ParamWithDefaultOrRequiredContext) Expr() IExprContext {
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

func (s *ParamWithDefaultOrRequiredContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterParamWithDefaultOrRequired(s)
	}
}

func (s *ParamWithDefaultOrRequiredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitParamWithDefaultOrRequired(s)
	}
}

func (s *ParamWithDefaultOrRequiredContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitParamWithDefaultOrRequired(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParamOptionalContext struct {
	ParamDeclContext
}

func NewParamOptionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParamOptionalContext {
	var p = new(ParamOptionalContext)

	InitEmptyParamDeclContext(&p.ParamDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParamDeclContext))

	return p
}

func (s *ParamOptionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamOptionalContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ParamOptionalContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(FigParserQUESTION, 0)
}

func (s *ParamOptionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterParamOptional(s)
	}
}

func (s *ParamOptionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitParamOptional(s)
	}
}

func (s *ParamOptionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitParamOptional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ParamDecl() (localctx IParamDeclContext) {
	localctx = NewParamDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FigParserRULE_paramDecl)
	var _la int

	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamWithDefaultOrRequiredContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(193)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(194)
				p.Expr()
			}

		}

	case 2:
		localctx = NewParamOptionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(197)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(FigParserQUESTION)
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
	p.EnterRule(localctx, 22, FigParserRULE_returnStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(FigParserTK_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(202)
			p.Expr()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(205)
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
	ID() antlr.TerminalNode
	STAR() antlr.TerminalNode

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

func (s *ImportStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ImportStmtContext) STAR() antlr.TerminalNode {
	return s.GetToken(FigParserSTAR, 0)
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
	p.EnterRule(localctx, 24, FigParserRULE_importStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(FigParserTK_IMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(FigParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(210)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FigParserSTAR || _la == FigParserID) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(213)
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
	p.EnterRule(localctx, 26, FigParserRULE_useStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(FigParserTK_USE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Match(FigParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(218)
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
	p.EnterRule(localctx, 28, FigParserRULE_fnArgs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Expr()
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(222)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Expr()
		}

		p.SetState(228)
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
	p.EnterRule(localctx, 30, FigParserRULE_forInit)
	var _la int

	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Match(FigParserTK_LET)
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
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(231)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(232)
				p.Expr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(238)
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
	p.EnterRule(localctx, 32, FigParserRULE_forStep)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
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
	p.EnterRule(localctx, 34, FigParserRULE_forStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(FigParserTK_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310025666560) != 0 {
		{
			p.SetState(249)
			p.ForInit()
		}

	}
	{
		p.SetState(252)
		p.Match(FigParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
		{
			p.SetState(253)
			p.Expr()
		}

	}
	{
		p.SetState(256)
		p.Match(FigParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
		{
			p.SetState(257)
			p.ForStep()
		}

	}
	{
		p.SetState(260)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
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

func (s *ForRangeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ForRangeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
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
	p.EnterRule(localctx, 36, FigParserRULE_forInStmt)
	var _la int

	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Match(FigParserTK_ENUMERATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Expr()
		}
		{
			p.SetState(271)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.Block()
		}

	case 2:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(277)
			p.Match(FigParserTK_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Expr()
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserCOMMA {
			{
				p.SetState(280)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(281)
				p.Expr()
			}
			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == FigParserCOMMA {
				{
					p.SetState(282)
					p.Match(FigParserCOMMA)
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

		}
		{
			p.SetState(288)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.Block()
		}

	case 3:
		localctx = NewForInContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(291)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)
			p.Expr()
		}
		{
			p.SetState(295)
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
	p.EnterRule(localctx, 38, FigParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(FigParserLBRACE)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275386155985432576) != 0 {
		{
			p.SetState(300)
			p.Statements()
		}

		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(306)
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
	p.EnterRule(localctx, 40, FigParserRULE_structDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(FigParserTK_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(310)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_FN || _la == FigParserID {
		{
			p.SetState(311)
			p.StructMember()
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(317)
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
	p.EnterRule(localctx, 42, FigParserRULE_structMember)
	var _la int

	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		localctx = NewStructFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(320)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(321)
				p.Expr()
			}

		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserSEMICOLON {
			{
				p.SetState(324)
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
			p.SetState(327)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(330)
				p.FnParams()
			}

		}
		{
			p.SetState(333)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
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

// IEnumDeclContext is an interface to support dynamic dispatch.
type IEnumDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_ENUM() antlr.TerminalNode
	ID() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllEnumMember() []IEnumMemberContext
	EnumMember(i int) IEnumMemberContext

	// IsEnumDeclContext differentiates from other interfaces.
	IsEnumDeclContext()
}

type EnumDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDeclContext() *EnumDeclContext {
	var p = new(EnumDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_enumDecl
	return p
}

func InitEmptyEnumDeclContext(p *EnumDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_enumDecl
}

func (*EnumDeclContext) IsEnumDeclContext() {}

func NewEnumDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclContext {
	var p = new(EnumDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_enumDecl

	return p
}

func (s *EnumDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclContext) TK_ENUM() antlr.TerminalNode {
	return s.GetToken(FigParserTK_ENUM, 0)
}

func (s *EnumDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *EnumDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACE, 0)
}

func (s *EnumDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACE, 0)
}

func (s *EnumDeclContext) AllEnumMember() []IEnumMemberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumMemberContext); ok {
			len++
		}
	}

	tst := make([]IEnumMemberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumMemberContext); ok {
			tst[i] = t.(IEnumMemberContext)
			i++
		}
	}

	return tst
}

func (s *EnumDeclContext) EnumMember(i int) IEnumMemberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumMemberContext); ok {
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

	return t.(IEnumMemberContext)
}

func (s *EnumDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterEnumDecl(s)
	}
}

func (s *EnumDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitEnumDecl(s)
	}
}

func (s *EnumDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitEnumDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) EnumDecl() (localctx IEnumDeclContext) {
	localctx = NewEnumDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FigParserRULE_enumDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(FigParserTK_ENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(339)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserID {
		{
			p.SetState(340)
			p.EnumMember()
		}

		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(346)
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

// IEnumMemberContext is an interface to support dynamic dispatch.
type IEnumMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsEnumMemberContext differentiates from other interfaces.
	IsEnumMemberContext()
}

type EnumMemberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumMemberContext() *EnumMemberContext {
	var p = new(EnumMemberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_enumMember
	return p
}

func InitEmptyEnumMemberContext(p *EnumMemberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_enumMember
}

func (*EnumMemberContext) IsEnumMemberContext() {}

func NewEnumMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumMemberContext {
	var p = new(EnumMemberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_enumMember

	return p
}

func (s *EnumMemberContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumMemberContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *EnumMemberContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *EnumMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumMemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterEnumMember(s)
	}
}

func (s *EnumMemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitEnumMember(s)
	}
}

func (s *EnumMemberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitEnumMember(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) EnumMember() (localctx IEnumMemberContext) {
	localctx = NewEnumMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FigParserRULE_enumMember)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(349)
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

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_LET() antlr.TerminalNode
	BindingTarget() IBindingTargetContext
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

func (s *VarDeclarationContext) BindingTarget() IBindingTargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindingTargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindingTargetContext)
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
	p.EnterRule(localctx, 48, FigParserRULE_varDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(FigParserTK_LET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(353)
		p.BindingTarget()
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserASSIGN {
		{
			p.SetState(354)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Expr()
		}

	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(358)
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
	BindingTarget() IBindingTargetContext
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

func (s *VarAtribuitionContext) BindingTarget() IBindingTargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindingTargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindingTargetContext)
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
	p.EnterRule(localctx, 50, FigParserRULE_varAtribuition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.BindingTarget()
	}
	{
		p.SetState(362)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.Expr()
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(364)
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

// IBindingTargetContext is an interface to support dynamic dispatch.
type IBindingTargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ArrayPattern() IArrayPatternContext
	ObjectPattern() IObjectPatternContext

	// IsBindingTargetContext differentiates from other interfaces.
	IsBindingTargetContext()
}

type BindingTargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindingTargetContext() *BindingTargetContext {
	var p = new(BindingTargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_bindingTarget
	return p
}

func InitEmptyBindingTargetContext(p *BindingTargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_bindingTarget
}

func (*BindingTargetContext) IsBindingTargetContext() {}

func NewBindingTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindingTargetContext {
	var p = new(BindingTargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_bindingTarget

	return p
}

func (s *BindingTargetContext) GetParser() antlr.Parser { return s.parser }

func (s *BindingTargetContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *BindingTargetContext) ArrayPattern() IArrayPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayPatternContext)
}

func (s *BindingTargetContext) ObjectPattern() IObjectPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectPatternContext)
}

func (s *BindingTargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindingTargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindingTargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterBindingTarget(s)
	}
}

func (s *BindingTargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitBindingTarget(s)
	}
}

func (s *BindingTargetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitBindingTarget(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) BindingTarget() (localctx IBindingTargetContext) {
	localctx = NewBindingTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FigParserRULE_bindingTarget)
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FigParserLBRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
			p.ArrayPattern()
		}

	case FigParserLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(369)
			p.ObjectPattern()
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

// IArrayPatternContext is an interface to support dynamic dispatch.
type IArrayPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllBindingElement() []IBindingElementContext
	BindingElement(i int) IBindingElementContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayPatternContext differentiates from other interfaces.
	IsArrayPatternContext()
}

type ArrayPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayPatternContext() *ArrayPatternContext {
	var p = new(ArrayPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_arrayPattern
	return p
}

func InitEmptyArrayPatternContext(p *ArrayPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_arrayPattern
}

func (*ArrayPatternContext) IsArrayPatternContext() {}

func NewArrayPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayPatternContext {
	var p = new(ArrayPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_arrayPattern

	return p
}

func (s *ArrayPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayPatternContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACKET, 0)
}

func (s *ArrayPatternContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, 0)
}

func (s *ArrayPatternContext) AllBindingElement() []IBindingElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindingElementContext); ok {
			len++
		}
	}

	tst := make([]IBindingElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindingElementContext); ok {
			tst[i] = t.(IBindingElementContext)
			i++
		}
	}

	return tst
}

func (s *ArrayPatternContext) BindingElement(i int) IBindingElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindingElementContext); ok {
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

	return t.(IBindingElementContext)
}

func (s *ArrayPatternContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ArrayPatternContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *ArrayPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterArrayPattern(s)
	}
}

func (s *ArrayPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitArrayPattern(s)
	}
}

func (s *ArrayPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitArrayPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ArrayPattern() (localctx IArrayPatternContext) {
	localctx = NewArrayPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FigParserRULE_arrayPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(FigParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&41095346599755776) != 0 {
		{
			p.SetState(373)
			p.BindingElement()
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(374)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(375)
				p.BindingElement()
			}

			p.SetState(380)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(383)
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

// IBindingElementContext is an interface to support dynamic dispatch.
type IBindingElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ArrayPattern() IArrayPatternContext
	ObjectPattern() IObjectPatternContext

	// IsBindingElementContext differentiates from other interfaces.
	IsBindingElementContext()
}

type BindingElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindingElementContext() *BindingElementContext {
	var p = new(BindingElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_bindingElement
	return p
}

func InitEmptyBindingElementContext(p *BindingElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_bindingElement
}

func (*BindingElementContext) IsBindingElementContext() {}

func NewBindingElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindingElementContext {
	var p = new(BindingElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_bindingElement

	return p
}

func (s *BindingElementContext) GetParser() antlr.Parser { return s.parser }

func (s *BindingElementContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *BindingElementContext) ArrayPattern() IArrayPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayPatternContext)
}

func (s *BindingElementContext) ObjectPattern() IObjectPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectPatternContext)
}

func (s *BindingElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindingElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindingElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterBindingElement(s)
	}
}

func (s *BindingElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitBindingElement(s)
	}
}

func (s *BindingElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitBindingElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) BindingElement() (localctx IBindingElementContext) {
	localctx = NewBindingElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FigParserRULE_bindingElement)
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FigParserLBRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(386)
			p.ArrayPattern()
		}

	case FigParserLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(387)
			p.ObjectPattern()
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

// IObjectPatternContext is an interface to support dynamic dispatch.
type IObjectPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsObjectPatternContext differentiates from other interfaces.
	IsObjectPatternContext()
}

type ObjectPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectPatternContext() *ObjectPatternContext {
	var p = new(ObjectPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_objectPattern
	return p
}

func InitEmptyObjectPatternContext(p *ObjectPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_objectPattern
}

func (*ObjectPatternContext) IsObjectPatternContext() {}

func NewObjectPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectPatternContext {
	var p = new(ObjectPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_objectPattern

	return p
}

func (s *ObjectPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectPatternContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACE, 0)
}

func (s *ObjectPatternContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACE, 0)
}

func (s *ObjectPatternContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(FigParserID)
}

func (s *ObjectPatternContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(FigParserID, i)
}

func (s *ObjectPatternContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ObjectPatternContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *ObjectPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterObjectPattern(s)
	}
}

func (s *ObjectPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitObjectPattern(s)
	}
}

func (s *ObjectPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitObjectPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ObjectPattern() (localctx IObjectPatternContext) {
	localctx = NewObjectPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FigParserRULE_objectPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID {
		{
			p.SetState(391)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(392)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(393)
				p.Match(FigParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(398)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(401)
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
	AllMemberName() []IMemberNameContext
	MemberName(i int) IMemberNameContext
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

func (s *MemberAssignContext) AllMemberName() []IMemberNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMemberNameContext); ok {
			len++
		}
	}

	tst := make([]IMemberNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMemberNameContext); ok {
			tst[i] = t.(IMemberNameContext)
			i++
		}
	}

	return tst
}

func (s *MemberAssignContext) MemberName(i int) IMemberNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberNameContext); ok {
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

	return t.(IMemberNameContext)
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
	p.EnterRule(localctx, 60, FigParserRULE_memberAssign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Expr()
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FigParserDOT || _la == FigParserLBRACKET {
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case FigParserLBRACKET:
			{
				p.SetState(404)
				p.Match(FigParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(405)
				p.Expr()
			}
			{
				p.SetState(406)
				p.Match(FigParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case FigParserDOT:
			{
				p.SetState(408)
				p.Match(FigParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(409)
				p.MemberName()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(414)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.Expr()
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(416)
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
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	SEMICOLON() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *PrintStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *PrintStmtContext) AllExpr() []IExprContext {
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

func (s *PrintStmtContext) Expr(i int) IExprContext {
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

func (s *PrintStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FigParserSEMICOLON, 0)
}

func (s *PrintStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *PrintStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
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
	p.EnterRule(localctx, 62, FigParserRULE_printStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(FigParserTK_PRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(420)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
		{
			p.SetState(421)
			p.Expr()
		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(422)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(423)
				p.Expr()
			}

			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(431)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(432)
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
	p.EnterRule(localctx, 64, FigParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
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
	p.EnterRule(localctx, 66, FigParserRULE_logicalOr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.LogicalAnd()
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserOR {
		{
			p.SetState(438)
			p.Match(FigParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(439)
			p.LogicalAnd()
		}

		p.SetState(444)
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
	p.EnterRule(localctx, 68, FigParserRULE_logicalAnd)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		p.Equality()
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserAND {
		{
			p.SetState(446)
			p.Match(FigParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Equality()
		}

		p.SetState(452)
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
	p.EnterRule(localctx, 70, FigParserRULE_equality)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.Comparison()
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserEQ || _la == FigParserNEQ {
		{
			p.SetState(454)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FigParserEQ || _la == FigParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(455)
			p.Comparison()
		}

		p.SetState(460)
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
	p.EnterRule(localctx, 72, FigParserRULE_comparison)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
		p.Term()
	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0 {
		{
			p.SetState(462)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(463)
			p.Term()
		}

		p.SetState(468)
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
	p.EnterRule(localctx, 74, FigParserRULE_term)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(469)
		p.Factor()
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(470)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FigParserPLUS || _la == FigParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(471)
				p.Factor()
			}

		}
		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 76, FigParserRULE_factor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.Unary()
	}
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4292608) != 0 {
		{
			p.SetState(478)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4292608) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(479)
			p.Unary()
		}

		p.SetState(484)
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
	p.EnterRule(localctx, 78, FigParserRULE_unary)
	var _la int

	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserMINUS, FigParserPLUSPLUS, FigParserMINUSMINUS, FigParserEXCLAM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(485)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1077248) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(486)
			p.Unary()
		}

	case FigParserLPAREN, FigParserTK_NULL, FigParserTK_FN, FigParserTK_THIS, FigParserTK_TRY, FigParserTK_MATCH, FigParserLBRACKET, FigParserLBRACE, FigParserBOOL, FigParserID, FigParserNUMBER, FigParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(487)
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
	AllMemberName() []IMemberNameContext
	MemberName(i int) IMemberNameContext
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

func (s *PostfixContext) AllMemberName() []IMemberNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMemberNameContext); ok {
			len++
		}
	}

	tst := make([]IMemberNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMemberNameContext); ok {
			tst[i] = t.(IMemberNameContext)
			i++
		}
	}

	return tst
}

func (s *PostfixContext) MemberName(i int) IMemberNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemberNameContext); ok {
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

	return t.(IMemberNameContext)
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
	p.EnterRule(localctx, 80, FigParserRULE_postfix)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.Primary()
	}
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(502)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case FigParserLBRACKET:
				{
					p.SetState(491)
					p.Match(FigParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(492)
					p.Expr()
				}
				{
					p.SetState(493)
					p.Match(FigParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case FigParserDOT:
				{
					p.SetState(495)
					p.Match(FigParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(496)
					p.MemberName()
				}

			case FigParserLPAREN:
				{
					p.SetState(497)
					p.Match(FigParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(499)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
					{
						p.SetState(498)
						p.FnArgs()
					}

				}
				{
					p.SetState(501)
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
		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
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

// IMemberNameContext is an interface to support dynamic dispatch.
type IMemberNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TK_MATCH() antlr.TerminalNode

	// IsMemberNameContext differentiates from other interfaces.
	IsMemberNameContext()
}

type MemberNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberNameContext() *MemberNameContext {
	var p = new(MemberNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_memberName
	return p
}

func InitEmptyMemberNameContext(p *MemberNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_memberName
}

func (*MemberNameContext) IsMemberNameContext() {}

func NewMemberNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberNameContext {
	var p = new(MemberNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_memberName

	return p
}

func (s *MemberNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberNameContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *MemberNameContext) TK_MATCH() antlr.TerminalNode {
	return s.GetToken(FigParserTK_MATCH, 0)
}

func (s *MemberNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterMemberName(s)
	}
}

func (s *MemberNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitMemberName(s)
	}
}

func (s *MemberNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitMemberName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) MemberName() (localctx IMemberNameContext) {
	localctx = NewMemberNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, FigParserRULE_memberName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FigParserTK_MATCH || _la == FigParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	MatchExpr() IMatchExprContext
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

func (s *PrimaryContext) MatchExpr() IMatchExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchExprContext)
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
	p.EnterRule(localctx, 84, FigParserRULE_primary)
	var _la int

	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(509)
			p.Match(FigParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(510)
			p.Match(FigParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(511)
			p.Match(FigParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(512)
			p.Match(FigParserTK_NULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(513)
			p.Match(FigParserTK_THIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(514)
			p.ArrayLiteral()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(515)
			p.ObjectLiteral()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(516)
			p.TryExpr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(517)
			p.MatchExpr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(518)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(519)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(521)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(520)
				p.FnParams()
			}

		}
		{
			p.SetState(523)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)
			p.Block()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(525)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(528)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
			{
				p.SetState(527)
				p.FnArgs()
			}

		}
		{
			p.SetState(530)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(531)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(533)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(532)
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

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(535)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(536)
			p.Expr()
		}
		{
			p.SetState(537)
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
	TK_ONERROR() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	Expr() IExprContext
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

func (s *TryExprContext) TK_ONERROR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_ONERROR, 0)
}

func (s *TryExprContext) AllBlock() []IBlockContext {
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

func (s *TryExprContext) Block(i int) IBlockContext {
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
	p.EnterRule(localctx, 86, FigParserRULE_tryExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(541)
		p.Match(FigParserTK_TRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(542)
			p.Expr()
		}

	case 2:
		{
			p.SetState(543)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(546)
		p.Match(FigParserTK_ONERROR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserLPAREN {
		{
			p.SetState(547)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(549)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(548)
				p.Match(FigParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(551)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(554)
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

// IMatchExprContext is an interface to support dynamic dispatch.
type IMatchExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_MATCH() antlr.TerminalNode
	Expr() IExprContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllMatchArm() []IMatchArmContext
	MatchArm(i int) IMatchArmContext

	// IsMatchExprContext differentiates from other interfaces.
	IsMatchExprContext()
}

type MatchExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchExprContext() *MatchExprContext {
	var p = new(MatchExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_matchExpr
	return p
}

func InitEmptyMatchExprContext(p *MatchExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_matchExpr
}

func (*MatchExprContext) IsMatchExprContext() {}

func NewMatchExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchExprContext {
	var p = new(MatchExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_matchExpr

	return p
}

func (s *MatchExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchExprContext) TK_MATCH() antlr.TerminalNode {
	return s.GetToken(FigParserTK_MATCH, 0)
}

func (s *MatchExprContext) Expr() IExprContext {
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

func (s *MatchExprContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACE, 0)
}

func (s *MatchExprContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACE, 0)
}

func (s *MatchExprContext) AllMatchArm() []IMatchArmContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatchArmContext); ok {
			len++
		}
	}

	tst := make([]IMatchArmContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatchArmContext); ok {
			tst[i] = t.(IMatchArmContext)
			i++
		}
	}

	return tst
}

func (s *MatchExprContext) MatchArm(i int) IMatchArmContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchArmContext); ok {
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

	return t.(IMatchArmContext)
}

func (s *MatchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterMatchExpr(s)
	}
}

func (s *MatchExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitMatchExpr(s)
	}
}

func (s *MatchExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitMatchExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) MatchExpr() (localctx IMatchExprContext) {
	localctx = NewMatchExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, FigParserRULE_matchExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(556)
		p.Match(FigParserTK_MATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(557)
		p.Expr()
	}
	{
		p.SetState(558)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(560)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0) {
		{
			p.SetState(559)
			p.MatchArm()
		}

		p.SetState(562)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(564)
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

// IMatchArmContext is an interface to support dynamic dispatch.
type IMatchArmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMatchArmContext differentiates from other interfaces.
	IsMatchArmContext()
}

type MatchArmContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchArmContext() *MatchArmContext {
	var p = new(MatchArmContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_matchArm
	return p
}

func InitEmptyMatchArmContext(p *MatchArmContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_matchArm
}

func (*MatchArmContext) IsMatchArmContext() {}

func NewMatchArmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchArmContext {
	var p = new(MatchArmContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_matchArm

	return p
}

func (s *MatchArmContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchArmContext) CopyAll(ctx *MatchArmContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MatchArmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchArmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MatchArmCaseContext struct {
	MatchArmContext
}

func NewMatchArmCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchArmCaseContext {
	var p = new(MatchArmCaseContext)

	InitEmptyMatchArmContext(&p.MatchArmContext)
	p.parser = parser
	p.CopyAll(ctx.(*MatchArmContext))

	return p
}

func (s *MatchArmCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchArmCaseContext) MatchPattern() IMatchPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchPatternContext)
}

func (s *MatchArmCaseContext) ARROW() antlr.TerminalNode {
	return s.GetToken(FigParserARROW, 0)
}

func (s *MatchArmCaseContext) Block() IBlockContext {
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

func (s *MatchArmCaseContext) Expr() IExprContext {
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

func (s *MatchArmCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterMatchArmCase(s)
	}
}

func (s *MatchArmCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitMatchArmCase(s)
	}
}

func (s *MatchArmCaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitMatchArmCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) MatchArm() (localctx IMatchArmContext) {
	localctx = NewMatchArmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, FigParserRULE_matchArm)
	localctx = NewMatchArmCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(566)
		p.MatchPattern()
	}
	{
		p.SetState(567)
		p.Match(FigParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(568)
			p.Block()
		}

	case 2:
		{
			p.SetState(569)
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

// IMatchPatternContext is an interface to support dynamic dispatch.
type IMatchPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsMatchPatternContext differentiates from other interfaces.
	IsMatchPatternContext()
}

type MatchPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchPatternContext() *MatchPatternContext {
	var p = new(MatchPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_matchPattern
	return p
}

func InitEmptyMatchPatternContext(p *MatchPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FigParserRULE_matchPattern
}

func (*MatchPatternContext) IsMatchPatternContext() {}

func NewMatchPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchPatternContext {
	var p = new(MatchPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FigParserRULE_matchPattern

	return p
}

func (s *MatchPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchPatternContext) AllExpr() []IExprContext {
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

func (s *MatchPatternContext) Expr(i int) IExprContext {
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

func (s *MatchPatternContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *MatchPatternContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *MatchPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterMatchPattern(s)
	}
}

func (s *MatchPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitMatchPattern(s)
	}
}

func (s *MatchPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitMatchPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) MatchPattern() (localctx IMatchPatternContext) {
	localctx = NewMatchPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, FigParserRULE_matchPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(572)
		p.Expr()
	}
	p.SetState(577)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(573)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(574)
			p.Expr()
		}

		p.SetState(579)
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

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func (s *ArrayLiteralContext) CopyAll(ctx *ArrayLiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayLiteralSimpleContext struct {
	ArrayLiteralContext
}

func NewArrayLiteralSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralSimpleContext {
	var p = new(ArrayLiteralSimpleContext)

	InitEmptyArrayLiteralContext(&p.ArrayLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLiteralContext))

	return p
}

func (s *ArrayLiteralSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralSimpleContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACKET, 0)
}

func (s *ArrayLiteralSimpleContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, 0)
}

func (s *ArrayLiteralSimpleContext) AllExpr() []IExprContext {
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

func (s *ArrayLiteralSimpleContext) Expr(i int) IExprContext {
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

func (s *ArrayLiteralSimpleContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ArrayLiteralSimpleContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *ArrayLiteralSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterArrayLiteralSimple(s)
	}
}

func (s *ArrayLiteralSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitArrayLiteralSimple(s)
	}
}

func (s *ArrayLiteralSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitArrayLiteralSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayCompForRangeContext struct {
	ArrayLiteralContext
}

func NewArrayCompForRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayCompForRangeContext {
	var p = new(ArrayCompForRangeContext)

	InitEmptyArrayLiteralContext(&p.ArrayLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLiteralContext))

	return p
}

func (s *ArrayCompForRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayCompForRangeContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACKET, 0)
}

func (s *ArrayCompForRangeContext) AllExpr() []IExprContext {
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

func (s *ArrayCompForRangeContext) Expr(i int) IExprContext {
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

func (s *ArrayCompForRangeContext) TK_FOR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FOR, 0)
}

func (s *ArrayCompForRangeContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ArrayCompForRangeContext) TK_IN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_IN, 0)
}

func (s *ArrayCompForRangeContext) TK_RANGE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_RANGE, 0)
}

func (s *ArrayCompForRangeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *ArrayCompForRangeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *ArrayCompForRangeContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, 0)
}

func (s *ArrayCompForRangeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ArrayCompForRangeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *ArrayCompForRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterArrayCompForRange(s)
	}
}

func (s *ArrayCompForRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitArrayCompForRange(s)
	}
}

func (s *ArrayCompForRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitArrayCompForRange(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayCompForEnumerateContext struct {
	ArrayLiteralContext
}

func NewArrayCompForEnumerateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayCompForEnumerateContext {
	var p = new(ArrayCompForEnumerateContext)

	InitEmptyArrayLiteralContext(&p.ArrayLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLiteralContext))

	return p
}

func (s *ArrayCompForEnumerateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayCompForEnumerateContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACKET, 0)
}

func (s *ArrayCompForEnumerateContext) AllExpr() []IExprContext {
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

func (s *ArrayCompForEnumerateContext) Expr(i int) IExprContext {
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

func (s *ArrayCompForEnumerateContext) TK_FOR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FOR, 0)
}

func (s *ArrayCompForEnumerateContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(FigParserID)
}

func (s *ArrayCompForEnumerateContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(FigParserID, i)
}

func (s *ArrayCompForEnumerateContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, 0)
}

func (s *ArrayCompForEnumerateContext) TK_IN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_IN, 0)
}

func (s *ArrayCompForEnumerateContext) TK_ENUMERATE() antlr.TerminalNode {
	return s.GetToken(FigParserTK_ENUMERATE, 0)
}

func (s *ArrayCompForEnumerateContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserLPAREN, 0)
}

func (s *ArrayCompForEnumerateContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *ArrayCompForEnumerateContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, 0)
}

func (s *ArrayCompForEnumerateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterArrayCompForEnumerate(s)
	}
}

func (s *ArrayCompForEnumerateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitArrayCompForEnumerate(s)
	}
}

func (s *ArrayCompForEnumerateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitArrayCompForEnumerate(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayCompForInContext struct {
	ArrayLiteralContext
}

func NewArrayCompForInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayCompForInContext {
	var p = new(ArrayCompForInContext)

	InitEmptyArrayLiteralContext(&p.ArrayLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ArrayLiteralContext))

	return p
}

func (s *ArrayCompForInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayCompForInContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserLBRACKET, 0)
}

func (s *ArrayCompForInContext) AllExpr() []IExprContext {
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

func (s *ArrayCompForInContext) Expr(i int) IExprContext {
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

func (s *ArrayCompForInContext) TK_FOR() antlr.TerminalNode {
	return s.GetToken(FigParserTK_FOR, 0)
}

func (s *ArrayCompForInContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
}

func (s *ArrayCompForInContext) TK_IN() antlr.TerminalNode {
	return s.GetToken(FigParserTK_IN, 0)
}

func (s *ArrayCompForInContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, 0)
}

func (s *ArrayCompForInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.EnterArrayCompForIn(s)
	}
}

func (s *ArrayCompForInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FigParserListener); ok {
		listenerT.ExitArrayCompForIn(s)
	}
}

func (s *ArrayCompForInContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FigParserVisitor:
		return t.VisitArrayCompForIn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FigParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, FigParserRULE_arrayLiteral)
	var _la int

	p.SetState(632)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 74, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrayLiteralSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(580)
			p.Match(FigParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(589)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
			{
				p.SetState(581)
				p.Expr()
			}
			p.SetState(586)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FigParserCOMMA {
				{
					p.SetState(582)
					p.Match(FigParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(583)
					p.Expr()
				}

				p.SetState(588)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(591)
			p.Match(FigParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewArrayCompForInContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(592)
			p.Match(FigParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(593)
			p.Expr()
		}
		{
			p.SetState(594)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(595)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(596)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(597)
			p.Expr()
		}
		{
			p.SetState(598)
			p.Match(FigParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewArrayCompForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(600)
			p.Match(FigParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)
			p.Expr()
		}
		{
			p.SetState(602)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(603)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(604)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(605)
			p.Match(FigParserTK_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(606)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(607)
			p.Expr()
		}
		p.SetState(614)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserCOMMA {
			{
				p.SetState(608)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(609)
				p.Expr()
			}
			p.SetState(612)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == FigParserCOMMA {
				{
					p.SetState(610)
					p.Match(FigParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(611)
					p.Expr()
				}

			}

		}
		{
			p.SetState(616)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(617)
			p.Match(FigParserRBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewArrayCompForEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(619)
			p.Match(FigParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(620)
			p.Expr()
		}
		{
			p.SetState(621)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(623)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(624)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(625)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(626)
			p.Match(FigParserTK_ENUMERATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(627)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(628)
			p.Expr()
		}
		{
			p.SetState(629)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)
			p.Match(FigParserRBRACKET)
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
	p.EnterRule(localctx, 96, FigParserRULE_objectLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(634)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(643)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID || _la == FigParserSTRING {
		{
			p.SetState(635)
			p.ObjectEntry()
		}
		p.SetState(640)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(636)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(637)
				p.ObjectEntry()
			}

			p.SetState(642)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(645)
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
	p.EnterRule(localctx, 98, FigParserRULE_objectEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(647)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FigParserID || _la == FigParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(648)
		p.Match(FigParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(649)
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
