package controller

import (
	"context"
	"github.com/nanakura/go-ramda"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/chatservice"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(ctx context.Context, c *app.RequestContext) {
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ChatResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ChatResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	_, err = rpc.ChatRpcClient.MessageAction(ctx, &chatservice.DouyinMessageActionRequest{
		Token:      c.Query("token"),
		ToUserId:   toUserId,
		ActionType: int32(actionType),
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
	toUserId, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ChatResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	resp, err := rpc.ChatRpcClient.MessageChat(ctx, &chatservice.DouyinMessageChatRequest{
		Token: token, ToUserId: toUserId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ChatResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: ramda.Map(func(t *chatservice.Message) Message {
		time, _ := strconv.ParseInt(*t.CreateTime, 10, 64)
		return Message{Id: t.Id, Content: t.Content, CreateTime: time}
	})(resp.GetMessageList())})
}
