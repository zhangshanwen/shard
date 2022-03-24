package model

type (
	Status int8
	Task   struct {
		BaseModel
		Name    string `json:"name"`     // 任务名称
		Spec    string `json:"spec"`     // 任务执行时间
		Remark  string `json:"remark"`   // 备注
		Status  Status `json:"status"`   // 状态 1:正常 0:停止
		BeginAt int64  `json:"begin_at"` // 定时任务开始时间
	}
)

const (
	StatusStop   Status = 2
	StatusNormal Status = 1
)

func (t *Task) Run() {

}
func (t *Task) GetSpec() string {
	return t.Spec

}
func (t *Task) GetId() int64 {
	return t.Id
}
