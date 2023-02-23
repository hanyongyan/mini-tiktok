package utils

import (
	"context"
	"fmt"
	"mini_tiktok/pkg/dal/query"
)

// LikeNumAdd 将点赞数加一
func LikeNumAdd(videoID int64) error {
	q := query.Q
	favorite := q.TVideo
	video, err := q.WithContext(context.Background()).TVideo.Where(favorite.ID.Eq(videoID)).First()
	if err != nil {
		return fmt.Errorf("数据库查询出错")
	}
	_, err = q.WithContext(context.Background()).TVideo.Where(favorite.ID.Eq(videoID)).Update(favorite.FavoriteCount, video.FavoriteCount+1)
	if err != nil {
		return err
	}
	return nil
}

// LikeNumDel 将点赞数减一
func LikeNumDel(videoID int64) error {
	q := query.Q
	favorite := q.TVideo
	video, err := q.WithContext(context.Background()).TVideo.Where(favorite.ID.Eq(videoID)).First()
	q.WithContext(context.Background()).TVideo.Where(favorite.ID.Eq(videoID)).Update(favorite.FavoriteCount, video.FavoriteCount-1)
	if err != nil {
		return fmt.Errorf("数据库查询出错")
	}
	return nil
}
