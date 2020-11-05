package parser

import (
	"testing"

	"github.com/fandan-nyc/all-interpretors/monkey/ast"
	"github.com/fandan-nyc/all-interpretors/monkey/lexer"
	"github.com/fandan-nyc/all-interpretors/monkey/token"
)

func TestLetStatements(t *testing.T) {
	input := `let x = 5;
let y = 10;
let foobar = 838383;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("parsing prorgam went nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("we have 3 different statements but got only %d.", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		checkParserErrors(t, p)
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) != 0 {
		t.Errorf("parser has %d errors", len(errors))
		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}
}

func testLetStatements(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Fatalf("the token type should be let, but get %s", stmt.TokenLiteral())
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if ok == false {
		t.Fatalf("cannot convert to let statement, got %T", stmt)
		return false
	}
	if letStmt.Name.Value != name {
		t.Fatalf("the token name is expected to be %s, but we got %s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Fatalf("the let statement token literal should be  %s, but got %s", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func TestReturnStatement(t *testing.T) {
	input := `return 5;
return 10;
return 992233;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		rtStat, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if rtStat.TokenLiteral() != "return" {
			t.Errorf("return statement.TokenLiteral is not 'return', but %s", rtStat.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	data := `foobar`

	l := lexer.New(data)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("length should be 1,but got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatal("cannot convert to expression statement")
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatal("cannot convert to identifier")
	}
	if ident.Token.Type != token.IDENT {
		t.Fatalf("wrong identifier, expect IDENT, got %s", ident.TokenLiteral())
	}
	if ident.Value != "foobar" {
		t.Fatalf("wrong value, got %s", ident.Value)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("length should be 1, got %d", len(program.Statements))
	}
	statement := program.Statements[0]
	es, ok := statement.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("cannot convert to ExpressionStatement")
	}
	literal, ok := es.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("cannot convert to IntegerLiteral")
	}
	if literal.Token.Type != token.INT {
		t.Fatalf("token type should be INT")
	}
	if literal.TokenLiteral() != "5" {
		t.Fatalf("token literal value should be 5")
	}
	if literal.Value != 5 {
		t.Fatalf("token value should be int 5")
	}
}
