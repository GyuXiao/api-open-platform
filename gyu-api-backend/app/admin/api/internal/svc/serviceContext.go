package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gyu-api-backend/app/admin/api/internal/config"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/app/admin/rpc/client/user"
)

type ServiceContext struct {
	Config           config.Config
	UserRpc          user.UserZrpcClient
	InterfaceInfoRpc interfaceinfo.InterfaceInfoZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		UserRpc:          user.NewUserZrpcClient(zrpc.MustNewClient(c.AdminRpcConf)),
		InterfaceInfoRpc: interfaceinfo.NewInterfaceInfoZrpcClient(zrpc.MustNewClient(c.AdminRpcConf)),
	}
}
