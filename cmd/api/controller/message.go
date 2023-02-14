package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
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
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var req MessageActionReq
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	_, err := rpc.ChatRpcClient.MessageAction(ctx, &chatservice.MessageActionReq{
		Token:      req.Token,
		ToUserKey:  req.ToUserId,
		ActionType: req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// MessageChat all users have same follow list
func MessageChat(ctx context.Context, c *app.RequestContext) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	resp, err := rpc.ChatRpcClient.MessageChat(ctx, &chatservice.MessageChatReq{
		Token: token, ToUserId: toUserId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ChatResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, utils.H{
		"status_code":  0,
		"message_list": resp.GetMessageList(),
	})
}

