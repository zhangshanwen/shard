package response

import (
	"encoding/json"

	"github.com/jinzhu/copier"
)

type (
	RolePermissionResponse struct {
		List []RolePermission `json:"list"`
	}
	RolePermission struct {
		Id          int64            `json:"id"`
		CreatedTime int64            `json:"created_time"`
		UpdatedTime int64            `json:"updated_time"`
		Name        string           `json:"name"`      // 权限名称
		ParentId    int64            `json:"parent_id"` // 父节点id
		Icon        string           `json:"icon"`      // 前端icon
		Key         string           `json:"key"`       // 前端唯一key
		Children    []RolePermission `json:"children"`  // 子页面
		Roles       []Role           `json:"roles"`     // 角色
	}
)

func (p *RolePermission) MarshalJSON() ([]byte, error) {

	if p.Children == nil {
		p.Children = []RolePermission{}
	}
	if p.Roles == nil {
		p.Roles = []Role{}
	}
	temp := struct {
		Id          int64            `json:"id"`
		CreatedTime int64            `json:"created_time"`
		UpdatedTime int64            `json:"updated_time"`
		Name        string           `json:"name"`      // 权限名称
		ParentId    int64            `json:"parent_id"` // 父节点id
		Icon        string           `json:"icon"`      // 前端icon
		Key         string           `json:"key"`       // 前端唯一key
		Children    []RolePermission `json:"children"`  // 子页面
		Roles       []Role           `json:"roles"`     // 角色
		IsChecked   bool             `json:"is_checked"`
	}{}
	_ = copier.Copy(&temp, p)
	temp.IsChecked = len(temp.Roles) > 0
	return json.Marshal(&temp)
}
