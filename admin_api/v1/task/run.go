package task

import (
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/initialize/task"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Run(c *service.AdminTxContext) (r service.Res) {
	p := param.UriId{}
	if r.Err = c.BindUri(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx = c.Tx
	)
	t := model.Task{}
	if r.Err = tx.First(&t, p.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}
	if t.Status != model.StatusRunning {
		r.TaskIsNotRunning()
		return
	}
	if r.Err = task.T.Run(t); r.Err != nil {
		r.TaskRunFailed()
		return
	}
	return
}
