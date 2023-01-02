package statement

import "yail/ast/node"

type Statement interface {
	node.Node
	statementNode()
}
