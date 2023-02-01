package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"mini_tiktok/cmd/api/model/api"
	"mini_tiktok/cmd/api/rpc"
	"mini_tiktok/kitex_gen/userservice"
	"strconv"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(_ context.Context, c *app.RequestContext) {
	var err error
	username := c.Query("username")
	password := c.Query("password")
	hlog.Info("start call login rpc api")
	hlog.Infof("name: %v, pass: %v", username, password)
	resp := &api.UserRegisterResp{}
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	registerResponse, err := rpc.UserRpcClient.Register(context.Background(), &userservice.DouyinUserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	resp = &api.UserRegisterResp{
		StatusCode:    int64(registerResponse.StatusCode),
		StatusMessage: registerResponse.StatusMsg,
		UserID:        registerResponse.UserId,
		Token:         registerResponse.Token,
	}

	c.JSON(consts.StatusOK, resp)
}

func Login(_ context.Context, c *app.RequestContext) {
	var err error
	username := c.Query("username")
	password := c.Query("password")
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	hlog.Info("start call login rpc api")
	hlog.Infof("name: %v, pass: %v", username, password)
	loginResponse, err := rpc.UserRpcClient.Login(context.Background(), &userservice.DouyinUserLoginRequest{
		Username: username,
		Password: password,
	})
	hlog.Info("call login rpc api end")
	if err != nil {
		hlog.Error("error occur", err)
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	if loginResponse == nil {
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
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

func UserInfo(_ context.Context, c *app.RequestContext) {
	var err error
	var req api.UserReq
	err = c.BindAndValidate(&req)
	resp := &api.UserResp{}
	if err != nil {
		c.JSON(consts.StatusBadRequest, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}

	userId, _ := strconv.Atoi(req.UserID)
	info, err := rpc.UserRpcClient.Info(context.Background(), &userservice.DouyinUserRequest{UserId: int64(userId), Token: req.Token})
	if err != nil {
		hlog.Infof("获取用户信息时error occur: %v", err)
		c.JSON(consts.StatusOK, utils.H{"status_code": 1, "status_msg": err.Error()})
		return
	}
	resp = &api.UserResp{
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
