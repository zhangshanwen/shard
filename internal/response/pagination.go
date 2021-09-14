package response

type Pagination struct {
	PageIndex int   `json:"page_index"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
}
