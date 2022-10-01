package live

import (
	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Create(c *service.AdminTxContext) (r service.Res) {
	p := param.CreateLiveRoom{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		m  model.LiveRoom
		tx = c.Tx
	)
	res := tx.First(&m, "`owner_id`=?", c.Admin.Id)
	if res.Error == nil {
		r.RoomExisted()
		return
	} else if res.Error != gorm.ErrRecordNotFound {
		r.DBError()
		return
	}
	m.Name = p.Name
	m.OwnerId = c.Admin.Id
	m.Status = model.LiveRoomStatusIdle
	if r.Err = tx.Save(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
