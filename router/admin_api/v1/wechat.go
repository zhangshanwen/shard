package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/admin_api/v1/wechat"
	"github.com/zhangshanwen/shard/admin_api/v1/wechat/rules"
	"github.com/zhangshanwen/shard/common"
)

func InitWechat(Router *gin.RouterGroup) {
	r := Router.Group(common.Wechat)
	{
		r.GET(common.UriEmpty, wh(wechat.Info))    // 获取自身信息
		r.GET(common.Status, wh(wechat.Status))    // 获取机器人状态
		r.GET(common.Friends, wh(wechat.Friends))  // 获取好友列表
		r.DELETE(common.Logout, wh(wechat.Logout)) // 登出

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
		rule.GET(common.Functions, jwt(rules.Functions)) // 获取可以调用的内部函数
		rule.GET(common.UriEmpty, jwtTx(rules.Rules))    // 获取用户添加规则
		rule.POST(common.UriEmpty, jwtTx(rules.Create))  // 用户创建新的规则
		rule.PUT(common.UriId, jwtTx(rules.Edit))        // 修改规则信息
		rule.DELETE(common.UriId, jwtTx(rules.Delete))   // 删除规则
	}

}
