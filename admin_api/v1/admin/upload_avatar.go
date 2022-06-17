package admin

import (
	"fmt"

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
	if r.Err = tx.Model(&c.Admin).Updates(&model.Admin{
		Avatar: p.Avatar,
	}).Error; r.Err != nil {
		r.DBError()
		return
	}

	c.SaveLog(tx, fmt.Sprintf("上传头像"), model.OperateLogTypeUpdate)
	return
}
