package user

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
	p := param.UserRecords{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	g := db.G.Model(&model.User{}).Preload("Wallet")
	if p.Username != "" {
		g = g.Where(model.User{Username: fmt.Sprintf("%%%s%%", p.Username)})
	}
	r := response.UsersResponse{}
	if resp.Err = db.FindByPagination(g, &p.Pagination, &r.Pagination); resp.Err != nil {
		return
	}
	var ms []model.User
	if resp.Err = g.Find(&ms).Error; resp.Err != nil {
		return
	}
	if resp.Err = copier.Copy(&r.List, &ms); resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
