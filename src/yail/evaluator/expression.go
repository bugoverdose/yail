package evaluator

import (
	"yail/ast"
	"yail/environment"
	"yail/object"
	"yail/token"
)

func evalExpression(node ast.Expression, env *environment.Environment) object.Object {
	switch node := node.(type) {
	case *ast.IdentifierExpression:
		return evalIdentifier(node, env)
	case *ast.IntegerLiteralExpression:
		return object.NewInteger(node.Value)
	case *ast.BooleanExpression:
		return object.GetPooledBooleanObject(node.Value)
	case *ast.NullExpression:
		return object.NULL
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)
	case *ast.InfixExpression:
		return evalInfixExpression(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.FunctionLiteral:
		return environment.NewFunction(node, env)
	case *ast.CallExpression:
		return evalFunctionCall(node, env)
	}
	return nil
}

func evalIdentifier(node *ast.IdentifierExpression, env *environment.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return object.NewError("identifier not found: " + node.Value)
	}
	return val
}

func evalPrefixExpression(node *ast.PrefixExpression, env *environment.Environment) object.Object {
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
	case object.TRUE:
		return object.FALSE
	case object.FALSE:
		return object.TRUE
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

func evalIfExpression(expression *ast.IfExpression, env *environment.Environment) object.Object {
	condition := Eval(expression.Condition, env)
	if isError(condition) {
		return condition
	}
	if condition == object.TRUE {
		return Eval(expression.Consequence, env)
	}
	if condition == object.FALSE && expression.Alternative != nil {
		return Eval(expression.Alternative, env)
	}
	return object.NULL
}
