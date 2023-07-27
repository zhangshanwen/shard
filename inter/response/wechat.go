package response

type (
	DefaultRule struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
	}
	DefaultRulesResponse struct {
		List []DefaultRule `json:"list"`
	}

	Rule struct {
		Id          int64 `json:"id"`
		CreatedTime int64 `json:"created_time"`
		UpdatedTime int64 `json:"updated_time"`
		DefaultRule
	}
	RulesResponse struct {
		List       []Rule     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
)
