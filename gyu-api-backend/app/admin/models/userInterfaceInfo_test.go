package models

import (
	"reflect"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

// 本地自测之前，先改正下面的 username 和 password

func Test_defaultUserInterfaceInfoModel_SearchUserInterfaceByUserIdAndInterfaceId(t *testing.T) {
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
		want    *UserInterfaceInfoModel
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				userId:          1,
				interfaceInfoId: 1,
			},
			want: &UserInterfaceInfoModel{
				Id:              1,
				UserId:          1,
				InterfaceInfoId: 1,
				TotalNum:        1,
				LeftNum:         4,
				Status:          0,
				CreateTime:      "2024-04-25T19:54:21+08:00",
				UpdateTime:      "2024-04-25T20:00:51+08:00",
				IsDelete:        0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := m.SearchUserInterfaceByUserIdAndInterfaceId(tt.args.userId, tt.args.interfaceInfoId)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. defaultUserInterfaceInfoModel.SearchUserInterfaceByUserIdAndInterfaceId() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. defaultUserInterfaceInfoModel.SearchUserInterfaceByUserIdAndInterfaceId() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// 本地自测之前，先改正下面的 username 和 password

func Test_defaultUserInterfaceInfoModel_GetTopInvokeInterfaceInfoList(t *testing.T) {
	dataSource := "username:password@tcp(127.0.0.1:3306)/api_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	m := NewDefaultUserInterfaceInfoModel(db)
	type args struct {
		limit uint64
	}
	tests := []struct {
		name    string
		args    args
		want    []*UserInterfaceInfoTopResultModel
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				limit: 2,
			},
			want: []*UserInterfaceInfoTopResultModel{
				{
					InterfaceInfoId: 28,
					TotalNum:        42,
				},
				{
					InterfaceInfoId: 3,
					TotalNum:        39,
				},
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				limit: 3,
			},
			want: []*UserInterfaceInfoTopResultModel{
				{
					InterfaceInfoId: 28,
					TotalNum:        42,
				},
				{
					InterfaceInfoId: 3,
					TotalNum:        39,
				},
				{
					InterfaceInfoId: 4,
					TotalNum:        32,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := m.GetTopInvokeInterfaceInfoList(tt.args.limit)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. defaultUserInterfaceInfoModel.GetTopInvokeInterfaceInfoList() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. defaultUserInterfaceInfoModel.GetTopInvokeInterfaceInfoList() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
