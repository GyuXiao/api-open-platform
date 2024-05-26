package svc

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gyu-api-backend/app/admin/api/internal/config"
	"gyu-api-backend/app/admin/rpc/client/user"
)

type ServiceContext struct {
	Config      config.Config
	DBEngin     *gorm.DB
	RedisClient *redis.Redis
	UserRpc     user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.MySQL.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, error=" + err.Error())
	}
	logc.Info(context.Background(), "connect MySQL database success")
	return &ServiceContext{
		Config:  c,
		DBEngin: db,
		RedisClient: redis.MustNewRedis(redis.RedisConf{
			Host: c.Redis.Host,
			Type: c.Redis.Type,
			Pass: c.Redis.Pass,
		}),
		UserRpc: user.NewUserZrpcClient(zrpc.MustNewClient(c.AdminRpcConf)),
	}
}
