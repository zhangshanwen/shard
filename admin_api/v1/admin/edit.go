package admin

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

func Edit(c *service.AdminContext) (resp service.Res) {
	pId := param.UriId{}
	if resp.Err = c.BindUri(&pId); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	p := param.AdminEdit{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	m := model.Admin{}
	g := db.G
	if resp.Err = g.First(&m, pId.Id).Error; resp.Err != nil {
		return
	}
	if resp.Err = copier.Copy(&m, &p); resp.Err != nil {
		return
	}
	if resp.Err = g.Model(&m).Updates(&m).Error; resp.Err != nil {
		return
	}
	resp.Data = m
	return
}
