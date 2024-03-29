package model

type (
	Rule struct {
		BaseModel
		Name        string `json:"name"`        // 名称
		Key         string `json:"key"`         // 关键词
		Reply       string `json:"reply"`       // 回复代码,可替换默认模板{{}}
		Description string `json:"description"` // 描述
		Uid         int64  `json:"uid"`         // 用户id
	}
	FriendsGroups struct {
		IsAllFriends   bool   `json:"is_all_friends"`  // 是否所有好友触发
		Friends        string `json:"friends"`         // 好友id , 分割
		ExcludeFriends string `json:"exclude_friends"` // 排除的好友id , 分割
		IsAllGroups    bool   `json:"is_all_groups"`   // 是否全部组群
		ExcludeGroups  string `json:"exclude_groups"`  // 排除掉的群聊(优先级高于选中的群聊)
		Groups         string `json:"groups"`          // 选中的群聊
	}
	ReplyBot struct {
		BaseModel
		FriendsGroups
		Name        string `json:"name"`
		Description string `json:"description"`
		Rules       []Rule `json:"rules"   gorm:"many2many:reply_bot_rule;"`
		Uid         int64  `json:"uid"`
	}
	ReplyBotRule struct {
		ReplyBotId int64 `json:"reply_bot_id"`
		RuleId     int64 `json:"rule_id"`
	}
	ChatMessage struct {
		BaseModel
		OwnerId    string `json:"owner_id"`    // 拥有者id
		SenderId   string `json:"sender_id"`   // 发送者id
		ReceiverId string `json:"receiver_id"` // 接收者id
		IsGroup    bool   `json:"is_group"`    // 是否为群组
		IsSuccess  bool   `json:"is_success"`  // 是否成功
		Msg        string `json:"msg"`         // 消息
		MsgType    int    `json:"msg_type"`    // 消息类型  文本 表情 图片 录音 视频
		FileId     int64  `json:"file_id"`     // 文件id
	}
	ChatFile struct {
		BaseModel
		Uri      string `json:"uri"`       // 文件地址
		FileType string `json:"file_type"` // 文件类型 图片 录音 视频
	}
	// TimerBot 定时机器人
	TimerBot struct {
		BaseModel
		FriendsGroups
		Uid         int64  `json:"uid"`
		Name        string `json:"name"`        // 机器人名称
		Description string `json:"description"` // 机器人描述
		Msg         string `json:"msg"`         // 发送内容
		Spec        string `json:"spec"`        // 定时规则
		Times       int    `json:"times"`       // <=0 为无限次 >0 为有限次数

	}
)
