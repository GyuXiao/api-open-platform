package models

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"sync"
)

type UserService interface {
	CreateUser(map[string]interface{}) error
	SearchUserByUsername(string) (*UserModel, error)
}

var userService UserService
var userOnce sync.Once

type defaultUserModel struct {
	*gorm.DB
}

func NewDefaultUserModel(db *gorm.DB) UserService {
	userOnce.Do(func() {
		userService = &defaultUserModel{db}
	})
	return userService
}

func (m *defaultUserModel) CreateUser(userMap map[string]interface{}) error {
	err := m.Table(constant.UserTableName).Model(&UserModel{}).Create(userMap).Error
	if err != nil {
		logc.Infof(ctx, "mysql create user err: %v", err)
		return xerr.NewErrCode(xerr.CreateUserError)
	}
	return nil
}

func (m *defaultUserModel) SearchUserByUsername(username string) (*UserModel, error) {
	user := UserModel{}
	err := m.Table(constant.UserTableName).Where("username = ? and isDelete = 0", username).Take(&user).Error
	switch {
	case err == nil:
		return &user, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Info(ctx, "mysql search user by username not found")
		return nil, xerr.NewErrCode(xerr.RecordNotFoundError)
	default:
		logc.Infof(ctx, "mysql search user by username err: %v", err)
		return nil, xerr.NewErrCode(xerr.SearchUserError)
	}
}
