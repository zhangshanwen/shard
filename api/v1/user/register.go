package user

import (
	"errors"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Register(c *service.Context) (r service.Res) {
	p := param.Register{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ResCode = code.ParamsError
		return
	}
	var (
		user = model.User{Username: p.Username}
		tx   = db.G.Begin()
		resp = response.UserInfo{}
	)

	defer func() {
		r.Data = resp
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	var count int64
	if r.Err = tx.Where(&user).First(&user).Count(&count).Error; r.Err != nil {
		return
	}
	if count > 0 {
		r.Err = errors.New("username is existed")
		r.ResCode = code.UsernameIsExisted
		return
	}
	if r.Err = copier.Copy(&user, &p); r.Err != nil {
		return
	}
	if r.Err = user.SetPassword(p.Password); r.Err != nil {
		return
	}
	if r.Err = tx.Create(&user).Error; r.Err != nil {
		return
	}

	if r.Err = copier.Copy(&resp, &user); r.Err != nil {
		return
	}
	return
}
