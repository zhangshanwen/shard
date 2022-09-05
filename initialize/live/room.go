package live

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/zhangshanwen/shard/initialize/node"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type (
	Room struct {
		clients []*websocket.Conn
		looker  map[string]*Client
		Owner   *websocket.Conn
		buf     bytes.Buffer
		mutex   sync.Mutex
	}
	Live struct {
		Room  map[string]*Room
		mutex sync.Mutex
	}

	Client struct {
		w      http.ResponseWriter
		RoomId string
		Uid    string
		End    chan bool
	}
)

var (
	L *Live
)

func NewClient(w http.ResponseWriter, roomId string) *Client {
	return &Client{
		RoomId: roomId,
		w:      w,
		Uid:    fmt.Sprintf("%v", node.N.Generate()),
	}
}
func (c *Client) leaveRoom() (err error) {
	return LeveRoom(c.RoomId, c.Uid)
}
func (c *Client) Wait() {
	var err error
	defer func() {
		if err = c.leaveRoom(); err != nil {
			logrus.Errorf("leave off room faield;%v", err)
		}
	}()
	select {
	case <-c.End:
		return
	}
}

func init() {
	once := sync.Once{}
	once.Do(func() {
		L = &Live{
			Room:  map[string]*Room{},
			mutex: sync.Mutex{},
		}
	})
}

func GetRoom(roomId string) (room *Room, err error) {
	L.mutex.Lock()
	defer L.mutex.Unlock()
	room = L.Room[roomId]
	if room == nil {
		err = errors.New("room is uncreated")
		return
	}
	return
}

func LeveRoom(roomId, uid string) (err error) {
	var room *Room
	if room, err = GetRoom(roomId); err != nil {
		return
	}
	room.LeaveRoom(uid)
	return
}

func (r *Room) LeaveRoom(uid string) {
	if r == nil {
		return
	}
	delete(r.looker, uid)
}

func (l *Live) CreateRoom(roomId string, conn *websocket.Conn) (room *Room, err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.Room[roomId] != nil {
		return nil, errors.New("room is created")
	}
	room = &Room{
		clients: []*websocket.Conn{},
		Owner:   conn,
		looker:  map[string]*Client{},
		mutex:   sync.Mutex{},
	}
	l.Room[roomId] = room
	return
}

func (l *Live) JoinInRoom(c *Client) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	room := l.Room[c.RoomId]
	if room == nil {
		return errors.New("room is uncreated")
	}
	room.looker[c.Uid] = c
	return
}

func (l *Live) JoinRoom(roomId string, conn *websocket.Conn) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	room := l.Room[roomId]
	if room == nil {
		return errors.New("room is uncreated")
	}
	room.clients = append(room.clients, conn)
	return
}

func (l *Live) LeaveRoom(roomId string, conn *websocket.Conn) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	room := l.Room[roomId]
	if room == nil {
		return errors.New("room is uncreated")
	}
	for index, c := range room.clients {
		if c == conn {
			room.clients = append(room.clients[:index], room.clients[index:]...)
		}
	}
	return
}

func (l *Live) DisposeRoom(roomId string) (err error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	room := l.Room[roomId]
	if room == nil {
		return errors.New("room is uncreated")
	}
	defer room.Owner.Close()
	l.Room[roomId] = nil
	return
}
func (r *Room) Run() {
	var (
		message     []byte
		err         error
		messageType int
	)
	for {
		time.Sleep(100 * time.Millisecond)
		if messageType, message, err = r.Owner.ReadMessage(); err != nil {
			logrus.Error("读取消息失败", err)
			break
		}
		for _, client := range r.clients {
			go client.WriteMessage(messageType, message)
		}
		for _, c := range r.looker {
			go func(w http.ResponseWriter) {
				_, _ = w.Write(message)
			}(c.w)
		}
	}
}

func (r *Room) JoinInRoom(c *Client) (err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.looker[c.Uid] = c
	return
}
