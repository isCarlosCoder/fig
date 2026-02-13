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
		4, 1, 59, 648, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
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
		1, 18, 3, 18, 285, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 3, 18, 296, 8, 18, 1, 19, 1, 19, 5, 19, 300, 8, 19, 10,
		19, 12, 19, 303, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20,
		311, 8, 20, 10, 20, 12, 20, 314, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 3, 21, 321, 8, 21, 1, 21, 3, 21, 324, 8, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 3, 21, 330, 8, 21, 1, 21, 1, 21, 3, 21, 334, 8, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 5, 22, 340, 8, 22, 10, 22, 12, 22, 343, 9, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 3, 23, 349, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 355,
		8, 24, 1, 24, 3, 24, 358, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 364,
		8, 25, 1, 26, 1, 26, 1, 26, 3, 26, 369, 8, 26, 1, 27, 1, 27, 1, 27, 1,
		27, 5, 27, 375, 8, 27, 10, 27, 12, 27, 378, 9, 27, 3, 27, 380, 8, 27, 1,
		27, 1, 27, 1, 28, 1, 28, 1, 28, 3, 28, 387, 8, 28, 1, 29, 1, 29, 1, 29,
		1, 29, 5, 29, 393, 8, 29, 10, 29, 12, 29, 396, 9, 29, 3, 29, 398, 8, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 4, 30, 409,
		8, 30, 11, 30, 12, 30, 410, 1, 30, 1, 30, 1, 30, 3, 30, 416, 8, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 423, 8, 31, 10, 31, 12, 31, 426,
		9, 31, 3, 31, 428, 8, 31, 1, 31, 1, 31, 3, 31, 432, 8, 31, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 5, 33, 439, 8, 33, 10, 33, 12, 33, 442, 9, 33, 1,
		34, 1, 34, 1, 34, 5, 34, 447, 8, 34, 10, 34, 12, 34, 450, 9, 34, 1, 35,
		1, 35, 1, 35, 5, 35, 455, 8, 35, 10, 35, 12, 35, 458, 9, 35, 1, 36, 1,
		36, 1, 36, 5, 36, 463, 8, 36, 10, 36, 12, 36, 466, 9, 36, 1, 37, 1, 37,
		1, 37, 5, 37, 471, 8, 37, 10, 37, 12, 37, 474, 9, 37, 1, 38, 1, 38, 1,
		38, 5, 38, 479, 8, 38, 10, 38, 12, 38, 482, 9, 38, 1, 39, 1, 39, 1, 39,
		3, 39, 487, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 3, 40, 498, 8, 40, 1, 40, 5, 40, 501, 8, 40, 10, 40, 12, 40,
		504, 9, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 520, 8, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 3, 42, 527, 8, 42, 1, 42, 1, 42, 1, 42, 3, 42, 532,
		8, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 538, 8, 42, 1, 43, 1, 43, 1,
		43, 3, 43, 543, 8, 43, 1, 43, 1, 43, 1, 43, 3, 43, 548, 8, 43, 1, 43, 3,
		43, 551, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 4, 44, 559, 8,
		44, 11, 44, 12, 44, 560, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45,
		569, 8, 45, 1, 46, 1, 46, 1, 46, 5, 46, 574, 8, 46, 10, 46, 12, 46, 577,
		9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 583, 8, 47, 10, 47, 12, 47, 586,
		9, 47, 3, 47, 588, 8, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 611, 8, 47, 1, 47, 1, 47, 1, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 3, 47, 629, 8, 47, 1, 48, 1, 48, 1, 48, 1, 48, 5, 48, 635,
		8, 48, 10, 48, 12, 48, 638, 9, 48, 3, 48, 640, 8, 48, 1, 48, 1, 48, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 0, 0, 50, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
		90, 92, 94, 96, 98, 0, 8, 1, 0, 5, 6, 1, 0, 1, 4, 1, 0, 11, 12, 2, 0, 15,
		16, 22, 22, 2, 0, 12, 14, 20, 20, 2, 0, 46, 46, 55, 55, 1, 0, 13, 14, 2,
		0, 55, 55, 57, 57, 706, 0, 103, 1, 0, 0, 0, 2, 126, 1, 0, 0, 0, 4, 128,
		1, 0, 0, 0, 6, 132, 1, 0, 0, 0, 8, 152, 1, 0, 0, 0, 10, 158, 1, 0, 0, 0,
		12, 167, 1, 0, 0, 0, 14, 171, 1, 0, 0, 0, 16, 175, 1, 0, 0, 0, 18, 184,
		1, 0, 0, 0, 20, 199, 1, 0, 0, 0, 22, 201, 1, 0, 0, 0, 24, 208, 1, 0, 0,
		0, 26, 216, 1, 0, 0, 0, 28, 221, 1, 0, 0, 0, 30, 239, 1, 0, 0, 0, 32, 245,
		1, 0, 0, 0, 34, 247, 1, 0, 0, 0, 36, 295, 1, 0, 0, 0, 38, 297, 1, 0, 0,
		0, 40, 306, 1, 0, 0, 0, 42, 333, 1, 0, 0, 0, 44, 335, 1, 0, 0, 0, 46, 346,
		1, 0, 0, 0, 48, 350, 1, 0, 0, 0, 50, 359, 1, 0, 0, 0, 52, 368, 1, 0, 0,
		0, 54, 370, 1, 0, 0, 0, 56, 386, 1, 0, 0, 0, 58, 388, 1, 0, 0, 0, 60, 401,
		1, 0, 0, 0, 62, 417, 1, 0, 0, 0, 64, 433, 1, 0, 0, 0, 66, 435, 1, 0, 0,
		0, 68, 443, 1, 0, 0, 0, 70, 451, 1, 0, 0, 0, 72, 459, 1, 0, 0, 0, 74, 467,
		1, 0, 0, 0, 76, 475, 1, 0, 0, 0, 78, 486, 1, 0, 0, 0, 80, 488, 1, 0, 0,
		0, 82, 505, 1, 0, 0, 0, 84, 537, 1, 0, 0, 0, 86, 539, 1, 0, 0, 0, 88, 554,
		1, 0, 0, 0, 90, 564, 1, 0, 0, 0, 92, 570, 1, 0, 0, 0, 94, 628, 1, 0, 0,
		0, 96, 630, 1, 0, 0, 0, 98, 643, 1, 0, 0, 0, 100, 102, 3, 2, 1, 0, 101,
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
		5, 57, 0, 0, 210, 212, 5, 55, 0, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1,
		0, 0, 0, 212, 214, 1, 0, 0, 0, 213, 215, 5, 21, 0, 0, 214, 213, 1, 0, 0,
		0, 214, 215, 1, 0, 0, 0, 215, 25, 1, 0, 0, 0, 216, 217, 5, 37, 0, 0, 217,
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
		19, 0, 273, 296, 1, 0, 0, 0, 274, 275, 5, 32, 0, 0, 275, 276, 5, 55, 0,
		0, 276, 277, 5, 38, 0, 0, 277, 278, 5, 39, 0, 0, 278, 279, 5, 17, 0, 0,
		279, 280, 3, 64, 32, 0, 280, 281, 5, 47, 0, 0, 281, 284, 3, 64, 32, 0,
		282, 283, 5, 47, 0, 0, 283, 285, 3, 64, 32, 0, 284, 282, 1, 0, 0, 0, 284,
		285, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 287, 5, 18, 0, 0, 287, 288,
		3, 38, 19, 0, 288, 296, 1, 0, 0, 0, 289, 290, 5, 32, 0, 0, 290, 291, 5,
		55, 0, 0, 291, 292, 5, 38, 0, 0, 292, 293, 3, 64, 32, 0, 293, 294, 3, 38,
		19, 0, 294, 296, 1, 0, 0, 0, 295, 263, 1, 0, 0, 0, 295, 274, 1, 0, 0, 0,
		295, 289, 1, 0, 0, 0, 296, 37, 1, 0, 0, 0, 297, 301, 5, 52, 0, 0, 298,
		300, 3, 2, 1, 0, 299, 298, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299,
		1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 301, 1, 0,
		0, 0, 304, 305, 5, 53, 0, 0, 305, 39, 1, 0, 0, 0, 306, 307, 5, 41, 0, 0,
		307, 308, 5, 55, 0, 0, 308, 312, 5, 52, 0, 0, 309, 311, 3, 42, 21, 0, 310,
		309, 1, 0, 0, 0, 311, 314, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 312, 313,
		1, 0, 0, 0, 313, 315, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 315, 316, 5, 53,
		0, 0, 316, 41, 1, 0, 0, 0, 317, 320, 5, 55, 0, 0, 318, 319, 5, 10, 0, 0,
		319, 321, 3, 64, 32, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		323, 1, 0, 0, 0, 322, 324, 5, 21, 0, 0, 323, 322, 1, 0, 0, 0, 323, 324,
		1, 0, 0, 0, 324, 334, 1, 0, 0, 0, 325, 326, 5, 34, 0, 0, 326, 327, 5, 55,
		0, 0, 327, 329, 5, 17, 0, 0, 328, 330, 3, 18, 9, 0, 329, 328, 1, 0, 0,
		0, 329, 330, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 332, 5, 18, 0, 0, 332,
		334, 3, 38, 19, 0, 333, 317, 1, 0, 0, 0, 333, 325, 1, 0, 0, 0, 334, 43,
		1, 0, 0, 0, 335, 336, 5, 42, 0, 0, 336, 337, 5, 55, 0, 0, 337, 341, 5,
		52, 0, 0, 338, 340, 3, 46, 23, 0, 339, 338, 1, 0, 0, 0, 340, 343, 1, 0,
		0, 0, 341, 339, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 344, 1, 0, 0, 0,
		343, 341, 1, 0, 0, 0, 344, 345, 5, 53, 0, 0, 345, 45, 1, 0, 0, 0, 346,
		348, 5, 55, 0, 0, 347, 349, 5, 21, 0, 0, 348, 347, 1, 0, 0, 0, 348, 349,
		1, 0, 0, 0, 349, 47, 1, 0, 0, 0, 350, 351, 5, 23, 0, 0, 351, 354, 3, 52,
		26, 0, 352, 353, 5, 10, 0, 0, 353, 355, 3, 64, 32, 0, 354, 352, 1, 0, 0,
		0, 354, 355, 1, 0, 0, 0, 355, 357, 1, 0, 0, 0, 356, 358, 5, 21, 0, 0, 357,
		356, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 49, 1, 0, 0, 0, 359, 360, 3,
		52, 26, 0, 360, 361, 5, 10, 0, 0, 361, 363, 3, 64, 32, 0, 362, 364, 5,
		21, 0, 0, 363, 362, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 51, 1, 0, 0,
		0, 365, 369, 5, 55, 0, 0, 366, 369, 3, 54, 27, 0, 367, 369, 3, 58, 29,
		0, 368, 365, 1, 0, 0, 0, 368, 366, 1, 0, 0, 0, 368, 367, 1, 0, 0, 0, 369,
		53, 1, 0, 0, 0, 370, 379, 5, 49, 0, 0, 371, 376, 3, 56, 28, 0, 372, 373,
		5, 47, 0, 0, 373, 375, 3, 56, 28, 0, 374, 372, 1, 0, 0, 0, 375, 378, 1,
		0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 380, 1, 0, 0,
		0, 378, 376, 1, 0, 0, 0, 379, 371, 1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380,
		381, 1, 0, 0, 0, 381, 382, 5, 50, 0, 0, 382, 55, 1, 0, 0, 0, 383, 387,
		5, 55, 0, 0, 384, 387, 3, 54, 27, 0, 385, 387, 3, 58, 29, 0, 386, 383,
		1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 386, 385, 1, 0, 0, 0, 387, 57, 1, 0,
		0, 0, 388, 397, 5, 52, 0, 0, 389, 394, 5, 55, 0, 0, 390, 391, 5, 47, 0,
		0, 391, 393, 5, 55, 0, 0, 392, 390, 1, 0, 0, 0, 393, 396, 1, 0, 0, 0, 394,
		392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396, 394,
		1, 0, 0, 0, 397, 389, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 1, 0,
		0, 0, 399, 400, 5, 53, 0, 0, 400, 59, 1, 0, 0, 0, 401, 408, 3, 64, 32,
		0, 402, 403, 5, 49, 0, 0, 403, 404, 3, 64, 32, 0, 404, 405, 5, 50, 0, 0,
		405, 409, 1, 0, 0, 0, 406, 407, 5, 9, 0, 0, 407, 409, 3, 82, 41, 0, 408,
		402, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 408,
		1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 413, 5, 10,
		0, 0, 413, 415, 3, 64, 32, 0, 414, 416, 5, 21, 0, 0, 415, 414, 1, 0, 0,
		0, 415, 416, 1, 0, 0, 0, 416, 61, 1, 0, 0, 0, 417, 418, 5, 24, 0, 0, 418,
		427, 5, 17, 0, 0, 419, 424, 3, 64, 32, 0, 420, 421, 5, 47, 0, 0, 421, 423,
		3, 64, 32, 0, 422, 420, 1, 0, 0, 0, 423, 426, 1, 0, 0, 0, 424, 422, 1,
		0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 428, 1, 0, 0, 0, 426, 424, 1, 0, 0,
		0, 427, 419, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429,
		431, 5, 18, 0, 0, 430, 432, 5, 21, 0, 0, 431, 430, 1, 0, 0, 0, 431, 432,
		1, 0, 0, 0, 432, 63, 1, 0, 0, 0, 433, 434, 3, 66, 33, 0, 434, 65, 1, 0,
		0, 0, 435, 440, 3, 68, 34, 0, 436, 437, 5, 8, 0, 0, 437, 439, 3, 68, 34,
		0, 438, 436, 1, 0, 0, 0, 439, 442, 1, 0, 0, 0, 440, 438, 1, 0, 0, 0, 440,
		441, 1, 0, 0, 0, 441, 67, 1, 0, 0, 0, 442, 440, 1, 0, 0, 0, 443, 448, 3,
		70, 35, 0, 444, 445, 5, 7, 0, 0, 445, 447, 3, 70, 35, 0, 446, 444, 1, 0,
		0, 0, 447, 450, 1, 0, 0, 0, 448, 446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0,
		449, 69, 1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 451, 456, 3, 72, 36, 0, 452,
		453, 7, 0, 0, 0, 453, 455, 3, 72, 36, 0, 454, 452, 1, 0, 0, 0, 455, 458,
		1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 71, 1, 0,
		0, 0, 458, 456, 1, 0, 0, 0, 459, 464, 3, 74, 37, 0, 460, 461, 7, 1, 0,
		0, 461, 463, 3, 74, 37, 0, 462, 460, 1, 0, 0, 0, 463, 466, 1, 0, 0, 0,
		464, 462, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465, 73, 1, 0, 0, 0, 466, 464,
		1, 0, 0, 0, 467, 472, 3, 76, 38, 0, 468, 469, 7, 2, 0, 0, 469, 471, 3,
		76, 38, 0, 470, 468, 1, 0, 0, 0, 471, 474, 1, 0, 0, 0, 472, 470, 1, 0,
		0, 0, 472, 473, 1, 0, 0, 0, 473, 75, 1, 0, 0, 0, 474, 472, 1, 0, 0, 0,
		475, 480, 3, 78, 39, 0, 476, 477, 7, 3, 0, 0, 477, 479, 3, 78, 39, 0, 478,
		476, 1, 0, 0, 0, 479, 482, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 480, 481,
		1, 0, 0, 0, 481, 77, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 483, 484, 7, 4,
		0, 0, 484, 487, 3, 78, 39, 0, 485, 487, 3, 80, 40, 0, 486, 483, 1, 0, 0,
		0, 486, 485, 1, 0, 0, 0, 487, 79, 1, 0, 0, 0, 488, 502, 3, 84, 42, 0, 489,
		490, 5, 49, 0, 0, 490, 491, 3, 64, 32, 0, 491, 492, 5, 50, 0, 0, 492, 501,
		1, 0, 0, 0, 493, 494, 5, 9, 0, 0, 494, 501, 3, 82, 41, 0, 495, 497, 5,
		17, 0, 0, 496, 498, 3, 28, 14, 0, 497, 496, 1, 0, 0, 0, 497, 498, 1, 0,
		0, 0, 498, 499, 1, 0, 0, 0, 499, 501, 5, 18, 0, 0, 500, 489, 1, 0, 0, 0,
		500, 493, 1, 0, 0, 0, 500, 495, 1, 0, 0, 0, 501, 504, 1, 0, 0, 0, 502,
		500, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 81, 1, 0, 0, 0, 504, 502, 1,
		0, 0, 0, 505, 506, 7, 5, 0, 0, 506, 83, 1, 0, 0, 0, 507, 538, 5, 56, 0,
		0, 508, 538, 5, 54, 0, 0, 509, 538, 5, 57, 0, 0, 510, 538, 5, 33, 0, 0,
		511, 538, 5, 43, 0, 0, 512, 538, 3, 94, 47, 0, 513, 538, 3, 96, 48, 0,
		514, 538, 3, 86, 43, 0, 515, 538, 3, 88, 44, 0, 516, 517, 5, 34, 0, 0,
		517, 519, 5, 17, 0, 0, 518, 520, 3, 18, 9, 0, 519, 518, 1, 0, 0, 0, 519,
		520, 1, 0, 0, 0, 520, 521, 1, 0, 0, 0, 521, 522, 5, 18, 0, 0, 522, 538,
		3, 38, 19, 0, 523, 524, 5, 55, 0, 0, 524, 526, 5, 17, 0, 0, 525, 527, 3,
		28, 14, 0, 526, 525, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 528, 1, 0,
		0, 0, 528, 538, 5, 18, 0, 0, 529, 531, 5, 55, 0, 0, 530, 532, 7, 6, 0,
		0, 531, 530, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 532, 538, 1, 0, 0, 0, 533,
		534, 5, 17, 0, 0, 534, 535, 3, 64, 32, 0, 535, 536, 5, 18, 0, 0, 536, 538,
		1, 0, 0, 0, 537, 507, 1, 0, 0, 0, 537, 508, 1, 0, 0, 0, 537, 509, 1, 0,
		0, 0, 537, 510, 1, 0, 0, 0, 537, 511, 1, 0, 0, 0, 537, 512, 1, 0, 0, 0,
		537, 513, 1, 0, 0, 0, 537, 514, 1, 0, 0, 0, 537, 515, 1, 0, 0, 0, 537,
		516, 1, 0, 0, 0, 537, 523, 1, 0, 0, 0, 537, 529, 1, 0, 0, 0, 537, 533,
		1, 0, 0, 0, 538, 85, 1, 0, 0, 0, 539, 542, 5, 44, 0, 0, 540, 543, 3, 64,
		32, 0, 541, 543, 3, 38, 19, 0, 542, 540, 1, 0, 0, 0, 542, 541, 1, 0, 0,
		0, 543, 544, 1, 0, 0, 0, 544, 550, 5, 45, 0, 0, 545, 547, 5, 17, 0, 0,
		546, 548, 5, 55, 0, 0, 547, 546, 1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548,
		549, 1, 0, 0, 0, 549, 551, 5, 18, 0, 0, 550, 545, 1, 0, 0, 0, 550, 551,
		1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 553, 3, 38, 19, 0, 553, 87, 1, 0,
		0, 0, 554, 555, 5, 46, 0, 0, 555, 556, 3, 64, 32, 0, 556, 558, 5, 52, 0,
		0, 557, 559, 3, 90, 45, 0, 558, 557, 1, 0, 0, 0, 559, 560, 1, 0, 0, 0,
		560, 558, 1, 0, 0, 0, 560, 561, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562,
		563, 5, 53, 0, 0, 563, 89, 1, 0, 0, 0, 564, 565, 3, 92, 46, 0, 565, 568,
		5, 48, 0, 0, 566, 569, 3, 38, 19, 0, 567, 569, 3, 64, 32, 0, 568, 566,
		1, 0, 0, 0, 568, 567, 1, 0, 0, 0, 569, 91, 1, 0, 0, 0, 570, 575, 3, 64,
		32, 0, 571, 572, 5, 47, 0, 0, 572, 574, 3, 64, 32, 0, 573, 571, 1, 0, 0,
		0, 574, 577, 1, 0, 0, 0, 575, 573, 1, 0, 0, 0, 575, 576, 1, 0, 0, 0, 576,
		93, 1, 0, 0, 0, 577, 575, 1, 0, 0, 0, 578, 587, 5, 49, 0, 0, 579, 584,
		3, 64, 32, 0, 580, 581, 5, 47, 0, 0, 581, 583, 3, 64, 32, 0, 582, 580,
		1, 0, 0, 0, 583, 586, 1, 0, 0, 0, 584, 582, 1, 0, 0, 0, 584, 585, 1, 0,
		0, 0, 585, 588, 1, 0, 0, 0, 586, 584, 1, 0, 0, 0, 587, 579, 1, 0, 0, 0,
		587, 588, 1, 0, 0, 0, 588, 589, 1, 0, 0, 0, 589, 629, 5, 50, 0, 0, 590,
		591, 5, 49, 0, 0, 591, 592, 3, 64, 32, 0, 592, 593, 5, 32, 0, 0, 593, 594,
		5, 55, 0, 0, 594, 595, 5, 38, 0, 0, 595, 596, 3, 64, 32, 0, 596, 597, 5,
		50, 0, 0, 597, 629, 1, 0, 0, 0, 598, 599, 5, 49, 0, 0, 599, 600, 3, 64,
		32, 0, 600, 601, 5, 32, 0, 0, 601, 602, 5, 55, 0, 0, 602, 603, 5, 38, 0,
		0, 603, 604, 5, 39, 0, 0, 604, 605, 5, 17, 0, 0, 605, 606, 3, 64, 32, 0,
		606, 607, 5, 47, 0, 0, 607, 610, 3, 64, 32, 0, 608, 609, 5, 47, 0, 0, 609,
		611, 3, 64, 32, 0, 610, 608, 1, 0, 0, 0, 610, 611, 1, 0, 0, 0, 611, 612,
		1, 0, 0, 0, 612, 613, 5, 18, 0, 0, 613, 614, 5, 50, 0, 0, 614, 629, 1,
		0, 0, 0, 615, 616, 5, 49, 0, 0, 616, 617, 3, 64, 32, 0, 617, 618, 5, 32,
		0, 0, 618, 619, 5, 55, 0, 0, 619, 620, 5, 47, 0, 0, 620, 621, 5, 55, 0,
		0, 621, 622, 5, 38, 0, 0, 622, 623, 5, 40, 0, 0, 623, 624, 5, 17, 0, 0,
		624, 625, 3, 64, 32, 0, 625, 626, 5, 18, 0, 0, 626, 627, 5, 50, 0, 0, 627,
		629, 1, 0, 0, 0, 628, 578, 1, 0, 0, 0, 628, 590, 1, 0, 0, 0, 628, 598,
		1, 0, 0, 0, 628, 615, 1, 0, 0, 0, 629, 95, 1, 0, 0, 0, 630, 639, 5, 52,
		0, 0, 631, 636, 3, 98, 49, 0, 632, 633, 5, 47, 0, 0, 633, 635, 3, 98, 49,
		0, 634, 632, 1, 0, 0, 0, 635, 638, 1, 0, 0, 0, 636, 634, 1, 0, 0, 0, 636,
		637, 1, 0, 0, 0, 637, 640, 1, 0, 0, 0, 638, 636, 1, 0, 0, 0, 639, 631,
		1, 0, 0, 0, 639, 640, 1, 0, 0, 0, 640, 641, 1, 0, 0, 0, 641, 642, 5, 53,
		0, 0, 642, 97, 1, 0, 0, 0, 643, 644, 7, 7, 0, 0, 644, 645, 5, 51, 0, 0,
		645, 646, 3, 64, 32, 0, 646, 99, 1, 0, 0, 0, 75, 103, 126, 130, 145, 150,
		165, 169, 173, 179, 189, 195, 199, 203, 206, 211, 214, 219, 226, 233, 239,
		245, 250, 254, 258, 284, 295, 301, 312, 320, 323, 329, 333, 341, 348, 354,
		357, 363, 368, 376, 379, 386, 394, 397, 408, 410, 415, 424, 427, 431, 440,
		448, 456, 464, 472, 480, 486, 497, 500, 502, 519, 526, 531, 537, 542, 547,
		550, 560, 568, 575, 584, 587, 610, 628, 636, 639,
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
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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
	p.EnterRule(localctx, 36, FigParserRULE_forInStmt)
	var _la int

	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
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
		{
			p.SetState(286)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.Block()
		}

	case 3:
		localctx = NewForInContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(289)
			p.Match(FigParserTK_FOR)
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
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Expr()
		}
		{
			p.SetState(293)
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
		p.SetState(297)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275386155985432576) != 0 {
		{
			p.SetState(298)
			p.Statements()
		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(304)
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
		p.SetState(306)
		p.Match(FigParserTK_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_FN || _la == FigParserID {
		{
			p.SetState(309)
			p.StructMember()
		}

		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(315)
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

	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		localctx = NewStructFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(318)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(319)
				p.Expr()
			}

		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserSEMICOLON {
			{
				p.SetState(322)
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
			p.SetState(325)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(328)
				p.FnParams()
			}

		}
		{
			p.SetState(331)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
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
		p.SetState(335)
		p.Match(FigParserTK_ENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(336)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserID {
		{
			p.SetState(338)
			p.EnumMember()
		}

		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(344)
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
		p.SetState(346)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(347)
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
		p.SetState(350)
		p.Match(FigParserTK_LET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.BindingTarget()
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserASSIGN {
		{
			p.SetState(352)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Expr()
		}

	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(356)
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
		p.SetState(359)
		p.BindingTarget()
	}
	{
		p.SetState(360)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(361)
		p.Expr()
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(362)
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
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(365)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FigParserLBRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(366)
			p.ArrayPattern()
		}

	case FigParserLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(367)
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
		p.SetState(370)
		p.Match(FigParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&41095346599755776) != 0 {
		{
			p.SetState(371)
			p.BindingElement()
		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(372)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
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
		}

	}
	{
		p.SetState(381)
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
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(383)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FigParserLBRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.ArrayPattern()
		}

	case FigParserLBRACE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(385)
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
		p.SetState(388)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID {
		{
			p.SetState(389)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(390)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
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
		}

	}
	{
		p.SetState(399)
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
		p.SetState(401)
		p.Expr()
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FigParserDOT || _la == FigParserLBRACKET {
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case FigParserLBRACKET:
			{
				p.SetState(402)
				p.Match(FigParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(403)
				p.Expr()
			}
			{
				p.SetState(404)
				p.Match(FigParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case FigParserDOT:
			{
				p.SetState(406)
				p.Match(FigParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(407)
				p.MemberName()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(412)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(413)
		p.Expr()
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(414)
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
		p.SetState(417)
		p.Match(FigParserTK_PRINT)
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
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
		{
			p.SetState(419)
			p.Expr()
		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(420)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
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
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(430)
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
		p.SetState(433)
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
		p.SetState(435)
		p.LogicalAnd()
	}
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserOR {
		{
			p.SetState(436)
			p.Match(FigParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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
		p.SetState(443)
		p.Equality()
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserAND {
		{
			p.SetState(444)
			p.Match(FigParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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
		p.SetState(451)
		p.Comparison()
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserEQ || _la == FigParserNEQ {
		{
			p.SetState(452)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FigParserEQ || _la == FigParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
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
		p.SetState(459)
		p.Term()
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0 {
		{
			p.SetState(460)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
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
		p.SetState(467)
		p.Factor()
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(468)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FigParserPLUS || _la == FigParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(469)
				p.Factor()
			}

		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
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
		p.SetState(475)
		p.Unary()
	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4292608) != 0 {
		{
			p.SetState(476)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4292608) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
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

	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserMINUS, FigParserPLUSPLUS, FigParserMINUSMINUS, FigParserEXCLAM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1077248) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(484)
			p.Unary()
		}

	case FigParserLPAREN, FigParserTK_NULL, FigParserTK_FN, FigParserTK_THIS, FigParserTK_TRY, FigParserTK_MATCH, FigParserLBRACKET, FigParserLBRACE, FigParserBOOL, FigParserID, FigParserNUMBER, FigParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(485)
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
		p.SetState(488)
		p.Primary()
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(500)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case FigParserLBRACKET:
				{
					p.SetState(489)
					p.Match(FigParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(490)
					p.Expr()
				}
				{
					p.SetState(491)
					p.Match(FigParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case FigParserDOT:
				{
					p.SetState(493)
					p.Match(FigParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(494)
					p.MemberName()
				}

			case FigParserLPAREN:
				{
					p.SetState(495)
					p.Match(FigParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(497)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
					{
						p.SetState(496)
						p.FnArgs()
					}

				}
				{
					p.SetState(499)
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
		p.SetState(504)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
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
		p.SetState(505)
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

	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(507)
			p.Match(FigParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(508)
			p.Match(FigParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(509)
			p.Match(FigParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(510)
			p.Match(FigParserTK_NULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(511)
			p.Match(FigParserTK_THIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(512)
			p.ArrayLiteral()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(513)
			p.ObjectLiteral()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(514)
			p.TryExpr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(515)
			p.MatchExpr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(516)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(517)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(518)
				p.FnParams()
			}

		}
		{
			p.SetState(521)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)
			p.Block()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(523)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(524)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
			{
				p.SetState(525)
				p.FnArgs()
			}

		}
		{
			p.SetState(528)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(529)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(531)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(530)
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
			p.SetState(533)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(534)
			p.Expr()
		}
		{
			p.SetState(535)
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
		p.SetState(539)
		p.Match(FigParserTK_TRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(540)
			p.Expr()
		}

	case 2:
		{
			p.SetState(541)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(544)
		p.Match(FigParserTK_ONERROR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserLPAREN {
		{
			p.SetState(545)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(546)
				p.Match(FigParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(549)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(552)
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
		p.SetState(554)
		p.Match(FigParserTK_MATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(555)
		p.Expr()
	}
	{
		p.SetState(556)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0) {
		{
			p.SetState(557)
			p.MatchArm()
		}

		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(562)
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
		p.SetState(564)
		p.MatchPattern()
	}
	{
		p.SetState(565)
		p.Match(FigParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(568)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(566)
			p.Block()
		}

	case 2:
		{
			p.SetState(567)
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
		p.SetState(570)
		p.Expr()
	}
	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(571)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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

func (s *ArrayCompForRangeContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FigParserCOMMA)
}

func (s *ArrayCompForRangeContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FigParserCOMMA, i)
}

func (s *ArrayCompForRangeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FigParserRPAREN, 0)
}

func (s *ArrayCompForRangeContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(FigParserRBRACKET, 0)
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

	p.SetState(628)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrayLiteralSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(578)
			p.Match(FigParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(587)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
			{
				p.SetState(579)
				p.Expr()
			}
			p.SetState(584)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FigParserCOMMA {
				{
					p.SetState(580)
					p.Match(FigParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
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
			}

		}
		{
			p.SetState(589)
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
			p.SetState(590)
			p.Match(FigParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(591)
			p.Expr()
		}
		{
			p.SetState(592)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(593)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(594)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(595)
			p.Expr()
		}
		{
			p.SetState(596)
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
			p.SetState(598)
			p.Match(FigParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(599)
			p.Expr()
		}
		{
			p.SetState(600)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(601)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(602)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(603)
			p.Match(FigParserTK_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(604)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(605)
			p.Expr()
		}
		{
			p.SetState(606)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(607)
			p.Expr()
		}
		p.SetState(610)
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

		}
		{
			p.SetState(612)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(613)
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
			p.SetState(615)
			p.Match(FigParserLBRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(616)
			p.Expr()
		}
		{
			p.SetState(617)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(618)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(619)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(620)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(622)
			p.Match(FigParserTK_ENUMERATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(623)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(624)
			p.Expr()
		}
		{
			p.SetState(625)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(626)
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
		p.SetState(630)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(639)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID || _la == FigParserSTRING {
		{
			p.SetState(631)
			p.ObjectEntry()
		}
		p.SetState(636)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(632)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(633)
				p.ObjectEntry()
			}

			p.SetState(638)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(641)
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
		p.SetState(643)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FigParserID || _la == FigParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(644)
		p.Match(FigParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(645)
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
