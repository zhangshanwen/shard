package file

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (r service.Res) {
	p := param.FileRecords{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m    = model.FileRecord{Uid: c.Admin.Id}
		tx   = db.G.Begin()
		resp = response.FileResponse{}
	)

	defer func() {
		r.Data = resp
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	g := tx.Model(&m).Where(&m)
	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Find(&resp.List).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("查看文件列表"), model.OperateLogTypeSelect)
	return
}
