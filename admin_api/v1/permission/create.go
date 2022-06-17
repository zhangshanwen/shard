package permission

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Create(c *service.AdminContext) (r service.Res) {
	p := param.Permission{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  = model.Permission{Name: p.Name}
		tx = db.G.Begin()
	)

	defer func() {
		r.Data = m
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	for _, i := range p.RouteIds {
		m.Routes = append(m.Routes, model.Route{
			BaseModel: model.BaseModel{Id: i},
		})
	}
	if r.Err = tx.Create(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("创建权限 id:%v name:%v", m.Id, m.Name), model.OperateLogTypeAdd)
	return
}
