package statement

import (
	"bytes"
	"yail/ast/expression"
	"yail/token"
)

type VariableBinding struct {
	Token token.Token
	Name  *expression.Identifier
	Value expression.Expression
}

func NewVariableBinding(keyword token.Token, name *expression.Identifier, value expression.Expression) *VariableBinding {
	if keyword.Type != token.VAR && keyword.Type != token.VAL {
		panic("Invalid implementation: var or val token expected.")
	}
	return &VariableBinding{
		Token: keyword,
		Name:  name,
		Value: value,
	}
}

func (statement *VariableBinding) statementNode() {}

func (statement *VariableBinding) TokenLiteral() string {
	return statement.Token.Literal
}

func (statement *VariableBinding) String() string {
	var out bytes.Buffer
	out.WriteString(statement.TokenLiteral() + " ") // var
	out.WriteString(statement.Name.String())        // a
	out.WriteString(" = ")
	out.WriteString(statement.Value.String()) // 10
	out.WriteString(";")
	return out.String() // var a = 10;
}
