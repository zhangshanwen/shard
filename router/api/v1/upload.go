package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/api/v1/upload"
	"github.com/zhangshanwen/shard/middleware"
)

func InitUpload(Router *gin.RouterGroup) {
	r := Router.Group("upload")
	v := middleware.Handel
	{
		r.POST("", v(upload.Upload)) // 上传代码
	}
}
