package host

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
)

// Room 创建房间
func Room(c *service.AdminContext) (r service.Res) {
	pId := param.Room{}
	if r.Err = c.Rebind(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	// 获取房间信息
	var (
		m      model.Host
		roomId int64
		res    = response.Room{}
		tx     = db.G.Begin()
	)

	defer func() {
		r.Data = res
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if r.Err = tx.First(&m, pId.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}
	key := fmt.Sprintf("%v_%v_%v", c.Admin.Id, pId.Id, time.Now().Unix())
	key = fmt.Sprintf("%x", sha256.Sum256([]byte(key)))
	// 查询房间是存在
	roomId, _ = db.R.Get(c.Context, key).Int64()
	if roomId > 0 {
		r.RoomExisted()
		return
	}
	if r.Err = db.R.SetNX(c.Context, key, m.Id, time.Minute*5).Err(); r.Err != nil {
		r.DBError()
		return
	}
	res.Id = key
	c.SaveLog(tx, fmt.Sprintf("创建房间(%v) id:%v name:%v connect_type:%v", key, m.Id, m.Name, m.ConnectType), model.OperateLogTypeAdd)
	return
}
