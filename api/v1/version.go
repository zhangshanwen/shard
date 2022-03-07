package v1

import (
	"github.com/zhangshanwen/shard/initialize/service"
)

var version string

func Version(c *service.Context) (resp service.Res) {
	resp.Data = version
	return
}
