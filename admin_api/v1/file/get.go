package file

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (resp service.Res) {
	p := param.FileRecords{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	m := model.FileRecord{Uid: c.Admin.Id}
	g := db.G.Model(&m).Where(&m)
	r := response.FileResponse{}
	if resp.Err = db.FindByPagination(g, &p.Pagination, &r.Pagination); resp.Err != nil {
		return
	}
	if resp.Err = g.Find(&r.List).Error; resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
