package wechat

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

type (
	Bot struct {
		*openwechat.Bot
		replies  []*Reply
		Self     *openwechat.Self
		Messages chan string
	}
	messageType string
)

const (
	messagePingType         messageType = "ping"
	messageLoginType        messageType = "login"
	messageMessageType      messageType = "message"
	messageMessageReplyType messageType = "messageReply"
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

	b.saveMessage(msg.Content, sender.ID(), b.Self.ID(), msg.CreateTime)
	b.SendMessage(messageMessageType, sender.ID(), common.MessageSplitSymbol, msg.Content)
	for k, v := range reply.Rules {
		if matched, err = regexp.Match(k, []byte(msg.Content)); err != nil && matched && b.checkReply(sender, reply) {
			if _, err = msg.ReplyText(v); err != nil {
				return
			}
			b.saveMessage(msg.Content, b.Self.ID(), sender.ID(), msg.CreateTime)
			b.SendMessage(messageMessageReplyType, sender.ID(), common.MessageSplitSymbol, msg.Content)

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
		return reply.IsAllFriends || b.checkFriendOrGroups(reply.Friends, reply.ExcludeFriends, sender.UserName)
	} else if sender.IsGroup() {
		return reply.IsAllGroups || b.checkFriendOrGroups(reply.Groups, reply.ExcludeGroups, sender.UserName)
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

func (b *Bot) FindFriend(friendId string) (friend *openwechat.Friend, err error) {

	var friends openwechat.Friends
	if friends, err = b.Friends(); err != nil {
		return
	}
	for _, f := range friends {
		if f.ID() == friendId {
			return f, nil
		}
	}
	return nil, errors.New("NotFoundFriendId")
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

// 保存历史信息
func (b *Bot) saveMessage(message, senderId, receiverId string, createdTime int64) {
	go func() {
		cm := model.ChatMessage{
			Msg:        message,
			SenderId:   senderId,
			ReceiverId: receiverId,
		}
		if createdTime != 0 {
			cm.CreatedTime = createdTime
		}
		db.G.Save(&cm)
	}()
}

// SendMessage 向客户端发送消息数据
func (b *Bot) SendMessage(t messageType, message ...string) {
	message = append([]string{string(t)}, message...)
	go func() {
		b.Messages <- strings.Join(message, common.MessageSplitSymbol)
	}()
}

func (b *Bot) ReceiveMessage(message []byte) {
	if len(message) <= 0 {
		return
	}
	var messages = strings.Split(string(message), common.MessageSplitSymbol)
	if len(messages) <= 1 {
		// 符合规范消息ping[::]ping
		return
	}
	switch messageType(messages[0]) {
	case messagePingType:
		b.SendMessage(messagePingType, "pong")
	default:
		b.Chat(messages[1:])
	}
}

func (b *Bot) CleanMessages() {
	b.Messages = make(chan string)
}
func (b *Bot) Chat(body []string) {
	if len(body) <= 1 {
		//符合聊天消息   message[::]friend_id[::]消息时间戳[::]hello world
		return
	}
	var (
		friendId    = body[0]
		createdTime int64
		msg         = strings.Join(body[2:], "")
		err         error
		friend      *openwechat.Friend
	)
	if friend, err = b.FindFriend(friendId); err != nil {
		logrus.Warning(err)
		return
	}
	if _, err = friend.SendText(msg); err != nil {
		logrus.Warning(err)
	}
	if createdTime, err = strconv.ParseInt(body[1], 10, 64); err != nil {
		return
	}
	b.saveMessage(msg, friendId, b.Self.ID(), createdTime)
}
