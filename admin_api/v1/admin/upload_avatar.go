package admin

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

func UploadAvatar(c *service.AdminContext) (resp service.Res) {
	p := param.AdminUploadAvatar{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	g := db.G
	if resp.Err = g.Model(&c.Admin).Updates(&model.Admin{
		Avatar: p.Avatar,
	}).Error; resp.Err != nil {
		return
	}
	resp.Data = c.Admin
	return
}
