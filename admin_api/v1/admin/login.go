package admin

import (
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func Login(c *service.AdminTxContext) (r service.Res) {
	p := param.AdminLogin{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		resp = response.AdminRolePermissionResponse{}
		tx   = c.Tx
		m    = model.Admin{Username: p.Username}
	)

	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	if r.Err = tx.Where(&m).First(&m).Error; r.Err != nil {
		r.DBError()
		return
	}

	if !m.CheckPassword(p.Password) {
		r.ActPWdError()
		return
	}
	if r.Err = tx.Model(&m).Where("id=?", m.Id).Updates(&model.Admin{
		LastLoginTime: time.Now().Unix(),
	}).Error; r.Err != nil {
		r.DBError()
		return
	}
	if r.Err = copier.Copy(&resp, &m); r.Err != nil {
		return
	}
	var token string
	if token, r.Err = tools.CreateToken(m.Id); r.Err != nil {
		r.LoginFailed()
		return
	}
	resp.Authorization = token
	if r.Err = rolePermission(tx, c, &m, &resp); r.Err != nil {
		r.DBError()
		return
	}
	oss, _ := tools.NewOss()
	resp.Avatar.Name = m.Avatar
	resp.Avatar.Url = oss.GetUrl(c, m.Avatar)
	c.Admin = m
	if r.Err = c.SaveLoginInfo(); r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLogSelect(tx, module, "login")
	return
}

func rolePermission(tx *gorm.DB, c *service.AdminTxContext, admin *model.Admin, resp *response.AdminRolePermissionResponse) (err error) {
	m := model.Role{}
	if err = tx.Preload("Permissions").Preload("Permissions.Routes").First(&m, admin.RoleId).Error; err != nil {
		return
	}
	routes := map[string]interface{}{}
	permissionMap := make(map[int64]*model.Permission)
	for _, item := range m.Permissions {
		for _, route := range item.Routes {
			routes[fmt.Sprintf("%s||%s", route.Method, route.Path)] = true
		}
		permissionMap[item.Id] = item
		if item.Key != "" {
			resp.Keys = append(resp.Keys, item.Key)
		}
	}

	if err = db.R.HSet(c.Context, fmt.Sprintf(common.RedisRoutesKey, admin.Id), routes).Err(); err != nil {
		return
	}
	var parents []*model.Permission
	for k, v := range permissionMap {
		if v.ParentId == 0 {
			parents = append(parents, v)
		} else {
			parent := permissionMap[v.ParentId]
			if parent.ParentId == 0 {
				parent.Children = append(permissionMap[v.ParentId].Children, *permissionMap[k])
			}
		}
	}
	if err = copier.Copy(&resp.List, &parents); err != nil {
		return
	}
	return
}
