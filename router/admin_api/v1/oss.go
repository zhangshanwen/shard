package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/oss"
	"github.com/zhangshanwen/shard/common"
)

func InitOss(Router *gin.RouterGroup) {
	r := Router.Group(common.Oss)
	{
		r.GET(common.Token, v(oss.Token)) // 获取token
	}
}
