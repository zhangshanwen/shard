package wechat

import (
	"strings"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools/wechat"
)

func QrCode(c *service.AdminTxContext) (r service.Res) {
	var (
		w            = wechat.NewWechat()
		replyBots    []model.ReplyBot
		timerBots    []model.TimerBot
		replies      []*wechat.Reply
		timerReplies []*wechat.TimerReply
		code         string
	)
	var (
		tx = c.Tx
	)
	defer func() {
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	tx.Where("uid=?", c.Admin.Id).Preload("Rules").Find(&replyBots)
	for _, i := range replyBots {
		var rules = make(map[string]string)
		for _, rule := range i.Rules {
			rules[rule.Key] = rule.Reply
		}
		replies = append(replies, &wechat.Reply{
			FriendsGroups: wechat.FriendsGroups{
				IsAllFriends:   i.IsAllFriends,
				ExcludeFriends: strings.Split(i.ExcludeFriends, ","),
				Friends:        strings.Split(i.Friends, ","),
				IsAllGroups:    i.IsAllGroups,
				ExcludeGroups:  strings.Split(i.ExcludeGroups, ","),
				Groups:         strings.Split(i.Groups, ","),
			},
			Rules: rules,
		})

	}
	tx.Where("uid=?", c.Admin.Id).Find(&timerBots)
	for _, i := range timerBots {
		timerReplies = append(timerReplies, &wechat.TimerReply{
			FriendsGroups: wechat.FriendsGroups{
				IsAllFriends:   i.IsAllFriends,
				ExcludeFriends: strings.Split(i.ExcludeFriends, ","),
				Friends:        strings.Split(i.Friends, ","),
				IsAllGroups:    i.IsAllGroups,
				ExcludeGroups:  strings.Split(i.ExcludeGroups, ","),
				Groups:         strings.Split(i.Groups, ","),
			},
			Msg:   i.Msg,
			Spec:  i.Spec,
			Times: i.Times,
		})

	}
	if code, r.Err = w.Qrcode(c.Admin.Id, replies, timerReplies); r.Err != nil {
		return
	}
	r.Data = code
	return
}
