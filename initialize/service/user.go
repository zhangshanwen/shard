package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

type Context struct {
	*gin.Context
	User model.User
}

func (c *Context) Rebind(obj interface{}) (err error) {
	if err = c.Bind(obj); err != nil {
		return
	}
	logrus.WithField("mod", "params").Info(obj)
	return
}

func (c *Context) getLoginInfoKey(uid int64) string {
	return fmt.Sprintf(common.UserLoginKey, uid)
}
func (c *Context) SaveLoginInfo() (err error) {
	var b []byte
	if b, err = json.Marshal(&c.User); err != nil {
		return
	}
	return db.R.SetNX(c.Context, c.getLoginInfoKey(c.User.Id), b, tools.GetExpireTime()).Err()
}
func (c *Context) GetLoginInfo(uid int64) (err error) {
	var b []byte
	if b, err = db.R.Get(c.Context, c.getLoginInfoKey(uid)).Bytes(); err != nil {
		return
	}
	return json.Unmarshal(b, &c.User)

}
