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
		return Eval(node.Expression, env)
	case *statement.VariableBinding:
		val := Eval(node.Value, env)
		// TODO: 할당하려는 우항 평가에서 문제 있는 경우에 대한 예외 처리 추가
		ok, err := evalVariableBinding(node, env, val)
		if !ok {
			return err
		}
	case *statement.Reassignment:
		val := Eval(node.Value, env)
		// TODO: 할당하려는 우항 평가에서 문제 있는 경우에 대한 예외 처리 추가
		ok, err := env.Reassign(node.Name.Value, val)
		if !ok {
			return err
		}
	}
	return nil
}

func evalVariableBinding(node *statement.VariableBinding, env *environment.Environment, val object.Object) (bool, *object.Error) {
	switch node.Token.Type {
	case token.VAL:
		return env.ImmutableAssign(node.Name.Value, val)
	case token.VAR:
		return env.MutableAssign(node.Name.Value, val)
	default:
		return false, object.NewError("unexpected token: %s", node.Token.Literal)
	}
}
