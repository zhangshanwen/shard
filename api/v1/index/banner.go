package index

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/response"
)

func Banner(c *service.Context) (resp service.Res) {
	r := response.BannerResponse{}
	if resp.Err = copier.Copy(&r, &c.User); resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
