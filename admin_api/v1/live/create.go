package live

import (
	"fmt"
	"time"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/tools"
)

func Post(c *service.AdminTxContext) (r service.Res) {
	p := param.CreateLiveRoom{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	key := fmt.Sprintf("%v_%v", c.Admin.Id, p.Name)
	hash := tools.Hash(key)
	r.Err = db.R.Set(c, hash, p.Name, 10*time.Minute).Err()
	r.Data = hash
	return
}
