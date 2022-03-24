package file

import (
	"crypto/md5"
	"fmt"

	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/conf"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

/*
1.检测文件是否上传过，如果无则创建，有这查询出文件
2.检测文件是否上传过该文件名，如果没有创建该记录，如果有则覆盖该记录
*/

func Upload(c *service.AdminContext) (resp service.Res) {
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
	fileRecord := model.FileRecord{Uid: c.Admin.Id, Name: p.FileName, FileType: p.FileType, FileId: file.Id}
	fileRecord.FileId = file.Id
	if resp.Err = g.Save(&fileRecord).Error; resp.Err != nil {
		return
	}
	resp.Data = fileRecord
	return
}

func GetHash(g *gorm.DB, f *model.File, b []byte, postfix, fileBody string) (err error) {
	Md5Inst := md5.New()
	Md5Inst.Write(b)
	hashByte := Md5Inst.Sum([]byte(""))
	f.Hash = fmt.Sprintf("%x%s", hashByte, postfix)
	err = g.Where(" hash=? ", f.Hash).First(&f).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 存储文件
			if err = tools.SaveFile(f.Hash, fileBody, ""); err != nil {
				return
			}
			// 文件信息录入数据库
			if err = g.Create(&f).Error; err != nil {
				return
			}
			return
		}
		return
	}
	// 读取文件
	var existFile string
	existFile, err = tools.FileToBase64(f.Hash, "")
	// 文件存在则不存储
	if fileBody == existFile {
		return
	}
	// 发生散列冲突，继续散列
	return GetHash(g, f, hashByte, postfix, fileBody)

}
