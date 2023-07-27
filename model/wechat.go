package model

type (
	Rule struct {
		BaseModel
		Key         string `json:"key"`         // 关键词
		Result      string `json:"result"`      // 回复代码,可替换默认模板{{}}
		Description string `json:"description"` // 描述
		Uid         int64  `json:"uid"`         // 用户id
	}
)
