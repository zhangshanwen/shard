package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/admin_api/v1/permission"
	"github.com/zhangshanwen/shard/common"
)

func InitPermission(Router *gin.RouterGroup) {
	r := Router.Group(common.Permissions)
	{
		r.GET(common.UriEmpty, jwtTx(permission.Get))     // 获取所有权限
		r.POST(common.UriEmpty, jwtTx(permission.Create)) // 创建权限
		r.PUT(common.UriId, jwtTx(permission.Edit))       // 修改权限信息
		r.DELETE(common.UriId, jwtTx(permission.Delete))  // 删除权限
	}
}
