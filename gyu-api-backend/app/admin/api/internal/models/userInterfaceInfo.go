package models

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"sync"
)

type UserInterfaceInfoService interface {
	CreateUserInterfaceInfo(map[string]interface{}) error
	SearchUserInterfaceByUserId(uint64) (*UserInterfaceInfoModel, error)
	SearchUserInterfaceByInterfaceId(uint64) (*UserInterfaceInfoModel, error)
	UpdateForInvokeSuccess(uint64, uint64) error
}

var userInterfaceInfoService UserInterfaceInfoService
var userInterfaceInfoOnce sync.Once

type defaultUserInterfaceInfoModel struct {
	*gorm.DB
}

func NewDefaultUserInterfaceInfoModel(db *gorm.DB) UserInterfaceInfoService {
	userInterfaceInfoOnce.Do(func() {
		userInterfaceInfoService = &defaultUserInterfaceInfoModel{db}
	})
	return userInterfaceInfoService
}

func (m *defaultUserInterfaceInfoModel) CreateUserInterfaceInfo(userInterfaceInfoMap map[string]interface{}) error {
	err := m.Table(constant.UserInterfaceInfoTableName).Model(&UserInterfaceInfoModel{}).Create(userInterfaceInfoMap).Error
	if err != nil {
		logc.Infof(ctx, "mysql create userInterfaceInfo err: %v", err)
		return xerr.NewErrCode(xerr.CreateUserInterfaceInfoError)
	}
	return nil
}

func (m *defaultUserInterfaceInfoModel) SearchUserInterfaceByUserId(userId uint64) (*UserInterfaceInfoModel, error) {
	userInterfaceInfo := UserInterfaceInfoModel{}
	err := m.Table(constant.UserInterfaceInfoTableName).Where("userId = ? and isDelete = 0", userId).Take(&userInterfaceInfo).Error
	switch {
	case err == nil:
		return &userInterfaceInfo, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Info(ctx, "mysql search userInterfaceInfo by userId not found")
		return nil, xerr.NewErrCode(xerr.RecordNotFoundError)
	default:
		logc.Infof(ctx, "mysql search userInterfaceInfo by userId err: %v", err)
		return nil, xerr.NewErrCode(xerr.SearchUserInterfaceInfoError)
	}
}

func (m *defaultUserInterfaceInfoModel) SearchUserInterfaceByInterfaceId(interfaceInfoId uint64) (*UserInterfaceInfoModel, error) {
	userInterfaceInfo := UserInterfaceInfoModel{}
	err := m.Table(constant.UserInterfaceInfoTableName).Where("interfaceInfoId = ? and isDelete = 0", interfaceInfoId).Take(&userInterfaceInfo).Error
	switch {
	case err == nil:
		return &userInterfaceInfo, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Info(ctx, "mysql search userInterfaceInfo by interfaceInfoId not found")
		return nil, xerr.NewErrCode(xerr.RecordNotFoundError)
	default:
		logc.Infof(ctx, "mysql search userInterfaceInfo by interfaceInfoId err: %v", err)
		return nil, xerr.NewErrCode(xerr.SearchUserInterfaceInfoError)
	}
}

func (m *defaultUserInterfaceInfoModel) UpdateForInvokeSuccess(userId uint64, interfaceInfoId uint64) error {
	// 这里还应该再优化，比如使用分布式事务来解决并发问题
	err := m.Table(constant.UserInterfaceInfoTableName).
		Model(&UserInterfaceInfoModel{}).
		Where("userId = ? and interfaceInfoId = ? and isDelete = 0 and leftNum > 0", userId, interfaceInfoId).
		Update("totalNum", gorm.Expr("totalNum + ?", 1)).
		Update("leftNum", gorm.Expr("leftNum - ?", 1)).Error
	if err != nil {
		return xerr.NewErrCode(xerr.InvokeSuccessUpdateError)
	}
	return nil
}
