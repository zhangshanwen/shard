package permission

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		ms   []model.Permission
		tx   = db.G.Begin()
		resp = response.RolePermissionResponse{}
		m    model.Permission
	)

	defer func() {
		r.Data = resp
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	g := tx.Model(&m).Where(" parent_id = 0 ")
	// 使用signal 防止其他查询也使用AfterFind 钩子
	if r.Err = g.Set(common.PermissionRoleFindChildrenId, pId.Id).Preload("Roles", "id = ? ", pId.Id).
		Find(&ms).Error; r.Err != nil {
		r.DBError()
		return
	}

	if r.Err = copier.Copy(&resp.List, &ms); r.Err != nil {
		r.CopierError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("获取角色(%v)权限", pId.Id), model.OperateLogTypeSelect)
	return
}
