package ast

import (
	"bytes"
	"yail/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func NewBlock(statements []Statement) *BlockStatement {
	return &BlockStatement{
		Token:      token.LEFT_BRACKET_TOKEN,
		Statements: statements,
	}
}

func (b *BlockStatement) statementNode() {}
func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
