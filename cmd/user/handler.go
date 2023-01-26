package main

import (
	"context"
	"fmt"
	"mini_tiktok/cmd/user/utils"
	userservice "mini_tiktok/kitex_gen/userservice"
	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
	jwtutil "mini_tiktok/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *userservice.DouyinUserLoginRequest) (resp *userservice.DouyinUserLoginResponse, err error) {
	resp = &userservice.DouyinUserLoginResponse{}
	q := query.Q
	checkRes, err2 := utils.CheckUser(q, req.Username, req.Password)
	if err2 != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err2.Error()
		return
	}
	token, err2 := jwtutil.CreateToken(checkRes.ID)
	if err2 != nil {
		resp.StatusCode = 1
		resp.StatusMsg = err2.Error()
		return
	}
	resp = &userservice.DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "登陆成功",
		UserId:     checkRes.ID,
		Token:      token,
	}
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userservice.DouyinUserRegisterRequest) (resp *userservice.DouyinUserRegisterResponse, err error) {
	q := query.Q
	checkRes, _ := utils.CheckUser(q, req.Username, req.Password)
	if checkRes != nil {
		err = fmt.Errorf("注册失败：用户已存在 %w", err)
		return
	}
	newUser := &model.TUser{Name: req.Username, Password: req.Password}
	err = q.WithContext(context.Background()).TUser.Create(newUser)
	if err != nil {
		err = fmt.Errorf("注册失败: %w", err)
		return
	}
	token, err := jwtutil.CreateToken(newUser.ID)
	if err != nil {
		err = fmt.Errorf("token生成失败: %w", err)
		return
	}
	resp = &userservice.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		UserId:     newUser.ID,
		Token:      token,
	}
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *userservice.DouyinUserRequest) (resp *userservice.DouyinUserResponse, err error) {
	u := query.Q.TUser
	userInfo, err := u.WithContext(context.Background()).Where(u.ID.Eq(req.UserId)).First()
	if err != nil {
		err = fmt.Errorf("查询失败: %w", err)
		return
	}
	userId := req.UserId
	tfollow := query.Q.TFollow
	claims, _ := jwtutil.CheckToken(req.Token)
	toUserId := claims.UserId
	findFollowRes, err := tfollow.WithContext(context.Background()).
		Where(tfollow.UserID.Eq(userId), tfollow.FollowerID.Eq(int64(toUserId))).
		First()
	if err != nil {
		return
	}
	isFollow := false
	if findFollowRes != nil {
		isFollow = true
	}
	resp = &userservice.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  "",
		User: &userservice.User{
			Id:            userInfo.ID,
			Name:          userInfo.Name,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			IsFollow:      isFollow,
		},
	}
	return
}

// Action implements the UserServiceImpl interface.
func (s *UserServiceImpl) Action(ctx context.Context, req *userservice.DouyinRelationActionRequest) (resp *userservice.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowList(ctx context.Context, req *userservice.DouyinRelationFollowListRequest) (resp *userservice.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowerList(ctx context.Context, req *userservice.DouyinRelationFollowerListRequest) (resp *userservice.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FriendList(ctx context.Context, req *userservice.DouyinRelationFriendListRequest) (resp *userservice.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
