package role

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Create(c *service.AdminContext) (r service.Res) {
	p := param.Role{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  = model.Role{Name: p.Name}
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
	if r.Err = tx.Model(&m).Where(&m).Count(&count).Error; r.Err != nil {
		r.DBError()
		return
	}
	if count > 0 {
		r.NameIsExisted()
		return
	}
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	if r.Err = tx.Create(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("创建角色 id:%v name:%v", m.Id, m.Name), model.OperateLogTypeAdd)
	return
}
