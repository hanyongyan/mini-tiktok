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
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *chatservice.DouyinMessageActionRequest) (resp *chatservice.DouyinMessageActionResponse, err error) {
	claims, ok := utils.CheckToken(req.GetToken())
	if !ok {
		err = errors.New("解析token失败")
		return
	}
	userId := claims.UserId
	toUserId := req.GetToUserId()
	msgDao := new(dao.MsgDao)
	err = msgDao.Insert(ctx, userId, toUserId, req.GetContent())
	if err != nil {
		err = fmt.Errorf("事务失败：%w", err)
		return
	}
	return
}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *chatservice.DouyinMessageChatRequest) (resp *chatservice.DouyinMessageChatResponse, err error) {
	claims, ok := utils.CheckToken(req.GetToken())
	if !ok {
		err = errors.New("解析token失败")
		return
	}
	userId := claims.UserId
	toUserId := req.GetToUserId()
	msgDao := new(dao.MsgDao)
	messages, err := msgDao.GetAllByChatKey(ctx, userId, toUserId)
	if err != nil {
		err = fmt.Errorf("事务失败：%w", err)
		return
	}
	resp = &chatservice.DouyinMessageChatResponse{MessageList: ramda.MapIndexed(func(t model.Message, idx int) *chatservice.Message {
		time := fmt.Sprintf("%v", t.CreateTime)
		return &chatservice.Message{Id: int64(idx) + 1, Content: t.Content, CreateTime: &time}
	})(messages)}

	return
}
