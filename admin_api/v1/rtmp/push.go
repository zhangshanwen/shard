package rtmp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/rtmp"
)

func Push(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Error("http flv handleConn panic: ", r)
		}
	}()
	var err error
	p := param.UriStrId{}
	if err = c.BindUri(&p); err != nil {
		c.String(http.StatusBadRequest, "invalid params")
		return
	}
	// 将视频流写入内存
	if err = rtmp.S.Push(c.Request.Body, p.Id); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

}
