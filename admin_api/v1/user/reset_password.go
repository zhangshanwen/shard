package user

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func ResetPassword(c *service.AdminContext) (resp service.Res) {
	pId := param.UriId{}
	if resp.Err = c.BindUri(&pId); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	user := model.User{}
	g := db.G
	if resp.Err = g.First(&user, pId.Id).Error; resp.Err != nil {
		return
	}
	if resp.Err = user.SetPassword(conf.C.ResetPassword); resp.Err != nil {
		return
	}
	if resp.Err = g.Model(&user).Updates(&model.User{
		Password: user.Password,
	}).Error; resp.Err != nil {
		return
	}
	resp.Data = response.PasswordResponse{Password: conf.C.ResetPassword}
	return
}
