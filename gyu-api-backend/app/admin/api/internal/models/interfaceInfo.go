package models

import (
	"gorm.io/gorm"
	"gyu-api-backend/common/xerr"
	"sync"
)

type InterfaceInfoService interface {
	FindListPage(string, uint64, uint64) ([]*InterfaceInfoModel, error)
}

var interfaceInfoService InterfaceInfoService
var interfaceInfoOnce sync.Once

type defaultInterfaceInfoModel struct {
	*gorm.DB
}

func (m *defaultInterfaceInfoModel) FindListPage(keyword string, pageNumber uint64, pageSize uint64) ([]*InterfaceInfoModel, error) {
	var interfaceInfoList []*InterfaceInfoModel
	offset := (pageNumber - 1) * pageSize
	// 根据什么字段进行查找，还需要再思考一下
	// 且如果设置索引的话，就不要使用左模糊查询了
	result := m.Table("interfaceInfo").Where("name LIKE ? AND isDelete = 0", "%"+keyword+"%").Offset(int(offset)).Limit(int(pageSize)).Find(&interfaceInfoList)
	if result.Error != nil {
		return nil, xerr.NewErrCode(xerr.SearchInterfaceInfoPageListError)
	}
	if len(interfaceInfoList) == 0 {
		return nil, nil
	}
	return interfaceInfoList, nil
}

func NewDefaultInterfaceInfoModel(db *gorm.DB) InterfaceInfoService {
	interfaceInfoOnce.Do(func() {
		interfaceInfoService = &defaultInterfaceInfoModel{db}
	})
	return interfaceInfoService
}
