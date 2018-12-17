package jlang

import (
	"testing"

	"fmt"

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

	assert.Equal(t, l.line, 7)
}

func TestLexer_NextToken3(t *testing.T) {
	input := `let five = 5;
			  let ten = 10;

			  let add = fn(x, y) {
			 	 x + y;
			  };

			  let result = add(five, ten);
			  !-/*5;
			  5 < 10 > 5;

			  if (5 < 10) {
				 return true;
			  } else {
				 return false;
			  }

			  10 == 10;
			  10 != 9;`

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
		{BANG, "!"},
		{MINUS, "-"},
		{SLASH, "/"},
		{ASTERISK, "*"},
		{INT, "5"},
		{SEMICOLON, ";"},
		{INT, "5"},
		{LT, "<"},
		{INT, "10"},
		{GT, ">"},
		{INT, "5"},
		{SEMICOLON, ";"},
		{IF, "if"},
		{LPAREN, "("},
		{INT, "5"},
		{LT, "<"},
		{INT, "10"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{TRUE, "true"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{ELSE, "else"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{FALSE, "false"},
		{SEMICOLON, ";"},
		{RBRACE, "}"},
		{INT, "10"},
		{EQ, "=="},
		{INT, "10"},
		{SEMICOLON, ";"},
		{INT, "10"},
		{NOT_EQ, "!="},
		{INT, "9"},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	l := New(input)
	for i, test := range tests {
		token := l.NextToken()

		if token.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, token.Type)
		}

		if token.Val != test.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedValue, token.Val)
		}
	}

	if l.line != 18 {
		t.Fatalf("num of line was wrong. expected=%q, got=%q",
			18, l.line)
	}
}

func TestLexer_NextToken4(t *testing.T) {
	input := `2*3+5`

	l := New(input)
	fmt.Print(l.NextToken().Val)
	fmt.Print(l.NextToken().Val)
	fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
	//fmt.Print(l.NextToken().Val)
}
