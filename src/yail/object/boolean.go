package object

import "fmt"

const BOOLEAN_OBJ = "BOOLEAN"

var (
	TRUE  = &Boolean{Value: true, hashKey: HashKey{Type: BOOLEAN_OBJ, Value: 1}}
	FALSE = &Boolean{Value: false, hashKey: HashKey{Type: BOOLEAN_OBJ, Value: 0}}
)

type Boolean struct {
	Value   bool
	hashKey HashKey
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

func (b *Boolean) HashKey() HashKey {
	return b.hashKey
}
