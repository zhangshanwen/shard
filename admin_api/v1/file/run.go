package file

import (
	"fmt"
	"os/exec"
	"path"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
)

func Run(c *service.AdminTxContext) (r service.Res) {
	p := param.FileRunParams{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		f   model.FileRecord
		tx  = c.Tx
		out []byte
	)
	defer func() {
		if r.Err == nil {
			r.Data = string(out)
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
	c.SaveLogAdd(tx, module, fmt.Sprintf("run id:%v name:%v,file_type:%v ,result;%v ", f.Id, f.Name, f.FileType, string(out)))
	return
}
