package operate_log

import (
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/model"
)

func Empty(c *service.AdminTxContext) (r service.Res) {
	var (
		tx = c.Tx
	)
	if r.Err = tx.Where("id>0").Delete(&model.OperateLog{}).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
