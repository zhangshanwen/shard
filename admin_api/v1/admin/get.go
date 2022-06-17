package admin

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (r service.Res) {
	p := param.AdminRecords{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.DBError()
		return
	}
	var (
		m    model.Admin
		ms   []model.Admin
		tx   = db.G.Begin()
		resp = response.AdminResponse{}
	)
	defer func() {
		r.Data = resp
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	g := tx.Model(&m)
	if p.Username != "" {
		m.Username = fmt.Sprintf("%%%s%%", p.Username)
		g = g.Where(&m)
	}

	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Preload("Role").Find(&ms).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = copier.Copy(&resp.List, &ms); r.Err != nil {
		r.CopierError()
		return
	}
	c.SaveLog(tx, "查询管理员列表", model.OperateLogTypeSelect)
	return
}
