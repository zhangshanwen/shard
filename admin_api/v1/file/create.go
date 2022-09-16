package file

import (
	"crypto/md5"
	"fmt"

	"gorm.io/gorm"

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

func Upload(c *service.AdminTxContext) (r service.Res) {
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
		fileRecord model.FileRecord
	)
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
	fileRecord = model.FileRecord{Uid: c.Admin.Id, Name: p.FileName, FileType: p.FileType, FileId: file.Id}
	fileRecord.FileId = file.Id
	if r.Err = tx.Save(&fileRecord).Error; r.Err != nil {
		r.DBError()
		return
	}
	c.SaveLogAdd(tx, fmt.Sprintf("上传文件 id:%v name:%v file_type%v", fileRecord.Id, fileRecord.Name, fileRecord.FileType))
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
