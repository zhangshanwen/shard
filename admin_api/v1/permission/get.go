package permission

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	p := param.PermissionRecords{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		ms   []model.Permission
		m    model.Permission
		resp = response.PermissionResponse{}
		tx   = c.Tx
	)

	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	g := tx.Model(&m).Where(" parent_id = 0 ")
	// 使用signal 防止其他查询也使用AfterFind 钩子
	if r.Err = g.Set(common.PermissionFindChildren, true).Preload("Routes").Find(&ms).Error; r.Err != nil {
		r.DBError()
		return
	}

	if r.Err = copier.Copy(&resp.List, &ms); r.Err != nil {
		r.CopierError()
		return
	}
	c.SaveLogSelect(tx, module, fmt.Sprintf("cat list"))
	return
}
