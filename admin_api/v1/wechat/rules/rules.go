package rules

import (
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

// Rules 返回用户返回规则
func Rules(c *service.AdminTxContext) (r service.Res) {
	p := param.Rule{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		resp response.RulesResponse
		tx   = c.Tx
		m    = model.Rule{Uid: c.Admin.Id}
	)
	var ()

	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	g := tx.Model(&m).Where(&m)
	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Find(&resp.List).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
