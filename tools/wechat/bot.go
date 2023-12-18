package wechat

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/jinzhu/copier"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"

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
		Messages    chan []byte
	}
	messageType string
	WsBody      struct {
		Data      interface{} `json:"data"`
		MsgType   messageType `json:"msg_type"`
		Timestamp int64       `json:"timestamp"`
	}
	ChatMsg struct {
		Msg        string `json:"msg"`
		SenderId   string `json:"sender_id"`
		ReceiverId string `json:"receiver_id"`
		OwnerId    string `json:"owner_id"`
		IsGroup    bool   `json:"is_group"`
	}

	syncReqMsgBody struct {
		FindId  string `json:"find_id"`
		IsGroup bool   `json:"is_group"`
		Count   int    `json:"count"`
	}
	syncResMsgBody struct {
		Messages  []model.ChatMessage `json:"messages"`
		LastCount int64               `json:"last_count"`
	}
	syncAvatarBody struct {
		FindId  string `form:"find_id"`
		IsGroup bool   `form:"is_group"`
	}
)

const (
	messagePingType         messageType = "ping"
	messageSyncFriendsType  messageType = "syncFriends"
	messageSyncGroupsType   messageType = "syncGroups"
	messageSyncMessagesType messageType = "syncMessages"
	messageLoginType        messageType = "login"
	messageChatType         messageType = "chat"
	selfInfoType            messageType = "selfInfo"
	AvatarType              messageType = "avatar"
)

func (b *Bot) replyMessage(msg *openwechat.Message, sender *openwechat.User, reply *Reply) {
	var (
		matched bool
		err     error
	)
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
		if !msg.IsText() {
			return
		}
		var (
			sender *openwechat.User
			err1   error
		)
		if sender, err1 = msg.Sender(); err1 != nil {
			return
		}
		if sender.IsGroup() {
			return
		}
		_ = b.dealMessage(sender.ID(), b.Self.ID(), msg.Content, sender.IsGroup(), false, msg.CreateTime)
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
		b.SendMessage(WsBody{
			MsgType: messageLoginType,
			Data:    "success",
		})
	case openwechat.LoginCodeScanned:
		b.SendMessage(WsBody{
			MsgType: messageLoginType,
			Data:    "scanned",
		})
	case openwechat.LoginCodeTimeout:
		b.SendMessage(WsBody{
			MsgType: messageLoginType,
			Data:    "timeout",
		})
	default:
		return
	}
}

func (b *Bot) ScanCallBack(body openwechat.CheckLoginResponse) {
	logrus.Info(string(body))
	b.SendMessage(WsBody{MsgType: messageLoginType, Data: "scanned"})
}

// 保存历史信息
func (b *Bot) saveMessage(message, ownerId, senderId, receiverId string, createdTime int64, isGroup bool, err error) (cm model.ChatMessage) {
	cm = model.ChatMessage{
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
	return cm
}

func (b *Bot) SendMessage(wb WsBody) {

	go func() {
		if wb.Data == nil {
			wb.Data = map[string]string{}
		}
		wb.Timestamp = time.Now().Unix()
		msgBytes, _ := json.Marshal(&wb)
		b.Messages <- msgBytes
	}()
}

func (b *Bot) ReceiveMessage(message []byte) (err error) {
	if len(message) <= 0 {
		return
	}
	var wb WsBody
	if err = json.Unmarshal(message, &wb); err != nil {
		return
	}

	switch wb.MsgType {
	case messagePingType:
		b.SendMessage(WsBody{
			MsgType: messagePingType,
			Data:    "pong",
		})
	case messageSyncFriendsType:
		return b.syncFriends()
	case messageSyncGroupsType:
		return b.syncGroups()
	case messageSyncMessagesType:
		return b.syncMessage(wb)
	case selfInfoType:
		return b.syncInfo()
	default:
		b.Chat(wb)
	}
	return
}
func (b *Bot) getSyncMessageTotalKey() string {
	return fmt.Sprintf("sync_message_total_%v", b.Self.ID())
}

func (b *Bot) CleanMessages() {
	b.Messages = make(chan []byte)
}
func (b *Bot) Chat(wb WsBody) {
	var (
		err     error
		chatMsg ChatMsg
		body    []byte
	)
	if body, err = json.Marshal(&wb.Data); err != nil {
		return
	}
	if err = json.Unmarshal(body, &chatMsg); err != nil {
		return
	}
	_ = b.dealMessage(b.Self.ID(), chatMsg.ReceiverId, chatMsg.Msg, chatMsg.IsGroup, true, wb.Timestamp)
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
	var msgReplyType = messageChatType
	if b.Self.ID() == sendId {
		// 自身为发送者,需要给好友/群组发送消息
		err = b.sendFriendGroupMsg(receiveId, msg, isGroup)
	}
	b.SendMessage(WsBody{
		MsgType: msgReplyType,
		Data:    b.saveMessage(msg, b.Self.ID(), sendId, receiveId, createdTime, isGroup, err),
	})
	return
}

func (b *Bot) syncFriends() (err error) {
	var (
		friends openwechat.Friends
		ot      []UserInfo
	)
	if friends, err = b.Friends(); err != nil {
		return
	}

	if err = copier.Copy(&ot, &friends); err != nil {
		return err
	}

	for i := 0; i < len(friends); i++ {
		var friend = friends[i]
		ot[i].Uin = friend.ID()
		ot[i].HeadImgBase64, _ = b.getHeadImgBase64(friend.GetAvatarResponse)
	}
	b.SendMessage(WsBody{
		MsgType: messageSyncFriendsType,
		Data:    ot,
	})
	return
}
func (b *Bot) syncGroups() (err error) {
	var groups openwechat.Groups
	if groups, err = b.Groups(); err != nil {
		return
	}
	b.SendMessage(WsBody{
		MsgType: messageSyncGroupsType,
		Data:    groups,
	})
	return
}

func (b *Bot) syncMessage(wb WsBody) (err error) {
	var (
		syb   syncReqMsgBody
		res   syncResMsgBody
		limit = 5
		body  []byte
	)
	if body, err = json.Marshal(&wb.Data); err != nil {
		return
	}
	if err = json.Unmarshal(body, &syb); err != nil {
		return
	}
	// 默认同步最近5条
	tx := db.G.Model(model.ChatMessage{}).Where("owner_id = ? and is_group = ? and  (sender_id = ? or receiver_id = ? )",
		b.Self.ID(), syb.IsGroup, syb.FindId, syb.FindId)
	if err = tx.Order(" created_time desc ").Offset(syb.Count).Limit(limit).Find(&res.Messages).Error; err != nil {
		return
	}
	res.Messages = tools.Reverse[model.ChatMessage](res.Messages)

	if err = tx.Count(&res.LastCount).Error; err != nil {
		return err
	}
	res.LastCount -= int64(syb.Count + limit)
	b.SendMessage(WsBody{
		MsgType: messageSyncMessagesType,
		Data:    res,
	})
	logrus.Infof("用户(%v:%v):消息%v条同步完成!,剩余%v条消息", b.Self.NickName, b.Self.ID(), len(res.Messages), res.LastCount)
	return
}

func (b *Bot) syncInfo() (err error) {
	var (
		ui UserInfo
	)
	if err = copier.Copy(&ui, b.Self); err != nil {
		return
	}
	ui.Uin = b.Self.ID()
	if ui.HeadImgBase64, err = b.getHeadImgBase64(b.Self.GetAvatarResponse); err != nil {
		return
	}
	b.SendMessage(WsBody{
		MsgType: selfInfoType,
		Data:    ui,
	})
	return
}

func (b *Bot) getHeadImgBase64(f func() (*http.Response, error)) (bs string, err error) {
	var (
		resp *http.Response
		body []byte
	)
	if resp, err = f(); err != nil {
		return
	}
	defer resp.Body.Close()
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}
	bs = base64.StdEncoding.EncodeToString(body)
	return
}

func (b *Bot) syncAvatar(wb WsBody) (err error) {
	var (
		p    syncAvatarBody
		bs   string
		body []byte
	)
	if body, err = json.Marshal(&wb.Data); err != nil {
		return
	}
	if err = json.Unmarshal(body, &p); err != nil {
		return
	}
	if p.FindId == b.Self.ID() {
		bs, err = b.getHeadImgBase64(b.Self.GetAvatarResponse)
	} else if p.IsGroup {
		var group *openwechat.Group
		if group, err = b.FindGroup(p.FindId); err != nil {
			return
		}
		bs, err = b.getHeadImgBase64(group.GetAvatarResponse)

	} else {
		var friend *openwechat.Friend
		if friend, err = b.FindFriend(p.FindId); err != nil {
			return
		}
		bs, err = b.getHeadImgBase64(friend.GetAvatarResponse)

	}
	b.SendMessage(WsBody{
		MsgType: AvatarType,
		Data:    bs,
	})
	return
}
