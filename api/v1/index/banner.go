package index

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/response"
)

func Banner(c *service.Context) (r service.Res) {
	resp := response.BannerResponse{}
	defer func() {
		r.Data = resp
	}()
	if r.Err = copier.Copy(&r, &c.User); r.Err != nil {
		return
	}
	return
}
