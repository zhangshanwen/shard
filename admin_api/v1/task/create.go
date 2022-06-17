package task

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/initialize/task"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Post(c *service.AdminContext) (r service.Res) {
	p := param.TaskSave{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	t := model.Task{}
	if r.Err = copier.Copy(&t, &p); r.Err != nil {
		r.CopierError()
		return
	}
	if r.Err = t.Verify(); r.Err != nil {
		r.TaskVerifyError()
		return
	}
	var (
		tx = db.G.Begin()
	)
	defer func() {
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	h := model.Host{}
	if r.Err = tx.First(&h, p.HostId).Error; r.Err != nil {
		r.NotFound()
		return
	}
	if r.Err = tx.Save(&t).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = task.T.Add(tx, t); r.Err != nil {
		r.TaskAddFailed()
	}
	return
}
