package environment

import (
	"testing"
	"yail/object"
	"yail/utils"
)

func TestGetNotFound(t *testing.T) {
	env := NewEnvironment()
	_, ok := env.Get("x")
	utils.ValidateValue(ok, false, t)
}

func TestImmutableAssignment(t *testing.T) {
	env := NewEnvironment()
	var value = &object.Integer{Value: 10}

	assignedOk, assignErr := env.ImmutableAssign("x", value)
	obj, ok := env.Get("x")
	reassignedOk, err := env.Reassign("x", &object.Integer{Value: 20})

	utils.ValidateValue(assignedOk, true, t)
	utils.ValidateValue(assignErr, nil, t)
	utils.ValidateObject(obj, value, t)
	utils.ValidateValue(ok, true, t)
	utils.ValidateValue(reassignedOk, false, t)
	utils.ValidateValue(err.Message, "can not reassign variables declared with 'val'", t)
}

func TestMutableAssignment(t *testing.T) {
	env := NewEnvironment()
	var value = &object.Integer{Value: 10}
	var updatedValue = &object.Integer{Value: 20}

	assignedOk, assignErr := env.MutableAssign("x", value)
	obj, getOk := env.Get("x")
	utils.ValidateValue(assignedOk, true, t)
	utils.ValidateValue(assignErr, nil, t)
	utils.ValidateObject(obj, value, t)
	utils.ValidateValue(getOk, true, t)

	reassignedOk, err := env.Reassign("x", updatedValue)
	updatedObj, updatedGetOk := env.Get("x")
	utils.ValidateValue(reassignedOk, true, t)
	utils.ValidateValue(err, nil, t)
	utils.ValidateObject(updatedObj, updatedValue, t)
	utils.ValidateValue(updatedGetOk, true, t)
}

func TestCanNotReassignWithAssignFunctions(t *testing.T) {
	env := NewEnvironment()
	var value = &object.Integer{Value: 10}

	assignedOk, assignErr := env.MutableAssign("x", value)
	reassignedOk, err := env.MutableAssign("x", &object.Integer{Value: 20})
	obj, _ := env.Get("x")

	utils.ValidateValue(assignedOk, true, t)
	utils.ValidateValue(assignErr, nil, t)
	utils.ValidateValue(reassignedOk, false, t)
	utils.ValidateValue(err.Message, "given identifier 'x' is already declared", t)
	utils.ValidateObject(obj, value, t)
}
