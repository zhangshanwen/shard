package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/host"
	"github.com/zhangshanwen/shard/common"
)

func InitHost(Router *gin.RouterGroup) {
	r := Router.Group(common.Host)
	{
		r.POST(common.UriEmpty, jwt(host.Post))  // 创建主机
		r.GET(common.UriEmpty, jwt(host.Get))    // 主机列表
		r.PUT(common.UriId, jwt(host.Edit))      // 修改主机
		r.DELETE(common.UriId, jwt(host.Delete)) // 删除主机
		r.POST(common.Room, jwt(host.Room))      // 创建房间
	}
	s := r.Group(common.Socket)
	{
		s.GET(common.UriAuthorization+"/"+common.UriId, socket(host.Socket)) // 主机socket
	}
}
