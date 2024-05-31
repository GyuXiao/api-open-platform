package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-gateway/config"
	"gyu-api-gateway/global"
	"gyu-api-gateway/rpc/pb"
	"strconv"
)

func InitDataMiddleware(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := global.BuildRpcConn(conf)

		// 通过 accessKey 拿到 user 对象
		accessKey := c.Request.Header.Get("accessKey")
		userClient := pb.NewUserClient(conn.Conn())
		var err error
		global.InvokeUserResp, err = userClient.GetInvokeUser(c.Request.Context(), &pb.GetInvokeUserReq{
			AccessKey: accessKey,
		})
		if err != nil || global.InvokeUserResp == nil {
			global.HandlerServerInternalError(c)
			c.Abort()
			return
		}

		// 通过 interfaceId 拿到 interface 对象
		interfaceInfoId := c.Request.Header.Get("itfId")
		if interfaceInfoId == "" {
			logc.Errorf(c.Request.Context(), "interfaceInfoId is empty")
			global.HandlerServerInternalError(c)
			return
		}
		id, _ := strconv.Atoi(interfaceInfoId)
		interfaceClient := pb.NewInterfaceInfoClient(conn.Conn())
		global.InterfaceInfoResp, err = interfaceClient.GetInterfaceInfo(c.Request.Context(), &pb.GetInterfaceInfoReq{
			Id: uint64(id),
		})
		if err != nil || global.InterfaceInfoResp == nil {
			logc.Errorf(c.Request.Context(), "get interface info error: %v", err)
			global.HandlerServerInternalError(c)
			return
		}

		initDataLog := logx.ContextWithFields(c.Request.Context(), logx.Field("initSuccess", true))
		logc.Info(initDataLog, "gyu-api-gateway 获取用户数据和接口数据成功")

		c.Next()
	}
}
