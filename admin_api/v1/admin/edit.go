package admin

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func Edit(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.AdminEdit{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  = model.Admin{}
		tx = c.Tx
	)
	defer func() {
		if r.Err == nil {
			r.Data = m
		}
	}()
	if r.Err = tx.First(&m, pId.Id).Error; r.Err != nil {
		r.DBError()
		return
	}
	diff := tools.DiffStruct(p, m, "json")
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	if r.Err = tx.Model(&m).Updates(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLogUpdate(tx, fmt.Sprintf("编辑管理员 id:%v %v ", m.Id, diff))
	return
}
