package file

import (
	"fmt"

	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

/*
1.检测文件是否上传过，如果无则创建，有这查询出文件
2.检测文件是否上传过该文件名，如果没有创建该记录，如果有则覆盖该记录
*/

func Update(c *service.AdminTxContext) (r service.Res) {
	pId := param.UriId{}
	if r.Err = c.BindUri(&pId); r.Err != nil {
		r.ParamsError()
		return
	}
	p := param.FileUploadParams{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var postFix string
	// python 文件
	if p.FileType == 1 {
		postFix = ".py"
	}

	var (
		file       model.File
		tx         = c.Tx
		fileRecord = model.FileRecord{}
	)
	// 开启事务

	defer func() {
		if r.Err == nil {
			r.Data = fileRecord
		}
	}()
	// hash文件内容
	file.Path = conf.C.File.Path
	if r.Err = GetHash(tx, &file, []byte(p.File), postFix, p.File); r.Err != nil {
		r.UploadFileFailed()
		return
	}

	// 查询用户文件记录
	if r.Err = tx.First(&fileRecord, pId.Id).Error; r.Err != nil {
		r.NotFound()
		return
	}
	if fileRecord.Uid != c.Admin.Id {
		r.NotOwner()
		return
	}
	c.SaveLog(tx, fmt.Sprintf("修改上传文件 id:%v %v ", fileRecord.Id, tools.DiffStruct(p, fileRecord, "json")), model.OperateLogTypeUpdate)
	fileRecord.FileType = p.FileType
	fileRecord.Name = p.FileName
	fileRecord.FileId = file.Id
	if r.Err = tx.Save(&fileRecord).Error; r.Err != nil {
		r.DBError()
		return
	}
	return
}
