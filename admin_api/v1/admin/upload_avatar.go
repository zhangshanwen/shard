package admin

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

func UploadAvatar(c *service.AdminContext) (resp service.Res) {
	pId := param.UriId{}
	if resp.Err = c.BindUri(&pId); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	p := param.AdminUploadAvatar{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	admin := model.Admin{}
	g := db.G
	if resp.Err = g.First(&admin, pId.Id).Error; resp.Err != nil {
		return
	}
	if resp.Err = g.Model(&admin).Updates(&model.Admin{
		Avatar: p.Avatar,
	}).Error; resp.Err != nil {
		return
	}
	resp.Data = admin
	return
}
