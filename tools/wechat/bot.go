package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/jinzhu/copier"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

	"github.com/zhangshanwen/shard/common"
	"github.com/zhangshanwen/shard/initialize/db"
	"github.com/zhangshanwen/shard/model"
	"github.com/zhangshanwen/shard/tools"
)

type (
	Bot struct {
		*openwechat.Bot
		replies     []*Reply
		timeReplies []*TimerReply
		Self        *openwechat.Self
		Messages    chan string
	}
	messageType string
)

const (
	messagePingType              messageType = "ping"
	messageSyncFriendsType       messageType = "syncFriends"
	messageSyncGroupsType        messageType = "syncGroups"
	messageSyncMessagesTotalType messageType = "syncMessagesTotal"
	messageSyncMessagesType      messageType = "syncMessages"
	messageLoginType             messageType = "login"
	messageMessageType           messageType = "message"
	messageMessageReplyType      messageType = "messageReply"
	selfInfoType                 messageType = "selfInfo"
)

func (b *Bot) replyMessage(msg *openwechat.Message, sender *openwechat.User, reply *Reply) {
	var (
		matched bool
		err     error
	)
	_ = b.dealMessage(sender.ID(), b.Self.ID(), msg.Content, sender.IsGroup(), false, msg.CreateTime)
	for k, replyMsg := range reply.Rules {
		if matched, err = regexp.Match(k, []byte(msg.Content)); err != nil && matched && b.checkReply(sender, reply.FriendsGroups) {
			_ = b.dealMessage(b.Self.ID(), sender.ID(), replyMsg, sender.IsGroup(), false, 0)
		}
	}
}

// AddReply 添加自动回复规则
func (b *Bot) AddReply(replies []*Reply) (err error) {
	b.replies = replies
	b.Bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() {
			return
		}
		var (
			sender *openwechat.User
			err1   error
		)
		if sender, err1 = msg.Sender(); err1 != nil {
			return
		}
		for _, reply := range b.replies {
			b.replyMessage(msg, sender, reply)
		}
	}
	return
}

// AddTimerReply 添加定时发送消息
func (b *Bot) AddTimerReply(timerReplies []*TimerReply) (err error) {
	b.timeReplies = timerReplies
	c := tools.NewCron()
	var (
		entryID cron.EntryID
	)
	for _, i := range timerReplies {
		if i.EntryId > 0 {
			if tools.FindInArray[cron.Entry, int](c.Entries(), i.EntryId, func(item cron.Entry, compare int) bool {
				return int(item.ID) == compare
			}) {
				continue
			}
		}
		if entryID, err = c.AddFunc(i.Spec, func() {
			b.sendFriendsGroupsMessages(i.FriendsGroups, i.Msg)
		}); err != nil {
			logrus.Warningf("添加定时任务失败")
			continue
		}
		i.EntryId = int(entryID)
	}
	return
}

func (b *Bot) checkReply(sender *openwechat.User, fg FriendsGroups) bool {
	if sender.IsFriend() {
		return fg.IsAllFriends || b.checkFriendOrGroups(fg.Friends, fg.ExcludeFriends, sender.UserName)
	} else if sender.IsGroup() {
		return fg.IsAllGroups || b.checkFriendOrGroups(fg.Groups, fg.ExcludeGroups, sender.UserName)
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
func (b *Bot) Groups() (groups openwechat.Groups, err error) {
	var (
		self *openwechat.Self
	)
	if self, err = b.GetCurrentUser(); err != nil {
		return
	}
	return self.Groups()
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

func (b *Bot) FindGroup(groupId string) (group *openwechat.Group, err error) {

	var groups openwechat.Groups
	if groups, err = b.Groups(); err != nil {
		return
	}
	for _, f := range groups {
		if f.ID() == groupId {
			return f, nil
		}
	}
	return nil, errors.New("NotFoundGroupId")
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
func (b *Bot) saveMessage(message, ownerId, senderId, receiverId string, createdTime int64, isGroup bool, err error) {
	go func() {
		cm := model.ChatMessage{
			Msg:        message,
			SenderId:   senderId,
			OwnerId:    ownerId,
			IsGroup:    isGroup,
			IsSuccess:  err == nil,
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

func (b *Bot) sendJsonMessages(msgType messageType, msg interface{}) (err error) {
	var (
		body []byte
	)
	if body, err = json.Marshal(&msg); err != nil {
		return
	}
	b.SendMessage(msgType, string(body))
	return
}

func (b *Bot) ReceiveMessage(message []byte) (err error) {
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
	case messageSyncFriendsType:
		return b.syncFriends()
	case messageSyncGroupsType:
		return b.syncGroups()
	case messageSyncMessagesTotalType:
		return b.syncMessageTotal()
	case messageSyncMessagesType:
		return b.syncMessage()
	case selfInfoType:
		return b.syncInfo()
	default:
		b.Chat(messages[1:])
	}
	return
}
func (b *Bot) getSyncMessageTotalKey() string {
	return fmt.Sprintf("sync_message_total_%v", b.Self.ID())
}

func (b *Bot) CleanMessages() {
	b.Messages = make(chan string)
}
func (b *Bot) Chat(body []string) {
	if len(body) <= 1 {
		//符合聊天消息   message[::]friend_id[::]is_group[::]消息时间戳[::]hello world
		return
	}
	var (
		friendId    = body[0]
		issGroup    bool
		createdTime int64
		msg         = strings.Join(body[3:], "")
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
	if createdTime, err = strconv.ParseInt(body[2], 10, 64); err != nil {
		return
	}
	if strings.ToLower(body[1]) == "true" {
		issGroup = true
	}
	_ = b.dealMessage(b.Self.ID(), friendId, msg, issGroup, true, createdTime)
}

func (b *Bot) sendFriendsGroupsMessages(fg FriendsGroups, msg string) {
	var (
		friends openwechat.Friends
		groups  openwechat.Groups
		err     error
	)
	if friends, err = b.Self.Friends(); err == nil {
		for _, friend := range friends {
			if !b.checkReply(friend.User, fg) {
				continue
			}
			if err = b.dealMessage(b.Self.ID(), friend.ID(), msg, false, false, 0); err != nil {
				logrus.Warningf("好友消息发送失败:%s-%s", friend.ID(), msg)
			}
		}
	}
	if groups, err = b.Self.Groups(); err == nil {
		for _, group := range groups {
			if !b.checkReply(group.User, fg) {
				continue
			}
			if err = b.dealMessage(b.Self.ID(), group.ID(), msg, true, false, 0); err != nil {
				logrus.Warningf("群组消息发送失败:%s-%s", group.ID(), msg)
			}
		}
	}
}
func (b *Bot) sendFriendGroupMsg(findId, msg string, isGroup bool) (err error) {
	if isGroup {
		var (
			group *openwechat.Group
		)
		if group, err = b.FindGroup(findId); err != nil {
			return
		}
		_, err = group.SendText(msg)

	} else {
		var (
			friend *openwechat.Friend
		)
		if friend, err = b.FindFriend(findId); err != nil {
			return
		}
		_, err = friend.SendText(msg)
	}
	return err
}

func (b *Bot) dealMessage(sendId, receiveId, msg string, isGroup, isClient bool, createdTime int64) (err error) {
	defer b.saveMessage(msg, b.Self.ID(), sendId, receiveId, createdTime, isGroup, err)
	var msgReplyType = messageMessageType
	if b.Self.ID() == sendId {
		msgReplyType = messageMessageReplyType
		// 自身为发送者,需要给好友/群组发送消息
		err = b.sendFriendGroupMsg(receiveId, msg, isGroup)
	}
	if !isClient {
		// 如果不是客户端发送的消息，需要向客户端发送此消息
		b.SendMessage(msgReplyType, sendId, receiveId, msg)
	}
	return
}

func (b *Bot) syncFriends() (err error) {
	var friends openwechat.Friends
	if friends, err = b.Friends(); err != nil {
		return
	}
	return b.sendJsonMessages(messageSyncFriendsType, friends)
}
func (b *Bot) syncGroups() (err error) {
	var groups openwechat.Groups
	if groups, err = b.Groups(); err != nil {
		return
	}
	return b.sendJsonMessages(messageSyncGroupsType, groups)
}
func (b *Bot) syncMessageTotal() (err error) {
	// 同步消息时,预先记录下记录条数,防止同步消息时,同步异常
	var (
		count int64
	)
	if err = db.G.Model(model.ChatMessage{}).Where(model.ChatMessage{OwnerId: b.Self.ID()}).Count(&count).Error; err != nil {
		return
	}
	// 暂定时间为10分钟
	if err = db.R.SetEX(b.Context(), b.getSyncMessageTotalKey(), count, time.Minute*10).Err(); err != nil {
		return
	}
	b.SendMessage(messageSyncMessagesTotalType, strconv.FormatInt(count, 10))
	return
}

func (b *Bot) syncMessage() (err error) {
	var (
		count int64
		page  int
		cm    []model.ChatMessage
	)
	if count, err = db.R.Get(b.Context(), b.getSyncMessageTotalKey()).Int64(); err != nil {
		return
	}
	defer db.R.Del(b.Context(), b.getSyncMessageTotalKey())
	// 给客户端发送消息时,限制消息条数为100条每次
	for {
		if int64(page)*100 > count {
			break
		}
		if err = db.G.Where(model.ChatMessage{OwnerId: b.Self.ID()}).Offset(page * 100).Limit(100).Find(&cm).Error; err != nil {
			return
		}
		if err = b.sendJsonMessages(messageMessageType, cm); err != nil {
			return
		}
	}
	logrus.Infof("用户(%v:%v):消息%v条同步完成!", b.Self.UserName, b.Self.ID(), count)
	return
}

func (b *Bot) syncInfo() (err error) {
	var (
		ui UserInfo
	)
	return copier.Copy(&ui, b.Self)
}
