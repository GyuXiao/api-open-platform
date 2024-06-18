package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/conf"
	"gyu-api-gateway/config"
	"gyu-api-gateway/middleware"
	"gyu-api-gateway/routes"
	"net/http"
)

func init() {
	middleware.IPLimiter = middleware.NewRateLimiter()
}

var configFile = flag.String("f", "etc/gateway-local.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	r := gin.Default()
	routes.Setup(r, c)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port),
		Handler: r,
	}
	server.ListenAndServe()
}
