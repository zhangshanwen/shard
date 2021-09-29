package response

import (
	"encoding/json"
)

type (
	AdminInfo struct {
		Authorization string `json:"authorization,omitempty"`
		Avatar        string `json:"avatar"`
		Username      string `json:"username"`
	}
	AdminResponse struct {
		List       []Admin    `json:"list"`
		Pagination Pagination `json:"pagination"`
	}
	Admin struct {
		Id            int64  `json:"id"`
		CreatedTime   int64  `json:"created_time"`
		UpdatedTime   int64  `json:"updated_time"`
		LastLoginTime int64  `json:"last_login_time"`
		Username      string `json:"username"`
		Role          Role   `json:"role"`
	}
	AdminRolePermissionResponse struct {
		Authorization string                `json:"authorization,omitempty"`
		Username      string                `json:"username"`
		List          []AdminRolePermission `json:"list"`
		Keys          []string              `json:"keys"`
		Avatar        Avatar                `json:"avatar"`
	}
	Avatar struct {
		Url  string `json:"url"`
		Name string `json:"name"`
	}
	AdminRolePermission struct {
		Id       int64                 `json:"id"`
		Name     string                `json:"name"`      // 权限名称
		ParentId int64                 `json:"parent_id"` // 父节点id
		Icon     string                `json:"icon"`      // 前端icon
		Key      string                `json:"key"`       // 前端唯一key
		Children []AdminRolePermission `json:"children"`  // 子页面
	}
)

func (p *AdminRolePermission) MarshalJSON() ([]byte, error) {

	if p.Children == nil {
		p.Children = []AdminRolePermission{}
	}
	temp := struct {
		Id       int64                 `json:"id"`
		Name     string                `json:"name"`      // 权限名称
		ParentId int64                 `json:"parent_id"` // 父节点id
		Icon     string                `json:"icon"`      // 前端icon
		Key      string                `json:"key"`       // 前端唯一key
		Children []AdminRolePermission `json:"children"`  // 子页面
	}{
		Id:       p.Id,
		Name:     p.Name,
		ParentId: p.ParentId,
		Icon:     p.Icon,
		Key:      p.Key,
		Children: p.Children,
	}
	return json.Marshal(&temp)
}
