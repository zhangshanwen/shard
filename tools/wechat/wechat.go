package wechat

import (
	"fmt"
	"sync"

	"github.com/eatmoreapple/openwechat"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/common"
)

/*
实现个自动回复，或者机器人之类的功能，可自动识别关键字,进行返回
*/

type (
	Wechat struct {
		bots map[int64]*Bot
	}
)

var (
	wechat *Wechat
	once   = sync.Once{}
)

func NewWechat() *Wechat {
	once.Do(func() {
		wechat = &Wechat{
			bots: make(map[int64]*Bot),
		}
	})
	return wechat
}

func (w *Wechat) Bot(uid int64) (weBot *Bot) {
	var (
		ok  bool
		bot *openwechat.Bot
	)
	if _, ok = w.bots[uid]; !ok {
		bot = openwechat.DefaultBot(openwechat.Desktop)
		w.bots[uid] = &Bot{
			bot,
			[]*Reply{},
			make(chan string),
		}
	}

	return w.bots[uid]
}

func (w *Wechat) Qrcode(uid int64, replies []*Reply) (code string, err error) {
	bot := w.Bot(uid)
	var getCallback = make(chan bool)
	bot.UUIDCallback = func(uuid string) {
		logrus.Info("进入UUIDCallback")
		defer func() { getCallback <- true }()
		code = openwechat.GetQrcodeUrl(uuid)
		bot.SendMessage(messageLoginType, "qrcode")
	}
	go func() {
		reloadStorage := openwechat.NewFileHotReloadStorage(fmt.Sprintf("%v_%v.json", common.WechatStorageFileName, uid))
		defer reloadStorage.Close()
		if err = bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
			logrus.Errorf("登录失败....%v", err)
			return
		}
		if err = bot.AddReply(replies); err != nil {
			logrus.Warning("添加规则失败", err)
		}
		bot.SendMessage(messageLoginType, "success")
		logrus.Info("登陆完成.......")
	}()
	<-getCallback
	return
}

func (w *Wechat) Login(uid int64) (err error) {
	bot := w.Bot(uid)
	if bot.Alive() {
		// 无需重新登陆
		return
	}
	reloadStorage := openwechat.NewFileHotReloadStorage(common.WechatStorageFileName)
	defer reloadStorage.Close()
	if err = bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		return
	}
	return
}

func (w *Wechat) GetCurrentUser(uid int64) (self *openwechat.Self, err error) {
	bot := w.Bot(uid)
	return bot.GetCurrentUser()
}

func (w *Wechat) Friends(uid int64) (friends openwechat.Friends, err error) {
	var (
		self *openwechat.Self
	)
	if self, err = w.GetCurrentUser(uid); err != nil {
		return
	}
	return self.Friends()
}

func (w *Wechat) Groups(uid int64) (groups openwechat.Groups, err error) {
	var (
		self *openwechat.Self
	)
	if self, err = w.GetCurrentUser(uid); err != nil {
		return
	}
	return self.Groups()
}
