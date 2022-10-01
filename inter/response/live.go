package response

type (
	LiveResponse struct {
		Url string `json:"url"`
	}
	LiveRoomResponse struct {
		List       []LiveRoom `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	LiveRoom struct {
		Id          int64  `json:"id"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		Name        string `json:"name"`       // 名称
		Hash        string `json:"hash"`       // 直播hash
		Status      int    `json:"status"`     // 房间状态   -1 直播结束  0 待直播 1 正在直播
		Owner       string `json:"owner"`      // 创建者
		StartTime   int64  `json:"start_time"` // 开播时间
		EndTime     int64  `json:"end_time"`   // 结束时间
	}
)
