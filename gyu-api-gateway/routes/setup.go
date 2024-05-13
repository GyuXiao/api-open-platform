package routes

import (
	"github.com/gin-gonic/gin"
	"gyu-api-gateway/middleware"
)

func Setup(r *gin.Engine) {

	// 5,判断请求的接口是否存在
	// 8,调用成功，接口次数 +1; 调用失败，返回业务码

	// 2,请求日志
	r.Use(middleware.LogFormatMiddleware())
	// 3,黑白名单
	r.Use(middleware.FilterWithAccessControlInWhiteIP())

	apiGroup := r.Group("/api/invoke")
	{
		apiGroup.Any("/*path",
			// 一系列的中间件操作，比如 统一鉴权，接口验证，路由转发，调用次数统计 等等
			// 4,用户鉴权（ak sk）
			middleware.FilterWithAuth(),

			// 6,请求转发，调用模拟接口
			middleware.TargetUrlMiddleware(),
		)
	}
	// 在管理者的校验
	//ManagerRouter(r)
}
