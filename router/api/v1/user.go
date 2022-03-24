package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/common"

	"github.com/zhangshanwen/shard/api/v1/user"
	"github.com/zhangshanwen/shard/middleware"
)

func InitUser(Router *gin.RouterGroup) {
	r := Router.Group(common.User)
	v := middleware.Handel
	jwt := middleware.JwtHandel
	{
		r.POST("", v(user.Register))   // 创建用户
		r.POST("login", v(user.Login)) // 登录用户
		r.GET("", jwt(user.Detail))    // 获取用户信息
	}
}
