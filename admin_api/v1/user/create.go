package user

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Create(c *service.AdminContext) (r service.Res) {
	p := param.Register{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ResCode = code.ParamsError
		return
	}
	var (
		m = model.User{
			Username: p.Username,
		}
		resp = &response.UserInfo{}
		tx   = db.G.Begin()
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
	if r.Err = tx.Model(&m).Where(&m).Count(&count).Error; r.Err != nil {
		return
	}
	if count > 0 {
		r.UsernameIsExisted()
		return
	}
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	if r.Err = m.SetPassword(p.Password); r.Err != nil {
		r.SetPasswordError()
		return
	}
	if r.Err = tx.Create(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = copier.Copy(resp, &m); r.Err != nil {
		r.CopierError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("创建用户 id:%v username:%v", m.Id, m.Username), model.OperateLogTypeAdd)
	return
}
