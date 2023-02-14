package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"mini_tiktok/cmd/chat/middleware/mongodb"
	"mini_tiktok/cmd/chat/rpc"
	chatservice "mini_tiktok/kitex_gen/chatservice/messageservice"
	"mini_tiktok/pkg/configs/config"
	"mini_tiktok/pkg/consts"
	"mini_tiktok/pkg/mw"
	"mini_tiktok/pkg/nacos"
	"net"
)

func Init() {
	config.Init()
	nacos.Init()
	rpc.Init()
	mongodb.Init()
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)
}

func main() {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ChatServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	Init()

	addr, err := net.ResolveTCPAddr(consts.TCP, fmt.Sprintf("127.0.0.1%v", consts.ChatServiceAddr))
	if err != nil {
		panic(err)
	}
	svr := chatservice.NewServer(new(MessageServiceImpl),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ChatServiceName}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithRegistry(registry.NewNacosRegistry(nacos.NacosClient)),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
