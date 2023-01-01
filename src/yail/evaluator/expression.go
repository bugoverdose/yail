package evaluator

import (
	"yail/ast/expression"
	"yail/environment"
	"yail/object"
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
	switch node.Operator {
	case "!":
		return evalNotOperatorExpression(right)
	case "-":
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
