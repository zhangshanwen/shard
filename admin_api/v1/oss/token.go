package oss

import (
	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/tools"
)

func Token(c *service.AdminContext) (resp service.Res) {
	oss := tools.NewOss()
	r := response.ImageToken{}
	r.Token = oss.NewToken(conf.C.Oss.AdminBuket)
	resp.Data = r
	return
}
