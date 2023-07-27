package wechat

import (
	"github.com/eatmoreapple/openwechat"

	"github.com/zhangshanwen/shard/initialize/service"
)

func Friends(c *service.AdminWechatContext) (r service.Res) {
	var (
		friends openwechat.Friends
	)
	if friends, r.Err = c.Bot.Friends(); r.Err != nil {
		return
	}
	r.Data = friends
	return
}
