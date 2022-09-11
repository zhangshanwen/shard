package user

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Delete(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  model.User
		tx = c.Tx
	)
	m.Id = pId.Id

	if r.Err = tx.Model(&m).Delete(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("删除用户 id:%v username:%v", m.Id, m.Username), model.OperateLogTypeDel)
	return
}
