package file

import (
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/internal/response"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func Detail(c *service.AdminContext) (resp service.Res) {
	p := param.UriId{}
	if resp.Err = c.BindUri(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	var fileRecord = model.FileRecord{}
	var r = response.FileDetail{}
	if resp.Err = db.G.Preload("File").First(&fileRecord, p.Id).Error; resp.Err != nil {
		return
	}
	r.Id = fileRecord.Id
	r.FileType = fileRecord.FileType
	r.Name = fileRecord.Name
	if fileRecord.File != nil {
		r.Code, _ = tools.FileToBase64(fileRecord.File.Hash, fileRecord.File.Path)
	}
	resp.Data = r
	return
}
