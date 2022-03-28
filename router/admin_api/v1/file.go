package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/file"
	"github.com/zhangshanwen/shard/common"
)

func InitFile(Router *gin.RouterGroup) {
	r := Router.Group(common.File)
	{
		r.POST(common.UriEmpty, jwt(file.Upload)) // 上传代码
		r.GET(common.UriId, jwt(file.Detail))     // 代码详情
		r.PUT(common.UriId, jwt(file.Update))     // 修改代码
		r.GET(common.UriEmpty, jwt(file.Get))     // 代码列表
		r.DELETE(common.UriId, jwt(file.Delete))  // 删除代码
	}
	run := r.Group(common.Run)
	{
		run.POST(common.UriEmpty, jwt(file.Run)) // 执行代码
	}
}
