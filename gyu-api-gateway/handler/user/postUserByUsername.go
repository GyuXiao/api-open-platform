package user

import (
	"github.com/gin-gonic/gin"
	"gyu-api-gateway/internal/response"
	"gyu-api-gateway/logic/user"
	"gyu-api-gateway/svc"
	"gyu-api-gateway/types"
)

func PostUserByUserNameHandle(c *gin.Context) {
	var req types.UserRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.HandleResponse(c, nil, err)
		return
	}
	resp, err := user.PostUserByUserName(svc.NewServiceContext(c), &req)
	response.HandleResponse(c, resp, err)
}
