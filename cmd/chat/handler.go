package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/nanakura/go-ramda"
	dao "mini_tiktok/cmd/chat/dao"
	"mini_tiktok/cmd/chat/model"
	chatservice "mini_tiktok/kitex_gen/chatservice"
	"mini_tiktok/pkg/utils"
	"strconv"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *chatservice.MessageActionReq) (resp *chatservice.MessageActionResp, err error) {
	claims, ok := utils.CheckToken(req.GetToken())
	if !ok {
		err = errors.New("解析token失败")
		return
	}
	userId := claims.UserId
	toUserId, err := strconv.ParseInt(req.GetToUserKey(), 10, 64)
	if err != nil {
		err = fmt.Errorf("用户id参数非法：%w", err)
		return
	}
	chatKey := genChatKey(userId, toUserId)
	msgDao := new(dao.MsgDao)
	err = msgDao.Insert(chatKey, req.GetContent())
	if err != nil {
		err = fmt.Errorf("事务失败：%w", err)
		return
	}
	return
}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *chatservice.MessageChatReq) (resp *chatservice.MessageChatResp, err error) {
	claims, ok := utils.CheckToken(req.GetToken())
	if !ok {
		err = errors.New("解析token失败")
		return
	}
	userId := claims.UserId
	toUserId, err := strconv.ParseInt(req.GetToUserId(), 10, 64)
	if err != nil {
		err = fmt.Errorf("用户id参数非法：%w", err)
		return
	}
	chatKey := genChatKey(userId, toUserId)
	msgDao := new(dao.MsgDao)
	messages, err := msgDao.GetAllByChatKey(chatKey)
	if err != nil {
		err = fmt.Errorf("事务失败：%w", err)
		return
	}
	resp = &chatservice.MessageChatResp{MessageList: ramda.Map(func(t model.Message) *chatservice.Message {
		return &chatservice.Message{Id: int64(t.ID.Pid()), Content: t.Content, CreateTime: t.CreateTime.Format("2006-01-02 15:04:05")}
	})(messages)}

	return
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
