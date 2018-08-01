// Code generated from Test.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 53, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 7, 9, 40, 10, 9, 12, 9, 14, 9, 43, 11, 9, 3, 10, 6, 10, 46, 10, 10,
	13, 10, 14, 10, 47, 3, 11, 3, 11, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 3, 2, 6, 4, 2, 67,
	92, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12,
	15, 15, 34, 34, 2, 54, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23,
	3, 2, 2, 2, 5, 25, 3, 2, 2, 2, 7, 27, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2, 11,
	31, 3, 2, 2, 2, 13, 33, 3, 2, 2, 2, 15, 35, 3, 2, 2, 2, 17, 37, 3, 2, 2,
	2, 19, 45, 3, 2, 2, 2, 21, 49, 3, 2, 2, 2, 23, 24, 7, 96, 2, 2, 24, 4,
	3, 2, 2, 2, 25, 26, 7, 44, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 49, 2, 2,
	28, 8, 3, 2, 2, 2, 29, 30, 7, 45, 2, 2, 30, 10, 3, 2, 2, 2, 31, 32, 7,
	47, 2, 2, 32, 12, 3, 2, 2, 2, 33, 34, 7, 42, 2, 2, 34, 14, 3, 2, 2, 2,
	35, 36, 7, 43, 2, 2, 36, 16, 3, 2, 2, 2, 37, 41, 9, 2, 2, 2, 38, 40, 9,
	3, 2, 2, 39, 38, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41,
	42, 3, 2, 2, 2, 42, 18, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 46, 9, 4, 2,
	2, 45, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 48,
	3, 2, 2, 2, 48, 20, 3, 2, 2, 2, 49, 50, 9, 5, 2, 2, 50, 51, 3, 2, 2, 2,
	51, 52, 8, 11, 2, 2, 52, 22, 3, 2, 2, 2, 5, 2, 41, 47, 3, 8, 2, 2,
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
	"", "'^'", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "ID", "DIGIT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ID", "DIGIT",
	"WS",
}

type TestLexer struct {
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

func NewTestLexer(input antlr.CharStream) *TestLexer {

	l := new(TestLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Test.G4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TestLexer tokens.
const (
	TestLexerT__0  = 1
	TestLexerT__1  = 2
	TestLexerT__2  = 3
	TestLexerT__3  = 4
	TestLexerT__4  = 5
	TestLexerT__5  = 6
	TestLexerT__6  = 7
	TestLexerID    = 8
	TestLexerDIGIT = 9
	TestLexerWS    = 10
)
