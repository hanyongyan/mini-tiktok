package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"mini_tiktok/cmd/user/utils"
	userservice "mini_tiktok/kitex_gen/userservice"
	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
	jwtutil "mini_tiktok/pkg/utils"
	utils2 "mini_tiktok/pkg/utils"
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
	pwd := utils.ScryptPwd(req.Password)
	newUser := &model.TUser{Name: req.Username, Password: pwd}
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
	if userInfo == nil {
		err = fmt.Errorf("查询失败, id为%v, error: %w", req.UserId, err)
		return
	}
	userId := req.UserId
	tfollow := query.Q.TFollow
	claims, _ := jwtutil.CheckToken(req.Token)
	toUserId := claims.UserId
	klog.Infof("userid: %v, touserId: %v", userId, toUserId)
	findFollowRes, _ := tfollow.WithContext(context.Background()).
		Where(tfollow.UserID.Eq(userId), tfollow.FollowerID.Eq(toUserId)).
		First()
	isFollow := false
	if findFollowRes != nil {
		isFollow = true
	}
	resp = &userservice.DouyinUserResponse{
		StatusCode: 0,
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
	// 关注操作
	queryFollow := query.Q.TFollow
	queryFriend := query.Q.TFriend
	queryUser := query.Q.TUser
	resp = &userservice.DouyinRelationActionResponse{}

	claims, flag := utils2.CheckToken(req.Token)
	// 解析 token 失败
	if !flag {
		err = errors.New("token is expired")
		return
	}
	follow := &model.TFollow{
		UserID:     claims.UserId,
		FollowerID: req.ToUserId,
	}
	if req.ActionType == 1 {
		// 关注操作
		resultFollow, err := queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(follow.UserID)).
			Where(queryFollow.FollowerID.Eq(follow.FollowerID)).First()
		// 说明还没有关注过
		if err != nil && err.Error() == "record not found" {
			err = queryFollow.WithContext(ctx).Create(follow)
			if err != nil {
				return nil, err
			}
			// 进行到此步说明 添加关注成功
			// 登录用户的关注数 +1
			_, _ = queryUser.WithContext(ctx).Where(queryUser.ID.Eq(claims.UserId)).Update(queryUser.FollowCount, queryUser.FollowCount.Add(1))
			// 关注用户的粉丝数 +1
			_, _ = queryUser.WithContext(ctx).Where(queryUser.ID.Eq(req.ToUserId)).Update(queryUser.FollowerCount, queryUser.FollowerCount.Add(1))
			// 查看关注用户是否关注了自己
			whetherToCare, _ := queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(follow.FollowerID)).
				Where(queryFollow.FollowerID.Eq(follow.UserID)).First()
			// 所关注的用户关注了自己
			// 添加好友数据
			if whetherToCare != nil {
				_ = queryFriend.WithContext(ctx).Create(&model.TFriend{
					UserID:   follow.UserID,
					FriendID: follow.FollowerID,
				})
				_ = queryFriend.WithContext(ctx).Create(&model.TFriend{
					UserID:   follow.FollowerID,
					FriendID: follow.UserID,
				})
			}

			resp.StatusCode = 0
			resp.StatusMsg = "关注成功"
			return resp, nil
		}
		// 说明已经关注过
		if resultFollow != nil {
			err = errors.New("请勿重复关注！")
			return nil, err
		}

	} else {
		// 取消关注操作
		// 先进行是否存在这样一种关注关系
		_, err := queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(follow.UserID)).Where(queryFollow.FollowerID.Eq(follow.FollowerID)).First()
		// 查询不到用户
		if err != nil && err.Error() == "record not found" {
			err = errors.New("请勿重复取消关注")
			return nil, err
		}

		// 进行删除数据库中的数据
		_, err = queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(follow.UserID)).Where(queryFollow.FollowerID.Eq(follow.FollowerID)).Delete()
		if err != nil {
			return nil, err
		}
		// 登录用户的关注数 -1
		_, _ = queryUser.WithContext(ctx).Where(queryUser.ID.Eq(claims.UserId)).Update(queryUser.FollowCount, queryUser.FollowCount.Sub(1))
		// 关注用户的粉丝数 -1
		_, _ = queryUser.WithContext(ctx).Where(queryUser.ID.Eq(req.ToUserId)).Update(queryUser.FollowerCount, queryUser.FollowerCount.Sub(1))

		// 查看是否存好友关系，如果存在好友关系，将好友关系从数据库中进行删除
		isFriend, _ := queryFriend.WithContext(ctx).Where(queryFriend.FriendID.Eq(follow.FollowerID), queryFriend.UserID.Eq(follow.UserID)).Find()
		// 说明存在好友关系
		if isFriend != nil {
			// 进行删除好友关系
			_, _ = queryFriend.WithContext(ctx).Where(queryFriend.FriendID.Eq(follow.FollowerID), queryFriend.UserID.Eq(follow.UserID)).Delete()
			_, _ = queryFriend.WithContext(ctx).Where(queryFriend.FriendID.Eq(follow.UserID), queryFriend.UserID.Eq(follow.FollowerID)).Delete()
		}
		resp.StatusMsg = "取消关注成功"
		resp.StatusCode = 0
		return resp, nil
	}
	return
}

// FollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowList(ctx context.Context, req *userservice.DouyinRelationFollowListRequest) (resp *userservice.DouyinRelationFollowListResponse, err error) {
	// 关注列表
	queryFollow := query.Q.TFollow
	// 进行查询用户
	queryUser := query.Q.TUser
	// 返回的用户信息
	// users := &[]model.TUser{}
	// 只进行查询关注用户的 id
	follows, err := queryFollow.WithContext(ctx).Select(queryFollow.FollowerID).Where(queryFollow.UserID.Eq(req.UserId)).Find()
	if err != nil {
		return
	}
	followerIds := make([]int64, len(follows))
	// 将关注用户的 id 进行提取出来
	for i, follow := range follows {
		followerIds[i] = follow.FollowerID
	}
	// 使用 select 进行规范查询的数据，使得不查询密码
	// 根据关注用户 id 查询到所有的 关注用户信息
	users, err := queryUser.WithContext(ctx).Select(queryUser.ID, queryUser.Name,
		queryUser.FollowCount, queryUser.FollowerCount).Where(queryUser.ID.In(followerIds...)).Find()
	if err != nil {
		return
	}
	// 将所有的关注用户信息进行添加到 返回值中
	var user userservice.User
	for _, tUser := range users {
		user.Id = tUser.ID
		user.Name = tUser.Name
		user.FollowerCount = tUser.FollowerCount
		user.FollowCount = tUser.FollowCount
		user.IsFollow = true
		resp.UserList = append(resp.UserList, &user)
	}
	resp.StatusCode = 0
	resp.StatusMsg = "the request succeeded"
	return
}

// FollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowerList(ctx context.Context, req *userservice.DouyinRelationFollowerListRequest) (resp *userservice.DouyinRelationFollowerListResponse, err error) {
	resp = &userservice.DouyinRelationFollowerListResponse{}
	// 用于查询粉丝的用户 id
	queryFollow := query.Q.TFollow
	// 根据粉丝的id查询出所有的粉丝
	queryUser := query.Q.TUser
	// 检索粉丝的id
	followers, err := queryFollow.WithContext(ctx).Select(queryFollow.UserID).Where(queryFollow.FollowerID.Eq(req.UserId)).Find()
	if err != nil {
		return
	}
	// 用来绑定粉丝id
	followersId := make([]int64, len(followers))
	for i, follower := range followers {
		followersId[i] = follower.UserID
	}
	// 进行查询粉丝用户信息
	t_users, err := queryUser.WithContext(ctx).Where(queryUser.ID.In(followersId...)).Find()
	if err != nil {
		return
	}
	for _, tUser := range t_users {
		var user userservice.User
		user.Id = tUser.ID
		user.FollowCount = tUser.FollowCount
		user.FollowerCount = tUser.FollowerCount
		user.Name = tUser.Name
		// 进行查询当前用户是否关注了此粉丝
		_, err := queryFollow.WithContext(ctx).Where(queryFollow.UserID.Eq(req.UserId)).Where(queryFollow.FollowerID.Eq(user.Id)).First()
		if err != nil && err.Error() == "record not found" {
			user.IsFollow = false
		} else {
			user.IsFollow = true
		}
		resp.UserList = append(resp.UserList, &user)
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询粉丝成功"
	return
}

// FriendList implements the UserServiceImpl interface. 好友列表
func (s *UserServiceImpl) FriendList(ctx context.Context, req *userservice.DouyinRelationFriendListRequest) (resp *userservice.DouyinRelationFriendListResponse, err error) {
	resp = &userservice.DouyinRelationFriendListResponse{}
	qUser := query.Q.TUser
	qFollow := query.Q.TFollow
	// 查询查看用户的关注列表
	followUsers, err := qFollow.WithContext(ctx).Where(qFollow.UserID.Eq(req.UserId)).Find()
	if err != nil {
		if err.Error() == "record not found" {
			resp.StatusCode = 0
			resp.StatusMsg = "用户没有好友"
			resp.UserList = nil
			return resp, nil
		}
	}
	// 抽离出用户关注的用户的 id
	followIds := make([]int64, len(followUsers))
	for i := 0; i < len(followIds); i++ {
		followIds[i] = followUsers[i].UserID
	}
	// 对关注的用户查询是否关注了登录用户
	find, _ := qFollow.WithContext(ctx).Where(qFollow.UserID.In(followIds...), qFollow.FollowerID.Eq(req.UserId)).Find()
	if find == nil {
		resp.UserList = nil
		resp.StatusCode = 0
		resp.StatusMsg = "该用户没有好友"
		return resp, nil
	}
	userIds := make([]int64, len(find))
	for i := 0; i < len(find); i++ {
		userIds[i] = find[i].UserID
	}
	// 进行查询用户信息
	users := make([]*userservice.User, len(userIds))
	tUsers, err := qUser.WithContext(ctx).Where(qUser.ID.In(userIds...)).Find()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(tUsers); i++ {
		first, _ := qFollow.WithContext(ctx).Where(qFollow.UserID.Eq(req.UserId), qFollow.FollowerID.Eq(tUsers[i].ID)).First()
		user := &userservice.User{
			Id:            tUsers[i].ID,
			Name:          tUsers[i].Name,
			FollowCount:   tUsers[i].FollowCount,
			FollowerCount: tUsers[i].FollowerCount,
			IsFollow:      first == nil,
		}
		users[i] = user
	}
	resp.UserList = users
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	return resp, nil

}
