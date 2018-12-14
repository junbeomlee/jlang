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

type BooleanLiteral struct {
	Token jlang.Token
	Value bool
}

func (b *BooleanLiteral) expressionNode() {}

func (b *BooleanLiteral) TokenValue() string {
	return b.Token.Val
}

func (b *BooleanLiteral) String() string {
	return b.Token.Val
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
		out.WriteString(rs.ReturnValue.String())
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

// <if> <condition> <blockstatements> <else> <blockstatements>
type IFExpression struct {
	Token       jlang.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IFExpression) TokenValue() string {
	return ie.Token.Val
}

func (ie *IFExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	if ie.Condition != nil {
		out.WriteString("(")
		out.WriteString(ie.Condition.String())
		out.WriteString(")")
	}

	if ie.Consequence != nil {
		out.WriteString("{")
		for _, stmt := range ie.Consequence.Statements {
			out.WriteString(stmt.String() + "\n")
		}
		out.WriteString("}")
	}

	out.WriteString("else")

	if ie.Alternative != nil {
		out.WriteString("{")
		for _, stmt := range ie.Alternative.Statements {
			out.WriteString(stmt.String() + "\n")
		}
		out.WriteString("}")
	}

	return out.String()
}

type FunctionExpression struct {
	Token jlang.Token
	Args  []*Identifier
	Body  *BlockStatement
}

func (f *FunctionExpression) TokenValue() string {
	return f.Token.Val
}

func (f *FunctionExpression) String() string {
	var out bytes.Buffer

	out.WriteString("fn")
	out.WriteString("(")
	for _, arg := range f.Args {
		out.WriteString(arg.String())
	}
	out.WriteString(")")
	out.WriteString("{")
	out.WriteString(f.Body.String())
	out.WriteString("}")

	return out.String()
}

type CallExpression struct {
	Token jlang.Token
	Args  []Expression
}

func (c *CallExpression) TokenValue() string {
	return c.Token.Val
}

func (c *CallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(c.Token.Val)
	out.WriteString("(")
	for _, arg := range c.Args {
		out.WriteString(arg.String())
	}
	out.WriteString(")")

	return out.String()
}

func (c *CallExpression) expressionNode() {
	panic("implement me")
}

func (f *FunctionExpression) expressionNode() {
	panic("implement me")
}

func (ie *IFExpression) expressionNode() {}

type BlockStatement struct {
	Token      jlang.Token
	Statements []Statement
}

func (b *BlockStatement) TokenValue() string {
	return b.Token.Val
}

func (b *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range b.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
