package user

import (
	"github.com/jinzhu/copier"
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
	"time"
)

func Login(c *service.Context) (resp service.Res) {
	p := param.Login{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	user := model.User{Username: p.Username}
	g := db.G
	if resp.Err = g.Where(&user).Preload("Wallet").First(&user).Error; resp.Err != nil {
		return
	}
	if !user.CheckPassword(p.Password) {
		resp.ResCode = code.ActPWdError
	}
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &user); resp.Err != nil {
		return
	}
	var token string
	token, resp.Err = tools.CreateToken(user.Id)
	if resp.Err != nil {
		return
	}
	if resp.Err = g.Model(&user).Updates(&model.User{
		LastLoginTime: time.Now().Unix(),
	}).Error; resp.Err != nil {
		return
	}
	if user.Wallet == nil {
		user.Wallet = &model.Wallet{Uid: user.Id}
		if resp.Err = g.Create(&user.Wallet).Error; resp.Err != nil {
			return
		}
	}
	r.Balance = user.Wallet.Balance
	r.Authorization = token
	resp.Data = r
	return
}
