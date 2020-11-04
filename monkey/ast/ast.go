package ast

import (
	"bytes"

	"github.com/fandan-nyc/all-interpretors/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stat := range p.Statements {
		out.WriteString(stat.String())
	}
	return out.String()
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

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) String() string  { return i.Value }

func (lt *LetStatement) StatementNode() {}

func (lt *LetStatement) TokenLiteral() string { return lt.Token.Literal }

func (lt *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(lt.TokenLiteral() + " ")
	out.WriteString(lt.Name.String())
	out.WriteString(" = ")
	if lt.Value != nil {
		out.WriteString(lt.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

// why the identifier is an Expression
// this is just to keep things simple. identifier in other parts do produce values

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Token.Literal + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

func (es *ExpressionStatement) StatementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
