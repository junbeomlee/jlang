package parser

import (
	"testing"

	"github.com/junbeomlee/jlang"
	"github.com/junbeomlee/jlang/ast"
)

func TestParser_Parse_LetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 5;
	let foobar = 838383;
`
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	l := jlang.New(input)
	p := New(l)

	program := p.Parse()
	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for i, tt := range tests {

		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestParser_Parse_ReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 12313;
	`

	l := jlang.New(input)
	p := New(l)

	program := p.Parse()
	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
		}

		if returnStmt.String() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.String())
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.String() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.String())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Ident.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Ident.Value)
		return false
	}

	if letStmt.Ident.Value != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, letStmt.Ident.Value)
		return false
	}

	return true
}
