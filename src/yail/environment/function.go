package environment

import (
	"bytes"
	"strings"
	"yail/ast"
	"yail/object"
)

const FUNCTION_OBJ = "FUNCTION"

// TODO: move to object package and handle import cycle
type Function struct {
	Parameters []*ast.IdentifierExpression
	Body       *ast.BlockStatement
	Env        *Environment
}

func NewFunction(node *ast.FunctionLiteral, env *Environment) *Function {
	return &Function{
		Parameters: node.Parameters,
		Body:       node.Body,
		Env:        env,
	}
}

func (f *Function) Type() object.ObjectType {
	return FUNCTION_OBJ
}
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}
