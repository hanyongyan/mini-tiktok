package controller

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"mini_tiktok/cmd/api/model/api"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/videoservice"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	videoId, _ := strconv.Atoi(req.VideoID)
	act, _ := strconv.Atoi(req.ActionType)
	// 删除操作
	if act == 2 {
		CommentId, _ := strconv.Atoi(*req.CommentID)
		info, err := rpc.VideoRpcClient.CommentAction(context.Background(), &videoservice.DouyinCommentActionRequest{
			Token:      req.Token,
			VideoId:    int64(videoId),
			ActionType: int32(act),
			CommentId:  int64(CommentId),
		})
		if err != nil {
			c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
			return
		}
		resp := &api.CommentActionResp{
			StatusCode:    int64(info.StatusCode),
			StatusMessage: info.StatusMsg,
		}
		c.JSON(consts.StatusOK, resp)
	} else {
		// 评论操作
		info, err := rpc.VideoRpcClient.CommentAction(context.Background(), &videoservice.DouyinCommentActionRequest{
			Token:       req.Token,
			VideoId:     int64(videoId),
			CommentText: *req.CommentText,
			ActionType:  int32(act),
		})
		if err != nil {
			c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
			return
		}

		user := &api.User{
			ID:            info.Comment.User.Id,
			Name:          info.Comment.User.Name,
			FollowCount:   info.Comment.User.FollowCount,
			FollowerCount: info.Comment.User.FollowerCount,
		}
		resp := &api.CommentActionResp{
			StatusCode:    int64(info.StatusCode),
			StatusMessage: info.StatusMsg,
			Comment: &api.Comment{
				ID:         info.Comment.Id,
				User:       user,
				Content:    info.Comment.Content,
				CreateDate: info.Comment.CreateDate,
			},
		}

		c.JSON(consts.StatusOK, resp)
	}
}

// CommentList all videos have same demo comment list
func CommentList(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	videoId, _ := strconv.Atoi(req.VideoID)
	resp, err := rpc.VideoRpcClient.CommentList(context.Background(), &videoservice.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: int64(videoId),
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	c.JSON(consts.StatusOK, resp)
}
