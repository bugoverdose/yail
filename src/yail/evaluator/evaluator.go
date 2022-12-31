package evaluator

import (
	"yail/ast"
	"yail/ast/expression"
	"yail/ast/node"
	"yail/ast/statement"
	"yail/environment"
	"yail/object"
	"yail/token"
)

func Eval(node node.Node, env *environment.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)

	// TODO: Statement 노드인지 체크하고 별도로 분기
	case *statement.ExpressionStatement:
		return Eval(node.Expression, env)
	case *statement.VariableBinding:
		val := Eval(node.Value, env)
		// TODO: 할당하려는 우항 평가에서 문제 있는 경우에 대한 예외 처리 추가
		var ok bool
		if node.Token.Type == token.VAL {
			ok = env.ImmutableAssign(node.Name.Value, val)
		}
		if node.Token.Type == token.VAR {
			ok = env.MutableAssign(node.Name.Value, val)
		}
		if !ok {
			return nil // TODO: 예외처리 로직 제대로 추가
		}

	// TODO: Expression 노드인지 체크하고 별도로 분기
	case *expression.Identifier:
		return evalIdentifier(node, env)
	case *expression.IntegerLiteral:
		return &object.Integer{Value: node.Value}
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

func evalIdentifier(node *expression.Identifier, env *environment.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return object.NewError("identifier not found: " + node.Value)
	}
	return val
}
