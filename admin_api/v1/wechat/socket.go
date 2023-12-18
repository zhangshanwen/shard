package wechat

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/zhangshanwen/shard/initialize/service"
)

// Socket 创建与前端页面的连接
func Socket(c *service.AdminWechatContext) {
	var (
		err     error
		conn    *websocket.Conn
		message []byte
	)

	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	if conn, err = upGrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	c.Bot.CleanMessages()

	defer func() { _ = conn.Close() }()
	go func() {
		for {
			select {
			case <-c.Done():
				logrus.Info("websocket断开连接")
				return
			case <-c.Bot.Context().Done():
				logrus.Info("机器人断开连接")
				return
			case m := <-c.Bot.Messages:
				if err = conn.WriteMessage(websocket.TextMessage, m); err != nil {
					logrus.Errorf("%v写入消息失败:", err)
					return
				}

			}
		}
	}()
	for {
		if _, message, err = conn.ReadMessage(); err != nil {
			logrus.Errorf("读取消息失败:%v", err)
			break
		}
		if err = c.Bot.ReceiveMessage(message); err != nil {
			logrus.Warningf("处理消息失败:%v", err)
		}
	}

}
