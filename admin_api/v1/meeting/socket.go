package meeting

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools/meeting"
)

func Socket(c *service.AdminContext) {
	var (
		err     error
		conn    *websocket.Conn
		message []byte
		p       param.SocketInt
		m       model.Meeting
		member  *meeting.Member
	)

	if err = c.BindUri(&p); err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("pid=%v", p.Id)

	if err = db.G.First(&m, p.Id).Error; err != nil {
		logrus.Error(err)
		return
	}
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	if member, err = meeting.NewMember(conn, c.Admin.Id, p.Id, c.Admin.Username, c.Admin.Id == m.Uid, c.Done); err != nil {
		logrus.Error(err)
		return
	}
	if m.Status == model.MeetingStatusPending {
		m.Status = model.MeetingStatusRunning
		m.StartTime = time.Now().Unix()
		db.G.Save(&m)
	}

	defer func() { _ = conn.Close() }()
	go func() {
		for {
			select {
			case <-c.Done():
				logrus.Info("websocket断开连接")
				return

			}
		}
	}()
	for {
		if _, message, err = conn.ReadMessage(); err != nil {
			logrus.Errorf("读取消息失败:%v", err)
			break
		}
		if err = member.ReceiveMessage(message); err != nil {
			logrus.Warningf("处理消息失败:%v", err)
		}
	}

}
