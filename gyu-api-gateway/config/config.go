package config

type Config struct {
	Server struct {
		Name string
		Host string
		Port int
	}
	RpcConfig struct {
		Target string
	}
	Etcd struct {
		Host string
		Key  string
	}
}
