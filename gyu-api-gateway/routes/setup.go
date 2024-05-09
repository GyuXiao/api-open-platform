package routes

import (
	"github.com/gin-gonic/gin"
	"gyu-api-gateway/middleware"
)

func Setup(r *gin.Engine) {
	apiGroup := r.Group("/api/invoke")
	{
		apiGroup.Any("/*path",
			// 一系列的中间件操作，比如 统一鉴权，接口验证，路由转发，调用次数统计 等等
			middleware.TargetUrlMiddleware,
		)
	}
	// 在管理者的校验
	//ManagerRouter(r)
}
