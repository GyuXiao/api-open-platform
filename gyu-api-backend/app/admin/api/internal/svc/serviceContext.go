package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gyu-api-backend/app/admin/api/internal/config"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/app/admin/rpc/client/user"
)

type ServiceContext struct {
	Config           config.Config
	RedisClient      *redis.Redis
	UserRpc          user.UserZrpcClient
	InterfaceInfoRpc interfaceinfo.InterfaceInfoZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.MustNewRedis(redis.RedisConf{
			Host: c.Redis.Host,
			Type: c.Redis.Type,
			Pass: c.Redis.Pass,
		}),
		UserRpc:          user.NewUserZrpcClient(zrpc.MustNewClient(c.AdminRpcConf)),
		InterfaceInfoRpc: interfaceinfo.NewInterfaceInfoZrpcClient(zrpc.MustNewClient(c.AdminRpcConf)),
	}
}
