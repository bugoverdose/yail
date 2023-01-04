package environment

import (
	"yail/object"
	"yail/token"
)

type Environment struct {
	dataStorage map[string]value
	outerScope  *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]value)
	return &Environment{dataStorage: s, outerScope: nil}
}

func NewInnerEnvironment(outer *Environment) *Environment {
	s := make(map[string]value)
	return &Environment{dataStorage: s, outerScope: outer}
}

func (e *Environment) Get(name string) (object.Object, bool) {
	obj, ok := e.dataStorage[name]
	if ok {
		return obj.data, ok
	}
	if e.outerScope != nil {
		return e.outerScope.Get(name)
	}
	return nil, ok
}

func (e *Environment) ImmutableAssign(name string, val object.Object) (bool, *object.Error) {
	if _, ok := e.dataStorage[name]; ok {
		return false, object.NewError("given identifier '%s' is already declared", name)
	}
	e.dataStorage[name] = newValue(val, false)
	return true, nil
}

func (e *Environment) MutableAssign(name string, val object.Object) (bool, *object.Error) {
	if _, ok := e.dataStorage[name]; ok {
		return false, object.NewError("given identifier '%s' is already declared", name)
	}
	e.dataStorage[name] = newValue(val, true)
	return true, nil
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
