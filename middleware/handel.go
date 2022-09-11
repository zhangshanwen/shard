package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
)

func Handel(fun func(c *service.Context) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		service.Json(c, fun(&service.Context{Context: c}))
	}
}

func AdminHandel(fun func(c *service.AdminContext) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		service.Json(c, fun(&service.AdminContext{Context: c}))
	}
}

func AdminTxHandel(fun func(c *service.AdminTxContext) service.Res) gin.HandlerFunc {
	return func(c *gin.Context) {
		sC := &service.AdminTxContext{}
		sC.Tx = db.G.Begin()
		sC.Context = c
		r := fun(sC)
		if r.Err == nil {
			sC.Tx.Commit()
		} else {
			sC.Tx.Rollback()
		}
		service.Json(c, r)
	}
}
