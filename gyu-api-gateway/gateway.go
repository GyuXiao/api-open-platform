package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gyu-api-gateway/routes"
	"net/http"
)

var release string

func init() {
	flag.StringVar(&release, "release", "local", "release model, optional local/dev/prod")
}

func main() {
	flag.Parse()

	//configFile := fmt.Sprintf("etc/gateway-demo-%s.yaml", release)

	r := gin.Default()

	routes.Setup(r)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", "127.0.0.1", 8123),
		Handler: r,
	}

	server.ListenAndServe()
}
