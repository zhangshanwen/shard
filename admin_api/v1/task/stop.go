package task

import (
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/initialize/task"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Stop(c *service.AdminTxContext) (r service.Res) {
	p := param.UriId{}
	if r.Err = c.BindUri(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  = model.Task{}
		tx = c.Tx
	)
	if r.Err = tx.First(&m, p.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}
	if r.Err = tx.Model(&m).Where("id=?", p.Id).Update("status", model.StatusStop).Error; r.Err != nil {
		r.DBError()
		return
	}
	task.T.StopAll(p.Id)
	return
}
