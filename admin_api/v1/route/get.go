package route

import (
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
)

func Get(c *service.AdminContext) (resp service.Res) {
	g := db.G.Model(&model.Route{})
	r := response.RouteResponse{}
	if resp.Err = g.Find(&r.List).Error; resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
