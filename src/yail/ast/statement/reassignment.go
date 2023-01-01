package statement

import (
	"bytes"
	"yail/ast/expression"
	"yail/token"
)

type Reassignment struct {
	Token token.Token
	Name  *expression.Identifier
	Value expression.Expression
}

func NewReassignment(keyword token.Token, name *expression.Identifier, value expression.Expression) *Reassignment {
	if keyword.Type != token.IDENTIFIER {
		panic("Invalid implementation: identifier token expected.")
	}
	return &Reassignment{
		Token: keyword,
		Name:  name,
		Value: value,
	}
}

func (statement *Reassignment) statementNode() {}

func (statement *Reassignment) TokenLiteral() string {
	return statement.Token.Literal
}

func (statement *Reassignment) String() string {
	var out bytes.Buffer
	out.WriteString(statement.Name.String())
	out.WriteString(" = ")
	out.WriteString(statement.Value.String())
	out.WriteString(";")
	return out.String()
}
