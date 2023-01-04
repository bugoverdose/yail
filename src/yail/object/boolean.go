package object

import "fmt"

const BOOLEAN_OBJ = "BOOLEAN"

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

type Boolean struct {
	Value bool
}

func GetPooledBooleanObject(input bool) *Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
