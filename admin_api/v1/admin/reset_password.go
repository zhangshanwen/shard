package admin

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func ResetPassword(c *service.AdminContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m    model.Admin
		tx   = db.G.Begin()
		resp = response.PasswordResponse{Password: conf.C.ResetPassword}
	)
	defer func() {
		r.Data = resp
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if r.Err = tx.First(&m, pId.Id).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = m.SetPassword(conf.C.ResetPassword); r.Err != nil {
		r.SetPasswordError()
		return
	}
	if r.Err = tx.Model(&m).Updates(m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("重置管理员id:%v,username:%v 密码 ->%v", m.Id,
		m.Username, conf.C.ResetPassword), model.OperateLogTypeUpdate)

	return
}
