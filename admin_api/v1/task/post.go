package task

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/initialize/task"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

func Post(c *service.AdminContext) (resp service.Res) {
	p := param.TaskCreate{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	t := model.Task{
		Name:   p.Name,
		Spec:   p.Spec,
		Status: model.StatusNormal,
	}
	g := db.G.Begin()
	defer func() {
		if resp.Err != nil {
			g.Rollback()
			return
		}
		g.Commit()
	}()
	if resp.Err = g.Save(&t).Error; resp.Err != nil {
		return
	}
	resp.Err = task.T.Add(&t)
	return
}
