package tools

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type (
	SshSocket struct {
		client  *ssh.Client
		session *ssh.Session
		buffer  bytes.Buffer
		mx      sync.Mutex
		quit    chan bool
		speed   int
	}
)

const (
	defaultSpeed = 100
)

func (s *SshSocket) Write(p []byte) (n int, err error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.buffer.Write(p)
}

func (s *SshSocket) SetSpeed(speed int) (err error) {
	if speed <= 0 {
		return errors.New("error speed")
	}
	return
}
func NewSshSocket(username, password, host string, port int) (ss *SshSocket, err error) {
	ss = &SshSocket{
		quit:  make(chan bool),
		speed: defaultSpeed,
	}
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if ss.client, err = ssh.Dial("tcp", fmt.Sprintf("%v:%v", host, port), config); err != nil {
		err = errors.New(fmt.Sprintf("创建连接失败:%v", err))
		return
	}
	return
}

func (s *SshSocket) Close() (err error) {
	if s.client != nil {
		if err = s.client.Close(); err != nil {
			return err
		}
	}
	if s.session != nil {
		return s.session.Close()
	}
	return
}
func (s *SshSocket) Session() (err error) {
	s.session, err = s.client.NewSession()
	return
}

func (s *SshSocket) Run(ws *websocket.Conn) {

	var (
		err       error
		message   []byte
		sessionIn io.WriteCloser
	)
	if err = s.Session(); err != nil {
		logrus.Error("创建session失败", err)
		return
	}
	defer s.Close()
	s.session.Stdout = s
	s.session.Stderr = s
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // 禁用回显（0禁用，1启动）
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	}
	if err = s.session.RequestPty("xterm", 32, 160, modes); err != nil {
		logrus.Error("创建平台", err)
		return
	}
	_ = s.session.Setenv("LANG", "zh_CN.UTF-8")
	if sessionIn, err = s.session.StdinPipe(); err != nil {
		logrus.Error("生成StdinPipe失败", err)
		return
	}
	if err = s.session.Shell(); err != nil {
		logrus.Error("创建shell失败", err)
		return
	}
	go func() {
		for {
			s.sleep()
			if _, message, err = ws.ReadMessage(); err != nil {
				logrus.Error("读取消息失败", err)
				break
			}
			if message == nil {
				continue
			}
			if _, err = sessionIn.Write(message); err != nil {
				logrus.Error("写入ssh消息失败", err)
				break
			}
		}
		s.quit <- true
	}()
	go func() {
		for {
			s.sleep()
			if s.buffer.Len() != 0 {
				if err = ws.WriteMessage(websocket.TextMessage, s.buffer.Bytes()); err != nil {
					logrus.Error("写入ws消息失败", err)
					break
				}
				s.buffer.Reset()
			}
		}
		s.quit <- true
	}()

	go func() {
		if err = s.session.Wait(); err != nil {
			logrus.Error("等待结束", err)
		}
		s.quit <- true

	}()

	<-s.quit
	return
}

func (s *SshSocket) sleep() {
	time.Sleep(time.Duration(s.speed) * time.Millisecond)
}

func (s *SshSocket) CombinedOutput(cmd string) (b []byte, err error) {
	if s.session == nil {
		return nil, errors.New("null session")
	}
	return s.session.CombinedOutput(cmd)
}
