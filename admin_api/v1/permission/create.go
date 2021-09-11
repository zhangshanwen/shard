package permission

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

func Create(c *service.AdminContext) (resp service.Res) {
	p := param.Permission{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	m := model.Permission{Name: p.Name}
	g := db.G
	g = g.Begin()
	defer func() {
		if resp.Err == nil {
			g.Commit()
		} else {
			g.Rollback()
		}
	}()
	if resp.Err = copier.Copy(&m, &p); resp.Err != nil {
		return
	}
	for _, i := range p.RouteIds {
		m.Routes = append(m.Routes, model.Route{
			BaseModel: model.BaseModel{Id: i},
		})
	}
	if resp.Err = g.Create(&m).Error; resp.Err != nil {
		return
	}

	resp.Data = m
	return
}
