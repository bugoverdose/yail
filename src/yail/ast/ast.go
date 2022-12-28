package ast

import (
	"bytes"
	"strings"
	"yail/ast/statement"
)

type Program struct {
	Statements []statement.Statement
}

func (p *Program) TokenLiteral() string {
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String() + " ")
	}
	return strings.TrimSpace(out.String())
}
