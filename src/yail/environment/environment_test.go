package environment

import (
	"testing"
	"yail/object"
)

func TestGetNotFound(t *testing.T) {
	env := NewEnvironment()
	_, ok := env.Get("x")
	validateValue(ok, false, t)
}

func TestImmutableAssignment(t *testing.T) {
	env := NewEnvironment()
	var value object.Object = object.Integer{Value: 10}

	assignedOk := env.ImmutableAssign("x", value)
	obj, ok := env.Get("x")
	reassignedOk := env.Reassign("x", object.Integer{Value: 20})

	validateValue(assignedOk, true, t)
	validateObject(obj, value, t)
	validateValue(ok, true, t)
	validateValue(reassignedOk, false, t)
}

func TestMutableAssignment(t *testing.T) {
	env := NewEnvironment()
	var value object.Object = object.Integer{Value: 10}
	var updatedValue object.Object = object.Integer{Value: 20}

	assignedOk := env.MutableAssign("x", value)
	obj, getOk := env.Get("x")
	validateValue(assignedOk, true, t)
	validateObject(obj, value, t)
	validateValue(getOk, true, t)

	reassignedOk := env.Reassign("x", updatedValue)
	updatedObj, updatedGetOk := env.Get("x")
	validateValue(reassignedOk, true, t)
	validateObject(updatedObj, updatedValue, t)
	validateValue(updatedGetOk, true, t)
}

func TestCanNotReassignWithAssginFunctions(t *testing.T) {
	env := NewEnvironment()
	var value object.Object = object.Integer{Value: 10}

	assignedOk := env.MutableAssign("x", value)
	reassignedOk := env.MutableAssign("x", object.Integer{Value: 20})
	obj, _ := env.Get("x")

	validateValue(assignedOk, true, t)
	validateValue(reassignedOk, false, t)
	validateObject(obj, value, t)
}

func validateObject(actual, expected object.Object, t *testing.T) {
	if actual.Type() != expected.Type() {
		t.Errorf("expected %+v to be %+v", actual, expected)
	}
	switch actual := actual.(type) {
	case object.Integer:
		validateValue(actual.Value, expected.(object.Integer).Value, t)
	}
}

func validateValue[T comparable](actual, expected T, t *testing.T) {
	if actual != expected {
		t.Errorf("expected %+v to be %+v", actual, expected)
	}
}
