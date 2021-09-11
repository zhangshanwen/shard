package model

type (
	// 角色表
	Role struct {
		BaseModel
		Name        string        `json:"name"` // 角色名称
		Permissions []*Permission `gorm:"many2many:role_permission;"`
	}
	RolePermission struct {
		RoleId       int64 `json:"role_id"`
		PermissionId int64 `json:"permission_id"`
	}
)
