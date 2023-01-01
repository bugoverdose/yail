package environment

import (
	"yail/object"
	"yail/token"
)

type Environment struct {
	dataStorage map[string]value
}

func NewEnvironment() *Environment {
	s := make(map[string]value)
	return &Environment{dataStorage: s}
}

func (e *Environment) Get(name string) (object.Object, bool) {
	obj, ok := e.dataStorage[name]
	if !ok {
		return nil, ok
	}
	return obj.data, ok
}

func (e *Environment) ImmutableAssign(name string, val object.Object) bool {
	if _, ok := e.dataStorage[name]; ok {
		return false
	}
	e.dataStorage[name] = newValue(val, false)
	return true
}

func (e *Environment) MutableAssign(name string, val object.Object) bool {
	if _, ok := e.dataStorage[name]; ok {
		return false
	}
	e.dataStorage[name] = newValue(val, true)
	return true
}

func (e *Environment) Reassign(name string, val object.Object) (bool, *object.Error) {
	data, ok := e.dataStorage[name]
	if !ok {
		return false, object.NewError("identifier not found: '%s'", name)
	}
	if !data.isMutable {
		return false, object.NewError("can not reassign variables declared with '%s'", token.VAL)
	}
	e.dataStorage[name] = newValue(val, true)
	return true, nil
}
