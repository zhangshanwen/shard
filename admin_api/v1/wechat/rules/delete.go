package rules

import (
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
		m  model.Rule
		tx = c.Tx
	)
	m.Id = pId.Id

	if r.Err = tx.First(&m, m.Id).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = tx.Model(&m).Delete(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
