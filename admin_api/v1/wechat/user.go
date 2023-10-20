package wechat

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/eatmoreapple/openwechat"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
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

func Avatar(c *service.AdminWechatContext) (r service.Res) {
	// 获取用户头像
	var (
		p    param.Avatar
		resp *http.Response
	)
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}

	if p.FindId == c.Bot.Self.ID() {
		resp, r.Err = c.Bot.Self.GetAvatarResponse()
	} else if p.IsGroup {
		var group *openwechat.Group
		if group, r.Err = c.Bot.FindGroup(p.FindId); r.Err != nil {
			return
		}
		resp, r.Err = group.GetAvatarResponse()

	} else {
		var friend *openwechat.Friend
		if friend, r.Err = c.Bot.FindFriend(p.FindId); r.Err != nil {
			return
		}
		resp, r.Err = friend.GetAvatarResponse()
	}
	if r.Err != nil {
		return
	}
	var body []byte
	defer resp.Body.Close()
	if body, r.Err = io.ReadAll(resp.Body); r.Err != nil {
		return
	}
	r.Data = base64.StdEncoding.EncodeToString(body)
	return
}
