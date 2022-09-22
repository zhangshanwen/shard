package operate_log

import (
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
	"strings"
)

func Del(c *service.AdminTxContext) (r service.Res) {
	p := param.LogDel{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx = c.Tx
	)
	if r.Err = tx.Delete(&model.OperateLog{}, " `id` in ? ", strings.Split(p.Ids, ",")).Error; r.Err != nil {
		r.CopierError()
		return
	}
	return
}
