package parser

import (
	"fmt"

	"github.com/junbeomlee/jlang"
	"github.com/junbeomlee/jlang/ast"
)

type Parser struct {
	l      *jlang.Lexer
	errors []string

	curToken  jlang.Token
	nextToken jlang.Token
}

func New(l *jlang.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.next()
	p.next()

	return p
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
	err := fmt.Sprintf("expected next token to be %s, got %s instead, line %d, col %d",
		t, p.nextToken.Type, p.nextToken.Line+1, p.nextToken.Column+1)
	p.errors = append(p.errors, err)
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
		return nil
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
