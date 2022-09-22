package model

type (
	// OperateLog 操作日志
	OperateLog struct {
		BaseModel
		OperateId int64          `json:"operate_id"`  // 操作人id
		Module    string         `json:"module"`      // 模块
		Log       string         `json:"operate_log"` // 操作日志
		Type      OperateLogType `json:"type"`        // 0增 1删 2改 3查
		Admin     *Admin         `gorm:"foreignkey:OperateId;rerences:Id;"`
	}
	OperateLogType uint8
)

const (
	OperateLogTypeAdd OperateLogType = iota
	OperateLogTypeDel
	OperateLogTypeUpdate
	OperateLogTypeSelect
)
