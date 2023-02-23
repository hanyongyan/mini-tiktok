package like

import (
	"context"
	"fmt"
	"mini_tiktok/pkg/dal/model"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"mini_tiktok/pkg/cache"
	"mini_tiktok/pkg/consts"
	"mini_tiktok/pkg/dal/query"
)

func Init() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)
	_, err := s.Every(1).Hours().Do(func() {
		go RedisLikeDateToMysql()
	})
	if err != nil {
		return
	}
	s.StartAsync()
}

// RedisLikeDateToMysql 定时将redis的点赞数据存入mysql
func RedisLikeDateToMysql() {
	fmt.Println("正在将redis数据存储进数据库")
	redis := cache.RedisCache.RedisClient
	//点赞数数据
	keys, _ := redis.Keys(context.Background(), consts.FavoriteCountPrefix+"*").Result()
	for _, v := range keys {
		split := strings.Split(v, consts.FavoriteCountPrefix)
		atoi, _ := strconv.Atoi(split[len(split)-1])
		result, _ := redis.Get(context.Background(), v).Result()
		i, _ := strconv.Atoi(result)
		redisVideoLikeNumToMysql(int64(atoi), int64(i))
	}
	//点赞数据
	result, _ := redis.Keys(context.Background(), consts.FavoriteActionPrefix+"*").Result()
	for _, v := range result {
		split := strings.Split(v, consts.FavoriteActionPrefix)
		atoi, _ := strconv.Atoi(split[len(split)-1])
		userIds, _ := redis.SMembers(context.Background(), v).Result()
		for _, v1 := range userIds {
			num, _ := strconv.Atoi(v1)
			b := selectVideoByUserIDAndVideoID(int64(atoi), int64(num))
			if b {
				update(int64(num), int64(atoi))
			} else {
				create(int64(num), int64(atoi))
			}
		}
	}
}

func update(userID, videoID int64) {
	q := query.Q
	tFavorite := q.TFavorite
	_, err := q.WithContext(context.Background()).TFavorite.Where(tFavorite.VideoID.Eq(videoID), tFavorite.UserID.Eq(userID)).Update(tFavorite.Status, true)
	if err != nil {
		return
	}
}

func create(userID, videoID int64) {
	q := query.Q
	newFav := &model.TFavorite{
		VideoID: videoID,
		UserID:  userID,
		Status:  true,
	}
	err := q.WithContext(context.Background()).TFavorite.Create(newFav)
	if err != nil {
		return
	}
}

func selectVideoByUserIDAndVideoID(videoID, userID int64) bool {
	q := query.Q
	tFavorite := q.TFavorite
	_, err := q.WithContext(context.Background()).TFavorite.Where(tFavorite.VideoID.Eq(videoID), tFavorite.UserID.Eq(userID)).First()
	if err != nil {
		return false
	}
	return true
}

// 将点赞数与videoid刷回数据库
func redisVideoLikeNumToMysql(videoId, count int64) {
	q := query.Q
	Tvideo := q.TVideo
	_, _ = q.WithContext(context.Background()).TVideo.Where(Tvideo.ID.Eq(videoId)).Update(Tvideo.FavoriteCount, count)
}
