package param

type (
	TaskSave struct {
		HostId     int64  `json:"host_id"       binding:"required"` // 主机id
		Name       string `json:"name"`                             // 任务名称
		Cmd        string `json:"cmd"`                              // 任务指令
		Spec       string `json:"spec"`                             // 执行时刻
		EffectTime int64  `json:"effect_time"   binding:"gte=0"`    // 生效时间
		ExpiryTime int64  `json:"expiry_time"   binding:"gte=0"`    // 失效时间
		Comment    string `json:"comment"`                          // 备注
	}
	Task struct {
		Pagination
	}
	TaskLog struct {
		Id int64 `form:"id"`
		Pagination
	}
	TaskDelete struct {
		All bool    `form:"all"`
		Ids []int64 `form:"ids"`
	}
)
