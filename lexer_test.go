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
