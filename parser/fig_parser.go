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
		"varDeclaration", "varAtribuition", "memberAssign", "printStmt", "expr",
		"logicalOr", "logicalAnd", "equality", "comparison", "term", "factor",
		"unary", "postfix", "memberName", "primary", "tryExpr", "matchExpr",
		"matchArm", "matchPattern", "arrayLiteral", "objectLiteral", "objectEntry",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 565, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1, 0, 5, 0, 94, 8,
		0, 10, 0, 12, 0, 97, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 119, 8, 1, 1, 2, 1, 2, 3, 2, 123, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 136, 8, 3, 10, 3, 12, 3,
		139, 9, 3, 1, 3, 1, 3, 3, 3, 143, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 158, 8, 5, 1, 6, 1,
		6, 3, 6, 162, 8, 6, 1, 7, 1, 7, 3, 7, 166, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 172, 8, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 5, 9, 180, 8, 9, 10,
		9, 12, 9, 183, 9, 9, 1, 10, 1, 10, 1, 10, 3, 10, 188, 8, 10, 1, 10, 1,
		10, 3, 10, 192, 8, 10, 1, 11, 1, 11, 3, 11, 196, 8, 11, 1, 11, 3, 11, 199,
		8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 204, 8, 12, 1, 12, 3, 12, 207, 8, 12,
		1, 13, 1, 13, 1, 13, 3, 13, 212, 8, 13, 1, 14, 1, 14, 1, 14, 5, 14, 217,
		8, 14, 10, 14, 12, 14, 220, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 226,
		8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 232, 8, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 238, 8, 16, 1, 17, 1, 17, 1, 17, 3, 17, 243, 8, 17, 1,
		17, 1, 17, 3, 17, 247, 8, 17, 1, 17, 1, 17, 3, 17, 251, 8, 17, 1, 17, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 277, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 3, 18, 288, 8, 18, 1, 19, 1, 19, 5, 19, 292, 8, 19,
		10, 19, 12, 19, 295, 9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5,
		20, 303, 8, 20, 10, 20, 12, 20, 306, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 3, 21, 313, 8, 21, 1, 21, 3, 21, 316, 8, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 3, 21, 322, 8, 21, 1, 21, 1, 21, 3, 21, 326, 8, 21, 1, 22, 1, 22,
		1, 22, 1, 22, 5, 22, 332, 8, 22, 10, 22, 12, 22, 335, 9, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 3, 23, 341, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24,
		347, 8, 24, 1, 24, 3, 24, 350, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25,
		356, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 4, 26, 365,
		8, 26, 11, 26, 12, 26, 366, 1, 26, 1, 26, 1, 26, 3, 26, 372, 8, 26, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 379, 8, 27, 10, 27, 12, 27, 382,
		9, 27, 3, 27, 384, 8, 27, 1, 27, 1, 27, 3, 27, 388, 8, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 5, 29, 395, 8, 29, 10, 29, 12, 29, 398, 9, 29, 1,
		30, 1, 30, 1, 30, 5, 30, 403, 8, 30, 10, 30, 12, 30, 406, 9, 30, 1, 31,
		1, 31, 1, 31, 5, 31, 411, 8, 31, 10, 31, 12, 31, 414, 9, 31, 1, 32, 1,
		32, 1, 32, 5, 32, 419, 8, 32, 10, 32, 12, 32, 422, 9, 32, 1, 33, 1, 33,
		1, 33, 5, 33, 427, 8, 33, 10, 33, 12, 33, 430, 9, 33, 1, 34, 1, 34, 1,
		34, 5, 34, 435, 8, 34, 10, 34, 12, 34, 438, 9, 34, 1, 35, 1, 35, 1, 35,
		3, 35, 443, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 3, 36, 454, 8, 36, 1, 36, 5, 36, 457, 8, 36, 10, 36, 12, 36,
		460, 9, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 476, 8, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 3, 38, 483, 8, 38, 1, 38, 1, 38, 1, 38, 3, 38, 488,
		8, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 494, 8, 38, 1, 39, 1, 39, 1,
		39, 3, 39, 499, 8, 39, 1, 39, 1, 39, 1, 39, 3, 39, 504, 8, 39, 1, 39, 3,
		39, 507, 8, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 4, 40, 515, 8,
		40, 11, 40, 12, 40, 516, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41,
		525, 8, 41, 1, 42, 1, 42, 1, 42, 5, 42, 530, 8, 42, 10, 42, 12, 42, 533,
		9, 42, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 539, 8, 43, 10, 43, 12, 43, 542,
		9, 43, 3, 43, 544, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 5,
		44, 552, 8, 44, 10, 44, 12, 44, 555, 9, 44, 3, 44, 557, 8, 44, 1, 44, 1,
		44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 0, 0, 46, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84,
		86, 88, 90, 0, 8, 1, 0, 5, 6, 1, 0, 1, 4, 1, 0, 11, 12, 2, 0, 15, 16, 22,
		22, 2, 0, 12, 14, 20, 20, 2, 0, 46, 46, 55, 55, 1, 0, 13, 14, 2, 0, 55,
		55, 57, 57, 615, 0, 95, 1, 0, 0, 0, 2, 118, 1, 0, 0, 0, 4, 120, 1, 0, 0,
		0, 6, 124, 1, 0, 0, 0, 8, 144, 1, 0, 0, 0, 10, 150, 1, 0, 0, 0, 12, 159,
		1, 0, 0, 0, 14, 163, 1, 0, 0, 0, 16, 167, 1, 0, 0, 0, 18, 176, 1, 0, 0,
		0, 20, 191, 1, 0, 0, 0, 22, 193, 1, 0, 0, 0, 24, 200, 1, 0, 0, 0, 26, 208,
		1, 0, 0, 0, 28, 213, 1, 0, 0, 0, 30, 231, 1, 0, 0, 0, 32, 237, 1, 0, 0,
		0, 34, 239, 1, 0, 0, 0, 36, 287, 1, 0, 0, 0, 38, 289, 1, 0, 0, 0, 40, 298,
		1, 0, 0, 0, 42, 325, 1, 0, 0, 0, 44, 327, 1, 0, 0, 0, 46, 338, 1, 0, 0,
		0, 48, 342, 1, 0, 0, 0, 50, 351, 1, 0, 0, 0, 52, 357, 1, 0, 0, 0, 54, 373,
		1, 0, 0, 0, 56, 389, 1, 0, 0, 0, 58, 391, 1, 0, 0, 0, 60, 399, 1, 0, 0,
		0, 62, 407, 1, 0, 0, 0, 64, 415, 1, 0, 0, 0, 66, 423, 1, 0, 0, 0, 68, 431,
		1, 0, 0, 0, 70, 442, 1, 0, 0, 0, 72, 444, 1, 0, 0, 0, 74, 461, 1, 0, 0,
		0, 76, 493, 1, 0, 0, 0, 78, 495, 1, 0, 0, 0, 80, 510, 1, 0, 0, 0, 82, 520,
		1, 0, 0, 0, 84, 526, 1, 0, 0, 0, 86, 534, 1, 0, 0, 0, 88, 547, 1, 0, 0,
		0, 90, 560, 1, 0, 0, 0, 92, 94, 3, 2, 1, 0, 93, 92, 1, 0, 0, 0, 94, 97,
		1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0,
		97, 95, 1, 0, 0, 0, 98, 99, 5, 0, 0, 1, 99, 1, 1, 0, 0, 0, 100, 119, 3,
		48, 24, 0, 101, 119, 3, 50, 25, 0, 102, 119, 3, 52, 26, 0, 103, 119, 3,
		54, 27, 0, 104, 119, 3, 6, 3, 0, 105, 119, 3, 8, 4, 0, 106, 119, 3, 10,
		5, 0, 107, 119, 3, 34, 17, 0, 108, 119, 3, 36, 18, 0, 109, 119, 3, 12,
		6, 0, 110, 119, 3, 14, 7, 0, 111, 119, 3, 16, 8, 0, 112, 119, 3, 22, 11,
		0, 113, 119, 3, 24, 12, 0, 114, 119, 3, 26, 13, 0, 115, 119, 3, 40, 20,
		0, 116, 119, 3, 44, 22, 0, 117, 119, 3, 4, 2, 0, 118, 100, 1, 0, 0, 0,
		118, 101, 1, 0, 0, 0, 118, 102, 1, 0, 0, 0, 118, 103, 1, 0, 0, 0, 118,
		104, 1, 0, 0, 0, 118, 105, 1, 0, 0, 0, 118, 106, 1, 0, 0, 0, 118, 107,
		1, 0, 0, 0, 118, 108, 1, 0, 0, 0, 118, 109, 1, 0, 0, 0, 118, 110, 1, 0,
		0, 0, 118, 111, 1, 0, 0, 0, 118, 112, 1, 0, 0, 0, 118, 113, 1, 0, 0, 0,
		118, 114, 1, 0, 0, 0, 118, 115, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118,
		117, 1, 0, 0, 0, 119, 3, 1, 0, 0, 0, 120, 122, 3, 56, 28, 0, 121, 123,
		5, 21, 0, 0, 122, 121, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 5, 1, 0,
		0, 0, 124, 125, 5, 25, 0, 0, 125, 126, 5, 17, 0, 0, 126, 127, 3, 56, 28,
		0, 127, 128, 5, 18, 0, 0, 128, 137, 3, 38, 19, 0, 129, 130, 5, 26, 0, 0,
		130, 131, 5, 17, 0, 0, 131, 132, 3, 56, 28, 0, 132, 133, 5, 18, 0, 0, 133,
		134, 3, 38, 19, 0, 134, 136, 1, 0, 0, 0, 135, 129, 1, 0, 0, 0, 136, 139,
		1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 142, 1, 0,
		0, 0, 139, 137, 1, 0, 0, 0, 140, 141, 5, 27, 0, 0, 141, 143, 3, 38, 19,
		0, 142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 7, 1, 0, 0, 0, 144,
		145, 5, 28, 0, 0, 145, 146, 5, 17, 0, 0, 146, 147, 3, 56, 28, 0, 147, 148,
		5, 18, 0, 0, 148, 149, 3, 38, 19, 0, 149, 9, 1, 0, 0, 0, 150, 151, 5, 29,
		0, 0, 151, 152, 3, 38, 19, 0, 152, 153, 5, 28, 0, 0, 153, 154, 5, 17, 0,
		0, 154, 155, 3, 56, 28, 0, 155, 157, 5, 18, 0, 0, 156, 158, 5, 21, 0, 0,
		157, 156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 11, 1, 0, 0, 0, 159, 161,
		5, 30, 0, 0, 160, 162, 5, 21, 0, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1,
		0, 0, 0, 162, 13, 1, 0, 0, 0, 163, 165, 5, 31, 0, 0, 164, 166, 5, 21, 0,
		0, 165, 164, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 15, 1, 0, 0, 0, 167,
		168, 5, 34, 0, 0, 168, 169, 5, 55, 0, 0, 169, 171, 5, 17, 0, 0, 170, 172,
		3, 18, 9, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 1, 0,
		0, 0, 173, 174, 5, 18, 0, 0, 174, 175, 3, 38, 19, 0, 175, 17, 1, 0, 0,
		0, 176, 181, 3, 20, 10, 0, 177, 178, 5, 47, 0, 0, 178, 180, 3, 20, 10,
		0, 179, 177, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181,
		182, 1, 0, 0, 0, 182, 19, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 187, 5,
		55, 0, 0, 185, 186, 5, 10, 0, 0, 186, 188, 3, 56, 28, 0, 187, 185, 1, 0,
		0, 0, 187, 188, 1, 0, 0, 0, 188, 192, 1, 0, 0, 0, 189, 190, 5, 55, 0, 0,
		190, 192, 5, 19, 0, 0, 191, 184, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192,
		21, 1, 0, 0, 0, 193, 195, 5, 35, 0, 0, 194, 196, 3, 56, 28, 0, 195, 194,
		1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 198, 1, 0, 0, 0, 197, 199, 5, 21,
		0, 0, 198, 197, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 23, 1, 0, 0, 0,
		200, 201, 5, 36, 0, 0, 201, 203, 5, 57, 0, 0, 202, 204, 5, 55, 0, 0, 203,
		202, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 207,
		5, 21, 0, 0, 206, 205, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 25, 1, 0,
		0, 0, 208, 209, 5, 37, 0, 0, 209, 211, 5, 57, 0, 0, 210, 212, 5, 21, 0,
		0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 27, 1, 0, 0, 0, 213,
		218, 3, 56, 28, 0, 214, 215, 5, 47, 0, 0, 215, 217, 3, 56, 28, 0, 216,
		214, 1, 0, 0, 0, 217, 220, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219,
		1, 0, 0, 0, 219, 29, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 221, 222, 5, 23,
		0, 0, 222, 225, 5, 55, 0, 0, 223, 224, 5, 10, 0, 0, 224, 226, 3, 56, 28,
		0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 232, 1, 0, 0, 0, 227,
		228, 5, 55, 0, 0, 228, 229, 5, 10, 0, 0, 229, 232, 3, 56, 28, 0, 230, 232,
		3, 56, 28, 0, 231, 221, 1, 0, 0, 0, 231, 227, 1, 0, 0, 0, 231, 230, 1,
		0, 0, 0, 232, 31, 1, 0, 0, 0, 233, 234, 5, 55, 0, 0, 234, 235, 5, 10, 0,
		0, 235, 238, 3, 56, 28, 0, 236, 238, 3, 56, 28, 0, 237, 233, 1, 0, 0, 0,
		237, 236, 1, 0, 0, 0, 238, 33, 1, 0, 0, 0, 239, 240, 5, 32, 0, 0, 240,
		242, 5, 17, 0, 0, 241, 243, 3, 30, 15, 0, 242, 241, 1, 0, 0, 0, 242, 243,
		1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 246, 5, 21, 0, 0, 245, 247, 3, 56,
		28, 0, 246, 245, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0,
		248, 250, 5, 21, 0, 0, 249, 251, 3, 32, 16, 0, 250, 249, 1, 0, 0, 0, 250,
		251, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 253, 5, 18, 0, 0, 253, 254,
		3, 38, 19, 0, 254, 35, 1, 0, 0, 0, 255, 256, 5, 32, 0, 0, 256, 257, 5,
		55, 0, 0, 257, 258, 5, 47, 0, 0, 258, 259, 5, 55, 0, 0, 259, 260, 5, 38,
		0, 0, 260, 261, 5, 40, 0, 0, 261, 262, 5, 17, 0, 0, 262, 263, 3, 56, 28,
		0, 263, 264, 5, 18, 0, 0, 264, 265, 3, 38, 19, 0, 265, 288, 1, 0, 0, 0,
		266, 267, 5, 32, 0, 0, 267, 268, 5, 55, 0, 0, 268, 269, 5, 38, 0, 0, 269,
		270, 5, 39, 0, 0, 270, 271, 5, 17, 0, 0, 271, 272, 3, 56, 28, 0, 272, 273,
		5, 47, 0, 0, 273, 276, 3, 56, 28, 0, 274, 275, 5, 47, 0, 0, 275, 277, 3,
		56, 28, 0, 276, 274, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 1, 0,
		0, 0, 278, 279, 5, 18, 0, 0, 279, 280, 3, 38, 19, 0, 280, 288, 1, 0, 0,
		0, 281, 282, 5, 32, 0, 0, 282, 283, 5, 55, 0, 0, 283, 284, 5, 38, 0, 0,
		284, 285, 3, 56, 28, 0, 285, 286, 3, 38, 19, 0, 286, 288, 1, 0, 0, 0, 287,
		255, 1, 0, 0, 0, 287, 266, 1, 0, 0, 0, 287, 281, 1, 0, 0, 0, 288, 37, 1,
		0, 0, 0, 289, 293, 5, 52, 0, 0, 290, 292, 3, 2, 1, 0, 291, 290, 1, 0, 0,
		0, 292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294,
		296, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 297, 5, 53, 0, 0, 297, 39,
		1, 0, 0, 0, 298, 299, 5, 41, 0, 0, 299, 300, 5, 55, 0, 0, 300, 304, 5,
		52, 0, 0, 301, 303, 3, 42, 21, 0, 302, 301, 1, 0, 0, 0, 303, 306, 1, 0,
		0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 307, 1, 0, 0, 0,
		306, 304, 1, 0, 0, 0, 307, 308, 5, 53, 0, 0, 308, 41, 1, 0, 0, 0, 309,
		312, 5, 55, 0, 0, 310, 311, 5, 10, 0, 0, 311, 313, 3, 56, 28, 0, 312, 310,
		1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 315, 1, 0, 0, 0, 314, 316, 5, 21,
		0, 0, 315, 314, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 326, 1, 0, 0, 0,
		317, 318, 5, 34, 0, 0, 318, 319, 5, 55, 0, 0, 319, 321, 5, 17, 0, 0, 320,
		322, 3, 18, 9, 0, 321, 320, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322, 323,
		1, 0, 0, 0, 323, 324, 5, 18, 0, 0, 324, 326, 3, 38, 19, 0, 325, 309, 1,
		0, 0, 0, 325, 317, 1, 0, 0, 0, 326, 43, 1, 0, 0, 0, 327, 328, 5, 42, 0,
		0, 328, 329, 5, 55, 0, 0, 329, 333, 5, 52, 0, 0, 330, 332, 3, 46, 23, 0,
		331, 330, 1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333,
		334, 1, 0, 0, 0, 334, 336, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 337,
		5, 53, 0, 0, 337, 45, 1, 0, 0, 0, 338, 340, 5, 55, 0, 0, 339, 341, 5, 21,
		0, 0, 340, 339, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 47, 1, 0, 0, 0,
		342, 343, 5, 23, 0, 0, 343, 346, 5, 55, 0, 0, 344, 345, 5, 10, 0, 0, 345,
		347, 3, 56, 28, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 349,
		1, 0, 0, 0, 348, 350, 5, 21, 0, 0, 349, 348, 1, 0, 0, 0, 349, 350, 1, 0,
		0, 0, 350, 49, 1, 0, 0, 0, 351, 352, 5, 55, 0, 0, 352, 353, 5, 10, 0, 0,
		353, 355, 3, 56, 28, 0, 354, 356, 5, 21, 0, 0, 355, 354, 1, 0, 0, 0, 355,
		356, 1, 0, 0, 0, 356, 51, 1, 0, 0, 0, 357, 364, 3, 56, 28, 0, 358, 359,
		5, 49, 0, 0, 359, 360, 3, 56, 28, 0, 360, 361, 5, 50, 0, 0, 361, 365, 1,
		0, 0, 0, 362, 363, 5, 9, 0, 0, 363, 365, 3, 74, 37, 0, 364, 358, 1, 0,
		0, 0, 364, 362, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0,
		366, 367, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 369, 5, 10, 0, 0, 369,
		371, 3, 56, 28, 0, 370, 372, 5, 21, 0, 0, 371, 370, 1, 0, 0, 0, 371, 372,
		1, 0, 0, 0, 372, 53, 1, 0, 0, 0, 373, 374, 5, 24, 0, 0, 374, 383, 5, 17,
		0, 0, 375, 380, 3, 56, 28, 0, 376, 377, 5, 47, 0, 0, 377, 379, 3, 56, 28,
		0, 378, 376, 1, 0, 0, 0, 379, 382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380,
		381, 1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 375,
		1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 387, 5, 18,
		0, 0, 386, 388, 5, 21, 0, 0, 387, 386, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0,
		388, 55, 1, 0, 0, 0, 389, 390, 3, 58, 29, 0, 390, 57, 1, 0, 0, 0, 391,
		396, 3, 60, 30, 0, 392, 393, 5, 8, 0, 0, 393, 395, 3, 60, 30, 0, 394, 392,
		1, 0, 0, 0, 395, 398, 1, 0, 0, 0, 396, 394, 1, 0, 0, 0, 396, 397, 1, 0,
		0, 0, 397, 59, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 399, 404, 3, 62, 31, 0,
		400, 401, 5, 7, 0, 0, 401, 403, 3, 62, 31, 0, 402, 400, 1, 0, 0, 0, 403,
		406, 1, 0, 0, 0, 404, 402, 1, 0, 0, 0, 404, 405, 1, 0, 0, 0, 405, 61, 1,
		0, 0, 0, 406, 404, 1, 0, 0, 0, 407, 412, 3, 64, 32, 0, 408, 409, 7, 0,
		0, 0, 409, 411, 3, 64, 32, 0, 410, 408, 1, 0, 0, 0, 411, 414, 1, 0, 0,
		0, 412, 410, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 63, 1, 0, 0, 0, 414,
		412, 1, 0, 0, 0, 415, 420, 3, 66, 33, 0, 416, 417, 7, 1, 0, 0, 417, 419,
		3, 66, 33, 0, 418, 416, 1, 0, 0, 0, 419, 422, 1, 0, 0, 0, 420, 418, 1,
		0, 0, 0, 420, 421, 1, 0, 0, 0, 421, 65, 1, 0, 0, 0, 422, 420, 1, 0, 0,
		0, 423, 428, 3, 68, 34, 0, 424, 425, 7, 2, 0, 0, 425, 427, 3, 68, 34, 0,
		426, 424, 1, 0, 0, 0, 427, 430, 1, 0, 0, 0, 428, 426, 1, 0, 0, 0, 428,
		429, 1, 0, 0, 0, 429, 67, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 431, 436, 3,
		70, 35, 0, 432, 433, 7, 3, 0, 0, 433, 435, 3, 70, 35, 0, 434, 432, 1, 0,
		0, 0, 435, 438, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 436, 437, 1, 0, 0, 0,
		437, 69, 1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 439, 440, 7, 4, 0, 0, 440, 443,
		3, 70, 35, 0, 441, 443, 3, 72, 36, 0, 442, 439, 1, 0, 0, 0, 442, 441, 1,
		0, 0, 0, 443, 71, 1, 0, 0, 0, 444, 458, 3, 76, 38, 0, 445, 446, 5, 49,
		0, 0, 446, 447, 3, 56, 28, 0, 447, 448, 5, 50, 0, 0, 448, 457, 1, 0, 0,
		0, 449, 450, 5, 9, 0, 0, 450, 457, 3, 74, 37, 0, 451, 453, 5, 17, 0, 0,
		452, 454, 3, 28, 14, 0, 453, 452, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454,
		455, 1, 0, 0, 0, 455, 457, 5, 18, 0, 0, 456, 445, 1, 0, 0, 0, 456, 449,
		1, 0, 0, 0, 456, 451, 1, 0, 0, 0, 457, 460, 1, 0, 0, 0, 458, 456, 1, 0,
		0, 0, 458, 459, 1, 0, 0, 0, 459, 73, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0,
		461, 462, 7, 5, 0, 0, 462, 75, 1, 0, 0, 0, 463, 494, 5, 56, 0, 0, 464,
		494, 5, 54, 0, 0, 465, 494, 5, 57, 0, 0, 466, 494, 5, 33, 0, 0, 467, 494,
		5, 43, 0, 0, 468, 494, 3, 86, 43, 0, 469, 494, 3, 88, 44, 0, 470, 494,
		3, 78, 39, 0, 471, 494, 3, 80, 40, 0, 472, 473, 5, 34, 0, 0, 473, 475,
		5, 17, 0, 0, 474, 476, 3, 18, 9, 0, 475, 474, 1, 0, 0, 0, 475, 476, 1,
		0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 478, 5, 18, 0, 0, 478, 494, 3, 38,
		19, 0, 479, 480, 5, 55, 0, 0, 480, 482, 5, 17, 0, 0, 481, 483, 3, 28, 14,
		0, 482, 481, 1, 0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 484,
		494, 5, 18, 0, 0, 485, 487, 5, 55, 0, 0, 486, 488, 7, 6, 0, 0, 487, 486,
		1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488, 494, 1, 0, 0, 0, 489, 490, 5, 17,
		0, 0, 490, 491, 3, 56, 28, 0, 491, 492, 5, 18, 0, 0, 492, 494, 1, 0, 0,
		0, 493, 463, 1, 0, 0, 0, 493, 464, 1, 0, 0, 0, 493, 465, 1, 0, 0, 0, 493,
		466, 1, 0, 0, 0, 493, 467, 1, 0, 0, 0, 493, 468, 1, 0, 0, 0, 493, 469,
		1, 0, 0, 0, 493, 470, 1, 0, 0, 0, 493, 471, 1, 0, 0, 0, 493, 472, 1, 0,
		0, 0, 493, 479, 1, 0, 0, 0, 493, 485, 1, 0, 0, 0, 493, 489, 1, 0, 0, 0,
		494, 77, 1, 0, 0, 0, 495, 498, 5, 44, 0, 0, 496, 499, 3, 56, 28, 0, 497,
		499, 3, 38, 19, 0, 498, 496, 1, 0, 0, 0, 498, 497, 1, 0, 0, 0, 499, 500,
		1, 0, 0, 0, 500, 506, 5, 45, 0, 0, 501, 503, 5, 17, 0, 0, 502, 504, 5,
		55, 0, 0, 503, 502, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 505, 1, 0, 0,
		0, 505, 507, 5, 18, 0, 0, 506, 501, 1, 0, 0, 0, 506, 507, 1, 0, 0, 0, 507,
		508, 1, 0, 0, 0, 508, 509, 3, 38, 19, 0, 509, 79, 1, 0, 0, 0, 510, 511,
		5, 46, 0, 0, 511, 512, 3, 56, 28, 0, 512, 514, 5, 52, 0, 0, 513, 515, 3,
		82, 41, 0, 514, 513, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 514, 1, 0,
		0, 0, 516, 517, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 519, 5, 53, 0, 0,
		519, 81, 1, 0, 0, 0, 520, 521, 3, 84, 42, 0, 521, 524, 5, 48, 0, 0, 522,
		525, 3, 38, 19, 0, 523, 525, 3, 56, 28, 0, 524, 522, 1, 0, 0, 0, 524, 523,
		1, 0, 0, 0, 525, 83, 1, 0, 0, 0, 526, 531, 3, 56, 28, 0, 527, 528, 5, 47,
		0, 0, 528, 530, 3, 56, 28, 0, 529, 527, 1, 0, 0, 0, 530, 533, 1, 0, 0,
		0, 531, 529, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 532, 85, 1, 0, 0, 0, 533,
		531, 1, 0, 0, 0, 534, 543, 5, 49, 0, 0, 535, 540, 3, 56, 28, 0, 536, 537,
		5, 47, 0, 0, 537, 539, 3, 56, 28, 0, 538, 536, 1, 0, 0, 0, 539, 542, 1,
		0, 0, 0, 540, 538, 1, 0, 0, 0, 540, 541, 1, 0, 0, 0, 541, 544, 1, 0, 0,
		0, 542, 540, 1, 0, 0, 0, 543, 535, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544,
		545, 1, 0, 0, 0, 545, 546, 5, 50, 0, 0, 546, 87, 1, 0, 0, 0, 547, 556,
		5, 52, 0, 0, 548, 553, 3, 90, 45, 0, 549, 550, 5, 47, 0, 0, 550, 552, 3,
		90, 45, 0, 551, 549, 1, 0, 0, 0, 552, 555, 1, 0, 0, 0, 553, 551, 1, 0,
		0, 0, 553, 554, 1, 0, 0, 0, 554, 557, 1, 0, 0, 0, 555, 553, 1, 0, 0, 0,
		556, 548, 1, 0, 0, 0, 556, 557, 1, 0, 0, 0, 557, 558, 1, 0, 0, 0, 558,
		559, 5, 53, 0, 0, 559, 89, 1, 0, 0, 0, 560, 561, 7, 7, 0, 0, 561, 562,
		5, 51, 0, 0, 562, 563, 3, 56, 28, 0, 563, 91, 1, 0, 0, 0, 67, 95, 118,
		122, 137, 142, 157, 161, 165, 171, 181, 187, 191, 195, 198, 203, 206, 211,
		218, 225, 231, 237, 242, 246, 250, 276, 287, 293, 304, 312, 315, 321, 325,
		333, 340, 346, 349, 355, 364, 366, 371, 380, 383, 387, 396, 404, 412, 420,
		428, 436, 442, 453, 456, 458, 475, 482, 487, 493, 498, 503, 506, 516, 524,
		531, 540, 543, 553, 556,
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
	FigParserRULE_memberAssign   = 26
	FigParserRULE_printStmt      = 27
	FigParserRULE_expr           = 28
	FigParserRULE_logicalOr      = 29
	FigParserRULE_logicalAnd     = 30
	FigParserRULE_equality       = 31
	FigParserRULE_comparison     = 32
	FigParserRULE_term           = 33
	FigParserRULE_factor         = 34
	FigParserRULE_unary          = 35
	FigParserRULE_postfix        = 36
	FigParserRULE_memberName     = 37
	FigParserRULE_primary        = 38
	FigParserRULE_tryExpr        = 39
	FigParserRULE_matchExpr      = 40
	FigParserRULE_matchArm       = 41
	FigParserRULE_matchPattern   = 42
	FigParserRULE_arrayLiteral   = 43
	FigParserRULE_objectLiteral  = 44
	FigParserRULE_objectEntry    = 45
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275386155985432576) != 0 {
		{
			p.SetState(92)
			p.Statements()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.VarDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.VarAtribuition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.MemberAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.PrintStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(104)
			p.IfStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(105)
			p.WhileStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(106)
			p.DoWhileStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(107)
			p.ForStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(108)
			p.ForInStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(109)
			p.BreakStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(110)
			p.ContinueStmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(111)
			p.FnDecl()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(112)
			p.ReturnStmt()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(113)
			p.ImportStmt()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(114)
			p.UseStmt()
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(115)
			p.StructDecl()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(116)
			p.EnumDecl()
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(117)
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
		p.SetState(120)
		p.Expr()
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(121)
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
		p.SetState(124)
		p.Match(FigParserTK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.Expr()
	}
	{
		p.SetState(127)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Block()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_ELIF {
		{
			p.SetState(129)
			p.Match(FigParserTK_ELIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Expr()
		}
		{
			p.SetState(132)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Block()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserTK_ELSE {
		{
			p.SetState(140)
			p.Match(FigParserTK_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
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
		p.SetState(144)
		p.Match(FigParserTK_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Expr()
	}
	{
		p.SetState(147)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
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
		p.SetState(150)
		p.Match(FigParserTK_DO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Block()
	}
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
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(156)
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
		p.SetState(159)
		p.Match(FigParserTK_BREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(160)
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
		p.SetState(163)
		p.Match(FigParserTK_CONTINUE)
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
		p.SetState(167)
		p.Match(FigParserTK_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID {
		{
			p.SetState(170)
			p.FnParams()
		}

	}
	{
		p.SetState(173)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
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
		p.SetState(176)
		p.ParamDecl()
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(177)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.ParamDecl()
		}

		p.SetState(183)
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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParamWithDefaultOrRequiredContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(185)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(186)
				p.Expr()
			}

		}

	case 2:
		localctx = NewParamOptionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
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
		p.SetState(193)
		p.Match(FigParserTK_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(194)
			p.Expr()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(197)
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
		p.SetState(200)
		p.Match(FigParserTK_IMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(FigParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(202)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
		p.SetState(208)
		p.Match(FigParserTK_USE)
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
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(210)
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
		p.SetState(213)
		p.Expr()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(214)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Expr()
		}

		p.SetState(220)
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

	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Match(FigParserTK_LET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserASSIGN {
			{
				p.SetState(223)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(224)
				p.Expr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
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
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
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
		p.SetState(239)
		p.Match(FigParserTK_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310025666560) != 0 {
		{
			p.SetState(241)
			p.ForInit()
		}

	}
	{
		p.SetState(244)
		p.Match(FigParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
		{
			p.SetState(245)
			p.Expr()
		}

	}
	{
		p.SetState(248)
		p.Match(FigParserSEMICOLON)
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

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
		{
			p.SetState(249)
			p.ForStep()
		}

	}
	{
		p.SetState(252)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
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

	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForEnumerateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(255)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Match(FigParserTK_ENUMERATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Expr()
		}
		{
			p.SetState(263)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(264)
			p.Block()
		}

	case 2:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(FigParserTK_RANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Expr()
		}
		{
			p.SetState(272)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Expr()
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserCOMMA {
			{
				p.SetState(274)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(275)
				p.Expr()
			}

		}
		{
			p.SetState(278)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Block()
		}

	case 3:
		localctx = NewForInContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(281)
			p.Match(FigParserTK_FOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Match(FigParserTK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Expr()
		}
		{
			p.SetState(285)
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
		p.SetState(289)
		p.Match(FigParserLBRACE)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275386155985432576) != 0 {
		{
			p.SetState(290)
			p.Statements()
		}

		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(296)
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
		p.SetState(298)
		p.Match(FigParserTK_STRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(299)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserTK_FN || _la == FigParserID {
		{
			p.SetState(301)
			p.StructMember()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(307)
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

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserID:
		localctx = NewStructFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Match(FigParserID)
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

		if _la == FigParserASSIGN {
			{
				p.SetState(310)
				p.Match(FigParserASSIGN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(311)
				p.Expr()
			}

		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserSEMICOLON {
			{
				p.SetState(314)
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
			p.SetState(317)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(318)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(320)
				p.FnParams()
			}

		}
		{
			p.SetState(323)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
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
		p.SetState(327)
		p.Match(FigParserTK_ENUM)
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
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserID {
		{
			p.SetState(330)
			p.EnumMember()
		}

		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(336)
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
		p.SetState(338)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
	p.EnterRule(localctx, 48, FigParserRULE_varDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(FigParserTK_LET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserASSIGN {
		{
			p.SetState(344)
			p.Match(FigParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.Expr()
		}

	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(348)
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
	p.EnterRule(localctx, 50, FigParserRULE_varAtribuition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(FigParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
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
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(354)
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
	p.EnterRule(localctx, 52, FigParserRULE_memberAssign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Expr()
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FigParserDOT || _la == FigParserLBRACKET {
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case FigParserLBRACKET:
			{
				p.SetState(358)
				p.Match(FigParserLBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(359)
				p.Expr()
			}
			{
				p.SetState(360)
				p.Match(FigParserRBRACKET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case FigParserDOT:
			{
				p.SetState(362)
				p.Match(FigParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(363)
				p.MemberName()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(368)
		p.Match(FigParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Expr()
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(370)
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
	p.EnterRule(localctx, 54, FigParserRULE_printStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(FigParserTK_PRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.Match(FigParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
		{
			p.SetState(375)
			p.Expr()
		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(376)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(377)
				p.Expr()
			}

			p.SetState(382)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(385)
		p.Match(FigParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserSEMICOLON {
		{
			p.SetState(386)
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
	p.EnterRule(localctx, 56, FigParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
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
	p.EnterRule(localctx, 58, FigParserRULE_logicalOr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.LogicalAnd()
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserOR {
		{
			p.SetState(392)
			p.Match(FigParserOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.LogicalAnd()
		}

		p.SetState(398)
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
	p.EnterRule(localctx, 60, FigParserRULE_logicalAnd)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Equality()
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserAND {
		{
			p.SetState(400)
			p.Match(FigParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(401)
			p.Equality()
		}

		p.SetState(406)
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
	p.EnterRule(localctx, 62, FigParserRULE_equality)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.Comparison()
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserEQ || _la == FigParserNEQ {
		{
			p.SetState(408)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FigParserEQ || _la == FigParserNEQ) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(409)
			p.Comparison()
		}

		p.SetState(414)
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
	p.EnterRule(localctx, 64, FigParserRULE_comparison)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(415)
		p.Term()
	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0 {
		{
			p.SetState(416)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(417)
			p.Term()
		}

		p.SetState(422)
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
	p.EnterRule(localctx, 66, FigParserRULE_term)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.Factor()
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(424)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FigParserPLUS || _la == FigParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(425)
				p.Factor()
			}

		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 68, FigParserRULE_factor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.Unary()
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4292608) != 0 {
		{
			p.SetState(432)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4292608) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(433)
			p.Unary()
		}

		p.SetState(438)
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
	p.EnterRule(localctx, 70, FigParserRULE_unary)
	var _la int

	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FigParserMINUS, FigParserPLUSPLUS, FigParserMINUSMINUS, FigParserEXCLAM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(439)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1077248) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(440)
			p.Unary()
		}

	case FigParserLPAREN, FigParserTK_NULL, FigParserTK_FN, FigParserTK_THIS, FigParserTK_TRY, FigParserTK_MATCH, FigParserLBRACKET, FigParserLBRACE, FigParserBOOL, FigParserID, FigParserNUMBER, FigParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(441)
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
	p.EnterRule(localctx, 72, FigParserRULE_postfix)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Primary()
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(456)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case FigParserLBRACKET:
				{
					p.SetState(445)
					p.Match(FigParserLBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(446)
					p.Expr()
				}
				{
					p.SetState(447)
					p.Match(FigParserRBRACKET)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case FigParserDOT:
				{
					p.SetState(449)
					p.Match(FigParserDOT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(450)
					p.MemberName()
				}

			case FigParserLPAREN:
				{
					p.SetState(451)
					p.Match(FigParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(453)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
					{
						p.SetState(452)
						p.FnArgs()
					}

				}
				{
					p.SetState(455)
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
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 74, FigParserRULE_memberName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(461)
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
	p.EnterRule(localctx, 76, FigParserRULE_primary)
	var _la int

	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.Match(FigParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(464)
			p.Match(FigParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(465)
			p.Match(FigParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(466)
			p.Match(FigParserTK_NULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(467)
			p.Match(FigParserTK_THIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(468)
			p.ArrayLiteral()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(469)
			p.ObjectLiteral()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(470)
			p.TryExpr()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(471)
			p.MatchExpr()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(472)
			p.Match(FigParserTK_FN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(474)
				p.FnParams()
			}

		}
		{
			p.SetState(477)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)
			p.Block()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(479)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
			{
				p.SetState(481)
				p.FnArgs()
			}

		}
		{
			p.SetState(484)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(485)
			p.Match(FigParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(486)
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
			p.SetState(489)
			p.Match(FigParserLPAREN)
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
	p.EnterRule(localctx, 78, FigParserRULE_tryExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		p.Match(FigParserTK_TRY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(496)
			p.Expr()
		}

	case 2:
		{
			p.SetState(497)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(500)
		p.Match(FigParserTK_ONERROR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserLPAREN {
		{
			p.SetState(501)
			p.Match(FigParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FigParserID {
			{
				p.SetState(502)
				p.Match(FigParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(505)
			p.Match(FigParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(508)
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
	p.EnterRule(localctx, 80, FigParserRULE_matchExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.Match(FigParserTK_MATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(511)
		p.Expr()
	}
	{
		p.SetState(512)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0) {
		{
			p.SetState(513)
			p.MatchArm()
		}

		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(518)
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
	p.EnterRule(localctx, 82, FigParserRULE_matchArm)
	localctx = NewMatchArmCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(520)
		p.MatchPattern()
	}
	{
		p.SetState(521)
		p.Match(FigParserARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(522)
			p.Block()
		}

	case 2:
		{
			p.SetState(523)
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
	p.EnterRule(localctx, 84, FigParserRULE_matchPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(526)
		p.Expr()
	}
	p.SetState(531)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FigParserCOMMA {
		{
			p.SetState(527)
			p.Match(FigParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(528)
			p.Expr()
		}

		p.SetState(533)
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
	p.EnterRule(localctx, 86, FigParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(534)
		p.Match(FigParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&275379310017277952) != 0 {
		{
			p.SetState(535)
			p.Expr()
		}
		p.SetState(540)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(536)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(537)
				p.Expr()
			}

			p.SetState(542)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(545)
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
	p.EnterRule(localctx, 88, FigParserRULE_objectLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(547)
		p.Match(FigParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FigParserID || _la == FigParserSTRING {
		{
			p.SetState(548)
			p.ObjectEntry()
		}
		p.SetState(553)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == FigParserCOMMA {
			{
				p.SetState(549)
				p.Match(FigParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(550)
				p.ObjectEntry()
			}

			p.SetState(555)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(558)
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
	p.EnterRule(localctx, 90, FigParserRULE_objectEntry)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(560)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FigParserID || _la == FigParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(561)
		p.Match(FigParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(562)
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
