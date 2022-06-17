package host

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Post(c *service.AdminContext) (r service.Res) {
	p := param.HostSave{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}

	m := model.Host{}
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
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
	if r.Err = tx.Save(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("创建主机 id:%v name:%v connect_type%v", m.Id, m.Name, m.ConnectType), model.OperateLogTypeAdd)
	return
}
