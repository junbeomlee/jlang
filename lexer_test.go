package jlang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer_NextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType  TokenType
		expectedValue string
	}{
		{ASSIGN, "="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	l := New(input)

	for _, test := range tests {
		token := l.NextToken()
		assert.Equal(t, token.Type, test.expectedType)
		assert.Equal(t, token.Val, test.expectedValue)
	}

	assert.Equal(t, l.line, 0)
	assert.Equal(t, l.next(), byte(0))
}

func TestLexer_NextToken2(t *testing.T) {
	input := `let five = 5;
 			  let ten = 10;

			  let add = fn(x, y) {
				  x + y;
			  };	
	
			  let result = add(five, ten);`

	tests := []struct {
		expectedType  TokenType
		expectedValue string
	}{
		{LET, "let"},
		{IDENT, "five"},
		{ASSIGN, "="},
		{INT, "5"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "ten"},
		{ASSIGN, "="},
		{INT, "10"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "add"},
		{ASSIGN, "="},
		{FUNCTION, "fn"},
		{LPAREN, "("},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "y"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "y"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{SEMICOLON, ";"},
		{LET, "let"},
		{IDENT, "result"},
		{ASSIGN, "="},
		{IDENT, "add"},
		{LPAREN, "("},
		{IDENT, "five"},
		{COMMA, ","},
		{IDENT, "ten"},
		{RPAREN, ")"},
		{SEMICOLON, ";"},
	}

	l := New(input)
	for _, test := range tests {
		token := l.NextToken()
		assert.Equal(t, token.Type, test.expectedType)
		assert.Equal(t, token.Val, test.expectedValue)
	}
}
