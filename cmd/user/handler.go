package main

import (
	"context"
	userService "mini_tiktok/cmd/user/kitex_gen/userService"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *userService.DouyinUserLoginRequest) (resp *userService.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *userService.DouyinUserRegisterRequest) (resp *userService.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *userService.DouyinUserRequest) (resp *userService.DouyinUserResponse, err error) {
	// TODO: Your code here...
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
