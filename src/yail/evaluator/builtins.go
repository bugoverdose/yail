package evaluator

import (
	"yail/object"
)

const (
	LEN      = "len"
	HEAD     = "head"
	TAIL     = "tail"
	PUSH     = "push"
	PUSHLEFT = "pushleft"
	POP      = "pop"
	POPLEFT  = "popleft"

	INVALID_TYPE_EXCEPTION_MESSAGE = "%s(%s) not supported"
	INVALID_ARGUMENT_COUNT_MESSAGE = "wrong number of arguments: expected %d, but received %d"
)

var builtinFunctions = map[string]*object.Builtin{
	LEN: {
		Fn: func(args ...object.Object) object.Object {
			ok, err := validateArgCount(args, 1)
			if !ok {
				return err
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return object.NewError(INVALID_TYPE_EXCEPTION_MESSAGE, LEN, arg.Type())
			}
		},
	},
	HEAD: {
		Fn: func(args ...object.Object) object.Object {
			ok, err := validateArrayFunctionArguments(HEAD, 1, args)
			if !ok {
				return err
			}
			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return object.NULL
		},
	},
	TAIL: {
		Fn: func(args ...object.Object) object.Object {
			ok, err := validateArrayFunctionArguments(TAIL, 1, args)
			if !ok {
				return err
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return object.NULL
		},
	},
	PUSH: {
		Fn: func(args ...object.Object) object.Object {
			ok, err := validateArrayFunctionArguments(PUSH, 2, args)
			if !ok {
				return err
			}
			arr := args[0].(*object.Array)
			arr.Elements = append(arr.Elements, args[1])
			return object.NULL
		},
	},
	PUSHLEFT: {
		Fn: func(args ...object.Object) object.Object {
			ok, err := validateArrayFunctionArguments(PUSHLEFT, 2, args)
			if !ok {
				return err
			}
			arr := args[0].(*object.Array)
			arr.Elements = append([]object.Object{args[1]}, arr.Elements...)
			return object.NULL
		},
	},
	POP: {
		Fn: func(args ...object.Object) object.Object {
			ok, err := validateArrayFunctionArguments(POP, 1, args)
			if !ok {
				return err
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				arr.Elements = arr.Elements[0 : length-1]
				return arr
			}
			return object.NULL
		},
	},
	POPLEFT: {
		Fn: func(args ...object.Object) object.Object {
			ok, err := validateArrayFunctionArguments(POPLEFT, 1, args)
			if !ok {
				return err
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				arr.Elements = arr.Elements[1:length]
				return arr
			}
			return object.NULL
		},
	},
}

func validateArrayFunctionArguments(functionName string, expectedArgCount int, args []object.Object) (bool, *object.Error) {
	ok, err := validateArgCount(args, expectedArgCount)
	if !ok {
		return false, err
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return false, object.NewError(INVALID_TYPE_EXCEPTION_MESSAGE, functionName, args[0].Type())
	}
	return true, nil
}

func validateArgCount(args []object.Object, expectedArgCount int) (bool, *object.Error) {
	if len(args) != expectedArgCount {
		return false, object.NewError(INVALID_ARGUMENT_COUNT_MESSAGE, expectedArgCount, len(args))
	}
	return true, nil
}
