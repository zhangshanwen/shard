package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/role"
	"github.com/zhangshanwen/shard/admin_api/v1/role/permission"
	"github.com/zhangshanwen/shard/common"
)

func InitRole(Router *gin.RouterGroup) {
	r := Router.Group(common.Roles)
	{
		r.GET(common.UriEmpty, jwt(role.Get))     // 获取角色
		r.POST(common.UriEmpty, jwt(role.Create)) // 创建角色
		r.PUT(common.UriId, jwt(role.Edit))       // 修改角色
		r.DELETE(common.UriId, jwt(role.Delete))  // 删除角色

		permissions := r.Group(common.Permissions)
		{
			permissions.GET(common.UriId, jwt(permission.Get))  // 获取角色权限
			permissions.PUT(common.UriId, jwt(permission.Edit)) // 修改角色权限
		}
	}
}
