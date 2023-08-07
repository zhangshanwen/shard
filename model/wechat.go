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
	ReplyBot struct {
		BaseModel
		Friends     string `json:"friends"` // 好友id , 分割
		Name        string `json:"name"`
		Description string `json:"description"`
		Rules       []Rule `json:"rules"   gorm:"many2many:reply_bot_rule;"`
		Uid         int64  `json:"uid"`
	}
	ReplyBotRule struct {
		ReplyBotId int64 `json:"reply_bot_id"`
		RuleId     int64 `json:"rule_id"`
	}
)
