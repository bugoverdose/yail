package environment

import "yail/object"

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

func (e *Environment) Reassign(name string, val object.Object) bool {
	data, ok := e.dataStorage[name]
	if !ok {
		return false
	}
	if !data.isMutable {
		return false
	}
	e.dataStorage[name] = newValue(val, true)
	return true
}
