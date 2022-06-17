package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/operate_log"
	"github.com/zhangshanwen/shard/common"
)

func InitOperateLog(Router *gin.RouterGroup) {
	r := Router.Group(common.Log)
	{
		r.GET(common.UriEmpty, jwt(operate_log.Get)) // 日志列表
	}

}
