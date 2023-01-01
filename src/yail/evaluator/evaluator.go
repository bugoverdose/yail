package evaluator

import (
	"yail/ast"
	"yail/ast/expression"
	"yail/ast/node"
	"yail/ast/statement"
	"yail/environment"
	"yail/object"
)

func Eval(node node.Node, env *environment.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case statement.Statement:
		return evalStatement(node, env)
	case expression.Expression:
		return evalExpression(node, env)
	}
	return nil
}

func evalProgram(program *ast.Program, env *environment.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)
	}
	return result
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
