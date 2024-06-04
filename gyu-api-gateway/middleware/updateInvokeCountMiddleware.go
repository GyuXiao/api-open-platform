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

func UpdateInvokeInterfaceCountMiddleware(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		invokeUser := global.InvokeUserResp
		itfId := c.Request.Header.Get("itfId")
		interfaceInfoId, _ := strconv.Atoi(itfId)

		conn := global.BuildRpcConn(conf)
		interfaceInfoClient := pb.NewInterfaceInfoClient(conn.Conn())
		resp, err := interfaceInfoClient.InvokeInterfaceCount(c.Request.Context(), &pb.UpdateInvokeInterfaceCountReq{
			InterfaceInfoId: uint64(interfaceInfoId),
			UserId:          invokeUser.Id,
		})
		if err != nil || !resp.IsUpdated {
			global.HandlerServerInternalError(c)
			c.Abort()
			return
		}

		updateCountLog := logx.ContextWithFields(c.Request.Context(), logx.Field("updateSuccess", true))
		logc.Info(updateCountLog, "gyu-api-gateway 更新用户接口调用次数成功")

		c.Next()
	}
}
