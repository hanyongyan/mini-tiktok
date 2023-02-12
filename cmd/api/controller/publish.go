package controller

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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
	title := c.FormValue("title")
	token := c.FormValue("token")
	f, _ := c.FormFile("data")
	data := make([]byte, f.Size)
	fd, _ := f.Open()
	fd.Read(data)
	fd.Close()
	if len(title) == 0 {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "标题不能为空"})
		return
	}

	hlog.Info("视频标题", string(title))

	hlog.Info("query", len(data))
	if len(data) == 0 {
		c.JSON(consts.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "视频不能为空"})
		return
	}

	resp := new(api.PublishActionResp)
	_, err = rpc.VideoRpcClient.PublishAction(context.Background(), &videoservice.DouyinPublishActionRequest{
		Token: string(token),
		Data:  data,
		Title: string(title),
	})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = fmt.Sprintf("上传失败: %v", err)
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMessage = "上传成功"
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
