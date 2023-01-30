// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "mini_tiktok/cmd/api/biz/model/api"
	"mini_tiktok/cmd/api/biz/rpc"
	userservice "mini_tiktok/kitex_gen/userservice"
	"mini_tiktok/kitex_gen/videoservice"
	"strconv"
)

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FeedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FeedResp)

	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /douyin/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	registerResponse, err := rpc.UserRpcClient.Register(context.Background(), &userservice.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"code": 0, "message": err.Error()})
		return
	}
	resp := &api.UserRegisterResp{
		StatusCode:    int64(registerResponse.StatusCode),
		StatusMessage: registerResponse.StatusMsg,
		UserID:        registerResponse.UserId,
		Token:         registerResponse.Token,
	}

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Info("start call login rpc api")
	loginResponse, err := rpc.UserRpcClient.Login(context.Background(), &userservice.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	hlog.Info("call login rpc api end")
	if err != nil {
		hlog.Error("error occur", err)
		c.JSON(consts.StatusOK, utils.H{"code": 0, "message": err.Error()})
		return
	}
	if loginResponse == nil {
		c.JSON(consts.StatusOK, utils.H{
			"status": "nil",
		})
		return
	}
	resp := &api.UserLoginResp{
		StatusCode:    int64(loginResponse.StatusCode),
		StatusMessage: loginResponse.StatusMsg,
		UserID:        loginResponse.UserId,
		Token:         loginResponse.Token,
	}
	hlog.Infof("get resp: %+v", loginResponse)

	c.JSON(consts.StatusOK, resp)
}

// User .
// @router /douyin/user [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	userId, _ := strconv.Atoi(req.UserID)
	info, err := rpc.UserRpcClient.Info(context.Background(), &userservice.DouyinUserRequest{UserId: int64(userId), Token: req.Token})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code":    0,
			"message": err.Error(),
		})
		return
	}
	resp := &api.UserResp{
		StatusCode:    int64(info.StatusCode),
		StatusMessage: info.StatusMsg,
		User: &api.User{
			ID:            info.User.Id,
			Name:          info.User.Name,
			FollowCount:   info.User.FollowCount,
			FollowerCount: info.User.FollowerCount,
			IsFollow:      info.User.IsFollow,
		},
	}

	c.JSON(consts.StatusOK, resp)
}

// PublishAction .
// @router /douyin/publish/action [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.PublishActionResp)

	c.JSON(consts.StatusOK, resp)
}

// PublishList .
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.PublishListResp)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteAction .
// @router /douyin/favorite/action [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FavoriteActionResp)

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.FavoriteListResp)

	c.JSON(consts.StatusOK, resp)
}

// CommentAction .
// @router /douyin/comment/action [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
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
			c.JSON(consts.StatusOK, utils.H{
				"code":    401,
				"message": err.Error(),
			})
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
			c.JSON(consts.StatusOK, utils.H{
				"code":    401,
				"message": err.Error(),
			})
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

// CommentList .
// @router /douyin/comment/list [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	videoId, _ := strconv.Atoi(req.VideoID)
	resp, err := rpc.VideoRpcClient.CommentList(context.Background(), &videoservice.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: int64(videoId),
	})

	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status": "nil"})
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// RelationAction .
// @router /douyin/relation/action [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {

	// 前面有中间件检测当前用户是否在登录
	var err error
	var req *userservice.DouyinRelationActionRequest
	// 将参数进行绑定
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 调用 userService 的函数
	action, err := rpc.UserRpcClient.Action(ctx, req)
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"code": 1, "message": err.Error()})
		return
	}
	if action == nil {
		c.JSON(consts.StatusOK, utils.H{"status": "nil"})
		return
	}

	resp := &api.RelationActionResp{
		StatusCode:    int64(action.StatusCode),
		StatusMessage: action.StatusMsg,
	}
	c.JSON(consts.StatusOK, resp)
}

// RelationFollowList .
// @router /douyin/relation/follow/list [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	// 关注列表
	var err error
	var req api.RelationFollowListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	hlog.Info("start call login rpc api")
	// 将 字符串 转为 int64
	userId, _ := strconv.ParseInt(req.UserID, 10, 64)
	resp, err := rpc.UserRpcClient.FollowList(ctx, &userservice.DouyinRelationFollowListRequest{
		UserId: userId,
		Token:  req.Token,
	})
	hlog.Info("call login rpc api end")
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status": "nil"})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// RelationFollowerList .
// @router /douyin/relation/follower/list [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFollowerListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RelationFollowerListResp)

	c.JSON(consts.StatusOK, resp)
}

// RelationFriendList .
// @router /douyin/relation/friend/list [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFriendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RelationFriendListResp)

	c.JSON(consts.StatusOK, resp)
}
