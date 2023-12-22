package meeting

import (
	"errors"
	"github.com/zhangshanwen/shard/rtmp"

	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

// Create 创建会议室
func Create(c *service.AdminTxContext) (r service.Res) {
	p := param.CreateMeeting{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx = c.Tx
		m  model.Meeting
	)
	if rtmp.S.IsRunning(c.Admin.Id) {
		r.RoomExisted()
		return
	}
	r.Err = tx.Where("status = ? and uid = ? ", model.MeetingStatusRunning, c.Admin.Id).First(&m).Error
	if r.Err != nil && !errors.Is(r.Err, gorm.ErrRecordNotFound) {
		r.DBError()
		return
	} else if r.Err == nil {
		r.Data = m
		r.RoomExisted()
		return
	}
	if r.Err = tx.Save(&model.Meeting{Uid: c.Admin.Id, Name: p.Name}).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
