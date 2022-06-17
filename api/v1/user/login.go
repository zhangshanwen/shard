package user

import (
	"github.com/jinzhu/copier"
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
	"time"
)

func Login(c *service.Context) (r service.Res) {
	p := param.Login{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ResCode = code.ParamsError
		return
	}
	var (
		m    = model.User{Username: p.Username}
		tx   = db.G
		resp = response.UserInfo{}
	)
	defer func() {
		r.Data = resp
	}()
	if r.Err = tx.Where(&m).Preload("Wallet").First(&m).Error; r.Err != nil {
		return
	}
	if !m.CheckPassword(p.Password) {
		r.ResCode = code.ActPWdError
	}
	if r.Err = copier.Copy(&r, &m); r.Err != nil {
		return
	}
	var token string
	token, r.Err = tools.CreateToken(m.Id)
	if r.Err != nil {
		return
	}
	if r.Err = tx.Model(&m).Updates(&model.User{
		LastLoginTime: time.Now().Unix(),
	}).Error; r.Err != nil {
		return
	}
	if m.Wallet == nil {
		m.Wallet = &model.Wallet{Uid: m.Id}
		if r.Err = tx.Create(&m.Wallet).Error; r.Err != nil {
			return
		}
	}
	resp.Balance = m.Wallet.Balance
	resp.Authorization = token
	return
}
