package rtmp

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/node"
	"github.com/zhangshanwen/shard/model"
)

type (
	Steam struct {
		mutex      sync.Mutex // 锁
		Publishers map[int64]*Publisher
	}
)

var (
	S *Steam
)

func (s *Steam) NewPublisher(uid int64) {
	s.Publishers[uid] = &Publisher{
		members:   []*Member{},
		pid:       node.N.Generate(),
		uid:       uid,
		isPushing: false,
		stop:      make(chan struct{}),
		movie:     make(chan []byte, 10),
	}
}

func (s *Steam) IsRunning(uid int64) bool {
	return s.Publishers[uid] != nil
}
func (s *Steam) IsPushing(uid int64) bool {
	return s.Publishers[uid] != nil && s.Publishers[uid].isPushing
}

func (s *Steam) Offset(uid int64) (offset int64) {
	if s.Publishers[uid] != nil {
		return s.Publishers[uid].pushTime - s.Publishers[uid].createdTime
	}
	return
}
func (s *Steam) Push(read io.ReadCloser, uid int64) (err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var body []byte
	p := s.Publishers[uid]
	if body, err = io.ReadAll(read); err != nil {
		return
	}
	p.isAlive = true
	go func() {
		p.movie <- body
	}()
	p.pushTime = time.Now().Unix()
	if !p.isPushing {
		p.createdTime = time.Now().Unix()
		db.G.Model(model.Meeting{}).Where("id=?", uid).Updates(map[string]interface{}{
			"start_time": p.createdTime,
			"status":     model.MeetingStatusRunning,
		})
		p.isPushing = true
		go func() {
			for {
				select {
				case <-time.After(time.Second * 30):
					//30s 检测一次
					if p.isAlive {
						p.isAlive = false
					} else {
						p.stop <- struct{}{}
						return
					}
				}
			}
		}()
		go func() {
			for {
				select {
				case <-p.stop:
					p.done()
					p.MeetingOver()
					break
				case v := <-p.movie:
					logrus.Infof("获取到数据,开始向每个成员推送 %v.......", len(v))
					p.movies = append(p.movies, v...)
					p.index += 1
					var i = 0

					for {
						if i >= len(p.members) {
							break
						}
						member := p.members[i]
						var pushData []byte
						if p.index != 1 && !member.isPushMovies {
							member.isPushMovies = true
							pushData = bytes.Clone(p.movies)
						} else if p.index == 1 {
							member.isPushMovies = true
							pushData = bytes.Clone(v)
						} else {
							pushData = bytes.Clone(v)
						}
						if member == nil {
							p.members = append(p.members[:i], p.members[i+1:]...)
							logrus.Infof("成员不存在")
							continue
						}
						logrus.Infof("向成员%v推送数据,%v", member.mid, len(pushData))
						if _, err = member.w.Write(pushData); err != nil {
							p.members = append(p.members[:i], p.members[i+1:]...)
							logrus.Infof("成员%v退出", member.mid)
						}
					}
				default:
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()
	}
	return
}

func (s *Steam) AddMember(uid int64, done func() <-chan struct{}, w http.ResponseWriter) (member *Member, err error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Transfer-Encoding", "chunked")
	p := s.Publishers[uid]
	if p == nil {
		err = errors.New("no Publisher")
		return
	}
	member = &Member{w: w, mid: node.N.Generate(), publisher: p, done: done}
	p.members = append(p.members, member)
	logrus.Infof("成员%v加入", member.mid)
	return
}

func init() {
	S = &Steam{
		mutex:      sync.Mutex{},
		Publishers: make(map[int64]*Publisher),
	}
}
