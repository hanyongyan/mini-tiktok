package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"mini_tiktok/cmd/api/model/api"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/videoservice"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}

	videoId, _ := strconv.Atoi(req.VideoID)
	actionType, _ := strconv.Atoi(req.ActionType)
	r, err := rpc.VideoRpcClient.FavoriteAction(context.Background(), &videoservice.DouyinFavoriteActionRequest{
		Token:      req.Token,
		VideoId:    int64(videoId),
		ActionType: int32(actionType),
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	resp := &api.FavoriteActionResp{
		StatusCode:    0,
		StatusMessage: r.StatusMsg,
	}

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList all users have same favorite video list
func FavoriteList(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	userId, _ := strconv.Atoi(req.UserID)
	r, err := rpc.VideoRpcClient.FavoriteList(context.Background(), &videoservice.DouyinFavoriteListRequest{
		UserId: int64(userId),
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}

	resp := &videoservice.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功",
		VideoList:  r.VideoList,
	}
	c.JSON(consts.StatusOK, resp)
}
