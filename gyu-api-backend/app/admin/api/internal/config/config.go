package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	// mysql 和 redis 的设置应该可以去掉
	MySQL struct {
		DataSource string
	}
	Redis struct {
		Host string
		Pass string
		Type string
	}
	AdminRpcConf zrpc.RpcClientConf
}
