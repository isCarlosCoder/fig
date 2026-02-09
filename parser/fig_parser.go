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
		"'='", "'+'", "'-'", "'++'", "'--'", "'*'", "'/'", "'('", "')'", "'!'",
		"';'", "'%'", "'let'", "'print'", "'if'", "'elif'", "'else'", "'while'",
		"'do'", "'break'", "'continue'", "'for'", "'null'", "'fn'", "'return'",
		"'import'", "'use'", "'in'", "'range'", "'enumerate'", "'struct'", "'this'",
		"'try'", "'onerror'", "'match'", "','", "'=>'", "'['", "']'", "':'",
		"'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "LT", "GT", "LE", "GE", "EQ", "NEQ", "AND", "OR", "DOT", "ASSIGN",
		"PLUS", "MINUS", "PLUSPLUS", "MINUSMINUS", "STAR", "SLASH", "LPAREN",
		"RPAREN", "EXCLAM", "SEMICOLON", "MOD", "TK_LET", "TK_PRINT", "TK_IF",
		"TK_ELIF", "TK_ELSE", "TK_WHILE", "TK_DO", "TK_BREAK", "TK_CONTINUE",
		"TK_FOR", "TK_NULL", "TK_FN", "TK_RETURN", "TK_IMPORT", "TK_USE", "TK_IN",
		"TK_RANGE", "TK_ENUMERATE", "TK_STRUCT", "TK_THIS", "TK_TRY", "TK_ONERROR",
		"TK_MATCH", "COMMA", "ARROW", "LBRACKET", "RBRACKET", "COLON", "LBRACE",
		"RBRACE", "BOOL", "ID", "NUMBER", "STRING", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "statements", "exprStmt", "ifStmt", "whileStmt", "doWhileStmt",
		"breakStmt", "continueStmt", "fnDecl", "fnParams", "returnStmt", "importStmt",
		"useStmt", "fnArgs", "forInit", "forStep", "forStmt", "forInStmt", "block",
		"structDecl", "structMember", "varDeclaration", "varAtribuition", "memberAssign",
		"printStmt", "expr", "logicalOr", "logicalAnd", "equality", "comparison",
		"term", "factor", "unary", "postfix", "memberName", "primary", "tryExpr",
		"matchExpr", "matchArm", "matchPattern", "arrayLiteral", "objectLiteral",
		"objectEntry",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 57, 531, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 5, 0, 88, 8, 0, 10, 0, 12, 0, 91, 9, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 112, 8, 1, 1, 2, 1, 2, 3, 2, 116, 8, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		129, 8, 3, 10, 3, 12, 3, 132, 9, 3, 1, 3, 1, 3, 3, 3, 136, 8, 3, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 151, 8, 5, 1, 6, 1, 6, 3, 6, 155, 8, 6, 1, 7, 1, 7, 3, 7, 159, 8,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 165, 8, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 5, 9, 173, 8, 9, 10, 9, 12, 9, 176, 9, 9, 1, 10, 1, 10, 3, 10,
		180, 8, 10, 1, 10, 3, 10, 183, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 188,
		8, 11, 1, 11, 3, 11, 191, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 196, 8, 12,
		1, 13, 1, 13, 1, 13, 5, 13, 201, 8, 13, 10, 13, 12, 13, 204, 9, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 3, 14, 210, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 216, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 222, 8, 15, 1, 16,
		1, 16, 1, 16, 3, 16, 227, 8, 16, 1, 16, 1, 16, 3, 16, 231, 8, 16, 1, 16,
		1, 16, 3, 16, 235, 8, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 261, 8, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 272,
		8, 17, 1, 18, 1, 18, 5, 18, 276, 8, 18, 10, 18, 12, 18, 279, 9, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 287, 8, 19, 10, 19, 12, 19,
		290, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 3, 20, 297, 8, 20, 1, 20,
		3, 20, 300, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 306, 8, 20, 1, 20,
		1, 20, 3, 20, 310, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 316, 8, 21,
		1, 21, 3, 21, 319, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 325, 8, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 4, 23, 334, 8, 23, 11,
		23, 12, 23, 335, 1, 23, 1, 23, 1, 23, 3, 23, 341, 8, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 5, 24, 348, 8, 24, 10, 24, 12, 24, 351, 9, 24, 3,
		24, 353, 8, 24, 1, 24, 1, 24, 3, 24, 357, 8, 24, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 5, 26, 364, 8, 26, 10, 26, 12, 26, 367, 9, 26, 1, 27, 1, 27,
		1, 27, 5, 27, 372, 8, 27, 10, 27, 12, 27, 375, 9, 27, 1, 28, 1, 28, 1,
		28, 5, 28, 380, 8, 28, 10, 28, 12, 28, 383, 9, 28, 1, 29, 1, 29, 1, 29,
		5, 29, 388, 8, 29, 10, 29, 12, 29, 391, 9, 29, 1, 30, 1, 30, 1, 30, 5,
		30, 396, 8, 30, 10, 30, 12, 30, 399, 9, 30, 1, 31, 1, 31, 1, 31, 5, 31,
		404, 8, 31, 10, 31, 12, 31, 407, 9, 31, 1, 32, 1, 32, 1, 32, 3, 32, 412,
		8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3,
		33, 423, 8, 33, 1, 33, 5, 33, 426, 8, 33, 10, 33, 12, 33, 429, 9, 33, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 3, 35, 445, 8, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 3, 35, 452, 8, 35, 1, 35, 1, 35, 1, 35, 3, 35, 457, 8, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 3, 35, 463, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		3, 36, 470, 8, 36, 1, 36, 3, 36, 473, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37,
		1, 37, 1, 37, 4, 37, 481, 8, 37, 11, 37, 12, 37, 482, 1, 37, 1, 37, 1,
		38, 1, 38, 1, 38, 1, 38, 3, 38, 491, 8, 38, 1, 39, 1, 39, 1, 39, 5, 39,
		496, 8, 39, 10, 39, 12, 39, 499, 9, 39, 1, 40, 1, 40, 1, 40, 1, 40, 5,
		40, 505, 8, 40, 10, 40, 12, 40, 508, 9, 40, 3, 40, 510, 8, 40, 1, 40, 1,
		40, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 518, 8, 41, 10, 41, 12, 41, 521,
		9, 41, 3, 41, 523, 8, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 0, 0, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 8, 1, 0, 5, 6, 1, 0, 1, 4, 1, 0,
		11, 12, 2, 0, 15, 16, 21, 21, 2, 0, 12, 14, 19, 19, 2, 0, 44, 44, 53, 53,
		1, 0, 13, 14, 2, 0, 53, 53, 55, 55, 578, 0, 89, 1, 0, 0, 0, 2, 111, 1,
		0, 0, 0, 4, 113, 1, 0, 0, 0, 6, 117, 1, 0, 0, 0, 8, 137, 1, 0, 0, 0, 10,
		143, 1, 0, 0, 0, 12, 152, 1, 0, 0, 0, 14, 156, 1, 0, 0, 0, 16, 160, 1,
		0, 0, 0, 18, 169, 1, 0, 0, 0, 20, 177, 1, 0, 0, 0, 22, 184, 1, 0, 0, 0,
		24, 192, 1, 0, 0, 0, 26, 197, 1, 0, 0, 0, 28, 215, 1, 0, 0, 0, 30, 221,
		1, 0, 0, 0, 32, 223, 1, 0, 0, 0, 34, 271, 1, 0, 0, 0, 36, 273, 1, 0, 0,
		0, 38, 282, 1, 0, 0, 0, 40, 309, 1, 0, 0, 0, 42, 311, 1, 0, 0, 0, 44, 320,
		1, 0, 0, 0, 46, 326, 1, 0, 0, 0, 48, 342, 1, 0, 0, 0, 50, 358, 1, 0, 0,
		0, 52, 360, 1, 0, 0, 0, 54, 368, 1, 0, 0, 0, 56, 376, 1, 0, 0, 0, 58, 384,
		1, 0, 0, 0, 60, 392, 1, 0, 0, 0, 62, 400, 1, 0, 0, 0, 64, 411, 1, 0, 0,
		0, 66, 413, 1, 0, 0, 0, 68, 430, 1, 0, 0, 0, 70, 462, 1, 0, 0, 0, 72, 464,
		1, 0, 0, 0, 74, 476, 1, 0, 0, 0, 76, 486, 1, 0, 0, 0, 78, 492, 1, 0, 0,
		0, 80, 500, 1, 0, 0, 0, 82, 513, 1, 0, 0, 0, 84, 526, 1, 0, 0, 0, 86, 88,
		3, 2, 1, 0, 87, 86, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5,
		0, 0, 1, 93, 1, 1, 0, 0, 0, 94, 112, 3, 42, 21, 0, 95, 112, 3, 44, 22,
		0, 96, 112, 3, 46, 23, 0, 97, 112, 3, 48, 24, 0, 98, 112, 3, 6, 3, 0, 99,
		112, 3, 8, 4, 0, 100, 112, 3, 10, 5, 0, 101, 112, 3, 32, 16, 0, 102, 112,
		3, 34, 17, 0, 103, 112, 3, 12, 6, 0, 104, 112, 3, 14, 7, 0, 105, 112, 3,
		16, 8, 0, 106, 112, 3, 20, 10, 0, 107, 112, 3, 22, 11, 0, 108, 112, 3,
		24, 12, 0, 109, 112, 3, 38, 19, 0, 110, 112, 3, 4, 2, 0, 111, 94, 1, 0,
		0, 0, 111, 95, 1, 0, 0, 0, 111, 96, 1, 0, 0, 0, 111, 97, 1, 0, 0, 0, 111,
		98, 1, 0, 0, 0, 111, 99, 1, 0, 0, 0, 111, 100, 1, 0, 0, 0, 111, 101, 1,
		0, 0, 0, 111, 102, 1, 0, 0, 0, 111, 103, 1, 0, 0, 0, 111, 104, 1, 0, 0,
		0, 111, 105, 1, 0, 0, 0, 111, 106, 1, 0, 0, 0, 111, 107, 1, 0, 0, 0, 111,
		108, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 3, 1,
		0, 0, 0, 113, 115, 3, 50, 25, 0, 114, 116, 5, 20, 0, 0, 115, 114, 1, 0,
		0, 0, 115, 116, 1, 0, 0, 0, 116, 5, 1, 0, 0, 0, 117, 118, 5, 24, 0, 0,
		118, 119, 5, 17, 0, 0, 119, 120, 3, 50, 25, 0, 120, 121, 5, 18, 0, 0, 121,
		130, 3, 36, 18, 0, 122, 123, 5, 25, 0, 0, 123, 124, 5, 17, 0, 0, 124, 125,
		3, 50, 25, 0, 125, 126, 5, 18, 0, 0, 126, 127, 3, 36, 18, 0, 127, 129,
		1, 0, 0, 0, 128, 122, 1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0,
		0, 0, 130, 131, 1, 0, 0, 0, 131, 135, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0,
		133, 134, 5, 26, 0, 0, 134, 136, 3, 36, 18, 0, 135, 133, 1, 0, 0, 0, 135,
		136, 1, 0, 0, 0, 136, 7, 1, 0, 0, 0, 137, 138, 5, 27, 0, 0, 138, 139, 5,
		17, 0, 0, 139, 140, 3, 50, 25, 0, 140, 141, 5, 18, 0, 0, 141, 142, 3, 36,
		18, 0, 142, 9, 1, 0, 0, 0, 143, 144, 5, 28, 0, 0, 144, 145, 3, 36, 18,
		0, 145, 146, 5, 27, 0, 0, 146, 147, 5, 17, 0, 0, 147, 148, 3, 50, 25, 0,
		148, 150, 5, 18, 0, 0, 149, 151, 5, 20, 0, 0, 150, 149, 1, 0, 0, 0, 150,
		151, 1, 0, 0, 0, 151, 11, 1, 0, 0, 0, 152, 154, 5, 29, 0, 0, 153, 155,
		5, 20, 0, 0, 154, 153, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 13, 1, 0,
		0, 0, 156, 158, 5, 30, 0, 0, 157, 159, 5, 20, 0, 0, 158, 157, 1, 0, 0,
		0, 158, 159, 1, 0, 0, 0, 159, 15, 1, 0, 0, 0, 160, 161, 5, 33, 0, 0, 161,
		162, 5, 53, 0, 0, 162, 164, 5, 17, 0, 0, 163, 165, 3, 18, 9, 0, 164, 163,
		1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 5, 18,
		0, 0, 167, 168, 3, 36, 18, 0, 168, 17, 1, 0, 0, 0, 169, 174, 5, 53, 0,
		0, 170, 171, 5, 45, 0, 0, 171, 173, 5, 53, 0, 0, 172, 170, 1, 0, 0, 0,
		173, 176, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175,
		19, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 177, 179, 5, 34, 0, 0, 178, 180,
		3, 50, 25, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182, 1,
		0, 0, 0, 181, 183, 5, 20, 0, 0, 182, 181, 1, 0, 0, 0, 182, 183, 1, 0, 0,
		0, 183, 21, 1, 0, 0, 0, 184, 185, 5, 35, 0, 0, 185, 187, 5, 55, 0, 0, 186,
		188, 5, 53, 0, 0, 187, 186, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 190,
		1, 0, 0, 0, 189, 191, 5, 20, 0, 0, 190, 189, 1, 0, 0, 0, 190, 191, 1, 0,
		0, 0, 191, 23, 1, 0, 0, 0, 192, 193, 5, 36, 0, 0, 193, 195, 5, 55, 0, 0,
		194, 196, 5, 20, 0, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196,
		25, 1, 0, 0, 0, 197, 202, 3, 50, 25, 0, 198, 199, 5, 45, 0, 0, 199, 201,
		3, 50, 25, 0, 200, 198, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1,
		0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 27, 1, 0, 0, 0, 204, 202, 1, 0, 0,
		0, 205, 206, 5, 22, 0, 0, 206, 209, 5, 53, 0, 0, 207, 208, 5, 10, 0, 0,
		208, 210, 3, 50, 25, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210,
		216, 1, 0, 0, 0, 211, 212, 5, 53, 0, 0, 212, 213, 5, 10, 0, 0, 213, 216,
		3, 50, 25, 0, 214, 216, 3, 50, 25, 0, 215, 205, 1, 0, 0, 0, 215, 211, 1,
		0, 0, 0, 215, 214, 1, 0, 0, 0, 216, 29, 1, 0, 0, 0, 217, 218, 5, 53, 0,
		0, 218, 219, 5, 10, 0, 0, 219, 222, 3, 50, 25, 0, 220, 222, 3, 50, 25,
		0, 221, 217, 1, 0, 0, 0, 221, 220, 1, 0, 0, 0, 222, 31, 1, 0, 0, 0, 223,
		224, 5, 31, 0, 0, 224, 226, 5, 17, 0, 0, 225, 227, 3, 28, 14, 0, 226, 225,
		1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 230, 5, 20,
		0, 0, 229, 231, 3, 50, 25, 0, 230, 229, 1, 0, 0, 0, 230, 231, 1, 0, 0,
		0, 231, 232, 1, 0, 0, 0, 232, 234, 5, 20, 0, 0, 233, 235, 3, 30, 15, 0,
		234, 233, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236,
		237, 5, 18, 0, 0, 237, 238, 3, 36, 18, 0, 238, 33, 1, 0, 0, 0, 239, 240,
		5, 31, 0, 0, 240, 241, 5, 53, 0, 0, 241, 242, 5, 45, 0, 0, 242, 243, 5,
		53, 0, 0, 243, 244, 5, 37, 0, 0, 244, 245, 5, 39, 0, 0, 245, 246, 5, 17,
		0, 0, 246, 247, 3, 50, 25, 0, 247, 248, 5, 18, 0, 0, 248, 249, 3, 36, 18,
		0, 249, 272, 1, 0, 0, 0, 250, 251, 5, 31, 0, 0, 251, 252, 5, 53, 0, 0,
		252, 253, 5, 37, 0, 0, 253, 254, 5, 38, 0, 0, 254, 255, 5, 17, 0, 0, 255,
		256, 3, 50, 25, 0, 256, 257, 5, 45, 0, 0, 257, 260, 3, 50, 25, 0, 258,
		259, 5, 45, 0, 0, 259, 261, 3, 50, 25, 0, 260, 258, 1, 0, 0, 0, 260, 261,
		1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 5, 18, 0, 0, 263, 264, 3, 36,
		18, 0, 264, 272, 1, 0, 0, 0, 265, 266, 5, 31, 0, 0, 266, 267, 5, 53, 0,
		0, 267, 268, 5, 37, 0, 0, 268, 269, 3, 50, 25, 0, 269, 270, 3, 36, 18,
		0, 270, 272, 1, 0, 0, 0, 271, 239, 1, 0, 0, 0, 271, 250, 1, 0, 0, 0, 271,
		265, 1, 0, 0, 0, 272, 35, 1, 0, 0, 0, 273, 277, 5, 50, 0, 0, 274, 276,
		3, 2, 1, 0, 275, 274, 1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0,
		0, 0, 277, 278, 1, 0, 0, 0, 278, 280, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0,
		280, 281, 5, 51, 0, 0, 281, 37, 1, 0, 0, 0, 282, 283, 5, 40, 0, 0, 283,
		284, 5, 53, 0, 0, 284, 288, 5, 50, 0, 0, 285, 287, 3, 40, 20, 0, 286, 285,
		1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0,
		0, 0, 289, 291, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 292, 5, 51, 0, 0,
		292, 39, 1, 0, 0, 0, 293, 296, 5, 53, 0, 0, 294, 295, 5, 10, 0, 0, 295,
		297, 3, 50, 25, 0, 296, 294, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 299,
		1, 0, 0, 0, 298, 300, 5, 20, 0, 0, 299, 298, 1, 0, 0, 0, 299, 300, 1, 0,
		0, 0, 300, 310, 1, 0, 0, 0, 301, 302, 5, 33, 0, 0, 302, 303, 5, 53, 0,
		0, 303, 305, 5, 17, 0, 0, 304, 306, 3, 18, 9, 0, 305, 304, 1, 0, 0, 0,
		305, 306, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0, 307, 308, 5, 18, 0, 0, 308,
		310, 3, 36, 18, 0, 309, 293, 1, 0, 0, 0, 309, 301, 1, 0, 0, 0, 310, 41,
		1, 0, 0, 0, 311, 312, 5, 22, 0, 0, 312, 315, 5, 53, 0, 0, 313, 314, 5,
		10, 0, 0, 314, 316, 3, 50, 25, 0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0,
		0, 0, 316, 318, 1, 0, 0, 0, 317, 319, 5, 20, 0, 0, 318, 317, 1, 0, 0, 0,
		318, 319, 1, 0, 0, 0, 319, 43, 1, 0, 0, 0, 320, 321, 5, 53, 0, 0, 321,
		322, 5, 10, 0, 0, 322, 324, 3, 50, 25, 0, 323, 325, 5, 20, 0, 0, 324, 323,
		1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 45, 1, 0, 0, 0, 326, 333, 3, 50,
		25, 0, 327, 328, 5, 47, 0, 0, 328, 329, 3, 50, 25, 0, 329, 330, 5, 48,
		0, 0, 330, 334, 1, 0, 0, 0, 331, 332, 5, 9, 0, 0, 332, 334, 3, 68, 34,
		0, 333, 327, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335,
		333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 338,
		5, 10, 0, 0, 338, 340, 3, 50, 25, 0, 339, 341, 5, 20, 0, 0, 340, 339, 1,
		0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 47, 1, 0, 0, 0, 342, 343, 5, 23, 0,
		0, 343, 352, 5, 17, 0, 0, 344, 349, 3, 50, 25, 0, 345, 346, 5, 45, 0, 0,
		346, 348, 3, 50, 25, 0, 347, 345, 1, 0, 0, 0, 348, 351, 1, 0, 0, 0, 349,
		347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 353, 1, 0, 0, 0, 351, 349,
		1, 0, 0, 0, 352, 344, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 354, 1, 0,
		0, 0, 354, 356, 5, 18, 0, 0, 355, 357, 5, 20, 0, 0, 356, 355, 1, 0, 0,
		0, 356, 357, 1, 0, 0, 0, 357, 49, 1, 0, 0, 0, 358, 359, 3, 52, 26, 0, 359,
		51, 1, 0, 0, 0, 360, 365, 3, 54, 27, 0, 361, 362, 5, 8, 0, 0, 362, 364,
		3, 54, 27, 0, 363, 361, 1, 0, 0, 0, 364, 367, 1, 0, 0, 0, 365, 363, 1,
		0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 53, 1, 0, 0, 0, 367, 365, 1, 0, 0,
		0, 368, 373, 3, 56, 28, 0, 369, 370, 5, 7, 0, 0, 370, 372, 3, 56, 28, 0,
		371, 369, 1, 0, 0, 0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373,
		374, 1, 0, 0, 0, 374, 55, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 381, 3,
		58, 29, 0, 377, 378, 7, 0, 0, 0, 378, 380, 3, 58, 29, 0, 379, 377, 1, 0,
		0, 0, 380, 383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0,
		382, 57, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384, 389, 3, 60, 30, 0, 385,
		386, 7, 1, 0, 0, 386, 388, 3, 60, 30, 0, 387, 385, 1, 0, 0, 0, 388, 391,
		1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 59, 1, 0,
		0, 0, 391, 389, 1, 0, 0, 0, 392, 397, 3, 62, 31, 0, 393, 394, 7, 2, 0,
		0, 394, 396, 3, 62, 31, 0, 395, 393, 1, 0, 0, 0, 396, 399, 1, 0, 0, 0,
		397, 395, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 61, 1, 0, 0, 0, 399, 397,
		1, 0, 0, 0, 400, 405, 3, 64, 32, 0, 401, 402, 7, 3, 0, 0, 402, 404, 3,
		64, 32, 0, 403, 401, 1, 0, 0, 0, 404, 407, 1, 0, 0, 0, 405, 403, 1, 0,
		0, 0, 405, 406, 1, 0, 0, 0, 406, 63, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0,
		408, 409, 7, 4, 0, 0, 409, 412, 3, 64, 32, 0, 410, 412, 3, 66, 33, 0, 411,
		408, 1, 0, 0, 0, 411, 410, 1, 0, 0, 0, 412, 65, 1, 0, 0, 0, 413, 427, 3,
		70, 35, 0, 414, 415, 5, 47, 0, 0, 415, 416, 3, 50, 25, 0, 416, 417, 5,
		48, 0, 0, 417, 426, 1, 0, 0, 0, 418, 419, 5, 9, 0, 0, 419, 426, 3, 68,
		34, 0, 420, 422, 5, 17, 0, 0, 421, 423, 3, 26, 13, 0, 422, 421, 1, 0, 0,
		0, 422, 423, 1, 0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 426, 5, 18, 0, 0, 425,
		414, 1, 0, 0, 0, 425, 418, 1, 0, 0, 0, 425, 420, 1, 0, 0, 0, 426, 429,
		1, 0, 0, 0, 427, 425, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 67, 1, 0,
		0, 0, 429, 427, 1, 0, 0, 0, 430, 431, 7, 5, 0, 0, 431, 69, 1, 0, 0, 0,
		432, 463, 5, 54, 0, 0, 433, 463, 5, 52, 0, 0, 434, 463, 5, 55, 0, 0, 435,
		463, 5, 32, 0, 0, 436, 463, 5, 41, 0, 0, 437, 463, 3, 80, 40, 0, 438, 463,
		3, 82, 41, 0, 439, 463, 3, 72, 36, 0, 440, 463, 3, 74, 37, 0, 441, 442,
		5, 33, 0, 0, 442, 444, 5, 17, 0, 0, 443, 445, 3, 18, 9, 0, 444, 443, 1,
		0, 0, 0, 444, 445, 1, 0, 0, 0, 445, 446, 1, 0, 0, 0, 446, 447, 5, 18, 0,
		0, 447, 463, 3, 36, 18, 0, 448, 449, 5, 53, 0, 0, 449, 451, 5, 17, 0, 0,
		450, 452, 3, 26, 13, 0, 451, 450, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452,
		453, 1, 0, 0, 0, 453, 463, 5, 18, 0, 0, 454, 456, 5, 53, 0, 0, 455, 457,
		7, 6, 0, 0, 456, 455, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 463, 1, 0,
		0, 0, 458, 459, 5, 17, 0, 0, 459, 460, 3, 50, 25, 0, 460, 461, 5, 18, 0,
		0, 461, 463, 1, 0, 0, 0, 462, 432, 1, 0, 0, 0, 462, 433, 1, 0, 0, 0, 462,
		434, 1, 0, 0, 0, 462, 435, 1, 0, 0, 0, 462, 436, 1, 0, 0, 0, 462, 437,
		1, 0, 0, 0, 462, 438, 1, 0, 0, 0, 462, 439, 1, 0, 0, 0, 462, 440, 1, 0,
		0, 0, 462, 441, 1, 0, 0, 0, 462, 448, 1, 0, 0, 0, 462, 454, 1, 0, 0, 0,
		462, 458, 1, 0, 0, 0, 463, 71, 1, 0, 0, 0, 464, 465, 5, 42, 0, 0, 465,
		466, 3, 50, 25, 0, 466, 472, 5, 43, 0, 0, 467, 469, 5, 17, 0, 0, 468, 470,
		5, 53, 0, 0, 469, 468, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 471, 1, 0,
		0, 0, 471, 473, 5, 18, 0, 0, 472, 467, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0,
		473, 474, 1, 0, 0, 0, 474, 475, 3, 36, 18, 0, 475, 73, 1, 0, 0, 0, 476,
		477, 5, 44, 0, 0, 477, 478, 3, 50, 25, 0, 478, 480, 5, 50, 0, 0, 479, 481,
		3, 76, 38, 0, 480, 479, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 480, 1,
		0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484, 485, 5, 51, 0,
		0, 485, 75, 1, 0, 0, 0, 486, 487, 3, 78, 39, 0, 487, 490, 5, 46, 0, 0,
		488, 491, 3, 36, 18, 0, 489, 491, 3, 50, 25, 0, 490, 488, 1, 0, 0, 0, 490,
		489, 1, 0, 0, 0, 491, 77, 1, 0, 0, 0, 492, 497, 3, 50, 25, 0, 493, 494,
		5, 45, 0, 0, 494, 496, 3, 50, 25, 0, 495, 493, 1, 0, 0, 0, 496, 499, 1,
		0, 0, 0, 497, 495, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 79, 1, 0, 0,
		0, 499, 497, 1, 0, 0, 0, 500, 509, 5, 47, 0, 0, 501, 506, 3, 50, 25, 0,
		502, 503, 5, 45, 0, 0, 503, 505, 3, 50, 25, 0, 504, 502, 1, 0, 0, 0, 505,
		508, 1, 0, 0, 0, 506, 504, 1, 0, 0, 0, 506, 507, 1, 0, 0, 0, 507, 510,
		1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 509, 501, 1, 0, 0, 0, 509, 510, 1, 0,
		0, 0, 510, 511, 1, 0, 0, 0, 511, 512, 5, 48, 0, 0, 512, 81, 1, 0, 0, 0,
		513, 522, 5, 50, 0, 0, 514, 519, 3, 84, 42, 0, 515, 516, 5, 45, 0, 0, 516,
		518, 3, 84, 42, 0, 517, 515, 1, 0, 0, 0, 518, 521, 1, 0, 0, 0, 519, 517,
		1, 0, 0, 0, 519, 520, 1, 0, 0, 0, 520, 523, 1, 0, 0, 0, 521, 519, 1, 0,
		0, 0, 522, 514, 1, 0, 0, 0, 522, 523, 1, 0, 0, 0, 523, 524, 1, 0, 0, 0,
		524, 525, 5, 51, 0, 0, 525, 83, 1, 0, 0, 0, 526, 527, 7, 7, 0, 0, 527,
		528, 5, 49, 0, 0, 528, 529, 3, 50, 25, 0, 529, 85, 1, 0, 0, 0, 62, 89,
		111, 115, 130, 135, 150, 154, 158, 164, 174, 179, 182, 187, 190, 195, 202,
		209, 215, 221, 226, 230, 234, 260, 271, 277, 288, 296, 299, 305, 309, 315,
		318, 324, 333, 335, 340, 349, 352, 356, 365, 373, 381, 389, 397, 405, 411,
		422, 425, 427, 444, 451, 456, 462, 469, 472, 482, 490, 497, 506, 509, 519,
		522,
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
	FigParserTK_MATCH     = 44
	FigParserCOMMA        = 45
	FigParserARROW        = 46
	FigParserLBRACKET     = 47
	FigParserRBRACKET     = 48
	FigParserCOLON        = 49
	FigParserLBRACE       = 50
	FigParserRBRACE       = 51
	FigParserBOOL         = 52
	FigParserID           = 53
	FigParserNUMBER       = 54
	FigParserSTRING       = 55
	FigParserWS           = 56
	FigParserCOMMENT      = 57
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
	FigParserRULE_memberName     = 34
	FigParserRULE_primary        = 35
	FigParserRULE_tryExpr        = 36
	FigParserRULE_matchExpr      = 37
	FigParserRULE_matchArm       = 38
	FigParserRULE_matchPattern   = 39
	FigParserRULE_arrayLiteral   = 40
	FigParserRULE_objectLiteral  = 41
	FigParserRULE_objectEntry    = 42
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
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68846057907974144) != 0 {
		{
			p.SetState(86)
			p.Statements()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.VarDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.VarAtribuition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.MemberAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(97)
			p.PrintStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(98)
			p.IfStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(99)
			p.WhileStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(100)
			p.DoWhileStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(101)
			p.ForStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(102)
			p.ForInStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(103)
			p.BreakStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(104)
			p.ContinueStmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(105)
			p.FnDecl()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(106)
			p.ReturnStmt()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(107)
			p.ImportStmt()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(108)
			p.UseStmt()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(109)
			p.StructDecl()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(110)
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
		p.SetState(113)
		p.Expr()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(114)
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
		p.SetState(117)
		p.Match(FigParserTK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)
		p.Expr()
	}
	{
		p.SetState(120)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Block()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_ELIF {
		{
			p.SetState(122)
			p.Match(FigParserTK_ELIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Expr()
		}
		{
			p.SetState(125)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Block()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserTK_ELSE {
		{
			p.SetState(133)
			p.Match(FigParserTK_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)
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
		p.SetState(137)
		p.Match(FigParserTK_WHILE)
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
		p.SetState(143)
		p.Match(FigParserTK_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Block()
	}
	{
		p.SetState(145)
		p.Match(FigParserTK_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Expr()
	}
	{
		p.SetState(148)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(149)
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
		p.SetState(152)
		p.Match(FigParserTK_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(153)
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
		p.SetState(156)
		p.Match(FigParserTK_CONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(157)
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
		p.SetState(160)
		p.Match(FigParserTK_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID {
		{
			p.SetState(163)
			p.FnParams()
		}

	}
	{
		p.SetState(166)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
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
		p.SetState(169)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(170)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(176)
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
		p.SetState(177)
		p.Match(FigParserTK_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(178)
			p.Expr()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(181)
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
	ID() antlr.TerminalNode
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

func (s *ImportStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(FigParserID, 0)
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
		p.SetState(184)
		p.Match(FigParserTK_IMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(FigParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(186)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(189)
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
		p.SetState(192)
		p.Match(FigParserTK_USE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(FigParserSTRING)
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

	if _la == FigParserSEMICOLON {
		{
			p.SetState(194)
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
		p.SetState(197)
		p.Expr()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(198)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Expr()
		}

		p.SetState(204)
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

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(FigParserTK_LET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(207)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(208)
				p.Expr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(214)
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
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
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
		p.SetState(223)
		p.Match(FigParserTK_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68844833951346688) != 0 {
		{
			p.SetState(225)
			p.ForInit()
		}

	}
	{
		p.SetState(228)
		p.Match(FigParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68844833947152384) != 0 {
		{
			p.SetState(229)
			p.Expr()
		}

	}
	{
		p.SetState(232)
		p.Match(FigParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68844833947152384) != 0 {
		{
			p.SetState(233)
			p.ForStep()
		}

	}
	{
		p.SetState(236)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
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

	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(240)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(FigParserTK_ENUMERATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Expr()
		}
		{
			p.SetState(247)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Block()
		}

	case 2:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(251)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)
			p.Match(FigParserTK_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Expr()
		}
		{
			p.SetState(256)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Expr()
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserCOMMA {
			{
				p.SetState(258)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(259)
				p.Expr()
			}

		}
		{
			p.SetState(262)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Block()
		}

	case 3:
		localctx = NewForInContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.Match(FigParserTK_FOR)
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
			p.Expr()
		}
		{
			p.SetState(269)
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
		p.SetState(273)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68846057907974144) != 0 {
		{
			p.SetState(274)
			p.Statements()
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(280)
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
		p.SetState(282)
		p.Match(FigParserTK_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(284)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_FN || _la == FigParserID {
		{
			p.SetState(285)
			p.StructMember()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(291)
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

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		localctx = NewStructFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(294)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(295)
				p.Expr()
			}

		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserSEMICOLON {
			{
				p.SetState(298)
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
			p.SetState(301)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(304)
				p.FnParams()
			}

		}
		{
			p.SetState(307)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
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
		p.SetState(311)
		p.Match(FigParserTK_LET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserASSIGN {
		{
			p.SetState(313)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Expr()
		}

	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(317)
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
		p.SetState(320)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Expr()
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(323)
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
	p.EnterRule(localctx, 46, FigParserRULE_memberAssign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Expr()
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FigParserDOT || _la == FigParserLBRACKET {
		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case FigParserLBRACKET:
			{
				p.SetState(327)
				p.Match(FigParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(328)
				p.Expr()
			}
			{
				p.SetState(329)
				p.Match(FigParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case FigParserDOT:
			{
				p.SetState(331)
				p.Match(FigParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(332)
				p.MemberName()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(337)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(338)
		p.Expr()
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(339)
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
	p.EnterRule(localctx, 48, FigParserRULE_printStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(FigParserTK_PRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68844833947152384) != 0 {
		{
			p.SetState(344)
			p.Expr()
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(345)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(346)
				p.Expr()
			}

			p.SetState(351)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(354)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(355)
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
		p.SetState(358)
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
		p.SetState(360)
		p.LogicalAnd()
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserOR {
		{
			p.SetState(361)
			p.Match(FigParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(362)
			p.LogicalAnd()
		}

		p.SetState(367)
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
		p.SetState(368)
		p.Equality()
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserAND {
		{
			p.SetState(369)
			p.Match(FigParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.Equality()
		}

		p.SetState(375)
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
		p.SetState(376)
		p.Comparison()
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserEQ || _la == FigParserNEQ {
		{
			p.SetState(377)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FigParserEQ || _la == FigParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(378)
			p.Comparison()
		}

		p.SetState(383)
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
		p.SetState(384)
		p.Term()
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0 {
		{
			p.SetState(385)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(386)
			p.Term()
		}

		p.SetState(391)
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
		p.SetState(392)
		p.Factor()
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(393)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FigParserPLUS || _la == FigParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(394)
				p.Factor()
			}

		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
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
		p.SetState(400)
		p.Unary()
	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2195456) != 0 {
		{
			p.SetState(401)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2195456) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(402)
			p.Unary()
		}

		p.SetState(407)
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

	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserMINUS, FigParserPLUSPLUS, FigParserMINUSMINUS, FigParserEXCLAM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(408)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&552960) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(409)
			p.Unary()
		}

	case FigParserLPAREN, FigParserTK_NULL, FigParserTK_FN, FigParserTK_THIS, FigParserTK_TRY, FigParserTK_MATCH, FigParserLBRACKET, FigParserLBRACE, FigParserBOOL, FigParserID, FigParserNUMBER, FigParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
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
	p.EnterRule(localctx, 66, FigParserRULE_postfix)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Primary()
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(425)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case FigParserLBRACKET:
				{
					p.SetState(414)
					p.Match(FigParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(415)
					p.Expr()
				}
				{
					p.SetState(416)
					p.Match(FigParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case FigParserDOT:
				{
					p.SetState(418)
					p.Match(FigParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(419)
					p.MemberName()
				}

			case FigParserLPAREN:
				{
					p.SetState(420)
					p.Match(FigParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(422)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68844833947152384) != 0 {
					{
						p.SetState(421)
						p.FnArgs()
					}

				}
				{
					p.SetState(424)
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
		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 68, FigParserRULE_memberName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
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
	p.EnterRule(localctx, 70, FigParserRULE_primary)
	var _la int

	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)
			p.Match(FigParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
			p.Match(FigParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(434)
			p.Match(FigParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(435)
			p.Match(FigParserTK_NULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(436)
			p.Match(FigParserTK_THIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(437)
			p.ArrayLiteral()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(438)
			p.ObjectLiteral()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(439)
			p.TryExpr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(440)
			p.MatchExpr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(441)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(442)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(443)
				p.FnParams()
			}

		}
		{
			p.SetState(446)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
			p.Block()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(448)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(449)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(451)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68844833947152384) != 0 {
			{
				p.SetState(450)
				p.FnArgs()
			}

		}
		{
			p.SetState(453)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(454)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(455)
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
			p.SetState(458)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.Expr()
		}
		{
			p.SetState(460)
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
	p.EnterRule(localctx, 72, FigParserRULE_tryExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(464)
		p.Match(FigParserTK_TRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(465)
		p.Expr()
	}
	{
		p.SetState(466)
		p.Match(FigParserTK_ONERROR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserLPAREN {
		{
			p.SetState(467)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(468)
				p.Match(FigParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(471)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(474)
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
	p.EnterRule(localctx, 74, FigParserRULE_matchExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(476)
		p.Match(FigParserTK_MATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(477)
		p.Expr()
	}
	{
		p.SetState(478)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68844833947152384) != 0) {
		{
			p.SetState(479)
			p.MatchArm()
		}

		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(484)
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
	p.EnterRule(localctx, 76, FigParserRULE_matchArm)
	localctx = NewMatchArmCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(486)
		p.MatchPattern()
	}
	{
		p.SetState(487)
		p.Match(FigParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(488)
			p.Block()
		}

	case 2:
		{
			p.SetState(489)
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
	p.EnterRule(localctx, 78, FigParserRULE_matchPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(492)
		p.Expr()
	}
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(493)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(494)
			p.Expr()
		}

		p.SetState(499)
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
	p.EnterRule(localctx, 80, FigParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		p.Match(FigParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&68844833947152384) != 0 {
		{
			p.SetState(501)
			p.Expr()
		}
		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(502)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(503)
				p.Expr()
			}

			p.SetState(508)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(511)
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
	p.EnterRule(localctx, 82, FigParserRULE_objectLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID || _la == FigParserSTRING {
		{
			p.SetState(514)
			p.ObjectEntry()
		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(515)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(516)
				p.ObjectEntry()
			}

			p.SetState(521)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(524)
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
	p.EnterRule(localctx, 84, FigParserRULE_objectEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FigParserID || _la == FigParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(527)
		p.Match(FigParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(528)
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
