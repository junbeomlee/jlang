package jlang

import (
	"bytes"
)

const eof = 0

type Lexer struct {
	filename string
	input    string
	start    int
	pos      int
	line     int
	tokench  chan Token
}

func New(input string) *Lexer {
	l := &Lexer{
		input:   input,
		tokench: make(chan Token, 2),
	}

	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *Lexer) run() {
	for state := lexInput; state != nil; {
		state = state(l)
	}
	close(l.tokench)
}

func (l *Lexer) next() byte {

	if int(l.pos) >= len(l.input) {
		return eof
	}

	ch := l.input[l.pos]
	if ch == '\n' {
		l.line++
	}

	l.pos += 1
	return ch
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.pos -= 1
	// Correct newline count.
	if l.input[l.pos] == '\n' {
		l.line--
	}
}

// peek returns but does not consume the next byte in the input.
func (l *Lexer) peek() byte {
	ch := l.next()
	l.backup()
	return ch
}

// emit passes an token back
func (l *Lexer) emit(t TokenType) {
	l.tokench <- Token{t, l.input[l.start:l.pos], l.start, l.pos, l.line}
	l.start = l.pos
}

// nextToken returns the next token from the input.
// Called by the parser, not in the lexing goroutine.
func (l *Lexer) NextToken() Token {
	return <-l.tokench
}

//
//// ignore skips over the pending input before this point.
//func (l *lexer) ignore() {
//	l.line += strings.Count(l.input[l.start:l.pos], "\n")
//	l.start = l.pos
//}
//
// accept consumes the next byte if it's from the valid set.
func (l *Lexer) accept(valid string) bool {

	ch := l.next()
	if bytes.Contains([]byte(valid), []byte{ch}) {
		return true
	}

	// eof
	if ch == 0 {
		return false
	}

	l.backup()
	return false
}

// acceptRun consumes a run of byte from the valid set.
func (l *Lexer) acceptRun(valid string) {

	var ch byte
	for {
		ch = l.next()
		if !bytes.Contains([]byte(valid), []byte{ch}) {
			break
		}
	}

	if ch == 0 {
		return
	}
	l.backup()
}

//
//// errorf returns an error token and terminates the scan by passing
//// back a nil pointer that will be the next state, terminating l.nextItem.
//func (l *lexer) errorf(format string, args ...interface{}) stateFn {
//	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...), l.line}
//	return nil
//}
//

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

// lexInput is a basic scanner that scans the elements
func lexInput(l *Lexer) stateFn {
	ch := l.next()

	switch ch {
	case '=':
		l.emit(ASSIGN)
	case ';':
		l.emit(SEMICOLON)
	case ')':
		l.emit(RPAREN)
	case '(':
		l.emit(LPAREN)
	case ',':
		l.emit(COMMA)
	case '+':
		l.emit(PLUS)
	case '{':
		l.emit(LBRACE)
	case '}':
		l.emit(RBRACE)
	case 0:
		l.emit(EOF)
	}

	return lexInput
}

//
//func lexNumber(l *Lexer) stateFn {
//
//}
