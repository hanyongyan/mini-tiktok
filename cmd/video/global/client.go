package global

import (
	"github.com/hibiken/asynq"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var (
	AsynqClient *asynq.Client
	CosClient *cos.Client
)
