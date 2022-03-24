package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/version"
	"github.com/zhangshanwen/shard/middleware"
)

func InitVersion(Router *gin.RouterGroup) {
	r := Router.Group("version")
	v := middleware.AdminHandel
	{
		r.GET("", v(version.Get)) // 获取当前版本
	}
}
