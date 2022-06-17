package file

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Run(c *service.AdminContext) (r service.Res) {
	p := param.FileRunParams{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		f   model.FileRecord
		tx  = db.G.Begin()
		out []byte
	)
	defer func() {
		r.Data = string(out)
		if r.Err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if r.Err = tx.Preload("File").First(&f, p.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}
	filePath := path.Join(f.File.Path, f.File.Hash)

	cmd := exec.Command(f.GetCmd(), filePath)
	if out, r.Err = cmd.CombinedOutput(); r.Err != nil {
		r.SystemError()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("执行文件 id:%v name:%v file_type%v ,结果 %v ", f.Id, f.Name, f.FileType, string(out)), model.OperateLogTypeSelect)
	return
}
