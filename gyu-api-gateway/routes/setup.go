package routes

import (
	"github.com/gin-gonic/gin"
	"gyu-api-gateway/middleware"
)

func Setup(r *gin.Engine) {

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

			// todo 正式发起请求之前，还要做的事情
			// 从数据库中查询模拟接口是否存在，以及请求方法是否匹配（还可以校验请求参数）gRPC
			// 5,判断请求的接口是否存在

			// 6,请求转发，调用模拟接口
			middleware.TargetUrlMiddleware(),

			// todo：调用成功后的业务逻辑，接口调用次数 +1
			// 8,调用成功，接口次数 +1; 调用失败，返回业务码

		)
	}
	// 在管理者的校验
	//ManagerRouter(r)
}
