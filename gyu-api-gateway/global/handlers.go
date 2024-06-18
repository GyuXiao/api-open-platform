package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerServerInternalError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": "服务器内部错误!"})
	c.Abort()
}

func HandlerInvokeError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "msg": "第三方API接口调用失败,请联系管理员!"})
	c.Abort()
}

// 禁止状态 403

func HandlerForbidden(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusForbidden, "msg": "禁止状态"})
	c.Abort()
}

// 未授权状态 401

func HandlerUnauthorized(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "msg": "鉴权未通过"})
	c.Abort()
}

func HandlerExceedLimit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusTooManyRequests, "msg": "rate limit exceed"})
	c.Abort()
}
