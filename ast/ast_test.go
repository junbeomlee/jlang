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

func TestFunctionExpression_String(t *testing.T) {
	fExp := FunctionExpression{
		Args: []*Identifier{
			{
				Value: "x",
				Token: jlang.Token{
					Val:  "x",
					Type: jlang.IDENT,
				},
			},
		},
		Body: &BlockStatement{
			Statements: []Statement{
				&ReturnStatement{
					Token: jlang.Token{
						Val:  "return",
						Type: jlang.RETURN,
					},
					ReturnValue: &IntegerLiteral{
						Value: 5,
						Token: jlang.Token{
							Val:  "5",
							Type: jlang.INT,
						},
					},
				},
			},
		},
		Token: jlang.Token{
			Val:  "fn",
			Type: jlang.FUNCTION,
		},
	}

	if fExp.String() != `fn(x){return 5;}` {
		t.Errorf("program.String() wrong. got=%q", fExp.String())
	}
}

func TestCallExpression_String(t *testing.T) {
	callExp := CallExpression{
		Args: []Expression{
			&Identifier{
				Value: "x",
				Token: jlang.Token{
					Val:  "x",
					Type: jlang.IDENT,
				},
			},
		},
		Function: &Identifier{
			Value: "add",
			Token: jlang.Token{
				Val:  "add",
				Type: jlang.IDENT,
			},
		},
	}

	if callExp.String() != `add(x)` {
		t.Errorf("program.String() wrong. got=%q", callExp.String())
	}
}
