package task

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/initialize/task"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

func Stop(c *service.AdminContext) (resp service.Res) {
	p := param.UriId{}
	if resp.Err = c.BindUri(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	g := db.G.Begin().Model(model.Task{})
	defer func() {
		if resp.Err != nil {
			g.Rollback()
			return
		}
		g.Commit()
	}()
	t := model.Task{}
	if resp.Err = g.First(&t, p.Id).Error; resp.Err != nil {
		return
	}
	if resp.Err = g.Where("id=?", p.Id).Update("status", model.StatusStop).Error; resp.Err != nil {
		return
	}
	resp.Err = task.T.Stop(p.Id)
	return
}
