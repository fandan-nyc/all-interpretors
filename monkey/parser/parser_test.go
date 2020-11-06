package parser

import (
	"fmt"
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

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("expect 1 statement, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("cannot convert to expression statement")
		}
		prefixExpression, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("cannot convert to *prefixExpression")
		}
		if prefixExpression.Operator != tt.operator {
			t.Fatalf("wrgong operator, expect %s , got %s", tt.operator, prefixExpression.Operator)
		}
		if !testIntegerLiteral(t, prefixExpression.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("cannot convert to *ast.IntegerLiteral")
		return false
	}
	if integ.Value != value {
		t.Errorf("integer value is wrong. expect %d, got %d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("token literal is wrong. expect %d, got %s", value, integ.TokenLiteral())
	}
	return true
}

func TestParsingInfixEXpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5+5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 /5;", 5, "/", 5},
		{"5 >5;", 5, ">", 5},
		{"5 <5;", 5, "<", 5},
		{"5 ==5;", 5, "==", 5},
		{"5!=5;", 5, "!=", 5},
		{"5!=5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1{
			t.Fail()
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fail()
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {t.Fail()}
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			t.Fail()
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {t.Fail()}
		if exp.Operator != tt.operator {
			t.Fail()
		}
	}
}
