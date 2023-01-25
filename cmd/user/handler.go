package main

import (
	"context"
	"fmt"
	"mini_tiktok/cmd/user/utils"
	"mini_tiktok/kitex_gen/userService"
	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
	jwtutil "mini_tiktok/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *userService.DouyinUserLoginRequest) (resp *userService.DouyinUserLoginResponse, err error) {
	q := query.Q
	checkRes, err := utils.CheckUser(q, req.Username, req.Password)
	token, err := jwtutil.CreateToken(checkRes.ID)
	if err != nil {
		return
	}
	resp = &userService.DouyinUserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "登陆成功",
		UserId:     checkRes.ID,
		Token:      token,
	}
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userService.DouyinUserRegisterRequest) (resp *userService.DouyinUserRegisterResponse, err error) {
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
	resp = &userService.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "登录成功",
		UserId:     newUser.ID,
		Token:      token,
	}
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *userService.DouyinUserRequest) (resp *userService.DouyinUserResponse, err error) {
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
	resp = &userService.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  "",
		User: &userService.User{
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
func (s *UserServiceImpl) Action(ctx context.Context, req *userService.DouyinRelationActionRequest) (resp *userService.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowList(ctx context.Context, req *userService.DouyinRelationFollowListRequest) (resp *userService.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowerList(ctx context.Context, req *userService.DouyinRelationFollowerListRequest) (resp *userService.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FriendList(ctx context.Context, req *userService.DouyinRelationFriendListRequest) (resp *userService.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}
