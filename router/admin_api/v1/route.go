package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/route"
	"github.com/zhangshanwen/shard/common"
)

func InitRoute(Router *gin.RouterGroup) {
	r := Router.Group(common.Routes)
	{
		r.GET(common.UriEmpty, jwt(route.Get)) // 获取所有路由
	}
}
