package model

import (
	"time"

	"github.com/zhangshanwen/shard/common"
)

type (
	Status int8
	Task   struct {
		BaseModel
		Name         string    `json:"name"`           // 任务名称
		HostId       int64     `json:"host_id"`        // 主机id
		Status       Status    `json:"status"`         // 状态 -3:失败 -2:过期 -1:停止 0:待生效 1:生效中
		Cmd          string    `json:"cmd"`            // 任务指令
		Spec         string    `json:"spec"`           // 执行时刻
		EffectTime   int64     `json:"effect_time"`    // 生效时间
		ExpiryTime   int64     `json:"expiry_time"`    // 失效时间
		NextExecTime int64     `json:"next_exec_time"` // 下次执行时间
		Comment      string    `json:"comment"`        // 备注
		TaskLogs     []TaskLog `gorm:"foreignKey:TaskId"`
		Host         Host      `gorm:"foreignkey:HostId;rerences:Id"`
	}
	TaskLog struct {
		BaseModel
		TaskId  int64  `json:"task_id"` // 执行记录
		Status  uint8  `json:"status"`  // 状态 0 失败 1 成功
		Comment string `json:"comment"` // 备注
	}
)

const (
	StatusFailed Status = iota - 3
	StatusExpiry
	StatusStop
	StatusIdle
	StatusRunning
)

func (t *Task) Verify() (err error) {
	t.Status = StatusIdle
	if t.ExpiryTime == 0 {
		t.Status = StatusRunning
		return
	}

	now := time.Now().Unix()
	if t.ExpiryTime+60*60*24*1 <= now {
		return common.ExpiryTimeShouldNotBeLessThanOneDayErr
	}
	if now >= t.ExpiryTime {
		return common.ExpiryTimeShouldNotBeLessThanCurrentTimeErr
	}
	if now >= t.EffectTime {
		t.Status = StatusRunning
	}
	return
}
