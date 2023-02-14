package controller

import (
	"context"
	"github.com/nanakura/go-ramda"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/chatservice"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(ctx context.Context, c *app.RequestContext) {
	_, err := rpc.ChatRpcClient.MessageAction(ctx, &chatservice.MessageActionReq{
		Token:      c.Query("token"),
		ToUserKey:  c.Query("to_user_id"),
		ActionType: c.Query("action_type"),
		Content:    c.Query("content"),
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
	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: ramda.Map(func(t *chatservice.Message) Message {
		return Message{Id: t.Id, Content: t.Content, CreateTime: t.CreateTime}
	})(resp.GetMessageList())})
}
