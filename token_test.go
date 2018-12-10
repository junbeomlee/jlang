package jlang

import (
	"fmt"
	"testing"
)

func TestToken_String(t *testing.T) {
	tok := Token{Val: "1", Type: NOT_EQ, Offset: 0, Line: 0, Column: 1}
	fmt.Printf("%+v", tok)
}
