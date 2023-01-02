package object

import "yail/token"

const NULL_OBJ = "NULL"

type Null struct {
}

var (
	NULL = &Null{}
)

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return token.NULL
}
