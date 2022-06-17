package model

type BaseModel struct {
	Id          int64 `json:"id"             gorm:"primaryKey;autoIncrement"`
	CreatedTime int64 `json:"created_time"   gorm:"autoCreateTime"` //autoUpdateTime:nano/milli 纳/毫秒
	UpdatedTime int64 `json:"updated_time"   gorm:"autoUpdateTime"`
	IsDeleted   bool  `json:"-"              gorm:"default:0"`
}

const (
	CommonStatusSuccess = 1
	CommonStatusFailed  = 0
)
