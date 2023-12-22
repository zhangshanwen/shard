package param

type (
	CreateMeeting struct {
		Name string `json:"name"`
	}
	MeetingRecords struct {
		Name   string `form:"name"`
		Status string `form:"status"`
		Pagination
	}
)
