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

		if returnStmt.TokenValue() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.String())
		}
	}
}

func TestParser_Parse_IdentityExpression(t *testing.T) {
	input := `footer;`

	l := jlang.New(input)
	p := New(l)
	program := p.Parse()

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statments. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not ast.ExpressionStatemet. got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression is not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "footer" {
		t.Errorf("ident.Value not %s, got=%s", "footer", ident.Value)
	}
}

func TestParser_Parse_LiteralExpression(t *testing.T) {
	input := `5;`

	l := jlang.New(input)
	p := New(l)
	program := p.Parse()

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statments. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statement[0] is not ast.ExpressionStatemet. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression is not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("IntegerLiteral.Value not %s, got=%s", "5", literal.Value)
	}
}

func TestParser_Parse_PrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		l := jlang.New(tt.input)
		p := New(l)
		program := p.Parse()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.RightExpression, tt.value) {
			return
		}
	}
}

func TestParser_Parse_InfixExpression(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := jlang.New(tt.input)
		p := New(l)
		program := p.Parse()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		opExp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Errorf("exp is not ast.InfixExpression. got=%T(%s)", stmt.Expression, stmt.Expression)
			return
		}

		if !testIntegerLiteral(t, opExp.LeftExpression, tt.leftValue) {
			return
		}

		if opExp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%q", tt.operator, opExp.Operator)
		}

		if !testIntegerLiteral(t, opExp.RightExpression, tt.rightValue) {
			return
		}
	}
}

func TestParser_Parse_OperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		l := jlang.New(tt.input)
		p := New(l)
		program := p.Parse()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
	integerLiteral, ok := exp.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("exp is not *ast.IntegerLiteral. got=%T", integerLiteral)
		return false
	}

	if integerLiteral.Value != value {
		t.Errorf("interger value is not %d. got=%T", integerLiteral.Value, value)
		return false
	}

	return true
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
