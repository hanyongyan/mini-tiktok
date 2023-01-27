package main

import (
	"context"
	"errors"
	videoservice "mini_tiktok/kitex_gen/videoservice"
	"mini_tiktok/pkg/cache"
	"mini_tiktok/pkg/consts"
	jwtutil "mini_tiktok/pkg/utils"
	"strconv"
	"time"
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

	// 将点赞存入redis
	cache.RedisCache.Set(ctx, consts.FavoriteActionPrefix+strconv.FormatInt(claims.UserId, 10), req.VideoId, time.Second*60*30)
	return
}

// FavoriteList 2023-1-27 @Auth by 李卓轩 version 1.0
// 喜欢列表
// FavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteList(ctx context.Context, req *videoservice.DouyinFavoriteListRequest) (resp *videoservice.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
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
