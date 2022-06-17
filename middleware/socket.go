package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/tools"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
)

func verifySocketJwt(c *gin.Context) (res service.Res, claims *tools.Claims) {
	h := param.UriAuthorization{}
	if res.Err = c.ShouldBindUri(&h); res.Err != nil {
		res.StatusCode = http.StatusUnauthorized
		res.ResCode = code.AuthFailed
		return
	}
	claims, res.Err = tools.VerifyToken(h.Authorization)
	return
}

func AdminSocketHandel(fun func(c *service.AdminContext)) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, claims := verifySocketJwt(c)
		logrus.Info("解析socket")
		if res.Err != nil {
			res.StatusCode = http.StatusUnauthorized
			res.ResCode = code.AuthFailed
			service.Json(c, res)
			return
		}
		logrus.Info("解析socket jwt成功", claims.Payload.Uid)
		sC := &service.AdminContext{Context: c}
		if res.Err = db.G.First(&sC.Admin, claims.Payload.Uid).Error; res.Err != nil {
			res.StatusCode = http.StatusUnauthorized
			res.ResCode = code.AuthFailed
			service.Json(c, res)
			return
		}
		logrus.Info("查询用户成功")
		fun(sC)
	}
}
