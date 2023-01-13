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

func evalHashMapLiteral(node *ast.HashMapLiteral, env *environment.Environment) object.Object {
	pairs := make(map[object.HashKey]*object.HashPair)
	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return object.NewError("%s can not be used as hash key", key.Type())
		}
		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}
		pairs[hashKey.HashKey()] = object.NewHashPair(key, value)
	}
	return object.NewHashMap(pairs)
}

func evalCollectionAccess(node *ast.CollectionAccessExpression, env *environment.Environment) object.Object {
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
		return evalArrayIndexAccessExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashKeyAccessExpression(left, index)
	default:
		return object.NewError("unsupported operation: %s[%s]", left.Type(), index.Type())
	}
}

func evalArrayIndexAccessExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)
	if idx < 0 || idx > max {
		return object.NULL
	}
	return arrayObject.Elements[idx]
}

func evalHashKeyAccessExpression(hashMap, index object.Object) object.Object {
	hashObject := hashMap.(*object.HashMap)
	key, ok := index.(object.Hashable)
	if !ok {
		return object.NewError("unusable as hash key: %s", index.Type())
	}
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return object.NULL
	}
	return pair.Value
}
