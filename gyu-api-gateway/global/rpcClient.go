package global

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"gyu-api-gateway/config"
	"gyu-api-gateway/rpc/pb"
	"sync"
)

var (
	rpcClientConn zrpc.Client
	once          sync.Once

	InvokeUserResp    *pb.GetInvokeUserResp
	InterfaceInfoResp *pb.GetInterfaceInfoResp
)

func BuildRpcConn(c config.Config) zrpc.Client {

	once.Do(func() {
		rpcClientConn = zrpc.MustNewClient(zrpc.RpcClientConf{
			Target: "dns:///" + c.RpcConfig.Target,
			Etcd: discov.EtcdConf{
				Hosts: []string{c.Etcd.Host},
				Key:   c.Etcd.Key,
			},
		})
	})

	return rpcClientConn
}
