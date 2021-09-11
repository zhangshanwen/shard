package param

type (
	Permission struct {
		Name     string  `json:"name"       binding:"required"` // 权限名称
		ParentId int     `json:"parent_id"`                     // 父节点id
		Icon     string  `json:"icon"`                          // 前端icon
		Key      string  `json:"key"`                           // 前端key
		RouteIds []int64 `json:"route_ids"`                     // 路由ids
	}
	PermissionRecords struct {
		Pagination
	}
)
