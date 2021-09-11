package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/user"
	"github.com/zhangshanwen/shard/common"
)

func InitUser(Router *gin.RouterGroup) {
	r := Router.Group(common.Users)
	{
		r.GET(common.UriEmpty, jwt(user.Get))     // 获取所有用户
		r.POST(common.UriEmpty, jwt(user.Create)) // 创建用户
		r.PUT(common.UriId, jwt(user.Edit))       // 修改用户信息
		r.DELETE(common.UriId, jwt(user.Delete))  // 删除用户
	}
}
