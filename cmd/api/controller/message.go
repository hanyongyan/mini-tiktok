package controller

import (
	"context"
	"fmt"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/chatservice"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

type MessageActionReq struct {
	Token      string `json:"token"`
	ToUserId   string `json:"to_user_id"`
	ActionType string `json:"action_type"`
	Content    string `json:"content"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(_ context.Context, c *app.RequestContext) {
	var req MessageActionReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// MessageChat all users have same follow list
func MessageChat(ctx context.Context, c *app.RequestContext) {
	//token := c.Query("token")
	//toUserId := c.Query("to_user_id")
	rpc.ChatRpcClient.MessageChat(ctx, chatservice.MessageChatReq{})
	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}})
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
