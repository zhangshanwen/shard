package file

import (
	"os/exec"
	"path"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

/*
 */

func Run(c *service.AdminContext) (resp service.Res) {
	p := param.FileRunParams{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	var f model.FileRecord
	if resp.Err = db.G.Preload("File").First(&f, p.Id).Error; resp.Err != nil {
		return
	}
	filePath := path.Join(f.File.Path, f.File.Hash)

	cmd := exec.Command(f.GetCmd(), filePath)
	var out []byte
	if out, resp.Err = cmd.CombinedOutput(); resp.Err != nil {
		return
	}
	resp.Data = string(out)
	return
}
