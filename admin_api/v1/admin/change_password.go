package admin

import (
	"fmt"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func ChangePassword(c *service.AdminTxContext) (r service.Res) {
	p := param.PasswordParam{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx   = c.Tx
		resp = response.PasswordResponse{}
	)
	resp.Password = p.Password
	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	if r.Err = c.Admin.SetPassword(p.Password); r.Err != nil {
		r.SetPasswordError()
		return
	}
	if r.Err = tx.Model(&c.Admin).Updates(&model.Admin{
		Password: c.Admin.Password,
	}).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLogUpdate(tx, module, fmt.Sprintf("change password:%v", p.Password))
	return
}
