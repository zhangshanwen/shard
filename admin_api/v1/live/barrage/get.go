package barrage

import (
	"github.com/sirupsen/logrus"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/initialize/service"
	"github.com/zhangshanwen/shard/inter/param"
)

func Get(c *service.AdminTxContext) (r service.Res) {
	//TODO 建立websoket连接
	p := param.GetBarrage{}
	if r.Err = c.Rebind(&p); r.Err != nil {
		r.ParamsError()
		return
	}
	var (
		err    error
		result []string
	)
	for {
		result, err = db.R.BLPop(c, 0, p.Hash).Result()
		if err != nil {
			break
		}
		logrus.Infof("接收到的弹幕是.....%v \n", result)
		//TODO 发送弹幕
		
	}

	return
}
