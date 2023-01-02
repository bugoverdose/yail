package object

import "fmt"

const ERROR_OBJ = "ERROR"

type Error struct {
	Message string
}

func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "[ERROR] " + e.Message
}
