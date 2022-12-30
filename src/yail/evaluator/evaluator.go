package evaluator

import (
	"yail/ast"
	"yail/ast/expression"
	"yail/ast/node"
	"yail/ast/statement"
	"yail/object"
)

func Eval(node node.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node)

	// TODO: Statement 노드인지 체크하고 별도로 분기
	case *statement.ExpressionStatement:
		return Eval(node.Expression)

	// TODO: Expression 노드인지 체크하고 별도로 분기
	case *expression.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

func evalProgram(program *ast.Program) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement)
	}
	return result
}
