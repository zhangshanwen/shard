package param

type (
	Rule struct {
		Pagination
	}
	SaveRule struct {
		Name        string `json:"name"         binding:"required"`
		Key         string `json:"key"          binding:"required"`
		Description string `json:"description"`
		Reply       string `json:"reply"        binding:"required"`
	}
	FriendsGroups struct {
		IsAllFriends   bool     `json:"is_all_friends"`  // 是否所有好友
		Friends        []string `json:"friends"`         // 好友
		ExcludeFriends []string `json:"exclude_friends"` // 排除掉的好友(优先级高于选中的好友)
		IsAllGroups    bool     `json:"is_all_groups"`   // 是否全部组群
		Groups         []string `json:"groups"`          // 选中的群聊
		ExcludeGroups  string   `json:"exclude_groups"`  // 排除掉的群聊(优先级高于选中的群聊)
	}

	SaveReplyBot struct {
		FriendsGroups
		Name        string  `json:"name"           binding:"required"`
		Description string  `json:"description"`
		RuleIds     []int64 `json:"rule_ids"       binding:"required"`
	}
	ReplyBot struct {
		Pagination
	}
	TimerBot struct {
		Pagination
	}
	SaveTimerBot struct {
		FriendsGroups
		Name        string `json:"name"           binding:"required"`
		Msg         string `json:"msg"            binding:"required"` // 发送信息
		Spec        string `json:"spec"           binding:"required"` // 定时规则
		Times       int    `json:"times"`                             // 次数
		Description string `json:"description"`                       // 描述
	}
	Avatar struct {
		FindId  string `form:"find_id"           binding:"required"`
		IsGroup bool   `form:"is_group"`
	}
)
