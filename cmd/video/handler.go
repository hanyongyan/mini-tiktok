package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/nanakura/go-ramda"
	"mini_tiktok/cmd/video/ftpUtil"
	videoservice "mini_tiktok/kitex_gen/videoservice"
	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
	"mini_tiktok/pkg/utils"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *videoservice.DouyinPublishActionRequest) (resp *videoservice.DouyinPublishActionResponse, err error) {
	data := bytes.NewBufferString(string(req.Data))
	uuidv4, _ := uuid.NewUUID()
	path := fmt.Sprintf("%s.mp4", uuidv4.String())
	tv := query.Q.TVideo
	cliams, _ := utils.CheckToken(req.Token)
	userId := cliams.UserId
	// TODO handle ftp read api
	playUrl := fmt.Sprintf("")
	err = tv.WithContext(context.Background()).
		Create(&model.TVideo{
			AuthorID:      userId,
			PlayURL:       playUrl,
			CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         req.Title,
			CreateDate:    time.Now(),
		})
	if err != nil {
		klog.Error("Error uploading file:", err)
		err = fmt.Errorf("视频保存失败：%w", err)
		return
	}
	if err = ftpUtil.FtpClient.Stor(path, data); err != nil {
		klog.Error("Error uploading file:", err)
		err = fmt.Errorf("视频保存失败：%w", err)
		return
	}

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

func CastQueryVideoListtoVideoServiceVideo(from []queryVideoListRes) []*videoservice.Video {
	return ramda.Map(func(model queryVideoListRes) *videoservice.Video {
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
	})(from)
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
		VideoList:  CastQueryVideoListtoVideoServiceVideo(resList),
	}
	return
}

// PublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishList(ctx context.Context, req *videoservice.DouyinPublishListRequest) (resp *videoservice.DouyinPublishListResponse, err error) {
	userId := req.UserId
	tv := query.Q.TVideo.As("tv")
	tu := query.Q.TUser
	tu2 := query.Q.TUser.As("tu2")
	var resList []queryVideoListRes
	qCtx := context.Background()
	err = tv.WithContext(qCtx).
		Select(
			tv.ID,
			tv.AuthorID,
			tu2.Name,
			tu2.Password,
			tu2.FollowCount,
			tu2.FollowerCount,
			tv.PlayURL, tv.CoverURL, tv.FavoriteCount,
			tv.CommentCount, tv.IsFavorite, tv.Title, tv.CreateDate,
		).
		LeftJoin(tu.WithContext(qCtx).Select(tu.ALL).Where(tu.ID.Eq(userId)).As("tu2"), tu2.ID.EqCol(tv.AuthorID)).
		Order(tv.CreateDate.Desc()).
		Limit(10).
		Scan(&resList)

	if resList == nil {
		resList = []queryVideoListRes{}
	}
	resp = &videoservice.DouyinPublishListResponse{
		StatusCode: 0,
		VideoList:  CastQueryVideoListtoVideoServiceVideo(resList),
	}
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
