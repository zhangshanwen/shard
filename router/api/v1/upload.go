package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/api/v1/file"
	"github.com/zhangshanwen/shard/middleware"
)

func InitUpload(Router *gin.RouterGroup) {
	r := Router.Group("file")
	v := middleware.Handel
	{
		r.POST("", v(file.Upload)) // 上传代码
	}
}
