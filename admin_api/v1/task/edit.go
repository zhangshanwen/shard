package task

import (
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/initialize/task"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Edit(c *service.AdminContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.TaskSave{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  = model.Task{}
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

	if r.Err = tx.First(&m, pId.Id).Error; r.Err != nil {
		r.NotChange()
		return
	}
	if m.Status >= model.StatusIdle {
		r.TaskIsRunning()
		return
	}
	if r.Err = m.Verify(); r.Err != nil {
		r.TaskVerifyError()
		return
	}
	if r.Err = tx.Model(&m).Updates(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = task.T.Add(tx, m); r.Err != nil {
		r.TaskAddFailed()
	}
	return
}
