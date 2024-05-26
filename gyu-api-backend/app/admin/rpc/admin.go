package main

import (
	"flag"
	"fmt"
	"gyu-api-backend/app/admin/rpc/internal/config"
	interfaceServer "gyu-api-backend/app/admin/rpc/internal/server/interfaceinfo"
	userServer "gyu-api-backend/app/admin/rpc/internal/server/user"
	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	uServer := userServer.NewUserServer(ctx)
	iServer := interfaceServer.NewInterfaceInfoServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, uServer)
		pb.RegisterInterfaceInfoServer(grpcServer, iServer)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
