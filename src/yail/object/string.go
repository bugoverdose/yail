package object

import "hash/fnv"

const STRING_OBJ = "STRING"

type String struct {
	Value   string
	hashKey HashKey
}

func NewString(value string) *String {
	h := fnv.New64a()
	h.Write([]byte(value))
	hashKey := HashKey{Type: STRING_OBJ, Value: h.Sum64()}
	return &String{Value: value, hashKey: hashKey}
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) HashKey() HashKey {
	return s.hashKey
}
