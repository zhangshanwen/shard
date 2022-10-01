package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/live"
	"github.com/zhangshanwen/shard/common"
)

func InitLive(Router *gin.RouterGroup) {
	r := Router.Group(common.Live)
	{
		r.POST(common.UriEmpty, jwtTx(live.Create)) // 创建房间
		r.POST(common.Start, jwtTx(live.Start))     // 获取直播hash
		r.GET(common.UriEmpty, jwtTx(live.Get))     // 获取直播列表
		r.GET(common.UriId, live.Flv)               // 获取直播数据
	}
}
