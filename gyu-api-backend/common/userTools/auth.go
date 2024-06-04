package userTools

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"strconv"
)

// 校验用户角色是否为管理员

func CheckUserIsAdminRole(client *redis.Redis, token string) error {
	tokenLogic := models.NewDefaultTokenModel(client)
	result, err := tokenLogic.CheckTokenExist(token)
	if err != nil {
		return err
	}

	userRole, _ := strconv.Atoi(result[1])
	if userRole != constant.AdminRole {
		return xerr.NewErrCode(xerr.PermissionDenied)
	}

	return nil
}
