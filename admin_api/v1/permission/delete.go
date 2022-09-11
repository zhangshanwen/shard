package permission

import (
	"fmt"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Delete(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ResCode = code.ParamsError
		return
	}
	var (
		m  model.Permission
		tx = c.Tx
	)
	m.Id = pId.Id
	if r.Err = tx.Select("Routes", "Roles").Delete(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("删除权限 id:%v name:%v", m.Id, m.Name), model.OperateLogTypeDel)
	return
}
