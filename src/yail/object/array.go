package object

import (
	"bytes"
	"strings"
)

const ARRAY_OBJ = "ARRAY"

type Array struct {
	Elements []Object
}

func NewArray(elements []Object) *Array {
	return &Array{
		Elements: elements,
	}
}

func (ao *Array) Type() ObjectType {
	return ARRAY_OBJ
}

func (ao *Array) Inspect() string {
	var out bytes.Buffer
	var elements []string
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}
