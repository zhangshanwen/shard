package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/live"
	"github.com/zhangshanwen/shard/common"
)

func InitLive(Router *gin.RouterGroup) {
	r := Router.Group(common.Live)
	{
		r.POST(common.UriEmpty, jwtTx(live.Post)) // 创建房间
		r.GET(common.UriId, live.Get)             // 获取房间
	}
}
