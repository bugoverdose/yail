package evaluator

import (
	"yail/ast"
	"yail/environment"
	"yail/object"
)

func evalExpression(node ast.Expression, env *environment.Environment) object.Object {
	switch node := node.(type) {
	case *ast.IdentifierExpression:
		return evalIdentifier(node, env)
	case *ast.IntegerLiteralExpression:
		return object.NewInteger(node.Value)
	case *ast.StringLiteralExpression:
		return object.NewString(node.Value)
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
	case *ast.ArrayLiteral:
		return evalArrayLiteral(node, env)
	case *ast.IndexAccessExpression:
		return evalIndexAccess(node, env)
	}
	return nil
}

func evalIdentifier(node *ast.IdentifierExpression, env *environment.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtinFunctions[node.Value]; ok {
		return builtin
	}
	return object.NewError("identifier not found: " + node.Value)
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

func evalExpressions(exps []ast.Expression, env *environment.Environment) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated} // return single error object
		}
		result = append(result, evaluated)
	}
	return result
}
