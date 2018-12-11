package ast

import (
	"bytes"

	"github.com/junbeomlee/jlang"
)

type Node interface {
	TokenValue() string
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

func (p *Program) TokenValue() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenValue()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) TokenValue() string {
	return ls.Token.Val
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.Token.Val + " ")
	out.WriteString(ls.Ident.Value + " = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String() + " ")
	}

	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token       jlang.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenValue() string {
	return rs.Token.Val
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Token.Val + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String() + " ")
	}

	out.WriteString(";")
	return out.String()
}
