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

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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
		var err *object.Error
		if node.Token.Type == token.VAL {
			ok, err = env.ImmutableAssign(node.Name.Value, val)
		}
		if node.Token.Type == token.VAR {
			ok, err = env.MutableAssign(node.Name.Value, val)
		}
		if !ok {
			return err
		}
		return nil
	case *statement.Reassignment:
		val := Eval(node.Value, env)
		// TODO: 할당하려는 우항 평가에서 문제 있는 경우에 대한 예외 처리 추가
		ok, err := env.Reassign(node.Name.Value, val)
		if !ok {
			return err
		}
		return nil

	// TODO: Expression 노드인지 체크하고 별도로 분기
	case *expression.Identifier:
		return evalIdentifier(node, env)
	case *expression.IntegerLiteral:
		return object.NewInteger(node.Value)
	case *expression.Boolean:
		return getPooledBooleanObject(node.Value)
	case *expression.Prefix:
		right := Eval(node.RightNode, env)
		// TODO: 할당하려는 우항 평가에서 문제 있는 경우에 대한 예외 처리 추가
		return evalPrefixExpression(node.Operator, right)
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

func getPooledBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalNotOperatorExpression(right)
	case "-":
		return evalNegativePrefixOperatorExpression(right)
	default:
		return object.NewError("unknown operator: %s%s", operator, right.Type())
	}
}

func evalNotOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	default:
		return object.NewError("unknown operator: !%s", right.Type())
	}
}

func evalNegativePrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return object.NewError("unknown operator: -%s", right.Type())
	}
	value := right.(*object.Integer).Value
	return object.NewInteger(-value)
}
