package parser

import (
	"fmt"

	"strconv"

	"github.com/junbeomlee/jlang"
	"github.com/junbeomlee/jlang/ast"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[jlang.TokenType]int{
	jlang.EQ:       EQUALS,
	jlang.NOT_EQ:   EQUALS,
	jlang.LT:       LESSGREATER,
	jlang.GT:       LESSGREATER,
	jlang.PLUS:     SUM,
	jlang.MINUS:    SUM,
	jlang.SLASH:    PRODUCT,
	jlang.ASTERISK: PRODUCT,
}

// Expression
//  1. IdentifierExpression:      <Ident>
//  2. IntegerLiteralExpression:  <Int>
// 	3. PrefixExpression: 		  <prefix operator><expression>
//  4. InfixExpression: 		  <expression><infix operator><expression>

type (
	prefixParsefn func() ast.Expression
	infixParsefn  func(expression ast.Expression) ast.Expression
)

type Parser struct {
	l      *jlang.Lexer
	errors []string

	curToken  jlang.Token
	nextToken jlang.Token

	prefixParsefns map[jlang.TokenType]prefixParsefn
	infixParsefns  map[jlang.TokenType]infixParsefn
}

func New(l *jlang.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.next()
	p.next()

	p.prefixParsefns = make(map[jlang.TokenType]prefixParsefn)
	p.registerPrefix(jlang.IDENT, p.parseIdentifier)
	p.registerPrefix(jlang.INT, p.parseIntegerLiteral)
	p.registerPrefix(jlang.BANG, p.parsePrefixExpression)
	p.registerPrefix(jlang.MINUS, p.parsePrefixExpression)

	p.infixParsefns = make(map[jlang.TokenType]infixParsefn)
	p.registerInfix(jlang.EQ, p.parseInfixExpression)
	p.registerInfix(jlang.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(jlang.PLUS, p.parseInfixExpression)
	p.registerInfix(jlang.MINUS, p.parseInfixExpression)
	p.registerInfix(jlang.SLASH, p.parseInfixExpression)
	p.registerInfix(jlang.ASTERISK, p.parseInfixExpression)
	p.registerInfix(jlang.LT, p.parseInfixExpression)
	p.registerInfix(jlang.GT, p.parseInfixExpression)

	return p
}

func (p *Parser) curPrecedence() int {
	if pe, ok := precedences[p.curToken.Type]; ok {
		return pe
	}

	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if pe, ok := precedences[p.nextToken.Type]; ok {
		return pe
	}

	return LOWEST
}

func (p *Parser) registerPrefix(tokenType jlang.TokenType, fn prefixParsefn) {
	p.prefixParsefns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType jlang.TokenType, fn infixParsefn) {
	p.infixParsefns[tokenType] = fn
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Val}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	integerLiteral := &ast.IntegerLiteral{Token: p.curToken}

	v, err := strconv.ParseInt(p.curToken.Val, 0, 64)
	if err != nil {
		p.Error(fmt.Sprintf("could not parse %q as integer", p.curToken.Val))
		return nil
	}

	integerLiteral.Value = v
	return integerLiteral
}

func (p *Parser) Error(msg string) {
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) next() {
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t jlang.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t jlang.TokenType) bool {
	return p.nextToken.Type == t
}

func (p *Parser) expectPeek(t jlang.TokenType) bool {
	if p.peekTokenIs(t) {
		p.next()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t jlang.TokenType) {
	p.Error(fmt.Sprintf("expected next token to be %s, got %s instead, line %d, col %d",
		t, p.nextToken.Type, p.nextToken.Line+1, p.nextToken.Column+1))
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != jlang.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.next()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case jlang.LET:
		return p.parseLetStatement()
	case jlang.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(jlang.IDENT) {
		// errors
		return nil
	}

	stmt.Ident = &ast.Identifier{Token: p.curToken, Value: p.curToken.Val}

	if !p.expectPeek(jlang.ASSIGN) {
		return nil
	}

	for !p.curTokenIs(jlang.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {

	// Define statement
	stmt := &ast.ReturnStatement{Token: p.curToken}

	// Check next
	// todo expression
	p.next()
	for !p.curTokenIs(jlang.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(jlang.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {

	prefix := p.prefixParsefns[p.curToken.Type]
	if prefix == nil {
		p.Error(fmt.Sprintf("no prefix parse function for %s found", p.curToken.Type))
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(jlang.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParsefns[p.nextToken.Type]
		if infix == nil {
			return leftExp
		}

		p.next()
		infixExp := infix(leftExp)
		return infixExp
	}

	return leftExp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Val,
	}

	p.next()
	expression.RightExpression = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:          p.curToken,
		Operator:       p.curToken.Val,
		LeftExpression: left,
	}

	precedence := p.curPrecedence()
	p.next()
	exp.RightExpression = p.parseExpression(precedence)

	return exp
}
