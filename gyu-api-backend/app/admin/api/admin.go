package main

import (
	"flag"
	"fmt"
	"gyu-api-backend/app/admin/api/internal/config"
	"gyu-api-backend/app/admin/api/internal/handler"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/common/constant"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Set(constant.AllowOrigin, constant.AllOrigins)
		w.Header().Set(constant.AllowHeaders, constant.AllOrigins)
		w.Header().Set(constant.AllowMethods, constant.Methods)
		w.Header().Set(constant.AllowExposeHeaders, constant.Headers)
		w.Header().Set(constant.AllowCredentials, constant.True)
	}, constant.AllOrigins))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
