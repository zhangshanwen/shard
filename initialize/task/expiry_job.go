package task

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
)

type (
	ExpiryJob struct {
		Id       int64
		ExecTime int64
	}
)

func (t *ExpiryJob) Run() {
	now := time.Now().Unix()
	mm := int64(60 * 60 * 24)
	if now-t.ExecTime > mm && now-t.ExecTime < -mm {
		logrus.Errorf("执行生效任务失败:生效时间:%v,当前时间:%v", t.ExecTime, now)
		return
	}
	logrus.Infof("开始执行生效任务")
	defer T.StopExpire(t.Id)
	var (
		err error
		m   model.Task
		tx  = db.G.Begin()
	)
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	// 删除开始、结束任务如果存在
	if err = tx.Model(&m).Where("id=?", t.Id).Update("status", model.StatusExpiry).Error; err != nil {
		logrus.Errorf("过期任务执行失败,任务id:%v,err:%v\n", t.Id, err.Error())
		return
	}
	logrus.Infof("过期任务执行完成,任务id:%v \n", t.Id)
}
