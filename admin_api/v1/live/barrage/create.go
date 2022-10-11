package barrage

import (
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
)

func Create(c *service.AdminTxContext) (r service.Res) {
	p := param.CreateBarrage{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	//向队列发送弹幕消息
	db.R.RPush(c, p.Hash, p.Content)
	return
}
