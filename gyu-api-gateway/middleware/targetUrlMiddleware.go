package middleware

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-gateway/constant"
	"gyu-api-gateway/global"
	"gyu-api-gateway/internal/response"
	"gyu-api-gateway/types"
	"gyu-api-gateway/utils"
	"io"
	"net/http"
	"time"
)

func TargetUrlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(constant.RequestTimeout)*time.Second)
		defer cancel()

		// todo: 这里的 targetUrl 应该是灵活的，不能写死，而是从外面传进来
		targetURL := "http://127.0.0.1:8090/api/user"
		queryRaw := c.Request.URL.RawQuery
		// 如果 query 不为空字符串，则将 query 添加到转发请求的 URL 中
		if queryRaw != "" {
			targetURL += "?" + queryRaw
		}
		// 从原请求中构造 request body
		body := utils.GetRequestBody(c)
		// 构造请求
		req, err := utils.BuildRequest(c.Request.Context(), c.Request.Method, targetURL, body)
		if err != nil {
			logc.Errorf(c.Request.Context(), "构建请求错误: %v", err)
			global.HandlerInvokeError(c)
			return
		}
		req = req.WithContext(ctx)
		// todo：后续可能还需要设置一些请求头信息（流量染色），待添加
		// 。。。

		// 发起请求
		resp, err := utils.DoRequest(req)
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
			global.HandlerServerInternalError(c)
			return
		}
		if resp.StatusCode != http.StatusOK {
			logc.Error(c.Request.Context(), "请求得到的响应码状态并非成功")
			global.HandlerInvokeError(c)
			return
		}

		var baseResponse types.BaseResponse
		err = json.Unmarshal(respBody, &baseResponse)
		if err != nil {
			logc.Errorf(c.Request.Context(), "解析响应内容错误: %v", err)
			global.HandlerServerInternalError(c)
			return
		}

		// 统一的返回处理
		response.HandleResponse(c, baseResponse.Data, nil)
		responseLog := logx.ContextWithFields(c.Request.Context(), logx.Field("response_pass", true))
		logx.Info(responseLog, "路由转发成功")
		c.Next()
	}
}
