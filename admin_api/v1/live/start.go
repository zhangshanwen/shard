package live

import (
	"fmt"
	"time"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func Start(c *service.AdminTxContext) (r service.Res) {
	var (
		m  model.LiveRoom
		tx = c.Tx
	)
	if r.Err = tx.First(&m, "`owner_id`=?", c.Admin.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}

	if m.Status < model.LiveRoomStatusIdle {
		if r.Err = tx.Model(&m).Update("status", model.LiveRoomStatusIdle).Error; r.Err != nil {
			r.DBError()
			return
		}
	} else if m.Status == model.LiveRoomStatusLiving {
		r.Living()
		return
	}
	key := fmt.Sprintf("%v_%v", c.Admin.Id, m.Id)
	hash := tools.Hash(key)
	r.Err = db.R.Set(c, hash, m.Id, 10*time.Minute).Err()
	r.Data = hash
	return
}
