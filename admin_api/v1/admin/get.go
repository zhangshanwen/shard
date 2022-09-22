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

func Get(c *service.AdminTxContext) (r service.Res) {
	p := param.AdminRecords{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.DBError()
		return
	}
	var (
		m    model.Admin
		ms   []model.Admin
		tx   = c.Tx
		resp = response.AdminResponse{}
	)
	defer func() {
		if r.Err == nil {
			r.Data = resp
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
	c.SaveLogSelect(tx, module, "cat list")
	return
}
