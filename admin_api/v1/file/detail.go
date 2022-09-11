package file

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/inter/response"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

func Detail(c *service.AdminTxContext) (r service.Res) {
	p := param.UriId{}
	if r.Err = c.BindUri(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		fileRecord = model.FileRecord{}
		resp       = response.FileDetail{}
		tx         = c.Tx
	)
	defer func() {
		if r.Err == nil {
			r.Data = resp
		}
	}()
	if r.Err = tx.Preload("File").First(&fileRecord, p.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}
	resp.Id = fileRecord.Id
	resp.FileType = fileRecord.FileType
	resp.Name = fileRecord.Name
	if fileRecord.File != nil {
		resp.Code, _ = tools.FileToBase64(fileRecord.File.Hash, fileRecord.File.Path)
	}
	c.SaveLog(tx, fmt.Sprintf("查看文件详情 id:%v name:%v file_type%v", resp.Id, resp.Name, resp.FileType), model.OperateLogTypeSelect)
	return
}
