package meeting

import (
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/rtmp"
)

func Push(c *service.AdminContext) (r service.Res) {
	if !rtmp.S.IsRunning(c.Admin.Id) {
		r.PushFailed()
		return
	}
	p := param.UriId{}
	if r.Err = c.BindUri(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	// 将视频流写入内存
	if r.Err = rtmp.S.Push(c.Request.Body, p.Id); r.Err != nil {
		r.PushFailed()
		return
	}
	return
}
