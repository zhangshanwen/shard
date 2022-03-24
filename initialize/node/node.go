package node

import "github.com/sirupsen/logrus"

type (
	Node interface {
		Generate() int64
	}
)

var (
	N       Node
	nodeNum = 1
)

func InitNode() {
	var err error
	if N, err = createSnowNode(nodeNum); err != nil {
		logrus.Panic(err)
	}
	return
}
