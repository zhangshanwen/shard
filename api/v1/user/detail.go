package user

import (
	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/response"
)

func Detail(c *service.Context) (r service.Res) {
	resp := response.UserInfo{}
	defer func() {
		r.Data = resp
	}()
	if r.Err = copier.Copy(&resp, &c.User); r.Err != nil {
		return
	}
	return
}

func Balance(c *service.Context) (r service.Res) {
	resp := response.UserInfo{}
	defer func() {
		r.Data = resp
	}()
	if r.Err = copier.Copy(&resp, &c.User); r.Err != nil {
		return
	}
	return
}
