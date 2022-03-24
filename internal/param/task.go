package param

type (
	TaskCreate struct {
		Name string `json:"name"`
		Spec string `json:"spec"`
	}
	Task struct {
		Pagination
	}
)
