package cos

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"mini_tiktok/cmd/video/global"
	"mini_tiktok/pkg/configs/config"
	"net/http"
	"net/url"
	"time"
)

var (
)

func Init() {
	//将<bucket>和<region>修改为真实的信息
	//bucket的命名规则为{name}-{appid} ，此处填写的存储桶名称必须为此格式
	cfg := config.GlobalConfigs.CosConfig
	u, _ := url.Parse(cfg.Url)
	b := &cos.BaseURL{BucketURL: u}
	global.CosClient = cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  cfg.SecretId,
			SecretKey: cfg.SecretKey,
		},
	})
}
