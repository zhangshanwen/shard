package meeting

import (
	"github.com/sirupsen/logrus"
	"github.com/zhangshanwen/shard/initialize/node"
	"sync"
)

type (
	Steam struct {
		Meetings map[int64]*Meeting
	}
	Meeting struct {
		Id      int64
		mutex   sync.Mutex // 锁
		Members []*Member
		stop    chan struct{}
		quitMid chan int64
	}
)

var (
	s           *Steam
	meetingOnce sync.Once
)

const (
	chanLen = 50
)

func newMeeting(mid int64) *Meeting {
	meetingOnce.Do(func() {
		s = &Steam{
			Meetings: make(map[int64]*Meeting),
		}
	})
	meeting := s.Meetings[mid]
	if meeting == nil {
		meeting = &Meeting{
			Id:      node.N.Generate(),
			mutex:   sync.Mutex{},
			stop:    make(chan struct{}),
			quitMid: make(chan int64),
		}
		s.Meetings[mid] = meeting
	}
	return meeting
}
func (m *Meeting) AddMember(member *Member) (err error) {

	member.meeting = m
	m.Members = append(m.Members, member)
	return
}

func (m *Meeting) hasCaptureMember(uid int64) bool {
	for _, member := range m.Members {
		if uid > 0 {
			if member.IsCapture && member.Mid != uid {
				return true
			}
		} else {
			if member.IsCapture {
				return true
			}
		}
	}
	return false
}
func (m *Meeting) hasOwnerMember(uid int64) bool {
	for _, member := range m.Members {
		if uid > 0 {
			if member.IsOwner && member.Mid != uid {
				return true
			}
		} else {
			if member.IsOwner {
				return true
			}
		}
	}
	return false
}

func (m *Meeting) checkMemberAlive() {
	go func() {
		for {
			select {
			case <-m.stop:
				return
			case quitMid := <-m.quitMid:
				for i := 0; i < len(m.Members); i++ {
					if m.Members[i].Mid == quitMid {
						logrus.Infof("成员%v 退出了", quitMid)
						m.Members = append(m.Members[:i], m.Members[i+1:]...)
						break
					}
				}
			}
		}
	}()
}
