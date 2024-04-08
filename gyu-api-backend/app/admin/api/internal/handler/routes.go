// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	interfaceInfo "gyu-api-backend/app/admin/api/internal/handler/interfaceInfo"
	user "gyu-api-backend/app/admin/api/internal/handler/user"
	"gyu-api-backend/app/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/gyu_api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/current",
				Handler: user.CurrentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/gyu_api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/logout",
				Handler: user.LogoutHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/gyu_api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/interfaceInfo/add",
				Handler: interfaceInfo.AddInterfaceInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/interfaceInfo/update",
				Handler: interfaceInfo.UpdateInterfaceInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/interfaceInfo/delete",
				Handler: interfaceInfo.DeleteInterfaceInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/interfaceInfo/list/page",
				Handler: interfaceInfo.GetPageListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/gyu_api/v1"),
	)
}
