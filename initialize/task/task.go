package task

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
)

const (
	MaxPageSize  = 1000
	MaxQueueSize = 10
	OnceFormat   = "04 15 02 01"
)
const (
	EffectType taskType = iota
	runType
	ExpiryType
)

var (
	T *Task
)

type (
	taskType int
	Task     struct {
		Cron     *cron.Cron
		m        sync.Map
		queue    chan bool
		nums     int
		EntryIDs []cron.EntryID // 存放事务的id列表
	}
)

func newTask() *Task {
	return &Task{
		Cron:     cron.New(),
		nums:     0,
		queue:    make(chan bool, MaxQueueSize),
		EntryIDs: []cron.EntryID{},
	}
}

func InitTask() {
	T = newTask()
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
	g := db.G.Begin()
	defer func() {
		if err != nil {
			g.Rollback()
		} else {
			g.Commit()
		}
	}()
	fg := g.Model(model.Task{}).Where("status>=?", model.StatusIdle)
	if err = fg.Count(&total).Error; err != nil {
		return
	}
	for int64((page-1)*pageSize) < total {
		if err = fg.Limit(pageSize).Offset((page - 1) * pageSize).Find(&taskList).Error; err != nil {
			err = errors.New(fmt.Sprintf("定时任务初始化#获取任务列表错误: %s", err.Error()))
			return
		}
		if len(taskList) == 0 {
			break
		}
		for _, item := range taskList {
			if err = t.Add(g, item); err != nil {
				return
			}
			t.nums += 1
		}
		page++
	}
	logrus.Infof("定时任务初始化完成, 共%d个定时任务添加到调度器", t.nums)
	return
}
func (t *Task) getOnceSpec(i int64) string {
	return time.Unix(i, 0).Format(OnceFormat)
}

func (t *Task) getMapKey(id int64, tType taskType) string {
	var prefix string
	switch tType {
	case runType:
		// 开始
		prefix = "run"
	case ExpiryType:
		// 过期
		prefix = "expiry"
	default:
		// 生效
		prefix = "effect"
	}
	return fmt.Sprintf("%s_%v", prefix, id)
}

// 添加`生效`定时任务
func (t *Task) pushEffectCron(p model.Task) (err error) {
	var jobId cron.EntryID
	job := EffectJob{
		Id:       p.Id,
		ExecTime: p.EffectTime,
	}
	if jobId, err = t.AddJob(t.getOnceSpec(p.EffectTime), &job); err != nil {
		return
	}
	t.m.Store(t.getMapKey(p.Id, EffectType), jobId)
	return
}

// 添加`开始`定时任务
func (t *Task) pushRunCron(p model.Task, tx *gorm.DB) (err error) {
	var jobId cron.EntryID
	job := RunJob{
		Id:     p.Id,
		Name:   p.Name,
		HostId: p.HostId,
		Cmd:    p.Cmd,
	}
	if jobId, err = t.Cron.AddJob(p.Spec, &job); err != nil {
		return
	}
	job.JobId = jobId
	tx.Model(model.Task{}).Where("id=?", p.Id).Update("next_exec_time", T.Cron.Entry(jobId).Next.Unix())

	t.m.Store(t.getMapKey(p.Id, runType), jobId)
	return
}

// 添加`过期`定时任务
func (t *Task) pushExpiryCron(p model.Task) (err error) {
	var jobId cron.EntryID
	job := ExpiryJob{
		Id:       p.Id,
		ExecTime: p.ExpiryTime,
	}
	if jobId, err = t.AddJob(t.getOnceSpec(p.ExpiryTime), &job); err != nil {
		return
	}
	t.m.Store(t.getMapKey(p.Id, ExpiryType), jobId)
	return
}

func (t *Task) addIdleTask(tx *gorm.DB, m model.Task, now int64) (err error) {
	t.Begin()
	if now < m.EffectTime {
		// 未到生效时间,添加生效、失效任务
		if err = t.pushEffectCron(m); err != nil {
			t.RollBack()
			return
		}
		if err = t.pushExpiryCron(m); err != nil {
			t.RollBack()
			return
		}

	} else if now >= m.EffectTime && now < m.ExpiryTime {
		// 到生效时间，但未生效,直接触发生效并添加失效任务，添加开始结束任务

		if err = t.pushExpiryCron(m); err != nil {
			t.RollBack()
			return
		}

		if err = t.pushRunCron(m, tx); err != nil {
			t.RollBack()
			return
		}

		if err = tx.Model(&m).Where("id=?", m.Id).Update("status", model.StatusRunning).Error; err != nil {
			T.RollBack()
			return
		}
	} else {
		if err = tx.Model(&m).Where("id=?", m.Id).Update("status", model.StatusExpiry).Error; err != nil {
			return
		}
		// 任务已过期，任务直接置为过期
	}
	t.Commit()
	return
}
func (t *Task) addIEffectTask(tx *gorm.DB, m model.Task, now int64) (err error) {
	t.Begin()
	if m.ExpiryTime == 0 {
		if err = t.pushRunCron(m, tx); err != nil {
			t.RollBack()
			return
		}
	} else if now < m.ExpiryTime {
		// 未到失效时间,添加失效任务，添加开始结束任务
		if err = t.pushExpiryCron(m); err != nil {
			t.RollBack()
			return
		}
		if err = t.pushRunCron(m, tx); err != nil {
			t.RollBack()
			return
		}
	} else {
		if err = tx.Model(&m).Where("id=?", m.Id).Update("status", model.StatusExpiry).Error; err != nil {
			return
		}
		// 任务已过期，任务直接置为过期
	}
	t.Commit()
	return

}

func (t *Task) Add(db *gorm.DB, p model.Task) (err error) {
	now := time.Now().Unix()
	switch p.Status {
	case model.StatusIdle:
		// 待生效
		return t.addIdleTask(db, p, now)
	default:
		// 已生效
		return t.addIEffectTask(db, p, now)
	}
}

func (t *Task) Stop(id int64, tType taskType) (err error) {
	if v, ok1 := t.m.Load(t.getMapKey(id, tType)); ok1 {
		if jobId, ok2 := v.(cron.EntryID); ok2 {
			t.Cron.Remove(jobId)
			logrus.Infof("任务:%v 停止成功", id)
			return
		}
	}
	logrus.Errorf("停止任务id:%v,taskType:%v", id, tType)
	return common.IdErr
}
func (t *Task) StopAll(id int64) {
	t.StopExpire(id)
	t.StopEffect(id)
	t.StopRun(id)
	return
}

func (t *Task) clean() {
	t.EntryIDs = []cron.EntryID{}
}
func (t *Task) Begin() {
	t.clean()
}
func (t *Task) AddJob(spec string, cmd cron.Job) (jobId cron.EntryID, err error) {
	if jobId, err = t.Cron.AddJob(spec, cmd); err != nil {
		return
	}
	t.EntryIDs = append(t.EntryIDs, jobId)
	return
}
func (t *Task) RollBack() {
	for _, jobId := range t.EntryIDs {
		t.Cron.Remove(jobId)
	}
	return
}

func (t *Task) Commit() {
	t.clean()
}
func (t *Task) StopEffect(id int64) {
	_ = t.Stop(id, EffectType)
}
func (t *Task) StopRun(id int64) {
	_ = t.Stop(id, runType)
}
func (t *Task) StopExpire(id int64) {
	_ = t.Stop(id, ExpiryType)
}

func (t *Task) Run(p model.Task) (err error) {
	if v, ok1 := t.m.Load(t.getMapKey(p.Id, runType)); ok1 {
		if jobId, ok2 := v.(cron.EntryID); ok2 {
			entry := t.Cron.Entry(jobId)
			entry.Job.Run()
			return
		}
	}
	return common.TaskRunErr

}
