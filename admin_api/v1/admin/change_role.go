package admin

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func ChangeRole(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.AdminChangeRole{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx = c.Tx
		m  = model.Admin{}
	)
	defer func() {
		if r.Err == nil {
			r.Data = m.Role
		}
	}()
	if pId.Id == c.Admin.Id {
		m = c.Admin
	} else {
		if r.Err = tx.First(&m, pId).Error; r.Err != nil {
			r.DBError()
			return
		}
	}
	role := model.Role{}
	if r.Err = tx.First(&role, p.RoleId).Error; r.Err != nil {
		r.DBError()
		return
	}
	m.RoleId = p.RoleId
	if r.Err = tx.Model(&m).Updates(&m).Error; r.Err != nil {
		r.DBError()
		return
	}

	c.SaveLog(tx, fmt.Sprintf("修改管理员 %v,角色%v->%v", m.Username, c.Admin.Role.Name, role.Name), model.OperateLogTypeUpdate)
	return
}
