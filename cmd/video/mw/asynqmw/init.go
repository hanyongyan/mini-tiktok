package asynqmw

import (
	"fmt"
	"github.com/hibiken/asynq"
	"log"
	"mini_tiktok/cmd/video/global"
	"mini_tiktok/pkg/configs/config"
)

func Init() {
	conf := config.GlobalConfigs.RedisConfig
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	global.AsynqClient = asynq.NewClient(asynq.RedisClientOpt{Addr: addr})
	go startupAsynqServer(addr)
}

func startupAsynqServer(addr string) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: addr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeFfmpegTask, HandleFfmpegTask)
	// ...register other handlers...
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
