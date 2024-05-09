package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-gateway/global"
	"io"
	"net/http"
)

func TargetUrlMiddleware(c *gin.Context) {

	targetURL := "http://127.0.0.1:8090/api/user"
	queryRaw := c.Request.URL.RawQuery
	if queryRaw != "" {
		targetURL += "?" + queryRaw
	}
	// 创建转发请求
	request, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		logc.Errorf(c.Request.Context(), "创建请求错误: %v", err)
		return
	}
	// 发起转发请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logc.Errorf(c.Request.Context(), "转发请求错误: %v", err)
		return
	}
	defer response.Body.Close()

	// 读取转发请求的响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logc.Errorf(c.Request.Context(), "读取响应内容错误: %v", err)
		return
	}

	if response.StatusCode == http.StatusOK {
		c.Writer.Write(body)
		c.Next()
	} else {
		global.HandlerInvokeError(c)
		return
	}
}
