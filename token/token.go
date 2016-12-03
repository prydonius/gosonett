package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

func New(tokenType TokenType, char byte) Token {
	return Token{Type: tokenType, Value: string(char)}
}

const (
	// Special Lexemes
	EOF = "EOF"

	// TODO: Symbols
	// {}[],.();
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COMMA     = ","
	DOT       = "."
	LPAREN    = "("
	RPAREN    = ")"
	SEMICOLON = ";"

	// Operators
	// !$:~+-&|^=<>*/%
	MINUS  = "-"
	PLUS   = "+"
	BANG   = "!"
	TILDE  = "~"
	DOLLAR = "$"
	COLON  = ":"
	AMP    = "&"
	PIPE   = "|"
	CARET  = "^"
	ASSIGN = "="
	LANGLE = "<"
	RANGLE = ">"
	STAR   = "*"
	SLASH  = "/"
	PERC   = "%"

	// Keywords
	ASSERT     = "assert"
	ERROR      = "error"
	IF         = "if"
	THEN       = "then"
	ELSE       = "else"
	TRUE       = "true"
	FALSE      = "false"
	FOR        = "for"
	FUNCTION   = "function"
	IMPORT     = "import"
	IMPORTSTR  = "importstr"
	TAILSTRICT = "tailstrict"
	IN         = "in"
	LOCAL      = "local"
	NULL       = "null"
	SELF       = "self"
	SUPER      = "super"
)
