package file

import (
	"errors"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
)

/*
1.检测文件是否上传过，如果无则创建，有这查询出文件
2.检测文件是否上传过该文件名，如果没有创建该记录，如果有则覆盖该记录
*/

func Update(c *service.AdminContext) (resp service.Res) {
	pId := param.UriId{}
	if resp.Err = c.BindUri(&pId); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	p := param.FileUploadParams{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	var postFix string
	// python 文件
	if p.FileType == 1 {
		postFix = ".py"
	}

	var file model.File
	// 开启事务
	g := db.G.Begin()
	defer func() {
		if resp.Err != nil {
			g.Rollback()
			return
		}
		g.Commit()
	}()
	// hash文件内容
	file.Path = conf.C.File.Path
	if resp.Err = GetHash(g, &file, []byte(p.File), postFix, p.File); resp.Err != nil {
		return
	}

	// 查询用户文件记录
	fileRecord := model.FileRecord{}
	if resp.Err = g.First(&fileRecord, pId.Id).Error; resp.Err != nil {
		return
	}
	if fileRecord.Uid != c.Admin.Id {
		resp.Err = errors.New("not owner")
		return
	}
	fileRecord.FileType = p.FileType
	fileRecord.Name = p.FileName
	fileRecord.FileId = file.Id
	if resp.Err = g.Save(&fileRecord).Error; resp.Err != nil {
		return
	}
	resp.Data = fileRecord
	return
}
