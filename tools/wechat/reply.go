package wechat

import (
	"fmt"
	"time"

	"github.com/eatmoreapple/openwechat"

	"github.com/zhangshanwen/shard/common"
)

type (
	Reply struct {
		IsAllFriends   bool              // 是否全部朋友
		ExcludeFriends []string          // 排除掉的朋友(优先级高于选中的朋友)
		Friends        []string          // 选中的朋友
		IsAllGroups    bool              // 是否全部组群
		ExcludeGroups  []string          // 排除掉的群聊(优先级高于选中的群聊)
		Groups         []string          // 选中的群聊
		Rules          map[string]string // 适配正则
	}
	template struct {
		Description string
		replyFunc   replyFunc
	}

	replyFunc func(*openwechat.User, ...interface{}) string
)

/*
回复模板
*/
var (
	DefaultTemplateReply = map[string]template{
		"now": {
			Description: fmt.Sprintf("系统时间,例如:%v", common.TimeFullFormat),
			replyFunc: func(sender *openwechat.User, i ...interface{}) string {
				return time.Now().Format(common.TimeFullFormat)
			},
		},
		"week": {
			Description: fmt.Sprintf("星期几,例如:%v", time.Now().Weekday().String()),
			replyFunc: func(sender *openwechat.User, i ...interface{}) string {
				return time.Now().Weekday().String()
			},
		},
		"woami": {
			Description: "微信昵称",
			replyFunc: func(sender *openwechat.User, i ...interface{}) string {
				return sender.NickName
			},
		},
	}
)
