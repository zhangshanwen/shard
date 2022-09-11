package task

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	p := param.Task{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx    = c.Tx
		resp  = response.TaskResponse{}
		m     model.Task
		tasks []model.Task
	)
	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	g := tx.Model(&m)
	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Preload("Host").Find(&tasks).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = copier.Copy(&resp.List, &tasks); r.Err != nil {
		r.CopierError()
		return
	}
	return
}
