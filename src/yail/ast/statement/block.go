package statement

import (
	"bytes"
	"yail/token"
)

type Block struct {
	Token      token.Token
	Statements []Statement
}

func NewBlock(statements []Statement) *Block {
	return &Block{
		Token:      token.LEFT_BRACKET_TOKEN,
		Statements: statements,
	}
}

func (b *Block) statementNode() {}

func (b *Block) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Block) String() string {
	var out bytes.Buffer
	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
