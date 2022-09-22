package admin

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func UploadAvatar(c *service.AdminTxContext) (r service.Res) {
	p := param.AdminUploadAvatar{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.DBError()
		return
	}
	var (
		tx = c.Tx
	)
	defer func() {
		if r.Err == nil {
			r.Data = c.Admin
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
	c.SaveLogUpdate(tx, module, "avatar")
	return
}
