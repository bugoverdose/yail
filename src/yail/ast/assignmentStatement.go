package ast

import (
	"bytes"
	"yail/token"
)

type VariableBindingStatement struct {
	Token token.Token
	Name  *IdentifierExpression
	Value Expression
}

func NewVariableBinding(keyword token.Token, name *IdentifierExpression, value Expression) *VariableBindingStatement {
	if keyword.Type != token.VAR && keyword.Type != token.VAL {
		panic("Invalid implementation: var or val token expected.")
	}
	return &VariableBindingStatement{
		Token: keyword,
		Name:  name,
		Value: value,
	}
}

func (statement *VariableBindingStatement) statementNode() {}
func (statement *VariableBindingStatement) TokenLiteral() string {
	return statement.Token.Literal
}
func (statement *VariableBindingStatement) String() string {
	var out bytes.Buffer
	out.WriteString(statement.TokenLiteral() + " ") // var
	out.WriteString(statement.Name.String())        // a
	out.WriteString(" = ")
	out.WriteString(statement.Value.String()) // 10
	out.WriteString(";")
	return out.String() // var a = 10;
}

type ReassignmentStatement struct {
	Token token.Token
	Name  *IdentifierExpression
	Value Expression
}

func NewReassignment(keyword token.Token, name *IdentifierExpression, value Expression) *ReassignmentStatement {
	if keyword.Type != token.IDENTIFIER {
		panic("Invalid implementation: identifier token expected.")
	}
	return &ReassignmentStatement{
		Token: keyword,
		Name:  name,
		Value: value,
	}
}

func (statement *ReassignmentStatement) statementNode() {}
func (statement *ReassignmentStatement) TokenLiteral() string {
	return statement.Token.Literal
}
func (statement *ReassignmentStatement) String() string {
	var out bytes.Buffer
	out.WriteString(statement.Name.String())
	out.WriteString(" = ")
	out.WriteString(statement.Value.String())
	out.WriteString(";")
	return out.String()
}
