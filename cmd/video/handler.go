package main

import (
	"context"
	"github.com/nanakura/go-ramda"
	videoservice "mini_tiktok/kitex_gen/videoservice"
	"mini_tiktok/pkg/dal/query"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videoservice.DouyinPublishActionRequest) (resp *videoservice.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

type queryVideoListRes struct {
	ID            int64 // 视频id
	AuthorID      int64 `sql:"author_id"`
	Name          string
	FollowCount   int64
	FollowerCount int64
	Password      string
	PlayURL       string    // 视频链接
	CoverURL      string    // 视频封面链接
	FavoriteCount int64     // 点赞数
	CommentCount  int64     // 评论数
	IsFavorite    bool      // 是否已点赞(0为未点赞, 1为已点赞)
	Title         string    // 视频标题
	CreateDate    time.Time // 视频上传时间
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *videoservice.DouyinFeedRequest) (resp *videoservice.DouyinFeedResponse, err error) {
	latestTime := req.LatestTime
	// 值为0（默认值）则说明不限制最新时间
	tv := query.Q.TVideo.As("v")
	tu := query.Q.TUser.As("u")
	var resList []queryVideoListRes
	if latestTime == 0 {
		err = tv.WithContext(context.Background()).
			Select(
				tv.ID,
				tv.AuthorID,
				tu.ALL,
				tv.PlayURL, tv.CoverURL, tv.FavoriteCount,
				tv.CommentCount, tv.IsFavorite, tv.Title, tv.CreateDate,
			).
			LeftJoin(tu, tu.ID.EqCol(tv.AuthorID)).
			Order(tv.CreateDate.Desc()).
			Limit(10).Scan(&resList)

		if err != nil {
			return
		}
	} else {
		t := time.Unix(latestTime/1000, 0)
		err = tv.WithContext(context.Background()).
			Select(
				tv.ID,
				tv.AuthorID,
				tu.Name,
				tu.Password,
				tu.FollowCount,
				tu.FollowerCount,
				tv.PlayURL, tv.CoverURL, tv.FavoriteCount,
				tv.CommentCount, tv.IsFavorite, tv.Title, tv.CreateDate,
			).
			LeftJoin(tu, tu.ID.EqCol(tv.AuthorID)).
			Where(tv.CreateDate.Lt(t)).
			Order(tv.CreateDate.Desc()).
			Limit(10).
			Scan(&resList)
		if err != nil {
			return
		}
	}
	if resList == nil {
		resList = []queryVideoListRes{}
	}
	resp = &videoservice.DouyinFeedResponse{
		StatusCode: 0,
		VideoList: ramda.Map(func(model queryVideoListRes) *videoservice.Video {
			return &videoservice.Video{
				Id: model.ID,
				Author: &videoservice.User{
					Id:            model.AuthorID,
					Name:          model.Name,
					FollowCount:   model.FollowCount,
					FollowerCount: model.FollowerCount,
				},
				PlayUrl:       model.PlayURL,
				CoverUrl:      model.CoverURL,
				FavoriteCount: model.FavoriteCount,
				CommentCount:  model.CommentCount,
				IsFavorite:    model.IsFavorite,
				Title:         model.Title,
			}
		})(resList),
	}
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *videoservice.DouyinPublishListRequest) (resp *videoservice.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteAction(ctx context.Context, req *videoservice.DouyinFavoriteActionRequest) (resp *videoservice.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

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
