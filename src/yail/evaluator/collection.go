package evaluator

import (
	"yail/ast"
	"yail/environment"
	"yail/object"
)

func evalArrayLiteral(node *ast.ArrayLiteral, env *environment.Environment) object.Object {
	elements := evalExpressions(node.Elements, env)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}
	return object.NewArray(elements)
}

func evalIndexAccess(node *ast.IndexAccessExpression, env *environment.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	index := Eval(node.Index, env)
	if isError(index) {
		return index
	}
	return evalIndexExpression(left, index)
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	default:
		return object.NewError("unsupported operation: %s[%s]", left.Type(), index.Type())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)
	if idx < 0 || idx > max {
		return object.NULL
	}
	return arrayObject.Elements[idx]
}
