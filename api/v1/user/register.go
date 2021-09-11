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

func Register(c *service.Context) (resp service.Res) {
	p := param.Register{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	user := model.User{Username: p.Username}
	g := db.G
	g = g.Begin()
	defer func() {
		if resp.Err == nil {
			g.Commit()
		} else {
			g.Rollback()
		}
	}()
	var count int64
	if resp.Err = g.Where(&user).First(&user).Count(&count).Error; resp.Err != nil {
		return
	}
	if count > 0 {
		resp.Err = errors.New("username is existed")
		resp.ResCode = code.UsernameIsExisted
		return
	}
	if resp.Err = copier.Copy(&user, &p); resp.Err != nil {
		return
	}
	if resp.Err = user.SetPassword(p.Password); resp.Err != nil {
		return
	}
	if resp.Err = g.Create(&user).Error; resp.Err != nil {
		return
	}
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &user); resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
