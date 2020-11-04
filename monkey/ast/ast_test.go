package ast

import (
	"testing"

	"github.com/fandan-nyc/all-interpretors/monkey/token"
)

func TestString(t *testing.T) {
	p := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "a",
					},
					Value: "a"},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "b",
					},
					Value: "b",
				},
			},
		},
	}
	if p.String() != "let a = b;" {
		t.Fatalf("string method for program failed. we got %s", p.String())
	}
}
