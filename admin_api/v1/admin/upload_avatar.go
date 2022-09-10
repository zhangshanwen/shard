package admin

import (
	"encoding/base64"
	"fmt"
	"github.com/zhangshanwen/shard/tools"
	"strings"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func UploadAvatar(c *service.AdminContext) (r service.Res) {
	p := param.AdminUploadAvatar{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.DBError()
		return
	}
	var (
		tx = db.G.Begin()
	)
	defer func() {
		r.Data = c.Admin
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	var b []byte
	ss := strings.Split(p.Avatar, ",")
	if len(ss) <= 1 {
		r.ParamsError()
		return
	}
	if b, r.Err = base64.StdEncoding.DecodeString(ss[1]); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		oss tools.Oss
		key string
	)
	oss, r.Err = tools.NewOss()
	if key, r.Err = oss.UploadFile(c, fmt.Sprintf("%s.jpeg", tools.Hash(string(b))), b); r.Err != nil {
		r.SystemError()
		return
	}
	c.Admin.Avatar = key
	if r.Err = tx.Model(&c.Admin).Updates(&model.Admin{
		Avatar: key,
	}).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("上传头像"), model.OperateLogTypeUpdate)
	return
}
