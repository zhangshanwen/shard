package model

type Wallet struct {
	BaseModel
	Uid     int64 `json:"uid"     gorm:"index"`
	Balance int64 `json:"balance"`
}
