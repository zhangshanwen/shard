package response

import (
	"encoding/json"
	"github.com/jinzhu/copier"
)

type (
	PermissionResponse struct {
		List []Permission `json:"list"`
	}
	Permission struct {
		Id          int64        `json:"id"`
		CreatedTime int64        `json:"created_time"`
		UpdatedTime int64        `json:"updated_time"`
		Name        string       `json:"name"`      // 权限名称
		ParentId    int64        `json:"parent_id"` // 父节点id
		Icon        string       `json:"icon"`      // 前端icon
		Key         string       `json:"key"`       // 前端唯一key
		Routes      []Route      `json:"routes"`    // 路由
		Children    []Permission `json:"children"`  // 子页面
	}
)

func (p *Permission) MarshalJSON() ([]byte, error) {

	if p.Children == nil {
		p.Children = []Permission{}
	}
	if p.Routes == nil {
		p.Routes = []Route{}
	}
	temp := struct {
		Id          int64        `json:"id"`
		CreatedTime int64        `json:"created_time"`
		UpdatedTime int64        `json:"updated_time"`
		Name        string       `json:"name"`      // 权限名称
		ParentId    int64        `json:"parent_id"` // 父节点id
		Icon        string       `json:"icon"`      // 前端icon
		Key         string       `json:"key"`       // 前端唯一key
		Routes      []Route      `json:"routes"`    // 路由
		Children    []Permission `json:"children"`  // 子页面
	}{}
	_ = copier.Copy(&temp, p)
	return json.Marshal(&temp)
}
