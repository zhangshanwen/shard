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
)
