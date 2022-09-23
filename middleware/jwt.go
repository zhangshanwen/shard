package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/header"
	"github.com/zhangshanwen/shard/tools"
)

func verifyJwt(c *gin.Context) (r service.Res, claims *tools.Claims) {
	h := header.Authorization{}
	if r.Err = c.ShouldBindHeader(&h); r.Err != nil {
		r.StatusCode = http.StatusUnauthorized
		r.ResCode = code.AuthFailed
		return
	}
	claims, r.Err = tools.VerifyToken(h.Authorization)
	return
}
func JwtHandel(fun func(*service.Context) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, claims := verifyJwt(c)
		if r.Err != nil {
			r.AuthFailed()
			service.Json(c, r)
			return
		}
		sC := &service.Context{Context: c}
		if r.Err = sC.GetLoginInfo(claims.Payload.Uid); r.Err != nil {
			r.AuthFailed()
			service.Json(c, r)
			return
		}
		service.Json(c, fun(sC))
	}
}

func AdminJwtHandel(fun func(ctx *service.AdminContext) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, claims := verifyJwt(c)
		if r.Err != nil {
			r.AuthFailed()
			service.Json(c, r)
			return
		}
		sC := &service.AdminContext{Context: c}
		if r.Err = sC.GetLoginInfo(claims.Payload.Uid); r.Err != nil {
			r.AuthFailed()
			service.Json(c, r)
			return
		}
		if r.Err = verifyPermission(sC); r.Err != nil {
			r.NoPermission()
			service.Json(c, r)
			return
		}

		service.Json(c, fun(sC))
	}
}

func AdminJwtTxHandel(fun func(ctx *service.AdminTxContext) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, claims := verifyJwt(c)
		if r.Err != nil {
			r.AuthFailed()
			service.Json(c, r)
			return
		}
		sC := &service.AdminTxContext{}
		sC.Context = c
		sC.Tx = db.G.Begin()
		if r.Err = sC.GetLoginInfo(claims.Payload.Uid); r.Err != nil {
			r.AuthFailed()
			service.Json(c, r)
			return
		}
		//if r.Err = verifyPermission(&sC.AdminContext); r.Err != nil {
		//	r.NoPermission()
		//	service.Json(c, r)
		//	return
		//}
		r = fun(sC)
		if r.Err == nil {
			sC.Tx.Commit()
		} else {
			sC.Tx.Rollback()
		}
		service.Json(c, r)
	}
}

func verifyPermission(c *service.AdminContext) (err error) {
	// verify last one  is number
	path := c.Request.URL.Path
	urlSplit := strings.Split(c.Request.URL.Path, common.Backlash)
	_, err = strconv.Atoi(urlSplit[len(urlSplit)-1])
	if err == nil {
		urlSplit[len(urlSplit)-1] = common.UriId
		path = strings.Join(urlSplit, common.Backlash)
	}
	key := fmt.Sprintf(common.RedisRoutesKey, c.Admin.Id)
	field := fmt.Sprintf("%s||%s", c.Request.Method, path)
	var val bool
	if val, err = db.R.HGet(c, key, field).Bool(); err != nil {
		return
	}
	if !val {
		return common.NoPermissionErr
	}
	return
}
