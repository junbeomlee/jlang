package ast

import (
	"github.com/junbeomlee/jlang"
)

type Node interface {
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) Value() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].String()
	} else {
		return ""
	}
}

type Identifier struct {
	Token jlang.Token
	Value string
}

type LetStatement struct {
	Token jlang.Token
	Ident *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) String() string {
	return ls.Token.Val
}

type ReturnStatement struct {
	Token       jlang.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) String() string {
	return rs.Token.Val
}
