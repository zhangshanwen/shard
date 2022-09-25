package live

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/app"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/live/protocol/flv"
)

func Get(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("http flv handleConn panic: ", r)
		}
	}()
	var err error
	p := param.UriStrId{}
	if err = c.BindUri(&p); err != nil {
		c.String(http.StatusBadRequest, "invalid params")
		return
	}
	if pos := strings.LastIndex(p.Id, "."); pos < 0 || p.Id[pos:] != ".flv" {
		c.String(http.StatusBadRequest, "invalid path")
		return
	}
	// 判断视屏流是否发布,如果没有发布,直接返回404
	msgs := flv.GetStreams(app.S)
	path := strings.Replace(p.Id, ".flv", "", -1)
	if msgs == nil || len(msgs.Publishers) == 0 {
		c.String(http.StatusNotFound, "invalid path")
		return
	} else {
		include := false
		for _, item := range msgs.Publishers {
			if item.Key == common.LiveAppName+"/"+path {
				include = true
				break
			}
		}
		if include == false {
			c.String(http.StatusNotFound, "invalid path")
			return
		}
	}
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer := flv.NewFLVWriter(common.Live, path, c.Request.URL.String(), c.Writer)
	app.S.HandleWriter(writer)
	writer.Wait()
}
