package object

type BuiltinFunction func(args ...Object) Object

const BUILTIN_OBJ = "BUILTIN"

type Builtin struct {
	Fn   BuiltinFunction
	name string
}

func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

func (b *Builtin) Inspect() string {
	return "builtin function " + b.name
}
