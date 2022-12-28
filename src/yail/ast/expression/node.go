package expression

import (
	"yail/ast/node"
)

type Expression interface {
	node.Node
	expressionNode()
}
