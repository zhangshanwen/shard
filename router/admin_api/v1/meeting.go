package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/admin_api/v1/meeting"
	"github.com/zhangshanwen/shard/common"
)

func InitMeeting(Router *gin.RouterGroup) {
	r := Router.Group(common.Meeting)
	{
		r.POST(common.Push+common.Backlash+common.UriId, jwt(meeting.Push))                                               // 推送视频
		r.GET(common.Join+common.Backlash+common.UriId, jwt(meeting.Join))                                                // 加入播放
		r.GET(common.Offset+common.Backlash+common.UriId, jwt(meeting.Offset))                                            // 获取播放偏移量
		r.POST(common.UriEmpty, jwtTx(meeting.Create))                                                                    // 创建房间
		r.GET(common.UriEmpty, jwtTx(meeting.Get))                                                                        // 房间列表
		r.GET(common.Socket+common.Backlash+common.UriAuthorization+common.Backlash+common.UriId, socket(meeting.Socket)) // 建立连接

	}
}
