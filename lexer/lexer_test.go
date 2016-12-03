package lexer

import (
	"github.com/owlci/gosonett/token"
	"testing"
)

type TokenMatcher struct {
	expectedType  token.TokenType
	expectedValue string
}

func runTokenMatches(t *testing.T, source string, tests []TokenMatcher) {
	testsLength := len(tests)
	lexer := New(source)

	for _, tm := range tests {
		tok := lexer.Tokenize()

		if tok.Type != tm.expectedType {
			t.Fatalf("Wrong token type: expected=%q, got=%q", tm.expectedType, tok.Type)
		}

		if tok.Value != tm.expectedValue {
			t.Fatalf("Wrong token value: expected=%q, got=%q", tm.expectedValue, tok.Value)
		}
	}

	tokensArrayLength := len(lexer.Tokens)

	if testsLength != tokensArrayLength {
		t.Fatalf("Wrong token array length: expected=%d, got=%d", testsLength, tokensArrayLength)
	}
}

func TestSymbols(t *testing.T) {
	// tests := []TokenMatcher {
	//   {token.ASSIGN, "="},
	//   {token.PLUS, "+"},
	//   {token.LPAREN, "("},
	//   {token.RPAREN, ")"},
	//   {token.LBRACE, "{"},
	//   {token.RBRACE, "}"},
	//   {token.COMMA, ","},
	//   {token.SEMICOLON, ";"},
	//   {token.EOF, ""},
	// }
}

func TestOperators(t *testing.T) {
	source := "!$:~+-&|^=<>*/%"

	tests := []TokenMatcher{
		{token.BANG, "!"},
		{token.DOLLAR, "$"},
		{token.COLON, ":"},
		{token.TILDE, "~"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.AMP, "&"},
		{token.PIPE, "|"},
		{token.CARET, "^"},
		{token.ASSIGN, "="},
		{token.LANGLE, "<"},
		{token.RANGLE, ">"},
		{token.STAR, "*"},
		{token.SLASH, "/"},
		{token.PERC, "%"},
	}

	runTokenMatches(t, source, tests)
}

func TestWhitepaceBehaviour(t *testing.T) {
	source := "! =        %"

	tests := []TokenMatcher{
		{token.BANG, "!"},
		{token.ASSIGN, "="},
		{token.PERC, "%"},
	}

	runTokenMatches(t, source, tests)
}

func TestKeywords(t *testing.T) {
}
