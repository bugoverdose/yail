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
		right := Eval(node.RightNode, env)
		// TODO: 할당하려는 우항 평가에서 문제 있는 경우에 대한 예외 처리 추가
		return evalPrefixExpression(node.Operator, right)
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

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalNotOperatorExpression(right)
	case "-":
		return evalNegativePrefixOperatorExpression(right)
	default:
		return object.NewError("unknown operator: %s%s", operator, right.Type())
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
