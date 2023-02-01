package utils

import (
	"context"
	"fmt"
	"mini_tiktok/pkg/cache"
	"mini_tiktok/pkg/dal/query"
	"strconv"
)

// LikeNumAdd 将点赞数加一
func LikeNumAdd(videoID int64) error {
	var redis = cache.RedisCache.RedisClient
	q := query.Q
	favorite := q.TVideo
	//查询对应的点赞数
	val, err2 := redis.HGet(context.Background(), "videos", strconv.FormatInt(videoID, 10)).Result()
	if err2 != nil {
		video, err := q.WithContext(context.Background()).TVideo.Where(favorite.ID.Eq(videoID)).First()
		if err != nil {
			return fmt.Errorf("数据库查询出错")
		}
		val = strconv.FormatInt(video.FavoriteCount, 10)
	}

	count, _ := strconv.Atoi(val)
	_, err := redis.HSet(context.Background(), "videos", strconv.FormatInt(videoID, 10), count+1).Result()
	if err != nil {
		return fmt.Errorf("redis 存储错误")
	}
	return nil
}

// LikeNumDel 将点赞数减一
func LikeNumDel(videoID int64) error {
	var redis = cache.RedisCache.RedisClient
	q := query.Q
	favorite := q.TVideo
	val, err2 := redis.HGet(context.Background(), "videos", strconv.FormatInt(videoID, 10)).Result()
	if err2 != nil {
		video, err := q.WithContext(context.Background()).TVideo.Where(favorite.ID.Eq(videoID)).First()
		if err != nil {
			return fmt.Errorf("数据库查询出错")
		}
		val = strconv.FormatInt(video.FavoriteCount, 10)
	}

	count, _ := strconv.Atoi(val)
	_, err := redis.HSet(context.Background(), "videos", strconv.FormatInt(videoID, 10), count-1).Result()
	if err != nil {
		return fmt.Errorf("redis 存储错误")
	}
	return nil
}
