package evaluator

import (
	"yail/ast"
	"yail/environment"
	"yail/object"
)

func Eval(node ast.Node, env *environment.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case ast.Statement:
		return evalStatement(node, env)
	case ast.Expression:
		return evalExpression(node, env)
	}
	return nil
}

func evalProgram(program *ast.Program, env *environment.Environment) object.Object {
	var result object.Object

	for _, stmt := range program.Statements {
		result = Eval(stmt, env)
		switch result := result.(type) {
		case *object.Error:
			return result
		case *object.ReturnValue:
			return result.Unwrap()
		}
	}
	return result
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
