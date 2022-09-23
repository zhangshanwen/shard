package live

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
)

//Get 获取直播推送地址并加入房间
func Get(c *service.AdminTxContext) (r service.Res) {
	p := param.LiveRoom{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ResCode = code.ParamsError
		return
	}
	var (
		resp response.LiveResponse
	)

	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	resp.Url = db.R.Get(c, p.Hash).String()
	return
}
