package user

import (
	"errors"
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func Create(c *service.AdminContext) (resp service.Res) {
	p := param.Register{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	m := model.User{Username: p.Username}
	g := db.G
	var count int64
	if resp.Err = g.Model(&m).Where(&m).Count(&count).Error; resp.Err != nil {
		return
	}
	if count > 0 {
		resp.Err = errors.New("username is existed")
		resp.ResCode = code.UsernameIsExisted
		return
	}
	if resp.Err = copier.Copy(&m, &p); resp.Err != nil {
		return
	}
	if resp.Err = m.SetPassword(p.Password); resp.Err != nil {
		return
	}
	if resp.Err = g.Create(&m).Error; resp.Err != nil {
		return
	}
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &m); resp.Err != nil {
		return
	}
	resp.Data = m
	return
}
