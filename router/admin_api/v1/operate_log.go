package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/operate_log"
	"github.com/zhangshanwen/shard/common"
)

func InitOperateLog(Router *gin.RouterGroup) {
	r := Router.Group(common.Log)
	{
		r.GET(common.UriEmpty, jwtTx(operate_log.Get))    // 日志列表
		r.DELETE(common.Empty, jwtTx(operate_log.Empty))  // 清空日志
		r.DELETE(common.UriEmpty, jwtTx(operate_log.Del)) // 删除日志
	}

}
