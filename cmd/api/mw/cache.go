package mw

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	"github.com/hertz-contrib/cache"
	"github.com/hertz-contrib/cache/persist"
	"mini_tiktok/pkg/configs/config"
	"time"
)

var (
	RedisStore *persist.RedisStore
)

func init() {
	conf := config.GlobalConfigs.RedisConfig
	RedisStore = persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    fmt.Sprintf("%s:%d", conf.Host, conf.Port),
	}))
	hlog.Info("simple cache init")
}

func CacheMw() app.HandlerFunc {
	return cache.NewCacheByRequestPath(RedisStore, 10*time.Minute)
}
