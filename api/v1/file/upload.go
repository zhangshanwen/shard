package file

import (
	"crypto/md5"
	"fmt"
	"github.com/zhangshanwen/shard/code"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/internal/param"
	"github.com/zhangshanwen/shard/model"
	"gorm.io/gorm"
)

func Upload(c *service.Context) (resp service.Res) {
	p := param.UploadParams{}
	if resp.Err = c.Rebind(&p); resp.Err != nil {
		resp.ResCode = code.ParamsError
		return
	}
	var hash string
	hash, resp.Err = getHash([]byte(p.File))
	if resp.Err != nil {
		return
	}
	resp.Data = hash
	// TODO 查询是否hash碰撞,录入数据库，文件存入存储目录

	return
}

func getHash(b []byte) (hash string, err error) {
	Md5Inst := md5.New()
	Md5Inst.Write(b)
	hashByte := Md5Inst.Sum([]byte(""))
	hash = fmt.Sprintf("%x", hashByte)
	var file model.File
	err = db.G.Where(" hash=? ", hash).First(&file).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return
	}
	// TODO 检测文件是否相同,如果相同则不写入该文件，只录入数据,如果不相同则发生hash碰撞,重新计算hash值
	return getHash(hashByte)

}
