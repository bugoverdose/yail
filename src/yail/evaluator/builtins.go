package evaluator

import (
	"yail/object"
)

var builtinFunctions = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewError("wrong number of arguments: expected 1, but received %d", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return object.NewError("len(%s) not supported", arg.Type())
			}
		},
	},
}
