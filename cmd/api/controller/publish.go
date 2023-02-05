package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/nanakura/go-ramda"
	"mini_tiktok/cmd/api/model/api"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/videoservice"
	utils2 "mini_tiktok/pkg/utils"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	resp := new(api.PublishActionResp)
	actionResponse, err := rpc.VideoRpcClient.PublishAction(context.Background(), &videoservice.DouyinPublishActionRequest{
		Token: req.Token,
		Data: ramda.Map(func(in int8) byte {
			return byte(in)
		})(req.Data),
		Title: req.Title,
	})
	if err != nil || actionResponse == nil {
		resp.StatusCode = 1
		resp.StatusMessage = fmt.Sprintf("上传失败: %v", err)
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.StatusCode = int64(actionResponse.StatusCode)
	resp.StatusMessage = actionResponse.StatusMsg
	c.JSON(consts.StatusOK, resp)
}

// PublishList all users have same publish video list
func PublishList(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}

	resp := new(api.PublishListResp)
	userId, err := strconv.Atoi(req.UserID)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	publishListResponse, err := rpc.VideoRpcClient.PublishList(context.Background(), &videoservice.DouyinPublishListRequest{
		Token:  req.Token,
		UserId: int64(userId),
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	if resp.VideoList == nil {
		resp.VideoList = []*api.Video{}
	}
	resp.VideoList = utils2.CastUserserviceVideoToApiVideo(publishListResponse.VideoList)
	c.JSON(consts.StatusOK, resp)
}
