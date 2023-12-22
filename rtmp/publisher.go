package rtmp

import (
	"fmt"
	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	Publisher struct {
		members     []*Member
		uid         int64
		pid         int64
		isPushing   bool
		isAlive     bool
		index       int
		stop        chan struct{}
		movie       chan []byte
		movies      []byte
		createdTime int64
		pushTime    int64
	}
)

func (p *Publisher) MeetingOver() {
	db.G.Model(model.Meeting{}).Where("id=?", p.uid).Updates(map[string]interface{}{
		"end_time": time.Now().Unix(),
		"status":   model.MeetingStatusEnd,
	})
	p.writeIntoFile()
	return
}

func (p *Publisher) writeIntoFile() {
	var (
		fileName = fmt.Sprintf("webm/%v/%v.webm", p.uid, time.Now().Format(common.TimeCrushFormat))
		file     *os.File
		err      error
	)
	defer func() {
		if err != nil {
			logrus.Errorf("写入文件失败%v", err)
		}
	}()
	if file, err = os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm); err != nil {
		return
	}
	_, err = file.Write(p.movies)
}

func (p *Publisher) removeMember(index int) {

}

func (p *Publisher) done() {
	for i := 0; i < len(p.members); i++ {
		p.members[i].stop <- struct{}{}
	}
}
