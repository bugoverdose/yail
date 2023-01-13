package object

import "fmt"

const INTEGER_OBJ = "INTEGER"

type Integer struct {
	Value   int64
	hashKey HashKey
}

func NewInteger(value int64) *Integer {
	hashKey := HashKey{Type: INTEGER_OBJ, Value: uint64(value)}
	return &Integer{Value: value, hashKey: hashKey}
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) HashKey() HashKey {
	return i.hashKey
}
