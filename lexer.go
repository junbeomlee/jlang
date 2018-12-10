package jlang

//
//type Position struct {
//	Filename string // filename, if any
//	Offset   int    // byte offset, starting at 0
//	Line     int    // line number, starting at 1
//	Column   int    // column number, starting at 1 (character count per line)
//}

type Lexer struct {
	filename string
	input    string
	start    int
	pos      int
	line     int
	tokench  chan Token
}

func New(input string) *Lexer {
	return &Lexer{
		input:   input,
		tokench: make(chan Token, 2),
	}
}

const eof = 0

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*Lexer) stateFn

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

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() byte {
	ch := l.next()
	l.backup()
	return ch
}

//// emit passes an item back to the client.
//func (l *lexer) emit(t itemType) {
//	l.items <- item{t, l.start, l.input[l.start:l.pos], l.line}
//	// Some items contain text internally. If so, count their newlines.
//	switch t {
//	case itemText, itemRawString, itemLeftDelim, itemRightDelim:
//		l.line += strings.Count(l.input[l.start:l.pos], "\n")
//	}
//	l.start = l.pos
//}
//
//// ignore skips over the pending input before this point.
//func (l *lexer) ignore() {
//	l.line += strings.Count(l.input[l.start:l.pos], "\n")
//	l.start = l.pos
//}
//
//// accept consumes the next rune if it's from the valid set.
//func (l *lexer) accept(valid string) bool {
//	if strings.ContainsRune(valid, l.next()) {
//		return true
//	}
//	l.backup()
//	return false
//}
//
//// acceptRun consumes a run of runes from the valid set.
//func (l *lexer) acceptRun(valid string) {
//	for strings.ContainsRune(valid, l.next()) {
//	}
//	l.backup()
//}
//
//// errorf returns an error token and terminates the scan by passing
//// back a nil pointer that will be the next state, terminating l.nextItem.
//func (l *lexer) errorf(format string, args ...interface{}) stateFn {
//	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...), l.line}
//	return nil
//}
//
//// nextItem returns the next item from the input.
//// Called by the parser, not in the lexing goroutine.
//func (l *lexer) nextItem() item {
//	return <-l.items
//}
