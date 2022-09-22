package permission

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Edit(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.RolePermissionEdit{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  = model.Role{}
		tx = c.Tx
	)

	defer func() {
		if r.Err == nil {
			r.Data = m
		}
	}()
	if r.Err = tx.First(&m, pId.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	var permissions []model.Permission
	for _, routeId := range p.PermissionIds {
		permissions = append(permissions, model.Permission{
			BaseModel: model.BaseModel{Id: routeId},
		})
	}
	if r.Err = tx.Model(&m).Association("Permissions").Replace(&permissions); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = tx.Save(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLogUpdate(tx, module, fmt.Sprintf("%v:%v", m.Id, m.Name))
	return
}
