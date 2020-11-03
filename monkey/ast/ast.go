package ast

import "github.com/fandan-nyc/all-interpretors/monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

// here we have statement and expression
// the difference basically (not 100% accurate) is that, statement does not generate a value, but expression does
// but this is language specific. in some languages the statement does generate value.
// in some languages, fn(x, y) { return x + y; } is just the def of the func
// in other languages, it could be used as input of another func

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token // this is LET
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // this is IDENTIFIER
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) ExpressionToken() {}

func (lt *LetStatement) StatementNode() {}

func (lt *LetStatement) TokenLiteral() string { return lt.Token.Literal }

// why the identifier is an Expression
// this is just to keep things simple. identifier in other parts do produce values
