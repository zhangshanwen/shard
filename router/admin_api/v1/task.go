package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/task"
	"github.com/zhangshanwen/shard/common"
)

func InitTask(Router *gin.RouterGroup) {
	r := Router.Group(common.Task)
	{
		r.POST(common.UriEmpty, jwt(task.Post)) // 定制任务
		r.GET(common.UriEmpty, jwt(task.Get))   // 获取任务
		r.DELETE(common.UriId, jwt(task.Stop))  // 停止任务
	}
}
