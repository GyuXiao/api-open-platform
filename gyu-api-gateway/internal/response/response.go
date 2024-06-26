package response

import (
	"github.com/gin-gonic/gin"
	"gyu-api-gateway/internal/translator"
)

// BaseResponse 统一返回
type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    any    `json:"data"`
}

// HandleResponse 统一返回处理
func HandleResponse(c *gin.Context, data any, err error) {
	if err != nil {
		c.JSON(200, BaseResponse{
			Code:    500,
			Data:    nil,
			Message: translator.Translate(err),
		})
		return
	}

	c.JSON(200, BaseResponse{
		Code:    200,
		Data:    data,
		Message: "成功",
	})
}

// HandleAbortResponse 统一 Abort 返回处理
func HandleAbortResponse(c *gin.Context, err string) {
	c.AbortWithStatusJSON(200, BaseResponse{
		Code:    500,
		Data:    nil,
		Message: err,
	})
}
