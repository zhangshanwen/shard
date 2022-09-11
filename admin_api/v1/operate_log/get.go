package operate_log

import (
	"strings"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	p := param.LogRecords{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx   = c.Tx
		ms   []model.OperateLog
		resp = response.LogResponse{}
		m    model.OperateLog
	)
	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	g := tx.Model(&m)
	if p.Types != "" {
		g = g.Where("type in ?", strings.Split(p.Types, ","))
	}
	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Preload("Admin").Preload("Admin.Role").Find(&ms).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = copier.Copy(&resp.List, &ms); r.Err != nil {
		r.CopierError()
		return
	}
	return
}
