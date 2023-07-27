package rtmp

import (
	"io"
	"sync"
)

type (
	Steam struct {
		m     map[string][]byte // 发布的视频流
		mutex sync.Mutex        // 锁
	}
)

var (
	S *Steam
)

func (s *Steam) Push(read io.ReadCloser, uid string) (err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var body []byte
	if body, err = io.ReadAll(read); err != nil {
		return
	}
	s.m[uid] = body
	return
}
func (s *Steam) Pull(uid string) (body []byte, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	body = s.m[uid]
	return
}
func init() {
	S = &Steam{
		m:     make(map[string][]byte),
		mutex: sync.Mutex{},
	}
}
