package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerInvokeError(c *gin.Context) {
	// http.StatusBadRequest
	c.JSON(http.StatusOK, gin.H{"result": -1, "msg": "第三方API接口调用失败,请联系管理员!"})
	c.Abort()
}
