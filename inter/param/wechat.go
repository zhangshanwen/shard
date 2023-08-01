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
)
