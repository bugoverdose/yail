package object

const RETURN_VALUE_OBJ = "RETURN_VALUE"

type ReturnValue struct {
	Value Object
}

func NewReturnValue(value Object) *ReturnValue {
	return &ReturnValue{Value: value}
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (rv *ReturnValue) Unwrap() Object {
	return rv.Value
}
