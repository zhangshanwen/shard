package response

type (
	MeetingResponse struct {
		List       []Meeting  `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	Meeting struct {
		Id          int64  `json:"id"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		Name        string `json:"name"`
		Status      int    `json:"status"`     // 会议状态 0进行中 1 结束
		EndTime     int64  `json:"end_time"`   // 结束时间
		StartTime   int64  `json:"start_time"` // 开始时间
		IsOwner     bool   `json:"is_owner"`
		Owner       string `json:"owner"`
	}
)
