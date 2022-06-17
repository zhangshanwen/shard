package response

type (
	TaskResponse struct {
		List       []Task     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	Task struct {
		Id           int64    `json:"id"`
		Name         string   `json:"name"`           // 任务名称
		HostId       int64    `json:"host_id"`        // 主机id
		Status       int8     `json:"status"`         // 状态 -3:失败 -2:过期 -1:停止 0:待生效 1:生效中
		Cmd          string   `json:"cmd"`            // 任务指令
		Spec         string   `json:"spec"`           // 执行时刻
		EffectTime   int64    `json:"effect_time"`    // 生效时间
		ExpiryTime   int64    `json:"expiry_time"`    // 失效时间
		NextExecTime int64    `json:"next_exec_time"` // 下次执行时间
		Comment      string   `json:"comment"`        // 备注
		Host         TaskHost `json:"host"`
		CreatedTime  int64    `json:"created_time"`
		UpdatedTime  int64    `json:"updated_time"`
	}
	TaskHost struct {
		Name string `json:"name"`
	}
	TaskLogResponse struct {
		List       []TaskLog  `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	TaskLog struct {
		Id          int64  `json:"id"`
		Status      uint8  `json:"status"`
		Comment     string `json:"comment"` // 备注
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
	}
)
