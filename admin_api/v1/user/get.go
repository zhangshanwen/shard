package user

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (r service.Res) {
	p := param.UserRecords{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx   = db.G.Begin()
		resp = response.UsersResponse{}
		m    model.User
		ms   []model.User
	)
	defer func() {
		r.Data = resp
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	g := tx.Model(&m).Preload("Wallet")
	if p.Username != "" {
		g = g.Where(model.User{Username: fmt.Sprintf("%%%s%%", p.Username)})
	}
	if r.Err = db.FindByPagination(g, &p.Pagination, &resp.Pagination); r.Err != nil {
		return
	}
	if r.Err = g.Find(&ms).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = copier.Copy(&resp.List, &ms); r.Err != nil {
		r.CopierError()
		return
	}
	c.SaveLog(tx, "查询用户列表", model.OperateLogTypeSelect)
	return
}
