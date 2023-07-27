package host

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func Socket(c *service.AdminContext) {
	var (
		err  error
		conn *websocket.Conn
		p    param.Socket
		id   int64
		host model.Host
		ss   *tools.SshSocket
	)
	if err = c.BindUri(&p); err != nil {
		logrus.Error(err)
		return
	}
	id, _ = db.R.Get(c, p.Id).Int64()
	if id <= 0 {
		logrus.Error("id失效")
		return
	}
	if err = db.G.First(&host, id).Error; err != nil {
		logrus.Error(err)
		return
	}

	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	var (
		m  model.Host
		tx = db.G.Begin()
	)

	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	defer conn.Close()
	if ss, err = tools.NewSshSocket(host.Username, host.Password, host.Host, host.Port, host.Id); err != nil {
		_ = tx.Model(&m).Where("id=? and status=?", host.Id, 1).Update("status", 0)
		logrus.Errorf("创建ssh连接失败")
		return
	}
	_ = tx.Model(&m).Where("id=? and status=?", host.Id, 0).Update("status", 1)
	c.SaveLogAdd(tx, module, fmt.Sprintf("terminal (%v) id:%v name:%v", p.Id, host.Id, host.Name))
	ss.Run(conn)
	logrus.Info("任务执行结束,断开连接")

}
