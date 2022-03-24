package response

type (
	TaskResponse struct {
		List       []Task     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	Task struct {
		Id      int64  `json:"id"`
		Name    string `json:"name"`
		StartAt string `json:"start_at"`
		EndAt   string `json:"end_at"`
		Status  int    `json:"status"`
	}
)
