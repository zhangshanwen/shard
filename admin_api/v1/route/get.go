package route

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (r service.Res) {
	var (
		resp = response.RouteResponse{}
		tx   = db.G.Begin()
		m    model.Route
	)
	defer func() {
		r.Data = resp
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if r.Err = tx.Model(&m).Find(&resp.List).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("获取路由列表"), model.OperateLogTypeSelect)
	return
}
