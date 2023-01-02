package ast

import (
	"strconv"
	"testing"
	"yail/token"
	"yail/utils"
)

func TestAstIntegration(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			newVariableAssignment(NewIdentifierFrom("x"), newIntegerLiteral("10")),
			newValueAssignment(NewIdentifierFrom("y"), newIntegerLiteral("20")),
			newExpressionStatement("z"),
		},
	}

	utils.ValidateValue(program.String(), "var x = 10; val y = 20; z;", t)
}

func newVariableAssignment(name *IdentifierExpression, value Expression) *VariableBindingStatement {
	return NewVariableBinding(token.New(token.VAR), name, value)
}

func newValueAssignment(name *IdentifierExpression, value Expression) *VariableBindingStatement {
	return NewVariableBinding(token.New(token.VAL), name, value)
}

func newIntegerLiteral(literal string) *IntegerLiteralExpression {
	value, err := strconv.ParseInt(literal, 0, 64)
	if err != nil {
		return nil
	}
	return &IntegerLiteralExpression{
		Token: token.NewInteger(literal),
		Value: value,
	}
}

func newExpressionStatement(identifier string) *ExpressionStatement {
	expr := NewIdentifierFrom(identifier)
	return NewExpressionStatement(expr.Token, expr)
}
