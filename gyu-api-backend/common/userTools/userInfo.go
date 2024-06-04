package userTools

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/constant"
)

// 从缓存中获取用户信息

func GetUserInfo(client *redis.Redis, token string) (map[string]string, error) {
	tokenLogic := models.NewDefaultTokenModel(client)
	result, err := tokenLogic.CheckTokenExist(token)
	if err != nil {
		return nil, err
	}

	mp := map[string]string{
		constant.KeyUserId:    result[0],
		constant.KeyUserRole:  result[1],
		constant.KeyUsername:  result[2],
		constant.KeyAvatarUrl: result[3],
	}
	return mp, nil
}
