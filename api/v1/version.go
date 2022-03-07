package v1

import (
	"github.com/zhangshanwen/shard/initialize/service"
)

var buildTime string
var git string

func Version(c *service.Context) (resp service.Res) {
	resp.Data = map[string]string{
		"build_time": buildTime,
		"git":        git,
	}
	return
}
