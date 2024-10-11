package main

import (
	"flag"
	"fmt"

	"Anitale/apps/anime/rpc/internal/config"
	"Anitale/apps/anime/rpc/internal/server"
	"Anitale/apps/anime/rpc/internal/svc"
	"Anitale/apps/anime/rpc/pb"
	"Anitale/pkg/interceptor"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/anime.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAnimeServer(grpcServer, server.NewAnimeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 错误处理
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)
	defer s.Stop()


	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
