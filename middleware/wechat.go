package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/tools/wechat"
)

func AdminWechatHandel(fun func(c *service.AdminWechatContext) service.Res) gin.HandlerFunc {
	c := &service.AdminWechatContext{}
	var (
		w = wechat.NewWechat()
	)
	m := AdminJwtHandel(func(ctx *service.AdminContext) service.Res {
		c.AdminContext = *ctx
		c.Bot = w.Bot(ctx.Admin.Id)
		res := fun(c)
		if c.Bot == nil {
			res.LoginWechatFailed()
		}
		return res
	})
	return m
}

func AdminWechatSocketHandel(fun func(c *service.AdminWechatContext)) gin.HandlerFunc {
	c := &service.AdminWechatContext{}
	var (
		w = wechat.NewWechat()
	)
	m := AdminSocketHandel(func(ctx *service.AdminContext) {
		c.AdminContext = *ctx
		c.Bot = w.Bot(ctx.Admin.Id)
		fun(c)
		return
	})
	return m

}
