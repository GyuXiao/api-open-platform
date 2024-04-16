package models

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"sync"
)

type InterfaceInfoService interface {
	AddInterfaceInfo(map[string]interface{}) error
	FindListPage(string, uint64, uint64) ([]*InterfaceInfoModel, int64, error)
	UpdateInterfaceInfo(uint64, map[string]interface{}) error
	UpdateInterfaceInfoStatus(uint8, uint64) error
	DeleteInterfaceInfo(uint64) error
	SearchInterfaceInfoByName(string) (*InterfaceInfoModel, error)
	SearchInterfaceInfoById(uint64) (*InterfaceInfoModel, error)
}

var interfaceInfoService InterfaceInfoService
var interfaceInfoOnce sync.Once

type defaultInterfaceInfoModel struct {
	*gorm.DB
}

func (m *defaultInterfaceInfoModel) SearchInterfaceInfoById(id uint64) (*InterfaceInfoModel, error) {
	interfaceInfo := InterfaceInfoModel{}
	err := m.Table(constant.InterfaceInfoTableName).Where("id = ? and isDelete = 0", id).Take(&interfaceInfo).Error
	switch {
	case err == nil:
		return &interfaceInfo, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Info(ctx, "mysql search interfaceInfo by name not found")
		return nil, xerr.NewErrCode(xerr.RecordNotFoundError)
	default:
		logc.Infof(ctx, "mysql search interfaceInfo by name err: ", err)
		return nil, xerr.NewErrCode(xerr.SearchInterfaceInfoError)
	}
}

func (m *defaultInterfaceInfoModel) DeleteInterfaceInfo(id uint64) error {
	_, err := m.SearchInterfaceInfoById(id)
	if err != nil {
		return err
	}
	err = m.Table(constant.InterfaceInfoTableName).Model(&InterfaceInfoModel{}).Where("isDelete = 0 and id = ?", id).Update("isDelete", 1).Error
	if err != nil {
		return xerr.NewErrCode(xerr.DeleteInterfaceInfoError)
	}
	return nil
}

func (m *defaultInterfaceInfoModel) UpdateInterfaceInfo(id uint64, interfaceInfoMap map[string]interface{}) error {
	interfaceInfo := InterfaceInfoModel{
		Name:           interfaceInfoMap["name"].(string),
		Description:    interfaceInfoMap["description"].(string),
		Url:            interfaceInfoMap["url"].(string),
		RequestHeader:  interfaceInfoMap["requestHeader"].(string),
		ResponseHeader: interfaceInfoMap["responseHeader"].(string),
		Method:         interfaceInfoMap["method"].(string),
	}
	err := m.Table(constant.InterfaceInfoTableName).Model(&InterfaceInfoModel{}).Where("id = ?", id).Updates(interfaceInfo).Error
	if err != nil {
		return xerr.NewErrCode(xerr.UpdateInterfaceInfoError)
	}
	return nil
}

func (m *defaultInterfaceInfoModel) UpdateInterfaceInfoStatus(status uint8, id uint64) error {
	err := m.Table(constant.InterfaceInfoTableName).Model(&InterfaceInfoModel{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return xerr.NewErrCode(xerr.UpdateInterfaceInfoStatusError)
	}
	return nil
}

func (m *defaultInterfaceInfoModel) SearchInterfaceInfoByName(name string) (*InterfaceInfoModel, error) {
	interfaceInfo := InterfaceInfoModel{}
	err := m.Table(constant.InterfaceInfoTableName).Where("name = ? and isDelete = 0", name).Take(&interfaceInfo).Error
	switch {
	case err == nil:
		return &interfaceInfo, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Info(ctx, "mysql search interfaceInfo by name not found")
		return nil, xerr.NewErrCode(xerr.RecordNotFoundError)
	default:
		logc.Infof(ctx, "mysql search interfaceInfo by name err: ", err)
		return nil, xerr.NewErrCode(xerr.SearchInterfaceInfoError)
	}
}

func (m *defaultInterfaceInfoModel) AddInterfaceInfo(interfaceInfoMap map[string]interface{}) error {
	err := m.Table(constant.InterfaceInfoTableName).Model(&InterfaceInfoModel{}).Create(interfaceInfoMap).Error
	if err != nil {
		return xerr.NewErrCode(xerr.AddInterfaceInfoError)
	}
	return nil
}

func (m *defaultInterfaceInfoModel) FindListPage(keyword string, pageNumber uint64, pageSize uint64) (interfaceInfoList []*InterfaceInfoModel, total int64, err error) {
	err = m.Table(constant.InterfaceInfoTableName).Where("isDelete = 0").Count(&total).Error
	if err != nil {
		return nil, 0, xerr.NewErrCode(xerr.RecordCountError)
	}
	// 根据什么字段进行查找，还需要再思考一下（暂时选择 name 字段）
	// 且如果设置索引的话，就不要使用左模糊查询了
	offset := (pageNumber - 1) * pageSize
	result := m.Table(constant.InterfaceInfoTableName).Where("name LIKE ? AND isDelete = 0", "%"+keyword+"%").Offset(int(offset)).Limit(int(pageSize)).Find(&interfaceInfoList)
	if result.Error != nil {
		return nil, 0, xerr.NewErrCode(xerr.SearchInterfaceInfoPageListError)
	}
	return interfaceInfoList, total, nil
}

func NewDefaultInterfaceInfoModel(db *gorm.DB) InterfaceInfoService {
	interfaceInfoOnce.Do(func() {
		interfaceInfoService = &defaultInterfaceInfoModel{db}
	})
	return interfaceInfoService
}
