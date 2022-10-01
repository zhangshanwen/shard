package model

type (
	// LiveRoom 直播房间
	// id 10000开始自增
	LiveRoom struct {
		BaseModel
		Name      string         `json:"name"`       // 房间名称
		Status    LiveRoomStatus `json:"status"`     // 房间状态   -1 直播结束  0 待直播 1 正在直播
		OwnerId   int64          `json:"owner_id"`   // 创建者
		StartTime int64          `json:"start_time"` // 开播时间
		EndTime   int64          `json:"end_time"`   // 结束时间
		Owner     *Admin         `gorm:"foreignkey:OwnerId;rerences:Id;"`
	}
	LiveRoomStatus int
)

const (
	LiveRoomStatusStop   LiveRoomStatus = -1 // 直播结束
	LiveRoomStatusIdle   LiveRoomStatus = 0  // 待直播
	LiveRoomStatusLiving LiveRoomStatus = 1  // 正在直播
)
