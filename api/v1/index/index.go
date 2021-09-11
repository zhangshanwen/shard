package index

import (
	"github.com/jinzhu/copier"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/response"
)

func Index(c *service.Context) (resp service.Res) {
	r := response.UserInfo{}
	if resp.Err = copier.Copy(&r, &c.User); resp.Err != nil {
		return
	}
	resp.Data = r
	return
}
