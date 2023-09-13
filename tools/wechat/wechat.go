package wechat

import (
	"fmt"
	"sync"

	"github.com/eatmoreapple/openwechat"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/common"
)

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
			[]*TimerReply{},
			nil,
			make(chan string),
		}
	}

	return w.bots[uid]
}

func (w *Wechat) Qrcode(uid int64, replies []*Reply, timerRelies []*TimerReply) (code string, err error) {
	bot := w.Bot(uid)
	var getCallback = make(chan bool)
	bot.UUIDCallback = func(uuid string) {
		logrus.Info("进入UUIDCallback")
		defer func() { getCallback <- true }()
		code = openwechat.GetQrcodeUrl(uuid)
		bot.SendMessage(messageLoginType, "qrcode")
	}
	go func() {
		reloadStorage := openwechat.NewFileHotReloadStorage(fmt.Sprintf("%v/%v.json", common.WechatStorageFileName, uid))
		defer reloadStorage.Close()
		if err = bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
			logrus.Errorf("登录失败....%v", err)
			return
		}
		if err = bot.AddReply(replies); err != nil {
			logrus.Warning("添加自动回复规则失败", err)
		}
		if err = bot.AddTimerReply(timerRelies); err != nil {
			logrus.Warning("添加定时发送消息规则失败", err)
		}
		bot.SendMessage(messageLoginType, "success")
		logrus.Info("登陆完成.......")
	}()
	<-getCallback
	return
}

func (w *Wechat) Login(uid int64) (err error) {
	bot := w.Bot(uid)
	defer func() {
		if bot.Self, err = bot.GetCurrentUser(); err != nil {
			return
		}
	}()
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
	bot := w.Bot(uid)
	return bot.Self.Friends()
}

func (w *Wechat) Groups(uid int64) (groups openwechat.Groups, err error) {
	bot := w.Bot(uid)
	return bot.Self.Groups()
}
