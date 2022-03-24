package version

import (
	"github.com/zhangshanwen/shard/initialize/service"
)

var buildTime string
var git string

func Get(c *service.AdminContext) (resp service.Res) {
	resp.Data = map[string]string{
		"build_time": buildTime,
		"git":        git,
	}
	return
}
