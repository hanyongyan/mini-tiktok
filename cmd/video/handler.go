package main

import (
	"context"
	"errors"
	"fmt"
	videoservice "mini_tiktok/kitex_gen/videoservice"
	"mini_tiktok/pkg/cache"
	"mini_tiktok/pkg/consts"
	"mini_tiktok/pkg/dal/query"
	jwtutil "mini_tiktok/pkg/utils"
	"strconv"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videoservice.DouyinPublishActionRequest) (resp *videoservice.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videoservice.DouyinFeedRequest) (resp *videoservice.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *videoservice.DouyinPublishListRequest) (resp *videoservice.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction 2023-1-27 @Auth by 李卓轩 version 1.0
// 赞操作
// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *videoservice.DouyinFavoriteActionRequest) (resp *videoservice.DouyinFavoriteActionResponse, err error) {
	// 通过 token 解析出当前用户
	claims, flag := jwtutil.CheckToken(req.Token)
	// 说明 token 已经过期
	if !flag {
		return nil, errors.New("token is expired")
	}

	//判断当前用户是否点赞
	result, err := cache.RedisCache.RedisClient.SIsMember(context.Background(), consts.FavoriteActionPrefix+strconv.FormatInt(claims.UserId, 10), req.VideoId).Result()
	if err != nil {
		err = fmt.Errorf("redis访问失败")
		return
	}

	//已点过赞，取消点赞
	if result {
		// redis数据库中删除关联
		_, err1 := cache.RedisCache.RedisClient.SRem(context.Background(), consts.FavoriteActionPrefix+strconv.FormatInt(claims.UserId, 10), req.VideoId).Result()
		if err1 != nil {
			err1 = fmt.Errorf("redis 取消点赞失败")
			return
		}
		resp = &videoservice.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "已取消点赞",
		}
		return
	}

	// 在数据库中查询点赞信息
	q := query.Q
	favorite := q.TFavorite
	first, err := q.WithContext(context.Background()).TFavorite.Where(favorite.UserID.Eq(claims.UserId), favorite.VideoID.Eq(req.VideoId)).First()

	if err != nil {
		err = fmt.Errorf("数据库获取数据失败")
	}

	// 查询为空
	if first == nil {
		// 将点赞存入redis
		cache.RedisCache.RedisClient.SAdd(context.Background(), consts.FavoriteActionPrefix+strconv.FormatInt(claims.UserId, 10), req.VideoId, 0)
		resp = &videoservice.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "已成功点赞",
		}
		return
	}

	// 查询数据库，数据库为已点赞，取消点赞
	if first.Status {
		_, err1 := q.WithContext(context.Background()).TFavorite.Update(favorite.Status, false)
		if err1 != nil {
			err1 = fmt.Errorf("更新数据库失败")
		}
		resp = &videoservice.DouyinFavoriteActionResponse{
			StatusCode: 0,
			StatusMsg:  "已取消点赞",
		}
		return
	}

	resp = &videoservice.DouyinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  "已成功点赞",
	}
	return
}

// FavoriteList 2023-1-27 @Auth by 李卓轩 version 1.0
// 喜欢列表
// FavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteList(ctx context.Context, req *videoservice.DouyinFavoriteListRequest) (resp *videoservice.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	// 通过 token 解析出当前用户
	claims, flag := jwtutil.CheckToken(req.Token)
	// 说明 token 已经过期
	if !flag {
		return nil, errors.New("token is expired")
	}

	q := query.Q
	favorite := q.TFavorite
	// 查询数据库得到喜欢列表
	data, err := q.WithContext(context.Background()).TFavorite.Where(favorite.UserID.Eq(claims.UserId)).Find()
	ids := make([]int64, 10)
	//得到喜欢视频的所有id
	for _, fav := range data {
		ids = append(ids, fav.VideoID)
	}

	//查询所有的喜欢视频信息
	video := q.TVideo
	find, err := q.WithContext(context.Background()).TVideo.Where(video.ID.In(ids...)).Find()
	if err != nil {
		err = fmt.Errorf("查询失败")
	}
	var videos []*videoservice.Video
	//通过用用户id查询用户
	Tuser := q.TUser
	for _, videosInfo := range find {
		var vid videoservice.Video
		var usr videoservice.User
		vid.FavoriteCount = videosInfo.FavoriteCount
		vid.Id = videosInfo.ID
		vid.CoverUrl = videosInfo.CoverURL
		vid.PlayUrl = videosInfo.PlayURL
		vid.IsFavorite = videosInfo.IsFavorite
		vid.Title = videosInfo.Title
		first, _ := q.WithContext(context.Background()).TUser.Where(Tuser.ID.Eq(videosInfo.AuthorID)).First()
		usr.Id = first.ID
		usr.Name = first.Name
		usr.FollowCount = first.FollowerCount
		usr.FollowerCount = first.FollowerCount
		vid.Author = &usr
		videos = append(videos, &vid)
	}

	resp = &videoservice.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "成功",
		VideoList:  videos,
	}
	return
}

// CommentAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentAction(ctx context.Context, req *videoservice.DouyinCommentActionRequest) (resp *videoservice.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentList(ctx context.Context, req *videoservice.DouyinCommentListRequest) (resp *videoservice.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
