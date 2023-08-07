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

	SaveReplyBot struct {
		Name        string   `json:"name"           binding:"required"`
		Description string   `json:"description"`
		Friends     []string `json:"friends"        binding:"required"`
		RuleIds     []int64  `json:"rule_ids"       binding:"required"`
	}
	ReplyBot struct {
		Pagination
	}
)
