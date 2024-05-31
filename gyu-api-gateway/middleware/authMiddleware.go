package middleware

import (
	"github.com/GyuXiao/gyu-api-sdk/sdk"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-gateway/global"
	"strconv"
	"time"
)

func FilterWithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		headers := c.Request.Header
		accessKey := headers.Get("AccessKey")
		nonce := headers.Get("Nonce")
		timestamp := headers.Get("Timestamp")
		sign := headers.Get("Sign")
		body := headers.Get("Body")

		invokeUser := global.InvokeUserResp

		// accessKey 校验
		if accessKey != invokeUser.AccessKey {
			logx.Info(c.Request.Context(), "accessKey 错误, 权限校验未通过")
			global.HandlerUnauthorized(c)
			return
		}

		// 时间和当前时间不能超过 5 分钟
		nowTime := time.Now().Unix()
		thatTime, _ := strconv.ParseInt(timestamp, 10, 64)
		fiveMinutes := int64(5 * 60)
		if nowTime-thatTime > fiveMinutes {
			logx.Info(c.Request.Context(), "超时 5 分钟, 权限校验未通过")
			global.HandlerUnauthorized(c)
			return
		}

		paramsMap := map[string]string{
			"title0": invokeUser.AccessKey,
			"title1": invokeUser.SecretKey,
			"title2": nonce,
			"title3": timestamp,
			"title4": body,
		}
		signature := sdk.GenSign(paramsMap, invokeUser.SecretKey)
		// 签名校验
		if signature != sign {
			logx.Info(c.Request.Context(), "签名错误, 权限校验未通过")
			global.HandlerUnauthorized(c)
			return
		}

		signatureLog := logx.ContextWithFields(c.Request.Context(), logx.Field("isPass", true))
		logc.Info(signatureLog, "gyu-api-gateway 统一鉴权-API权限验证通过")

		c.Next()
	}
}
