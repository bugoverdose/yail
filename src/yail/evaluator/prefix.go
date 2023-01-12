package evaluator

import (
	"yail/ast"
	"yail/environment"
	"yail/object"
	"yail/token"
)

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
