package wechat

import (
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/tools/wechat"
)

func QrCode(c *service.AdminContext) (r service.Res) {
	var (
		w    = wechat.NewWechat()
		code string
	)
	if code, r.Err = w.Qrcode(c.Admin.Id); r.Err != nil {
		return
	}
	r.Data = code
	return
}
