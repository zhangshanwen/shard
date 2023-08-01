package response

type (
	DefaultRule struct {
		FunctionName string `json:"function_name"`
		Description  string `json:"description"` // 描述
	}
	DefaultRulesResponse struct {
		List []DefaultRule `json:"list"`
	}

	Rule struct {
		Id          int64  `json:"id"`
		Name        string `json:"name"`        // 名称
		Key         string `json:"key"`         // 关键词
		Reply       string `json:"reply"`       // 回复代码,可替换默认模板{{}}
		Description string `json:"description"` // 描述
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
	}
	RulesResponse struct {
		List       []Rule     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
)
