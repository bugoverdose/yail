package evaluator

import (
	"yail/ast/statement"
	"yail/environment"
	"yail/object"
	"yail/token"
)

func evalStatement(node statement.Statement, env *environment.Environment) object.Object {
	switch node := node.(type) {
	case *statement.ExpressionStatement:
		return Eval(node, env)
	case *statement.VariableBinding:
		return evalVariableBinding(node, env)
	case *statement.Reassignment:
		return evalReassignment(node, env)
	}
	return nil
}

func evalVariableBinding(node *statement.VariableBinding, env *environment.Environment) object.Object {
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

func assignNewVariable(node *statement.VariableBinding, env *environment.Environment, val object.Object) (bool, *object.Error) {
	switch node.Token.Type {
	case token.VAL:
		return env.ImmutableAssign(node.Name.Value, val)
	case token.VAR:
		return env.MutableAssign(node.Name.Value, val)
	default:
		return false, object.NewError("unexpected token: %s", node.Token.Literal)
	}
}

func evalReassignment(node *statement.Reassignment, env *environment.Environment) object.Object {
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
