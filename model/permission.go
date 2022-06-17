package model

import (
	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/common"
)

type (

	// 权限表
	Permission struct {
		BaseModel
		Name     string       `json:"name"`                                                                                             // 权限名称
		ParentId int64        `json:"parent_id"`                                                                                        // 父节点id
		Key      string       `json:"key"`                                                                                              // 前端唯一key
		Routes   []Route      `json:"routes"                gorm:"many2many:permission_route;"`                                         // 路由
		Children []Permission `json:"children"              gorm:"foreignKey:parent_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 子页面
		Roles    []Role       `json:"roles"                 gorm:"many2many:role_permission;"`
	}
)

func (p *Permission) AfterFind(tx *gorm.DB) (err error) {
	if signal, ok := tx.Get(common.PermissionFindChildren); ok {
		if signal.(bool) {
			return tx.Set(common.PermissionFindChildren, true).Where("parent_id = ? ", p.Id).Preload("Routes").Find(&p.Children).Error
		}
	}
	if signal, ok := tx.Get(common.PermissionRoleFindChildrenId); ok {
		roleId := signal.(int64)
		if roleId > 0 {
			return tx.Set(common.PermissionRoleFindChildrenId, roleId).Where("parent_id = ? ", p.Id).Preload("Roles", "id = ? ", roleId).Find(&p.Children).Error
		}
	}
	return
}

func (p *Permission) BeforeDelete(tx *gorm.DB) (err error) {
	if p.Id == 0 {
		return
	}
	return tx.Where("parent_id = ? ", p.Id).Delete(&Permission{}).Error
}
