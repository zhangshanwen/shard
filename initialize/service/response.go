package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/node"
)

type Res struct {
	StatusCode int
	ResCode    int
	ReqId      int64
	Data       interface{}
	Err        error
}

type res struct {
	Code  int         `json:"code"`
	ReqId int64       `json:"req_id"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func success(c *gin.Context, r Res) {
	c.JSON(http.StatusOK, res{
		Code:  code.Success,
		Msg:   "ok",
		Data:  r.Data,
		ReqId: r.ReqId,
	})
}

func failed(c *gin.Context, r Res) {
	var msg string
	if gin.IsDebugging() {
		msg = r.Err.Error()
	}
	if r.ResCode == 0 {
		r.ResCode = code.BaseFailedCode
	}
	c.JSON(http.StatusBadRequest, res{
		Code:  r.ResCode,
		Data:  nil,
		Msg:   msg,
		ReqId: r.ReqId,
	})
}

func Json(c *gin.Context, r Res) {
	r.ReqId = node.N.Generate()
	if r.StatusCode == 0 {
		if r.Err == nil {
			success(c, r)
		} else {
			failed(c, r)
		}
		return
	}
	var msg string
	if r.Err != nil {
		logrus.Error(r.Err)
		msg = r.Err.Error()
	}
	c.JSON(r.StatusCode, res{
		Code:  r.ResCode,
		Data:  r.Data,
		Msg:   msg,
		ReqId: r.ReqId,
	})
}
