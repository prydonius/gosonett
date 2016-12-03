package lexer

import (
	"github.com/owlci/gosonett/token"
)

type LexerPosition struct {
	line     int
	lineChar int
}

func (lp *LexerPosition) NextLine() {
	lp.line++
	lp.lineChar = 0
}

func (lp *LexerPosition) NextChar() {
	lp.lineChar++
}

type Lexer struct {
	Source       string
	Tokens       []token.Token
	Position     LexerPosition // represents position of char with new lines, good for debugging.
	index        int           // hold the current index of currentChar within the whole input string
	currentChar  byte
	sourceLength int
}

func New(source string) *Lexer {
	return &Lexer{
		Source:       source,
		Position:     LexerPosition{line: 0, lineChar: 0},
		index:        0,
		sourceLength: len(source),
	}
}

func (l *Lexer) NextChar() byte {
	// TODO: Check overflow using sourceLength here
	l.currentChar = l.Source[l.index]
	l.index++

	if l.currentChar == '\n' {
		l.Position.NextLine()
	} else {
		l.Position.NextChar()
	}

	return l.currentChar
}

// Returns the next valid token in the input stream
func (l *Lexer) Tokenize() token.Token {
	var tok token.Token

	l.NextChar()

	switch l.currentChar {
	case '{':
		tok = token.New(token.LBRACE, l.currentChar)
	case '}':
		tok = token.New(token.RBRACE, l.currentChar)
	case '[':
		tok = token.New(token.LBRACKET, l.currentChar)
	case ']':
		tok = token.New(token.RBRACKET, l.currentChar)
	case ',':
		tok = token.New(token.COMMA, l.currentChar)
	case '.':
		tok = token.New(token.DOT, l.currentChar)
	case '(':
		tok = token.New(token.LPAREN, l.currentChar)
	case ')':
		tok = token.New(token.RPAREN, l.currentChar)
	case ';':
		tok = token.New(token.SEMICOLON, l.currentChar)
	case '!':
		tok = token.New(token.BANG, l.currentChar)
	case '$':
		tok = token.New(token.DOLLAR, l.currentChar)
	case ':':
		tok = token.New(token.COLON, l.currentChar)
	case '~':
		tok = token.New(token.TILDE, l.currentChar)
	case '+':
		tok = token.New(token.PLUS, l.currentChar)
	case '-':
		tok = token.New(token.MINUS, l.currentChar)
	case '&':
		tok = token.New(token.AMP, l.currentChar)
	case '|':
		tok = token.New(token.PIPE, l.currentChar)
	case '^':
		tok = token.New(token.CARET, l.currentChar)
	case '=':
		tok = token.New(token.ASSIGN, l.currentChar)
	case '<':
		tok = token.New(token.LANGLE, l.currentChar)
	case '>':
		tok = token.New(token.RANGLE, l.currentChar)
	case '*':
		tok = token.New(token.STAR, l.currentChar)
	case '/':
		tok = token.New(token.SLASH, l.currentChar)
	case '%':
		tok = token.New(token.PERC, l.currentChar)
	}

	// Store the token
	l.Tokens = append(l.Tokens, tok)

	return tok
}

// Chews up insignificant whitespace
func (l *Lexer) eatWhitespace() {

}
