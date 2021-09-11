package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/admin_api/v1/permission"
	"github.com/zhangshanwen/shard/common"
)

func InitPermission(Router *gin.RouterGroup) {
	r := Router.Group(common.Permissions)
	{
		r.GET(common.UriEmpty, jwt(permission.Get))     // 获取所有权限
		r.POST(common.UriEmpty, jwt(permission.Create)) // 创建权限
		r.PUT(common.UriId, jwt(permission.Edit))       // 修改权限信息
		r.DELETE(common.UriId, jwt(permission.Delete))  // 删除权限
	}
}
