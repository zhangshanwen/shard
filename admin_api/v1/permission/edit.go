package permission

import (
	"github.com/jinzhu/copier"
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

func Edit(c *service.AdminContext) (resp service.Res) {
	pId := param.UriId{}
	if resp.Err = c.BindUri(&pId); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	p := param.Permission{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	m := model.Permission{}
	g := db.G
	g = g.Begin()
	defer func() {
		if resp.Err == nil {
			g.Commit()
		} else {
			g.Rollback()
		}
	}()
	if resp.Err = g.First(&m, pId.Id).Error; resp.Err != nil {
		return
	}
	if resp.Err = copier.Copy(&m, &p); resp.Err != nil {
		return
	}
	var routes []model.Route
	for _, routeId := range p.RouteIds {
		routes = append(routes, model.Route{
			BaseModel: model.BaseModel{Id: routeId},
		})
	}
	if resp.Err = g.Model(&m).Association("Routes").Replace(&routes); resp.Err != nil {
		return
	}
	if resp.Err = g.Save(&m).Error; resp.Err != nil {
		return
	}
	resp.Data = m
	return
}
