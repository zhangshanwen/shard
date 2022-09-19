package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/robfig/cron/v3"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

type (
	RunJob struct {
		Id     int64
		HostId int64
		Name   string
		Cmd    string
		JobId  cron.EntryID
	}
)

func (t *RunJob) Run() {
	var (
		err error
		msg string
	)
	defer func() {
		tl := model.TaskLog{
			TaskId: t.Id,
		}
		if err == nil {
			tl.Comment = msg
			tl.Status = model.CommonStatusSuccess
			db.G.Model(model.Task{}).Where("id=?", t.Id).Update("next_exec_time", T.Cron.Entry(t.JobId).Next.Unix())
		} else {
			T.StopAll(t.Id)
			db.G.Model(model.Task{}).Where("id=?", t.Id).Update("status", model.StatusFailed)
			db.G.Model(model.Host{}).Where("id=?", t.HostId).Update("status", model.CommonStatusSuccess)
			tl.Comment = err.Error()
			tl.Status = model.CommonStatusFailed
		}
		db.G.Save(&tl)
	}()

	var (
		host model.Host
		b    []byte
	)

	if err = db.G.First(&host, t.HostId).Error; err != nil {
		return
	}
	if host.ConnectType == model.ConnectTypeSSh {
		var (
			ss *tools.SshSocket
		)
		if ss, err = tools.NewSshSocket(host.Username, host.Password, host.Host, host.Port); err != nil {
			return
		}
		if err = ss.Session(); err != nil {
			return
		}
		b, err = ss.CombinedOutput(t.Cmd)
	} else {
		var (
			resp *http.Response
			u    *url.URL
		)
		uri := host.Host
		if host.ConnectType == model.ConnectTypeHttp {
			if !strings.Contains(uri, common.HttpPrefix) {
				uri = common.HttpPrefix + uri
			}
		} else {
			if !strings.Contains(uri, common.HttpsPrefix) {
				uri = common.HttpsPrefix + uri
			}
		}
		uri = fmt.Sprintf("%v:%v", uri, host.Port)

		if u, err = url.Parse(uri); err != nil {
			return
		}
		u.Path = path.Join(u.Path, t.Cmd)
		if resp, err = http.Get(u.String()); err != nil {
			return
		}
		defer resp.Body.Close()
		if b, err = ioutil.ReadAll(resp.Body); err != nil {
			return
		}
		if resp.StatusCode>>2 != http.StatusOK>>2 {
			err = common.RequestErr
			return
		}
	}
	msg = string(b)
}
