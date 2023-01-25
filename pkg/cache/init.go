package cache

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redis/redis/v8"
	"github.com/hertz-contrib/cache"
	"github.com/hertz-contrib/cache/persist"
	"github.com/spf13/viper"
	"mini_tiktok/pkg/consts"
	"time"
)

var (
	RedisCache      *persist.RedisStore
	cacheMiddleware app.HandlerFunc
)

func Init() {
	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: consts.TCP,
		Addr:    fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
	}))
	RedisCache = redisStore
	cacheMiddleware = cache.NewCacheByRequestURI(redisStore, 10*time.Minute)
}

func RedisMiddleware() app.HandlerFunc {
	return cacheMiddleware
}
