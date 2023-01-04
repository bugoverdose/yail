package object

import "yail/token"

const NULL_OBJ = "NULL"

var (
	NULL = &Null{}
)

type Null struct {
}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return token.NULL
}
