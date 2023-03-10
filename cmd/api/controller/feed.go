package controller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"mini_tiktok/cmd/api/model/api"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/videoservice"
	utils2 "mini_tiktok/pkg/utils"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.FeedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	var latestDate int64
	if req.LatestTime == nil {
		latestDate = time.Now().Unix() * 1000
	} else {
		latestDate, _ = strconv.ParseInt(*req.LatestTime, 10, 64)
	}
	var token string
	if req.Token == nil {
		token = ""
	} else {
		token = *req.Token
	}
	feedResponse, err := rpc.VideoRpcClient.Feed(context.Background(),
		&videoservice.DouyinFeedRequest{
			LatestTime: latestDate,
			Token:      token,
		},
	)
	resp := new(api.FeedResp)
	if err != nil {
		c.JSON(consts.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	resp.StatusCode = 0
	resp.NextTime = time.Now().Unix()
	if feedResponse == nil || feedResponse.VideoList == nil {
		resp.StatusMessage = fmt.Sprintf("获取视频失败, error：%v", err)
		resp.VideoList = []*api.Video{}
	} else {
		resp.VideoList = utils2.CastUserserviceVideoToApiVideo(feedResponse.VideoList)
	}
	hlog.Infof("feed: %+v", resp)
	c.JSON(consts.StatusOK, resp)
}
