package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-gateway/constant"
	"gyu-api-gateway/utils"
	"time"
)

// 网关统一的请求日志和响应日志记录

func LogFormatMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		// 请求唯一标识
		requestId, err := uuid.NewV7()
		if err != nil {
			logc.Error(c.Request.Context(), "生成请求唯一标识失败："+err.Error())
		}
		c.Set(constant.UniqueSessionID, requestId.String())

		// 请求来源域名
		domain, err := utils.GetDomainFromReferer(c.Request.Referer())
		if err != nil {
			logc.Error(c.Request.Context(), "获取请求来源域名失败："+err.Error())
		}
		// 请求来源 ip
		sourceIP := utils.GetRequestIP(c)
		requestLog := logx.ContextWithFields(
			c.Request.Context(),
			logx.Field("请求路径", c.Request.RequestURI),
			logx.Field("请求方法", c.Request.Method),
			logx.Field("请求参数", c.Request.URL.RawQuery),
			logx.Field("请求来源域名", domain),
			logx.Field("请求来源 ip", sourceIP),
		)
		logc.Info(requestLog, "gyu-api-gateway 请求日志")

		c.Next()

		endTime := time.Now()
		responseStatus := c.Writer.Status()
		responseLog := logx.ContextWithFields(
			c.Request.Context(),
			logx.Field("响应码", responseStatus),
			logx.Field("总耗时", endTime.Sub(startTime)),
		)
		logc.Info(responseLog, "gyu-api-gateway 响应日志")
	}
}
