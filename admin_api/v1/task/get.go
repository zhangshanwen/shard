package task

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (resp service.Res) {
	p := param.Task{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	g := db.G.Model(&model.Task{})
	r := response.TaskResponse{}
	if resp.Err = db.FindByPagination(g, &p.Pagination, &r.Pagination); resp.Err != nil {
		return
	}
	if resp.Err = g.Find(&r.List).Error; resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
