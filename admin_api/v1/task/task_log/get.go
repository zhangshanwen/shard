package task_log

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.ShouldBindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.TaskLog{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		resp = response.TaskLogResponse{}
		tx   = c.Tx
		m    model.TaskLog
		ms   []model.TaskLog
	)
	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	g := tx.Model(&m).Where("task_id=?", pId.Id)

	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Find(&ms).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = copier.Copy(&resp.List, &ms); r.Err != nil {
		r.DBError()
		return
	}
	return
}
