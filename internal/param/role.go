package param

type (
	Role struct {
		Name string `json:"name"       binding:"required"` // 角色名称
	}
	RoleRecords struct {
		Pagination
	}
	RolePermissionEdit struct {
		PermissionIds []int64 `json:"permission_ids"`
	}
)
