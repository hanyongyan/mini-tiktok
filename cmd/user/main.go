package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"mini_tiktok/cmd/user/rpc"
	"mini_tiktok/kitex_gen/userService/userservice"
	"mini_tiktok/pkg/configs/config"
	"mini_tiktok/pkg/consts"
	"mini_tiktok/pkg/dal"
	"net"
)

func Init() {
	config.Init()
	rpc.Init()
	dal.Init()
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)
}

func main() {

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.UserServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

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
		Username:            "nacos",
		Password:            "nacos",
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, fmt.Sprintf("127.0.0.1%v", consts.UserServiceAddr))
	if err != nil {
		panic(err)
	}
	Init()
	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		//server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		//server.WithMiddleware(mw.CommonMiddleware),
		//server.WithMiddleware(mw.ServerMiddleware),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
