package environment

import "yail/object"

type value struct {
	data      object.Object
	isMutable bool
}

func newValue(data object.Object, mutable bool) value {
	return value{data: data, isMutable: mutable}
}
