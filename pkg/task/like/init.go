package like

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"mini_tiktok/pkg/cache"
	"mini_tiktok/pkg/consts"
	"mini_tiktok/pkg/dal/model"
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
	// todo ...
	fmt.Println("正在将redis数据存储进数据库")
	redis := cache.RedisCache.RedisClient

	result, err := redis.HGetAll(context.Background(), "videos").Result()
	if err != nil {
		err = fmt.Errorf("redis获取所有数据失败" + err.Error())
	}

	// 根据视频id更新点赞数量
	for k, v := range result {
		split := strings.Split(k, consts.FavoriteActionPrefix)
		vid, _ := strconv.Atoi(split[len(split)-1])
		count, _ := strconv.Atoi(v)
		redisVideoLikeNumToMysql(int64(vid), int64(count))
	}

	// 清空redis数据
	strs, err := redis.HKeys(context.Background(), "videos").Result()
	for _, k := range strs {
		redis.HDel(context.Background(), "videos", k)
	}

	// 将点赞数据批量存进数据库
	// 获取post_set中的所有键
	set, err := redis.Keys(context.Background(), "post_set:*").Result()
	for _, str := range set {
		userIds := redis.SMembers(context.Background(), "post_set:"+str).Val()
		split := strings.Split(str, consts.FavoriteActionPrefix)
		for _, userid := range userIds {
			vid, _ := strconv.Atoi(split[len(split)-1])
			uid, _ := strconv.Atoi(userid)
			redisDateToMysql(int64(vid), int64(uid))
		}

	}
}

// 将点赞数与videoid刷回数据库
func redisVideoLikeNumToMysql(videoId, count int64) {
	q := query.Q
	Tvideo := q.TVideo
	_, _ = q.WithContext(context.Background()).TVideo.Where(Tvideo.ID.Eq(videoId)).Update(Tvideo.FavoriteCount, count)
}

// 点赞数据传进数据库
func redisDateToMysql(videoId, userID int64) {
	q := query.Q
	newFav := &model.TFavorite{
		VideoID: videoId,
		UserID:  userID,
		Status:  true,
	}
	_ = q.WithContext(context.Background()).TFavorite.Create(newFav)
}
