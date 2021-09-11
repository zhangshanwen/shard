package admin

import (
	"fmt"
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (resp service.Res) {
	p := param.AdminRecords{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	m := model.Admin{}
	var ms []model.Admin
	g := db.G.Model(&m)
	if p.Username != "" {
		m.Username = fmt.Sprintf("%%%s%%", p.Username)
		g = g.Where(&m)
	}
	r := response.AdminResponse{}
	if resp.Err = db.FindByPagination(g, &p.Pagination, &r.Pagination); resp.Err != nil {
		return
	}
	if resp.Err = g.Preload("Role").Find(&ms).Error; resp.Err != nil {
		return
	}
	if resp.Err = copier.Copy(&r.List, &ms); resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
