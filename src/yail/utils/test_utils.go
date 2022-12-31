package utils

import (
	"testing"
	"yail/object"
)

func ValidateObject(actual, expected object.Object, t *testing.T) {
	ValidateValue(actual.Type(), expected.Type(), t)
	switch actual := actual.(type) {
	case *object.Integer:
		ValidateValue(actual.Value, expected.(*object.Integer).Value, t)
	case *object.Boolean:
		ValidateValue(actual.Value, expected.(*object.Boolean).Value, t)
	case *object.Error:
		ValidateValue(actual.Message, expected.(*object.Error).Message, t)
	}
}

func ValidateMatchAnyValue[T comparable](actual T, expectedList []T, t *testing.T) {
	for _, expected := range expectedList {
		if actual == expected {
			return
		}
	}
	t.Errorf("expected %+v to be one of %+v", actual, expectedList)
}

func ValidateValue[T comparable](actual, expected T, t *testing.T) {
	if actual != expected {
		t.Errorf("expected %+v to be %+v", actual, expected)
	}
}
