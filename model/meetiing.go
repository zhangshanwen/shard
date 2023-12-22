package model

type (
	MeetingStatus int
	Meeting       struct {
		BaseModel
		Uid       int64         `json:"uid"`
		Name      string        `json:"name"`       // 房间名
		Status    MeetingStatus `json:"status"`     // 会议状态 0待开始 1 开始 2 结束
		EndTime   int64         `json:"end_time"`   // 结束时间
		StartTime int64         `json:"start_time"` // 开始时间
		Owner     *Admin        `gorm:"foreignkey:Uid;rerences:Id;"`
	}
)

const (
	MeetingStatusPending = iota
	MeetingStatusRunning
	MeetingStatusEnd
)
