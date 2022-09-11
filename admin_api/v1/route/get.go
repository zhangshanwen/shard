package route

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	var (
		resp = response.RouteResponse{}
		tx   = c.Tx
		m    model.Route
	)
	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	if r.Err = tx.Model(&m).Find(&resp.List).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("获取路由列表"), model.OperateLogTypeSelect)
	return
}
