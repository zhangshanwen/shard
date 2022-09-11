package admin

import (
	"fmt"
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Create(c *service.AdminTxContext) (r service.Res) {
	p := param.Register{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m    = model.Admin{Username: p.Username}
		tx   = c.Tx
		resp response.AdminInfo
	)
	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	var count int64
	if r.Err = tx.Model(&m).Where("id=?", m.Id).Count(&count).Error; r.Err != nil {
		r.DBError()
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
	if r.Err = copier.Copy(&resp, &m); r.Err != nil {
		r.CopierError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("创建管理员 %v", m.Username), model.OperateLogTypeAdd)
	return
}
