package utils

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"mini_tiktok/pkg/cache"
	"mini_tiktok/pkg/consts"
	"mini_tiktok/pkg/dal/query"
)

// LikeNumAdd 将点赞数加一
func LikeNumAdd(videoID int64) error {
	redis := cache.RedisCache.RedisClient
	q := query.Q
	favorite := q.TVideo
	// 查询对应的点赞数
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
	redis := cache.RedisCache.RedisClient
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

// GetRedisVideoID 通过用户id获取redis中点赞过的视频id
func GetRedisVideoID(userID string) (vids []int64, err error) {
	redis := cache.RedisCache.RedisClient
	set, err := redis.Keys(context.Background(), "post_set:*").Result()
	arr := make([]int64, 10)
	for _, vid := range set {
		result, err1 := redis.SIsMember(context.Background(), vid, userID).Result()
		if err1 != nil {
			err1 = fmt.Errorf("redis error:" + err.Error())
			return nil, err1
		}

		if result {
			split := strings.Split(vid, consts.FavoriteActionPrefix)
			atoi, _ := strconv.Atoi(split[len(split)-1])
			arr = append(arr, int64(atoi))
		}
	}

	return arr, nil
}
