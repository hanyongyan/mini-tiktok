package controller

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"mini_tiktok/cmd/api/model/api"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/userservice"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(_ context.Context, c *app.RequestContext) {
	// 关注操作
	var err error
	var req api.RelationActionReq
	var resp api.RelationActionResp
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.StatusMessage = err.Error()
		resp.StatusCode = 1
		return
	}
	toUserId, toUserIdErr := strconv.ParseInt(req.ToUserID, 10, 64)
	actionType, actionTypeErr := strconv.ParseInt(req.ActionType, 10, 32)
	if toUserIdErr != nil || actionTypeErr != nil {
		resp.StatusCode = 0
		resp.StatusMessage = "传入的参数错误"
		c.JSON(consts.StatusOK, resp)
		return
	}
	// userService 完成 关注操作
	result, err := rpc.UserRpcClient.Action(context.Background(), &userservice.DouyinRelationActionRequest{
		Token:      req.Token,
		ToUserId:   toUserId,
		ActionType: int32(actionType),
	})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
		c.JSON(consts.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMessage = result.StatusMsg
	c.JSON(consts.StatusOK, resp)
}

// FollowList all users have same follow list
func FollowList(_ context.Context, c *app.RequestContext) {
	// 关注列表
	var err error
	var req api.RelationFollowListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	// 将 字符串 转为 int64
	userId, _ := strconv.ParseInt(req.UserID, 10, 64)
	resp, err := rpc.UserRpcClient.FollowList(context.Background(), &userservice.DouyinRelationFollowListRequest{
		UserId: userId,
		Token:  req.Token,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// FollowerList all users have same follower list
func FollowerList(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFollowerListReq
	var resp api.RelationFollowerListResp
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
	}
	result, err := rpc.UserRpcClient.FollowerList(context.Background(), &userservice.DouyinRelationFollowerListRequest{
		UserId: userId,
		Token:  req.Token,
	})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = "请求失败"
		c.JSON(consts.StatusOK, resp)
	}

	c.JSON(consts.StatusOK, result)
}

// FriendList all users have same friend list
func FriendList(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.RelationFriendListReq
	var resp api.RelationFriendListResp
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	userId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	result, err := rpc.UserRpcClient.FriendList(context.Background(), &userservice.DouyinRelationFriendListRequest{
		UserId: userId,
		Token:  req.Token,
	})
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMessage = err.Error()
		resp.UserList = nil
		c.JSON(consts.StatusOK, resp)
	}
	c.JSON(consts.StatusOK, result)
}
