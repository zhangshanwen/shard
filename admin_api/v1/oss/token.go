package oss

import (
	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/tools"
)

func Token(c *service.AdminContext) (r service.Res) {
	oss := tools.NewOss()
	resp := response.ImageToken{}
	defer func() {
		r.Data = resp
	}()
	resp.Token = oss.NewToken(conf.C.Oss.AdminBuket)
	return
}
