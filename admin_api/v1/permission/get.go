package permission

import (
	"github.com/jinzhu/copier"
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (resp service.Res) {
	p := param.PermissionRecords{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	m := model.Permission{}
	var ms []model.Permission
	g := db.G.Model(&m).Where(" parent_id = 0 ")
	r := response.PermissionResponse{}
	// 使用signal 防止其他查询也使用AfterFind 钩子
	if resp.Err = g.Set(common.PermissionFindChildren, true).Preload("Routes").Find(&ms).Error; resp.Err != nil {
		return
	}

	if resp.Err = copier.Copy(&r.List, &ms); resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
