// Code generated from Practica4.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 34, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 5, 5, 21, 10, 5, 3, 5, 6, 5, 24, 10, 5,
	13, 5, 14, 5, 25, 3, 6, 6, 6, 29, 10, 6, 13, 6, 14, 6, 30, 3, 6, 3, 6,
	2, 2, 7, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11,
	12, 15, 15, 34, 34, 2, 36, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 15,
	3, 2, 2, 2, 7, 17, 3, 2, 2, 2, 9, 20, 3, 2, 2, 2, 11, 28, 3, 2, 2, 2, 13,
	14, 7, 42, 2, 2, 14, 4, 3, 2, 2, 2, 15, 16, 7, 43, 2, 2, 16, 6, 3, 2, 2,
	2, 17, 18, 7, 46, 2, 2, 18, 8, 3, 2, 2, 2, 19, 21, 7, 47, 2, 2, 20, 19,
	3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 23, 3, 2, 2, 2, 22, 24, 9, 2, 2, 2,
	23, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3,
	2, 2, 2, 26, 10, 3, 2, 2, 2, 27, 29, 9, 3, 2, 2, 28, 27, 3, 2, 2, 2, 29,
	30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 3, 2, 2,
	2, 32, 33, 8, 6, 2, 2, 33, 12, 3, 2, 2, 2, 6, 2, 20, 25, 30, 3, 8, 2, 2,
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
	"", "'('", "')'", "','",
}

var lexerSymbolicNames = []string{
	"", "PARA", "PARC", "COMA", "NUMBER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"PARA", "PARC", "COMA", "NUMBER", "WHITESPACE",
}

type Practica4Lexer struct {
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

func NewPractica4Lexer(input antlr.CharStream) *Practica4Lexer {

	l := new(Practica4Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Practica4.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Practica4Lexer tokens.
const (
	Practica4LexerPARA       = 1
	Practica4LexerPARC       = 2
	Practica4LexerCOMA       = 3
	Practica4LexerNUMBER     = 4
	Practica4LexerWHITESPACE = 5
)