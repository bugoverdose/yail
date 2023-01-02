package evaluator

import (
	"yail/ast"
	"yail/environment"
	"yail/object"
	"yail/token"
)

func evalStatement(node ast.Statement, env *environment.Environment) object.Object {
	switch node := node.(type) {
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.VariableBindingStatement:
		return evalVariableBinding(node, env)
	case *ast.ReassignmentStatement:
		return evalReassignment(node, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	}
	return nil
}

func evalVariableBinding(node *ast.VariableBindingStatement, env *environment.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	ok, err := assignNewVariable(node, env, val)
	if !ok {
		return err
	}
	return nil
}

func assignNewVariable(node *ast.VariableBindingStatement, env *environment.Environment, val object.Object) (bool, *object.Error) {
	switch node.Token.Type {
	case token.VAL:
		return env.ImmutableAssign(node.Name.Value, val)
	case token.VAR:
		return env.MutableAssign(node.Name.Value, val)
	default:
		return false, object.NewError("unexpected token: %s", node.Token.Literal)
	}
}

func evalReassignment(node *ast.ReassignmentStatement, env *environment.Environment) object.Object {
	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}
	ok, err := env.Reassign(node.Name.Value, val)
	if !ok {
		return err
	}
	return nil
}

func evalBlockStatement(block *ast.BlockStatement, env *environment.Environment) object.Object {
	var result object.Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if isError(result) {
			return result
		}
	}
	return result
}
