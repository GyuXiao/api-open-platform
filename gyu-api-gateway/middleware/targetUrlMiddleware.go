package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-gateway/global"
	"gyu-api-gateway/internal/response"
	"io"
	"net/http"
)

func TargetUrlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL := "http://127.0.0.1:8090/api/user"
		queryRaw := c.Request.URL.RawQuery
		// 如果 query 不为空字符串，则将 query 添加到转发请求的 URL 中
		if queryRaw != "" {
			targetURL += "?" + queryRaw
		}
		// 创建转发请求
		req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
		if err != nil {
			logc.Errorf(c.Request.Context(), "创建请求错误: %v", err)
			return
		}
		// 发起转发请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logc.Errorf(c.Request.Context(), "转发请求错误: %v", err)
			return
		}

		defer resp.Body.Close()
		// 读取转发请求的响应内容
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logc.Errorf(c.Request.Context(), "读取响应内容错误: %v", err)
			return
		}

		if resp.StatusCode == http.StatusOK {
			// 统一返回响应
			response.HandleResponse(c, string(body), nil)
			// todo 在这里记录统一日志，合适嘛？
			logc.Field("response_pass", true)
			logc.Info(c.Request.Context(), "路由转发成功")
			c.Next()
		} else {
			global.HandlerInvokeError(c)
			return
		}
	}
}
