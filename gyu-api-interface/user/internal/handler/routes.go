// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "gyu-api-interface/user/internal/handler/user"
	"gyu-api-interface/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user",
				Handler: user.GetUsernameByPostHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)
}
