package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/admin_api/v1/task/task_log"

	"github.com/zhangshanwen/shard/admin_api/v1/task"
	"github.com/zhangshanwen/shard/common"
)

func InitTask(Router *gin.RouterGroup) {
	r := Router.Group(common.Task)
	{
		r.POST(common.UriEmpty, jwt(task.Post))  // 定制任务
		r.GET(common.UriEmpty, jwt(task.Get))    // 获取任务
		r.PUT(common.UriId, jwt(task.Edit))      // 编辑任务
		r.DELETE(common.UriId, jwt(task.Delete)) // 删除任务
	}
	stop := r.Group(common.Stop)
	{
		stop.DELETE(common.UriId, jwt(task.Stop)) // 停止任务
	}
	log := r.Group(common.Log)
	{
		log.GET(common.UriId, jwt(task_log.Get))       // 任务日志
		log.DELETE(common.UriId, jwt(task_log.Delete)) // 删除任务日志
	}
	run := r.Group(common.Run)
	{
		run.GET(common.UriId, jwt(task.Run)) // 执行任务
	}
}
