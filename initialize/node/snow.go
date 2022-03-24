package node

import (
	"github.com/bwmarrin/snowflake"
)

type (
	SnowNode struct {
		node *snowflake.Node
	}
)

func createSnowNode(node int) (s *SnowNode, err error) {
	s = new(SnowNode)
	s.node, err = snowflake.NewNode(int64(node))
	return
}

func (s SnowNode) Generate() int64 {
	return s.node.Generate().Int64()
}
