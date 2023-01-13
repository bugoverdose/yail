package evaluator

import (
	"yail/ast"
	"yail/environment"
	"yail/object"
	"yail/token"
)

func evalInfixExpression(node *ast.InfixExpression, env *environment.Environment) object.Object {
	left := Eval(node.LeftNode, env)
	if isError(left) {
		return left
	}
	right := Eval(node.RightNode, env)
	if isError(right) {
		return right
	}
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(node.Token, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(node.Token, left, right)
	case node.Token.Type == token.EQUAL:
		return object.GetPooledBooleanObject(left == right)
	case node.Token.Type == token.NOT_EQUAL:
		return object.GetPooledBooleanObject(left != right)
	case left.Type() != right.Type():
		return object.NewError("type mismatch: %s %s %s", left.Type(), node.Operator, right.Type())
	default:
		return object.NewError("unknown operator: %s %s %s", left.Type(), node.Operator, right.Type())
	}
}

func evalIntegerInfixExpression(infixToken token.Token, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch infixToken.Literal {
	case token.PLUS:
		return object.NewInteger(leftVal + rightVal)
	case token.MINUS:
		return object.NewInteger(leftVal - rightVal)
	case token.MULTIPLY:
		return object.NewInteger(leftVal * rightVal)
	case token.DIVIDE:
		return object.NewInteger(leftVal / rightVal)
	case token.MODULO:
		return object.NewInteger(leftVal % rightVal)
	case token.LESS_THAN:
		return object.GetPooledBooleanObject(leftVal < rightVal)
	case token.GREATER_THAN:
		return object.GetPooledBooleanObject(leftVal > rightVal)
	case token.EQUAL:
		return object.GetPooledBooleanObject(leftVal == rightVal)
	case token.NOT_EQUAL:
		return object.GetPooledBooleanObject(leftVal != rightVal)
	case token.LESS_OR_EQUAL:
		return object.GetPooledBooleanObject(leftVal <= rightVal)
	case token.GREATER_OR_EQUAL:
		return object.GetPooledBooleanObject(leftVal >= rightVal)
	default:
		return object.NewError("unknown operator: %s %s %s", left.Type(), infixToken.Literal, right.Type())
	}
}

func evalStringInfixExpression(infixToken token.Token, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	switch infixToken.Literal {
	case token.PLUS:
		return object.NewString(leftVal + rightVal)
	case token.EQUAL:
		return object.GetPooledBooleanObject(left == right)
	case token.NOT_EQUAL:
		return object.GetPooledBooleanObject(left != right)
	default:
		return object.NewError("unknown operator: %s %s %s", left.Type(), infixToken.Literal, right.Type())
	}
}
