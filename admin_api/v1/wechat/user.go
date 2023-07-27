package wechat

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/zhangshanwen/shard/initialize/service"
)

func Info(c *service.AdminWechatContext) (r service.Res) {
	// 获取用户信息
	var (
		self *openwechat.Self
	)
	if self, r.Err = c.Bot.GetCurrentUser(); r.Err != nil {
		return
	}
	r.Data = self.User
	return
}

func Status(c *service.AdminWechatContext) (r service.Res) {
	// 查询机器人是否激活
	r.Data = map[string]interface{}{
		"active": c.Bot.Alive(),
	}
	return
}
