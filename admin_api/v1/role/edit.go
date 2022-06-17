package role

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func Edit(c *service.AdminContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.Role{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  = model.Role{}
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
	var count int64
	if r.Err = tx.Model(&m).Where("id != ? ", pId.Id).Count(&count).Error; r.Err != nil {
		r.DBError()
		return
	}
	if count > 0 {
		r.NameIsExisted()
		return
	}
	if r.Err = tx.First(&m, pId.Id).Error; r.Err != nil {
		r.DBError()
		return
	}
	if p.Name == m.Name {
		r.NotChange()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("修改角色 id:%v %v ", m.Id, tools.DiffStruct(p, m, "json")), model.OperateLogTypeUpdate)
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	if r.Err = tx.Model(&m).Updates(&m).Error; r.Err != nil {
		r.DBError()
		return
	}

	return
}
