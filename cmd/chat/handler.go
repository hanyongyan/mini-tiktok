package main

import (
	"context"
	"fmt"
	chatservice "mini_tiktok/kitex_gen/chatservice"
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

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
