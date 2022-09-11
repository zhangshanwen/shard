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
		r.GET(common.UriEmpty, jwtTx(role.Get))     // 获取角色
		r.POST(common.UriEmpty, jwtTx(role.Create)) // 创建角色
		r.PUT(common.UriId, jwtTx(role.Edit))       // 修改角色
		r.DELETE(common.UriId, jwtTx(role.Delete))  // 删除角色

		permissions := r.Group(common.Permissions)
		{
			permissions.GET(common.UriId, jwtTx(permission.Get))  // 获取角色权限
			permissions.PUT(common.UriId, jwtTx(permission.Edit)) // 修改角色权限
		}
	}
}
