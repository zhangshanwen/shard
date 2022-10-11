package param

type (
	CreateLiveRoom struct {
		Name string `json:"name"`
	}
	LiveRoom struct {
		Pagination
	}

	CreateBarrage struct {
		Hash    string `json:"hash"    binding:"requeired"`
		Content string `json:"content" binding:"requeired"`
	}
	GetBarrage  struct {
		Hash    string `form:"hash"    binding:"requeired"`
	}
)
