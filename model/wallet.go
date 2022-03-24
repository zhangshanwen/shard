package model

type (
	Wallet struct {
		BaseModel
		Uid     int64 `json:"uid"     gorm:"index"`
		Balance int64 `json:"balance"`
	}
	WalletRecordType int
	WalletRecord     struct {
		BaseModel
		WalletId   int64            `json:"wallet_id"     gorm:"index"`
		Amount     int64            `json:"amount"`
		RecordType WalletRecordType `json:"record_type"` // -1后台扣除 1 后台增加
	}
)

const (
	WalletRecordTypeBackendIncrease WalletRecordType = 1
	WalletRecordTypeBackendDecrease WalletRecordType = -1
)
