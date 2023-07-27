package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/admin_api/v1/wechat/rules"

	"github.com/zhangshanwen/shard/admin_api/v1/wechat"
	"github.com/zhangshanwen/shard/common"
)

func InitWechat(Router *gin.RouterGroup) {
	r := Router.Group(common.Wechat)
	{
		r.GET(common.UriEmpty, wh(wechat.Info))   // 获取自身信息
		r.GET(common.Status, wh(wechat.Status))   // 获取机器人状态
		r.GET(common.Friends, wh(wechat.Friends)) // 获取好友列表
	}
	g := r.Group(common.UriLogin)
	{
		g.GET(common.Qrcode, jwt(wechat.QrCode)) // 获取微信登录二维码
	}
	s := r.Group(common.Socket)
	{
		s.GET(common.UriAuthorization, socketWechat(wechat.Socket)) // 建立连接
	}
	rule := r.Group(common.Rules)
	{
		rule.GET(common.Default, jwt(rules.DefaultRules)) /// 获取默认规则
		rule.GET(common.UriEmpty, jwtTx(rules.Rules))     /// 获取用户添加规则
	}

}
