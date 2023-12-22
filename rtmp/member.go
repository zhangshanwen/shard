package rtmp

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type (
	Member struct {
		mid          int64
		w            http.ResponseWriter
		publisher    *Publisher
		isPushMovies bool
		done         func() <-chan struct{}
		stop         chan struct{}
	}
)

func (m *Member) Wait() {
	select {
	case <-m.stop:
	case <-m.done():
	}
}

func (m *Member) Quit() {
	logrus.Infof("成员%v退出", m.mid)
	for i := 0; i < len(m.publisher.members); i++ {
		if m.publisher.members[i].mid == m.mid {
			m.publisher.members = append(m.publisher.members[:i], m.publisher.members[i+1:]...)
			break
		}
	}
}
