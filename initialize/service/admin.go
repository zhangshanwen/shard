package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
	"github.com/zhangshanwen/shard/tools/wechat"
)

type AdminContext struct {
	*gin.Context
	Admin model.Admin
}
type AdminWechatContext struct {
	AdminContext
	Bot *wechat.Bot
}

type AdminTxContext struct {
	AdminContext
	Tx *gorm.DB
}

func (c *AdminContext) Rebind(obj interface{}) (err error) {
	if err = c.Bind(obj); err != nil {
		return
	}
	logrus.WithField("mod", "params").Info(obj)
	return
}

func (c *AdminContext) SaveLogSelect(tx *gorm.DB, module, log string) {
	c.saveLog(tx, module, log, model.OperateLogTypeSelect)
}
func (c *AdminContext) SaveLogAdd(tx *gorm.DB, module, log string) {
	c.saveLog(tx, module, log, model.OperateLogTypeAdd)
}
func (c *AdminContext) SaveLogDel(tx *gorm.DB, module, log string) {
	c.saveLog(tx, module, log, model.OperateLogTypeDel)
}
func (c *AdminContext) SaveLogUpdate(tx *gorm.DB, module, log string) {
	c.saveLog(tx, module, log, model.OperateLogTypeUpdate)
}

func (c *AdminContext) saveLog(tx *gorm.DB, module, log string, operateLogType model.OperateLogType) {
	var err error
	ol := model.OperateLog{
		OperateId: c.Admin.Id,
		Module:    module,
		Log:       log,
		Type:      operateLogType,
	}
	if err = tx.Save(&ol).Error; err != nil {
		logrus.Errorf("记录操作日志失败,id=%v username=%v err=%v", c.Admin.Id, c.Admin.Username, err.Error())
	}
}

func (c *AdminContext) getLoginInfoKey(uid int64) string {
	return fmt.Sprintf(common.AdminLoginKey, uid)
}
func (c *AdminContext) SaveLoginInfo() (err error) {
	var b []byte
	if b, err = json.Marshal(&c.Admin); err != nil {
		return
	}
	return db.R.SetNX(c.Context, c.getLoginInfoKey(c.Admin.Id), b, tools.GetExpireTime()).Err()
}
func (c *AdminContext) GetLoginInfo(uid int64) (err error) {
	var b []byte
	if b, err = db.R.Get(c.Context, c.getLoginInfoKey(uid)).Bytes(); err != nil {
		return
	}
	return json.Unmarshal(b, &c.Admin)

}
