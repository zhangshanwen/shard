package response

type Pagination struct {
	PageIndex int   `form:"page_index"`
	PageSize  int   `from:"page_size"`
	Total     int64 `json:"total"`
}
