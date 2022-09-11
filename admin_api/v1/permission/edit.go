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
	p := param.Permission{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  model.Permission
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
	var routes []model.Route
	for _, routeId := range p.RouteIds {
		routes = append(routes, model.Route{
			BaseModel: model.BaseModel{Id: routeId},
		})
	}
	if r.Err = tx.Model(&m).Association("Routes").Replace(&routes); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = tx.Save(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("修改权限 id:%v name:%v", m.Id, m.Name), model.OperateLogTypeUpdate)
	return
}
