package role

import (
	"errors"
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
	p := param.Role{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	m := model.Role{}
	g := db.G
	g = g.Begin()
	defer func() {
		if resp.Err == nil {
			g.Commit()
		} else {
			g.Rollback()
		}
	}()
	var count int64
	if resp.Err = g.Model(&m).Where("id != ? ", pId.Id).Count(&count).Error; resp.Err != nil {
		return
	}
	if count > 0 {
		resp.ResCode = code.NameIsExisted
		resp.Err = errors.New("NameIsExisted")
		return
	}
	if resp.Err = g.First(&m, pId.Id).Error; resp.Err != nil {
		return
	}
	if p.Name == m.Name {
		resp.ResCode = code.NotChange
		resp.Err = errors.New("NotChange")
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
