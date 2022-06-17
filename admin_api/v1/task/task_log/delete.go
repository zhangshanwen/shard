package task_log

import (
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Delete(c *service.AdminContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.TaskDelete{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  model.TaskLog
		tx = db.G.Begin()
	)
	defer func() {
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if p.All {
		r.Err = tx.Where("task_id=?", pId.Id).Delete(&m).Error
	} else {
		r.Err = tx.Where("id in ?", p.Ids).Delete(&m).Error
	}
	if r.Err != nil {
		r.DBError()
		return
	}
	return
}
