package admin

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func ChangePassword(c *service.AdminContext) (resp service.Res) {
	p := param.PasswordParam{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	g := db.G
	if resp.Err = c.Admin.SetPassword(p.Password); resp.Err != nil {
		return
	}
	if resp.Err = g.Model(&c.Admin).Updates(&model.Admin{
		Password: c.Admin.Password,
	}).Error; resp.Err != nil {
		return
	}
	resp.Data = response.PasswordResponse{Password: p.Password}
	return
}
