package wechat

import (
	"github.com/zhangshanwen/shard/initialize/service"
)

func Logout(c *service.AdminWechatContext) (r service.Res) {
	r.Err = c.Bot.Logout()
	return
}
