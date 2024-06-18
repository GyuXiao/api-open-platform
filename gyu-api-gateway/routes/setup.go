package routes

import (
	"github.com/gin-gonic/gin"
	"gyu-api-gateway/config"
	"gyu-api-gateway/middleware"
)

func Setup(r *gin.Engine, c config.Config) {

	// 请求日志
	r.Use(middleware.LogFormatMiddleware())

	// 黑白名单
	r.Use(middleware.FilterWithAccessControlInWhiteIP())

	// ip 限流
	r.Use(middleware.RateLimiterMiddleware(c))

	// 路由转发
	apiGroup := r.Group("/api/invoke")
	{
		apiGroup.Any("/*path",
			middleware.InitDataMiddleware(c),

			// 用户鉴权（ak sk）
			middleware.FilterWithAuth(),

			// 请求转发，调用模拟接口
			middleware.TargetUrlMiddleware(),

			// 更新调用次数
			middleware.UpdateInvokeInterfaceCountMiddleware(c),
		)
	}
	// 在管理者的校验
	//ManagerRouter(r)
}
