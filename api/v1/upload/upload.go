package upload

import (
	"crypto/md5"
	"fmt"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
)

func Upload(c *service.Context) (resp service.Res) {
	p := param.UploadParams{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(p.File))
	p.File = fmt.Sprintf("%x", Md5Inst.Sum([]byte("")))
	resp.Data = p
	// TODO 查询是否hash碰撞,录入数据库，文件存入存储目录

	return
}
