package object

const STRING_OBJ = "STRING"

type String struct {
	Value string
}

func NewString(value string) *String {
	return &String{Value: value}
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}
