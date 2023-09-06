package reply_bot

import (
	"strings"

	"github.com/jinzhu/copier"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

// Create 创建新的定时发送消息机器人
func Create(c *service.AdminTxContext) (r service.Res) {
	p := param.SaveTimerBot{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		tx = c.Tx
		m  = model.TimerBot{
			Name: p.Name,
			Uid:  c.Admin.Id,
		}
	)
	var count int64
	if r.Err = tx.Model(&m).Where(&m).Count(&count).Error; r.Err != nil {
		return
	}
	if count > 0 {
		r.NameIsExisted()
		return
	}
	if r.Err = copier.Copy(&m, &p); r.Err != nil {
		r.CopierError()
		return
	}
	m.Friends = strings.Join(p.Friends, ",")
	m.Groups = strings.Join(p.Groups, ",")
	if r.Err = tx.Create(&m).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
