package ast

import (
	"testing"

	"github.com/junbeomlee/jlang"
)

func TestLetStatement_String_String(t *testing.T) {
	lstmt := LetStatement{
		Token: jlang.Token{
			Type: jlang.LET,
			Val:  "let",
		},
		Ident: &Identifier{
			Token: jlang.Token{
				Type: jlang.IDENT,
				Val:  "x",
			},
			Value: "x",
		},
		Value: &Identifier{
			Token: jlang.Token{
				Type: jlang.IDENT,
				Val:  "5",
			},
			Value: "5",
		},
	}

	if lstmt.String() != "let x = 5;" {
		t.Errorf("program.String() wrong. got=%q", lstmt.String())
	}
}
