package main

import (
	"context"
	"fmt"
	"mini_tiktok/cmd/video/mw/cos"
	"mini_tiktok/pkg/nacos"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/registry"
	"mini_tiktok/cmd/video/rpc"
	"mini_tiktok/kitex_gen/videoservice/videoservice"
	"mini_tiktok/pkg/cache"
	"mini_tiktok/pkg/configs/config"
	"mini_tiktok/pkg/consts"
	"mini_tiktok/pkg/dal"
	"mini_tiktok/pkg/mw"
)

func Init() {
	// 配置的初始化要放在最前面
	config.Init()
	nacos.Init()
	cache.Init()
	rpc.Init()
	dal.Init()
	cos.Init()
	//task.Init()
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.VideoServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	Init()

	addr, err := net.ResolveTCPAddr(consts.TCP, fmt.Sprintf("127.0.0.1%v", consts.VideoServiceAddr))
	if err != nil {
		panic(err)
	}
	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		//server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithRegistry(registry.NewNacosRegistry(nacos.NacosClient)),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
