// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"mini_tiktok/cmd/api/biz/rpc"
	"mini_tiktok/pkg/configs/config"
	"mini_tiktok/pkg/consts"
)

func Init() {
	config.Init()
	rpc.Init()
	hlog.SetLogger(hertzlogrus.NewLogger())
	hlog.SetLevel(hlog.LevelInfo)
}

func main() {
	Init()

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(consts.NacosAddr, consts.NacosPort),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
	}
	nacoscli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	r := nacos.NewNacosRegistry(nacoscli)

	tracer, cfg := tracing.NewServerTracer()
	addr := "0.0.0.0:8080"
	h := server.New(
		server.WithHostPorts(addr),
		server.WithHandleMethodNotAllowed(true),
		server.WithRegistry(r, &registry.Info{
			ServiceName: consts.ApiServiceName,
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
		tracer,
	)
	pprof.Register(h)
	h.Use(tracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
