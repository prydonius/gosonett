package lexer

import (
	"errors"
	"fmt"
	"github.com/owlci/gosonett/token"
	"unicode"
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

func (l *Lexer) willOverflow() bool {
	fmt.Printf("  willOverflow -> %v %v\n", l.index, l.sourceLength)
	return l.index+1 >= l.sourceLength
}

// NOTE: This might need to rune, depending on what character set jsonet supports.
func (l *Lexer) CurrentChar() byte {
	return l.Source[l.index]
}

func (l *Lexer) NextChar() (byte, error) {
	if l.willOverflow() {
		// TODO: Zero valued bytes, how?
		return '\x00', errors.New("Lexer will overflow")
	}

	char := l.CurrentChar()

	fmt.Printf("Token: %q  , Index: %d  , Length: %d\n", char, l.index, l.sourceLength)

	l.index++

	if char == '\n' {
		l.Position.NextLine()
	} else {
		l.Position.NextChar()
	}

	return char, nil
}

// Returns the next lookahead character without advancing the lexer
func (l *Lexer) Peek() (byte, error) {
	if l.willOverflow() {
		// TODO: Zero valued bytes, how?
		return '\x00', errors.New("Lexer will overflow")
	}

	return l.Source[l.index+1], nil
}

// Returns the next valid token in the input stream
func (l *Lexer) Tokenize() token.Token {
	var tok token.Token

	l.eatWhitespace()
	char := l.CurrentChar()

	switch char {
	case '!':
		tok = token.New(token.BANG, char)
	case '$':
		tok = token.New(token.DOLLAR, char)
	case ':':
		tok = token.New(token.COLON, char)
	case '~':
		tok = token.New(token.TILDE, char)
	case '+':
		tok = token.New(token.PLUS, char)
	case '-':
		tok = token.New(token.MINUS, char)
	case '&':
		tok = token.New(token.AMP, char)
	case '|':
		tok = token.New(token.PIPE, char)
	case '^':
		tok = token.New(token.CARET, char)
	case '=':
		tok = token.New(token.ASSIGN, char)
	case '<':
		tok = token.New(token.LANGLE, char)
	case '>':
		tok = token.New(token.RANGLE, char)
	case '*':
		tok = token.New(token.STAR, char)
	case '/':
    peekedChar, err := l.Peek()

    // Maybe Peek should just return os.EOF constant or something
    if err != nil {
      panic("Out of bounds")
    }

    // Single-line comment
    if peekedChar == '/' {
      l.eatCurrentLine()
      return l.Tokenize()
    }

    // Multi-line comment
    if peekedChar == '*' {
      // TODO: Handle multi-line-comments
      // l.eatMultiLineComment()
      // return l.Tokenize()
    }

    // Must be a single token acting as an operator
		tok = token.New(token.SLASH, char)
	case '%':
		tok = token.New(token.PERC, char)
	case '#':
    l.eatCurrentLine()
    return l.Tokenize()
	}

  fmt.Printf("Appending token => %q %q", tok.Type, tok.Value)

	// Store the token
	l.Tokens = append(l.Tokens, tok)

	// End of token, advance to next byte
	l.NextChar()

	return tok
}

// Chews up insignificant whitespace up until the next potential token
func (l *Lexer) eatWhitespace() {
	// TODO: More idiomatic way to do this
	for unicode.IsSpace(rune(l.CurrentChar())) {
		l.NextChar()
	}
}

func (l *Lexer) eatUntil(untilChar byte) {
	for l.CurrentChar() != untilChar {
		l.NextChar()
	}

  // Point to the byte after our until
  l.NextChar()
}

func (l *Lexer) eatCurrentLine() {
  l.eatUntil('\n')
}

func (l *Lexer) eatMultiLineComment() {
}
