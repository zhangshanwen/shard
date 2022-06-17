package task

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
)

type (
	EffectJob struct {
		Id       int64
		ExecTime int64
	}
)

func (t *EffectJob) Run() {
	now := time.Now().Unix()
	mm := int64(60 * 60 * 24)
	if now-t.ExecTime > mm && now-t.ExecTime < -mm {
		logrus.Errorf("执行生效任务失败:生效时间:%v,当前时间:%v", t.ExecTime, now)
		return
	}
	logrus.Infof("开始执行生效任务")
	defer T.StopEffect(t.Id)
	// 添加 开始、结束执行任务
	T.Begin()
	var (
		err error
		tx  = db.G.Begin()
		m   model.Task
	)
	defer func() {
		if err == nil {
			tx.Commit()
			T.Commit()
		} else {
			tx.Rollback()
			T.RollBack()
		}
	}()
	if err = tx.First(&m, t.Id).Error; err != nil {
		logrus.Errorf("生效任务执行失败,任务id:%v,err:%v\n", t.Id, err.Error())
		// 任务添加失败
		return
	}
	if err = T.pushRunCron(m, tx); err != nil {
		// 任务添加失败
		logrus.Errorf("生效任务执行失败,任务id:%v,err:%v\n", t.Id, err.Error())
		return
	}
	if err = tx.Model(&m).Where("id=?", t.Id).Update("status", model.StatusRunning).Error; err != nil {
		logrus.Errorf("生效任务执行失败,任务id:%v,err:%v\n", t.Id, err.Error())
		return
	}
	logrus.Infof("生效任务执行完成,任务id:%v \n", t.Id)
}
