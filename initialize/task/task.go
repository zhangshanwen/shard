package task

import (
	"errors"
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
)

const (
	MaxPageSize  = 1000
	MaxQueueSize = 10
)

var (
	T *Task
)

type (
	Task struct {
		Cron  *cron.Cron
		m     sync.Map
		queue chan bool
		nums  int
	}
	RunTask interface {
		Run()
		GetId() int64
		GetSpec() string
	}
)

func InitTask() {
	T = &Task{
		Cron:  cron.New(),
		nums:  0,
		queue: make(chan bool, MaxQueueSize),
	}
	_ = T.Initialize()
}

// Initialize 初始化任务, 从数据库取出所有任务, 添加到定时任务并运行
func (t *Task) Initialize() (err error) {
	t.Cron.Start()
	logrus.Info("开始初始化定时任务")
	var taskList []model.Task
	page := 1
	pageSize := MaxPageSize
	var total int64
	g := db.G.Model(model.Task{}).Where("status=?", model.StatusNormal)
	if err = g.Count(&total).Error; err != nil {
		return
	}
	for int64((page-1)*pageSize) < total {
		if err = g.Limit(pageSize).Offset((page - 1) * pageSize).Find(&taskList).Error; err != nil {
			err = errors.New(fmt.Sprintf("定时任务初始化#获取任务列表错误: %s", err.Error()))
			return
		}
		if len(taskList) == 0 {
			break
		}
		for _, item := range taskList {
			if err = t.Add(&item); err != nil {
				return
			}
		}
		page++
	}
	logrus.Infof("定时任务初始化完成, 共%d个定时任务添加到调度器", t.nums)
	return
}

func (t *Task) Add(p RunTask) (err error) {
	var jobId cron.EntryID
	if jobId, err = t.Cron.AddJob(p.GetSpec(), p); err != nil {
		return
	}
	t.nums += 1
	t.m.Store(p.GetId(), jobId)
	return
}

func (t *Task) Stop(id int64) (err error) {
	if v, ok1 := t.m.Load(id); ok1 {
		if jobId, ok2 := v.(cron.EntryID); ok2 {
			t.Cron.Remove(jobId)
			logrus.Infof("任务:%v 停止成功", id)
			t.nums--
			return
		}
	}
	return errors.New("错误id")
}
