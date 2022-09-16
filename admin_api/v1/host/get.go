package host

import (
	"fmt"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	p := param.Host{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ResCode = code.ParamsError
		return
	}
	var (
		tx   = c.Tx
		resp response.HostResponse
		m    model.Host
	)

	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	g := tx.Model(&m)
	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = g.Find(&resp.List).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLogSelect(tx, fmt.Sprintf("查看主机列表"))
	return
}
