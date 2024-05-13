package middleware

import (
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"gyu-api-gateway/global"
	"gyu-api-gateway/internal/response"
	"gyu-api-gateway/types"
	"gyu-api-gateway/utils"
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
		// 获取请求 body，格式是 map[string]interface{}
		body := utils.GetRequestBody(c)
		// 构造请求
		var req types.Request
		err := maputil.MapTo(body, &req)
		if err != nil {
			logc.Errorf(c.Request.Context(), "map 转换为 struct 错误: %v", err)
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

		response.HandleResponse(c, string(respBody), nil)
		responseLog := logx.ContextWithFields(c.Request.Context(), logx.Field("response_pass", true))
		logx.Info(responseLog, "路由转发成功")
		c.Next()
	}
}
