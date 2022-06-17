package response

type (
	RoleResponse struct {
		List       []Role     `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	Role struct {
		Id          int64  `json:"id"`
		CreatedTime int64  `json:"created_time"`
		UpdatedTime int64  `json:"updated_time"`
		Name        string `json:"name"` // 角色名称
	}
)
