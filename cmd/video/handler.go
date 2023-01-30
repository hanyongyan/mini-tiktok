package main

import (
	"context"
	"errors"
	videoservice "mini_tiktok/kitex_gen/videoservice"
	"mini_tiktok/pkg/dal/model"
	"mini_tiktok/pkg/dal/query"
	jwtutil "mini_tiktok/pkg/utils"
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
	// 评论操作
	queryUser := query.Q.TUser
	queryVideo := query.Q.TVideo
	queryComment := query.Q.TComment
	timeLayoutStr := "2006-01-02 15:04:05"
	// 解析 token 拿取用户id
	claims, flag := jwtutil.CheckToken(req.Token)
	if !flag {
		return nil, errors.New("token is expired")
	}
	// 判断视频是否存在
	_, err = queryVideo.WithContext(ctx).Where(queryVideo.ID.Eq(req.VideoId)).First()
	if err != nil {
		return nil, errors.New("video does not exist")
	}

	// 发布评论
	if req.ActionType == 1 {
		comment := &model.TComment{
			UserID:     claims.UserId,
			Content:    req.CommentText,
			VideoID:    req.VideoId,
			CreateDate: time.Now(),
		}

		err := queryComment.WithContext(ctx).Create(comment)
		user, _ := queryUser.WithContext(ctx).Select(queryUser.ID, queryUser.Name).
			Where(queryUser.ID.Eq(claims.UserId)).First()
		if err != nil {
			return nil, errors.New("add failure")
		}
		resp = &videoservice.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "评论成功",
			Comment: &videoservice.Comment{
				Id: comment.ID,
				User: &videoservice.User{
					Id:            user.ID,
					Name:          user.Name,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
				},
				Content:    comment.Content,
				CreateDate: comment.CreateDate.Format(timeLayoutStr),
			},
		}
		// 删除评论
	} else if req.ActionType == 2 {
		// 用户是否有此条评论
		_, err := queryComment.WithContext(ctx).Where(queryComment.ID.Eq(req.CommentId)).
			Where(queryComment.UserID.Eq(claims.UserId)).Delete()
		if err != nil {
			return nil, errors.New("comment does not exist")
		}
		resp = &videoservice.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "删除成功",
		}
	} else {
		return nil, errors.New("operation error")
	}

	return
}

// CommentList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CommentList(ctx context.Context, req *videoservice.DouyinCommentListRequest) (resp *videoservice.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	// 获取评论
	queryUser := query.Q.TUser
	queryComment := query.Q.TComment
	// 数据库取出的 result
	type Result struct {
		Content    string
		CreateDate string
		ID         int64
		userID     int64
		Name       string
	}
	var result []Result
	// 解析 token
	_, flag := jwtutil.CheckToken(req.Token)
	// 登录后以查看评论
	if !flag {
		return nil, errors.New("log in to view the comments")
	}
	// 查询视频下的评论
	// 运用 left join 联合查询
	err = queryComment.WithContext(ctx).
		Select(queryComment.Content, queryComment.CreateDate, queryUser.ID, queryUser.Name).LeftJoin(&queryUser, queryUser.ID.EqCol(queryComment.UserID)).Where(queryComment.VideoID.Eq(req.VideoId)).Scan(&result)
	if err != nil {
		return
	}
	var comment videoservice.Comment
	var comments []*videoservice.Comment
	for _, com := range result {
		// 序列化
		user := videoservice.User{
			Id:   com.userID,
			Name: com.Name,
		}
		comment.User = &user
		comment.Content = com.Content
		comment.CreateDate = com.CreateDate
		// 这里要再创建一个干净的变量，要不然会只传最后一个
		var com videoservice.Comment
		com = comment
		comments = append(comments, &com)
	}

	resp = &videoservice.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "the request succeeded",
		CommentList: comments,
	}

	return
}
