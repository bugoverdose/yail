package evaluator

import (
	"yail/ast/expression"
	"yail/environment"
	"yail/object"
	"yail/token"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func evalExpression(node expression.Expression, env *environment.Environment) object.Object {
	switch node := node.(type) {
	case *expression.Identifier:
		return evalIdentifier(node, env)
	case *expression.IntegerLiteral:
		return object.NewInteger(node.Value)
	case *expression.Boolean:
		return getPooledBooleanObject(node.Value)
	case *expression.Prefix:
		return evalPrefixExpression(node, env)
	case *expression.Infix:
		return evalInfixExpression(node, env)
	}
	return nil
}

func evalIdentifier(node *expression.Identifier, env *environment.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return object.NewError("identifier not found: " + node.Value)
	}
	return val
}

func getPooledBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(node *expression.Prefix, env *environment.Environment) object.Object {
	right := Eval(node.RightNode, env)
	if isError(right) {
		return right
	}
	switch node.Token.Literal {
	case token.NOT:
		return evalNotOperatorExpression(right)
	case token.MINUS:
		return evalNegativePrefixOperatorExpression(right)
	default:
		return object.NewError("unknown operator: %s%s", node.Operator, right.Type())
	}
}

func evalNotOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	default:
		return object.NewError("unknown operator: !%s", right.Type())
	}
}

func evalNegativePrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return object.NewError("unknown operator: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return object.NewInteger(-value)
}

func evalInfixExpression(node *expression.Infix, env *environment.Environment) object.Object {
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
	case node.Token.Type == token.EQUAL:
		return getPooledBooleanObject(left == right)
	case node.Token.Type == token.NOT_EQUAL:
		return getPooledBooleanObject(left != right)
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
		return getPooledBooleanObject(leftVal < rightVal)
	case token.GREATER_THAN:
		return getPooledBooleanObject(leftVal > rightVal)
	case token.EQUAL:
		return getPooledBooleanObject(leftVal == rightVal)
	case token.NOT_EQUAL:
		return getPooledBooleanObject(leftVal != rightVal)
	case token.LESS_OR_EQUAL:
		return getPooledBooleanObject(leftVal <= rightVal)
	case token.GREATER_OR_EQUAL:
		return getPooledBooleanObject(leftVal >= rightVal)
	default:
		return object.NewError("unknown operator: %s %s %s", left.Type(), infixToken.Literal, right.Type())
	}
}
