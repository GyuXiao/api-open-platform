package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// 本地自测之前，先改正下面的 username 和 password

func Test_defaultUserInterfaceInfoModel_UpdateForInvokeSuccess(t *testing.T) {
	dataSource := "username:password@tcp(127.0.0.1:3306)/api_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	m := NewDefaultUserInterfaceInfoModel(db)
	type args struct {
		userId          uint64
		interfaceInfoId uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				userId:          1,
				interfaceInfoId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := m.UpdateForInvokeSuccess(tt.args.userId, tt.args.interfaceInfoId); (err != nil) != tt.wantErr {
			t.Errorf("%q. defaultUserInterfaceInfoModel.UpdateForInvokeSuccess() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
