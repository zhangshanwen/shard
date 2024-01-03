package meeting

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

type (
	Member struct {
		Mid            int64                `json:"mid"`
		NickName       string               `json:"nick_name"`
		IsOwner        bool                 `json:"is_owner"`
		IsCapture      bool                 `json:"is_capture"`
		meeting        *Meeting             // 会议
		captures       []string             // 屏幕捕捉 整段影片
		capturesOffset int                  // 屏幕捕捉 偏移量
		cameras        []string             // 摄像头 整段影片
		camerasOffset  int                  // 摄像头 偏移量
		voices         []string             // 语音 整段影片
		media          chan MessageRespBody // 媒体 管道
		voicesOffset   int                  // 语音 偏移量
		voiceMap       map[int64]bool       // 判断是否id 是否给 发过语音
		captureMap     map[int64]bool       // 判断是否id 是否给 发过屏幕捕捉
		cameraMap      map[int64]bool       // 判断是否id 是否给 发过摄像头
		conn           *websocket.Conn
		mutex          sync.Mutex
		done           func() <-chan struct{}
	}
)

func (m *Member) ReceiveMessage(message []byte) (err error) {

	var (
		mb MessageReqBody
	)
	if err = json.Unmarshal(message, &mb); err != nil {
		return err
	}
	switch mb.MsgType {
	case "stop":
	case "ping":
		err = m.sendMessage(&MessageRespBody{
			MsgType: "ping",
			SendId:  m.Mid,
		})
	case Capture:
		err = m.getCaptures(mb.Data)
	case Camera:
		err = m.getCameras(mb.Data)
	case Voice:
		err = m.getVoices(mb.Data)
	}

	return err
}

func NewMember(conn *websocket.Conn, mid, roomId int64, nickName string, isOwner bool, done func() <-chan struct{}) (*Member, error) {
	member := &Member{
		Mid:        mid,
		NickName:   nickName,
		IsOwner:    isOwner,
		conn:       conn,
		mutex:      sync.Mutex{},
		done:       done,
		media:      make(chan MessageRespBody, chanLen),
		voiceMap:   make(map[int64]bool),
		captureMap: make(map[int64]bool),
		cameraMap:  make(map[int64]bool),
	}
	logrus.Infof("roomID=%v", roomId)
	meeting := newMeeting(roomId)
	defer member.checkAlive()
	logrus.Infof("meeting_id=%v", meeting.Id)
	return member, meeting.AddMember(member)
}

func (m *Member) stop() (err error) {
	// 如果主持人在仅主持人可以终止会员,其他人只可退出,所有人退出自动终止
	return
}

func (m *Member) getCaptures(mb MediaBody) (err error) {
	m.capturesOffset += mb.Offset
	logrus.Infof("上传总时长 %v", m.capturesOffset/1000)
	m.captures = append(m.captures, mb.Data)
	for _, member := range m.others() {
		go func(member *Member) {
			if !member.captureMap[m.Mid] {
				mb.Data = strings.Join(m.captures, "")
				mb.Offset = m.capturesOffset
				member.captureMap[m.Mid] = true
			}
			member.media <- MessageRespBody{
				MsgType: Capture,
				Data:    mb,
				SendId:  member.Mid}
		}(member)
	}
	return
}
func (m *Member) getCameras(mb MediaBody) (err error) {
	m.camerasOffset += mb.Offset
	m.cameras = append(m.cameras, mb.Data)
	for _, member := range m.others() {
		go func(member *Member) {
			var body = mb.Data
			if !member.cameraMap[m.Mid] {
				body = strings.Join(m.cameras, "")
				member.cameraMap[m.Mid] = true
			}
			member.media <- MessageRespBody{
				MsgType: Camera,
				Data:    MediaBody{Data: body, Offset: m.camerasOffset},
				SendId:  member.Mid}
		}(member)
	}
	return
}

func (m *Member) getVoices(mb MediaBody) (err error) {
	m.voicesOffset += mb.Offset
	m.voices = append(m.voices, mb.Data)
	for _, member := range m.others() {
		go func(member *Member) {
			var body = mb.Data
			if !member.voiceMap[m.Mid] {
				body = strings.Join(m.voices, "")
				member.voiceMap[m.Mid] = true
			}
			member.media <- MessageRespBody{
				MsgType: Voice,
				Data:    MediaBody{Data: body, Offset: m.voicesOffset},
				SendId:  member.Mid}
		}(member)
	}
	return
}

func (m *Member) others() []*Member {
	var members []*Member
	for _, member := range m.meeting.Members {
		if member.Mid != m.Mid {
			members = append(members, member)
		}
	}
	return members
}

func (m *Member) sendMessage(resp *MessageRespBody) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	var (
		body []byte
	)
	if body, err = json.Marshal(resp); err != nil {
		return
	}
	return m.conn.WriteMessage(websocket.TextMessage, body)
}

func (m *Member) checkAlive() {
	var err error
	go func() {
		for {
			select {
			case mrb := <-m.media:
				if err = m.sendMessage(&mrb); err != nil {
					goto quit
				}
			case <-m.done():
				goto quit
			}
		}
	quit:
		m.meeting.quitMid <- m.Mid

	}()
}
