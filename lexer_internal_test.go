package jlang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNext_lexer(t *testing.T) {
	input := "hello world! \n hello world2"

	tests := []struct {
		expectedByte byte
	}{
		{byte('h')},
		{byte('e')},
		{byte('l')},
		{byte('l')},
		{byte('o')},
		{byte(' ')},
		{byte('w')},
		{byte('o')},
		{byte('r')},
		{byte('l')},
		{byte('d')},
		{byte('!')},
		{byte(' ')},
		{byte('\n')},
		{byte(' ')},
		{byte('h')},
		{byte('e')},
		{byte('l')},
		{byte('l')},
		{byte('o')},
		{byte(' ')},
		{byte('w')},
		{byte('o')},
		{byte('r')},
		{byte('l')},
		{byte('d')},
		{byte('2')},
		{byte(0)},
	}

	lex := &Lexer{
		input: input,
	}

	for _, test := range tests {
		ch := lex.next()
		assert.Equal(t, test.expectedByte, ch)
	}

	assert.Equal(t, lex.line, 1)
}

func TestBackup_lexer(t *testing.T) {
	input := "he\na"
	lex := &Lexer{
		input: input,
	}

	assert.Equal(t, lex.next(), byte('h'))

	// test backup
	lex.backup()

	assert.Equal(t, lex.next(), byte('h'))
	assert.Equal(t, lex.next(), byte('e'))
	assert.Equal(t, lex.next(), byte('\n'))
	assert.Equal(t, lex.line, 1)

	// test backup decrease line number when meet '\n'
	lex.backup()
	assert.Equal(t, lex.line, 0)
}

func TestPeek_lexer(t *testing.T) {
	input := "he\na"
	lex := &Lexer{
		input: input,
	}

	// pick return next char but does not increase pos
	assert.Equal(t, lex.peek(), byte('h'))
	assert.Equal(t, lex.start, 0)
	assert.Equal(t, lex.pos, 0)
}

func TestEmit_lexer(t *testing.T) {
	input := "he\na"
	lex := &Lexer{
		input:   input,
		tokench: make(chan Token),
	}
	go func() {
		lex.next()
		lex.emit(IDENT)
	}()

	ch := <-lex.tokench

	assert.Equal(t, ch, Token{
		Val:    "h",
		Type:   IDENT,
		Column: 0,
		Line:   0,
		Offset: 1,
	})
}

func TestAccept_lexer(t *testing.T) {
	input := "he\na"
	lex := &Lexer{
		input: input,
	}

	assert.True(t, lex.accept("h"))
	assert.False(t, lex.accept("h"))
	assert.Equal(t, lex.next(), byte('e'))
}

func TestAcceptRun_lexer(t *testing.T) {

	input := "he\nak"
	lex := New(input)

	lex.acceptRun("he\na")
	assert.Equal(t, lex.next(), byte('k'))
	assert.Equal(t, lex.next(), byte(eof))
}
