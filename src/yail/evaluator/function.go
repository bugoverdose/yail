package evaluator

import (
	"yail/ast"
	"yail/environment"
	"yail/object"
)

var builtinFunctions = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewError("wrong number of arguments: expected 1, but received %d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return object.NewError("len(%s) not supported", arg.Type())
			}
		},
	},
}

func evalFunctionCall(node *ast.CallExpression, env *environment.Environment) object.Object {
	boundFunctionFromEnv := Eval(node.Function, env)
	if isError(boundFunctionFromEnv) {
		return boundFunctionFromEnv
	}
	args := evalArguments(node, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0] // error object
	}
	return applyFunction(boundFunctionFromEnv, args)
}

func evalArguments(node *ast.CallExpression, env *environment.Environment) []object.Object {
	var result []object.Object
	for _, e := range node.Arguments {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated} // return single error object
		}
		result = append(result, evaluated)
	}
	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch function := fn.(type) {
	case *environment.Function:
		innerEnv := createInnerScopeEnvironment(function, args)
		evaluated := Eval(function.Body, innerEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return function.Fn(args...)
	default:
		return object.NewError("failed to invoke %s as a function", fn.Type())
	}
}

func createInnerScopeEnvironment(fn *environment.Function, args []object.Object) *environment.Environment {
	env := environment.NewInnerEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.MutableAssign(param.Value, args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Unwrap()
	}
	return obj
}
