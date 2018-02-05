// Generated from Yql.g4 by ANTLR 4.7.

package grammar

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 24, 146,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19,
	6, 19, 102, 10, 19, 13, 19, 14, 19, 103, 3, 20, 3, 20, 7, 20, 108, 10,
	20, 12, 20, 14, 20, 111, 11, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 5,
	22, 118, 10, 22, 3, 22, 6, 22, 121, 10, 22, 13, 22, 14, 22, 122, 3, 23,
	5, 23, 126, 10, 23, 3, 23, 6, 23, 129, 10, 23, 13, 23, 14, 23, 130, 3,
	23, 3, 23, 7, 23, 135, 10, 23, 12, 23, 14, 23, 138, 11, 23, 3, 24, 6, 24,
	141, 10, 24, 13, 24, 14, 24, 142, 3, 24, 3, 24, 3, 109, 2, 25, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 2, 43,
	22, 45, 23, 47, 24, 3, 2, 6, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 4, 2,
	45, 45, 47, 47, 5, 2, 11, 12, 15, 15, 34, 34, 2, 152, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45,
	3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 49, 3, 2, 2, 2, 5, 53, 3, 2, 2, 2, 7,
	56, 3, 2, 2, 2, 9, 58, 3, 2, 2, 2, 11, 60, 3, 2, 2, 2, 13, 62, 3, 2, 2,
	2, 15, 65, 3, 2, 2, 2, 17, 67, 3, 2, 2, 2, 19, 69, 3, 2, 2, 2, 21, 72,
	3, 2, 2, 2, 23, 75, 3, 2, 2, 2, 25, 78, 3, 2, 2, 2, 27, 80, 3, 2, 2, 2,
	29, 84, 3, 2, 2, 2, 31, 86, 3, 2, 2, 2, 33, 89, 3, 2, 2, 2, 35, 94, 3,
	2, 2, 2, 37, 101, 3, 2, 2, 2, 39, 105, 3, 2, 2, 2, 41, 114, 3, 2, 2, 2,
	43, 117, 3, 2, 2, 2, 45, 125, 3, 2, 2, 2, 47, 140, 3, 2, 2, 2, 49, 50,
	7, 99, 2, 2, 50, 51, 7, 112, 2, 2, 51, 52, 7, 102, 2, 2, 52, 4, 3, 2, 2,
	2, 53, 54, 7, 113, 2, 2, 54, 55, 7, 116, 2, 2, 55, 6, 3, 2, 2, 2, 56, 57,
	7, 42, 2, 2, 57, 8, 3, 2, 2, 2, 58, 59, 7, 43, 2, 2, 59, 10, 3, 2, 2, 2,
	60, 61, 7, 63, 2, 2, 61, 12, 3, 2, 2, 2, 62, 63, 7, 35, 2, 2, 63, 64, 7,
	63, 2, 2, 64, 14, 3, 2, 2, 2, 65, 66, 7, 64, 2, 2, 66, 16, 3, 2, 2, 2,
	67, 68, 7, 62, 2, 2, 68, 18, 3, 2, 2, 2, 69, 70, 7, 64, 2, 2, 70, 71, 7,
	63, 2, 2, 71, 20, 3, 2, 2, 2, 72, 73, 7, 62, 2, 2, 73, 74, 7, 63, 2, 2,
	74, 22, 3, 2, 2, 2, 75, 76, 7, 107, 2, 2, 76, 77, 7, 112, 2, 2, 77, 24,
	3, 2, 2, 2, 78, 79, 7, 46, 2, 2, 79, 26, 3, 2, 2, 2, 80, 81, 7, 35, 2,
	2, 81, 82, 7, 107, 2, 2, 82, 83, 7, 112, 2, 2, 83, 28, 3, 2, 2, 2, 84,
	85, 7, 8747, 2, 2, 85, 30, 3, 2, 2, 2, 86, 87, 7, 35, 2, 2, 87, 88, 7,
	8747, 2, 2, 88, 32, 3, 2, 2, 2, 89, 90, 7, 118, 2, 2, 90, 91, 7, 116, 2,
	2, 91, 92, 7, 119, 2, 2, 92, 93, 7, 103, 2, 2, 93, 34, 3, 2, 2, 2, 94,
	95, 7, 104, 2, 2, 95, 96, 7, 99, 2, 2, 96, 97, 7, 110, 2, 2, 97, 98, 7,
	117, 2, 2, 98, 99, 7, 103, 2, 2, 99, 36, 3, 2, 2, 2, 100, 102, 9, 2, 2,
	2, 101, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103,
	104, 3, 2, 2, 2, 104, 38, 3, 2, 2, 2, 105, 109, 7, 41, 2, 2, 106, 108,
	11, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 110, 3, 2,
	2, 2, 109, 107, 3, 2, 2, 2, 110, 112, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2,
	112, 113, 7, 41, 2, 2, 113, 40, 3, 2, 2, 2, 114, 115, 9, 3, 2, 2, 115,
	42, 3, 2, 2, 2, 116, 118, 9, 4, 2, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3,
	2, 2, 2, 118, 120, 3, 2, 2, 2, 119, 121, 5, 41, 21, 2, 120, 119, 3, 2,
	2, 2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2,
	123, 44, 3, 2, 2, 2, 124, 126, 9, 4, 2, 2, 125, 124, 3, 2, 2, 2, 125, 126,
	3, 2, 2, 2, 126, 128, 3, 2, 2, 2, 127, 129, 5, 41, 21, 2, 128, 127, 3,
	2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2,
	2, 131, 132, 3, 2, 2, 2, 132, 136, 7, 48, 2, 2, 133, 135, 5, 41, 21, 2,
	134, 133, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136,
	137, 3, 2, 2, 2, 137, 46, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 141, 9,
	5, 2, 2, 140, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 140, 3, 2, 2,
	2, 142, 143, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 145, 8, 24, 2, 2, 145,
	48, 3, 2, 2, 2, 11, 2, 103, 109, 117, 122, 125, 130, 136, 142, 3, 8, 2,
	2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'and'", "'or'", "'('", "')'", "'='", "'!='", "'>'", "'<'", "'>='",
	"'<='", "'in'", "','", "'!in'", "'\u2229'", "'!\u2229'", "'true'", "'false'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TRUE",
	"FALSE", "FIELDNAME", "STRING", "INT", "FLOAT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "TRUE", "FALSE", "FIELDNAME",
	"STRING", "DIGIT", "INT", "FLOAT", "WS",
}

type YqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewYqlLexer(input antlr.CharStream) *YqlLexer {

	l := new(YqlLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Yql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// YqlLexer tokens.
const (
	YqlLexerT__0      = 1
	YqlLexerT__1      = 2
	YqlLexerT__2      = 3
	YqlLexerT__3      = 4
	YqlLexerT__4      = 5
	YqlLexerT__5      = 6
	YqlLexerT__6      = 7
	YqlLexerT__7      = 8
	YqlLexerT__8      = 9
	YqlLexerT__9      = 10
	YqlLexerT__10     = 11
	YqlLexerT__11     = 12
	YqlLexerT__12     = 13
	YqlLexerT__13     = 14
	YqlLexerT__14     = 15
	YqlLexerTRUE      = 16
	YqlLexerFALSE     = 17
	YqlLexerFIELDNAME = 18
	YqlLexerSTRING    = 19
	YqlLexerINT       = 20
	YqlLexerFLOAT     = 21
	YqlLexerWS        = 22
)
