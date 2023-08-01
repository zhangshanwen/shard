package wechat

import (
	"fmt"
	"time"

	"github.com/eatmoreapple/openwechat"

	"github.com/zhangshanwen/shard/common"
)

type (
	Reply struct {
		isAllFriends   bool              // 是否全部朋友
		excludeFriends []string          // 排除掉的朋友(优先级高于选中的朋友)
		friends        []string          // 选中的朋友
		isAllGroups    bool              // 是否全部组群
		excludeGroups  []string          // 排除掉的群聊(优先级高于选中的群聊)
		groups         []string          // 选中的群聊
		rules          map[string]string // 适配正则
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
	templateReply = map[string]replyFunc{
		"当前时间": func(sender *openwechat.User, i ...interface{}) string {
			return time.Now().Format(common.TimeFullFormat)
		},
		"星期几": func(sender *openwechat.User, i ...interface{}) string {
			return time.Now().Weekday().String()
		},
		"周几": func(sender *openwechat.User, i ...interface{}) string {
			return time.Now().Weekday().String()
		},
		"我是谁": func(sender *openwechat.User, i ...interface{}) string {
			return sender.NickName
		},
	}
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
