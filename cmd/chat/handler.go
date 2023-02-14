package main

import (
	"context"
	chatservice "mini_tiktok/kitex_gen/service/ChatService"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *chatservice.MessageActionReq) (resp *chatservice.MessageActionResp, err error) {
	// TODO: Your code here...
	return
}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *chatservice.MessageChatReq) (resp *chatservice.MessageChatResp, err error) {
	// TODO: Your code here...
	return
}
