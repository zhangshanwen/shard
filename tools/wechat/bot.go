package wechat

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"regexp"

	"github.com/eatmoreapple/openwechat"

	"github.com/zhangshanwen/shard/tools"
)

type (
	Bot struct {
		*openwechat.Bot
		replies  []*Reply
		Messages chan string
	}
	messageType string
)

const (
	messagePingType    messageType = "ping"
	messageLoginType   messageType = "login"
	messageMessageType messageType = "message"
)

func (b *Bot) replyMessage(msg *openwechat.Message, reply *Reply) (err error) {
	if !msg.IsText() {
		return
	}
	var (
		matched bool
		sender  *openwechat.User
	)
	if sender, err = msg.Sender(); err != nil {
		return
	}
	for k, v := range reply.rules {
		if matched, err = regexp.Match(k, []byte(msg.Content)); err != nil && matched && b.checkReply(sender, reply) {
			if _, err = msg.ReplyText(v); err != nil {
				return
			}
		}
	}
	return
}

func (b *Bot) AddReply(replies []*Reply) (err error) {

	b.replies = replies
	b.Bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() {
			for _, reply := range b.replies {
				if err = b.replyMessage(msg, reply); err != nil {
					return
				}
			}
		}
	}
	return
}
func (b *Bot) checkReply(sender *openwechat.User, reply *Reply) bool {
	if sender.IsFriend() {
		return reply.isAllFriends || b.checkFriendOrGroups(reply.friends, reply.excludeFriends, sender.UserName)
	} else if sender.IsGroup() {
		return reply.isAllGroups || b.checkFriendOrGroups(reply.groups, reply.excludeGroups, sender.UserName)
	}
	return false
}

func (b *Bot) checkFriendOrGroups(includeArray, excludeArray []string, compare string) bool {
	exclude := tools.FindInArray[string, string](excludeArray, compare, func(item, compare string) bool {
		return item == compare
	})
	include := tools.FindInArray[string, string](includeArray, compare, func(item, compare string) bool {
		return item == compare
	})
	return include && !exclude
}

func (b *Bot) Friends() (friends openwechat.Friends, err error) {
	var (
		self *openwechat.Self
	)
	if self, err = b.GetCurrentUser(); err != nil {
		return
	}
	return self.Friends()
}
func (b *Bot) LoginCallBack(body openwechat.CheckLoginResponse) {
	var (
		loginCode openwechat.LoginCode
		err       error
	)
	if loginCode, err = body.Code(); err != nil {
		logrus.Errorf("登陆失败:%v", err)
		return
	}
	switch loginCode {
	case openwechat.LoginCodeSuccess:
		b.SendMessage(messageLoginType, "success")
	case openwechat.LoginCodeScanned:
		b.SendMessage(messageLoginType, "scanned")
	case openwechat.LoginCodeTimeout:
		b.SendMessage(messageLoginType, "timeout")

	default:
		return
	}
}

func (b *Bot) MessageHandler(msg *openwechat.Message) {
	if msg.IsText() && msg.StatusNotifyCode == 0 {
		if f, ok := templateReply[msg.Content]; ok {
			_, _ = msg.ReplyText(f(msg.Sender()))
		}
	}
}
func (b *Bot) SendMessage(t messageType, message string) {
	fmt.Println("------------------")
	go func() {
		b.Messages <- fmt.Sprintf("%v:%v", t, message)
	}()
}

func (b *Bot) ReceiveMessage(message []byte) {
	switch string(message) {
	case "ping":
		b.SendMessage(messagePingType, "pong")
	}
}

func (b *Bot) CleanMessages() {
	b.Messages = make(chan string)
}
