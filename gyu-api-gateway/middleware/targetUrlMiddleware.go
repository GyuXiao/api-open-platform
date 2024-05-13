package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"gyu-api-gateway/global"
	"gyu-api-gateway/internal/response"
	"gyu-api-gateway/types"
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
		// 构造请求
		var req types.Request
		err := c.BindJSON(&req)
		if err != nil {
			logc.Errorf(c.Request.Context(), "获取请求体错误: %v", err)
			return
		}
		// 发起请求
		// todo: 这里的 targetUrl 和 req 都应该是灵活的，不能写死
		resp, err := httpc.Do(c.Request.Context(), c.Request.Method, targetURL, req)
		if err != nil {
			logc.Errorf(c.Request.Context(), "请求转发错误: %v", err)
			global.HandlerInvokeError(c)
			return
		}

		defer resp.Body.Close()
		// 读取转发请求的响应内容
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			logc.Errorf(c.Request.Context(), "读取响应内容错误: %v", err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			global.HandlerInvokeError(c)
			return
		}

		var baseResponse types.BaseResponse
		err = json.Unmarshal(respBody, &baseResponse)
		if err != nil {
			logc.Errorf(c.Request.Context(), "解析响应内容错误: %v", err)
			return
		}

		// 统一的返回处理
		response.HandleResponse(c, baseResponse.Data, nil)
		responseLog := logx.ContextWithFields(c.Request.Context(), logx.Field("response_pass", true))
		logx.Info(responseLog, "路由转发成功")
		c.Next()
	}
}
