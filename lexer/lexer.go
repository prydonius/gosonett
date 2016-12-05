package lexer

import (
  "fmt"
	"errors"
	"unicode"
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
	return l.index+1 >= l.sourceLength
}

// NOTE: This might need to rune, depending on what character set jsonet supports.
func (l *Lexer) CurrentChar() byte {
	return l.Source[l.index]
}

func (l *Lexer) NextChar() (byte, error) {
	var char byte

	if l.willOverflow() {
		// TODO: Zero valued bytes, how?
		return char, errors.New("Lexer will overflow")
	}

	char = l.CurrentChar()
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

// Advances through the whole string source and tokenizes every lexeme
// func (l *Lexer) Lex() ([]token.Token, error) {
//   for r := l.Tokenize(); r != token.EOF; r = l.Tokenize()() {}

//   return l.Tokens, nil
// }

// Returns the next valid token in the input stream
func (l *Lexer) Tokenize() token.Token {
	var tok token.Token
  // var err error

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
  // case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
  // token, _ := l.lexNumber()
  default:
    if isIdentifierFirst(rune(char)) {
      // NOTE: Error handling
      tok, _ = l.lexIdentifier()
    } else {
      // TODO: Use the LexerPosition struct to print out something nice here
      panic("Unknown lexing character")
    }
	}

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

func (l *Lexer) lexIdentifier() (token.Token, error) {
  // var tok token.Token

  startIndex := l.index

  fmt.Println("Starting to lex identifier")

  // for unicode.IsLetter(rune(l.CurrentChar())) {
  //   fmt.Printf("Have identifier char : %q", l.CurrentChar())
  //   l.NextChar()
  // }

  ident := l.Source[startIndex:l.index]

  fmt.Printf("Have identifier : %q \n", ident)

  // matchKeyword and return keyword token

  // else return valid identifier token
  // return token.Token{Type: token.IDENT, Value: "if"}, nil
  return token.Token{Type: "IF", Value: "if"}, nil

}

// NOTE: Taken from here https://github.com/google/go-jsonnet/blob/master/lexer.go#L189
func isIdentifierFirst(r rune) bool {
	return unicode.IsUpper(r) || unicode.IsLower(r) || r == '_'
}

func isIdentifier(r rune) bool {
	return isIdentifierFirst(r) || unicode.IsNumber(r)
}
