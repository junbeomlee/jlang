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

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenValue() string {
	return i.Token.Val
}

func (i *Identifier) String() string {
	return i.Token.Val
}

type IntegerLiteral struct {
	Token jlang.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}

func (i *IntegerLiteral) TokenValue() string {
	return i.Token.Val
}

func (i *IntegerLiteral) String() string {
	return i.Token.Val
}

// PrefixExpression form will be <operator> <right expression>
type PrefixExpression struct {
	Token           jlang.Token
	Operator        string
	RightExpression Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenValue() string {
	return pe.Token.Val
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)

	if pe.RightExpression != nil {
		out.WriteString(pe.RightExpression.String())
	}
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token           jlang.Token
	Operator        string
	LeftExpression  Expression
	RightExpression Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenValue() string {
	return ie.Token.Val
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	if ie.LeftExpression != nil {
		out.WriteString(ie.LeftExpression.String())
	}
	out.WriteString(" " + ie.Operator + " ")
	if ie.RightExpression != nil {
		out.WriteString(ie.RightExpression.String())
	}
	out.WriteString(")")
	return out.String()
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
		out.WriteString(ls.Value.String())
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

type ExpressionStatement struct {
	Token      jlang.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenValue() string {
	return es.Token.Val
}

func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	if es.Expression != nil {
		out.WriteString(es.Expression.String())
	}

	return out.String()
}
