package models

import "time"

type UserModel struct {
	Id         uint64    `gorm:"column:id;primaryKey" json:"id"`
	Username   string    `gorm:"column:username;index" json:"username"`
	Password   string    `gorm:"column:password" json:"password"`
	AvatarUrl  string    `gorm:"column:avatarUrl" json:"avatarUrl"`
	Email      string    `gorm:"column:email;index" json:"email"`
	Phone      string    `gorm:"column:phone;index" json:"phone"`
	UserRole   uint8     `gorm:"column:userRole" json:"userRole"`
	IsDelete   uint8     `gorm:"column:isDelete" json:"isDelete"`
	CreateTime time.Time `gorm:"column:createTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:updateTime" json:"updateTime"`
}

type InterfaceInfoModel struct {
	Id             uint64 `gorm:"column:id;primaryKey" json:"id"`
	Name           string `gorm:"column:name" json:"name"`
	Description    string `gorm:"column:description" json:"description"`
	Url            string `gorm:"column:url" json:"url"`
	RequestHeader  string `gorm:"column:requestHeader" json:"requestHeader"`
	ResponseHeader string `gorm:"column:responseHeader" json:"responseHeader"`
	Status         uint8  `gorm:"column:status" json:"status"`
	Method         string `gorm:"column:method" json:"method"`
	UserId         uint64 `gorm:"column:userId" json:"userId"`
	CreateTime     string `gorm:"column:createTime" json:"createTime"`
	UpdateTime     string `gorm:"column:updateTime" json:"updateTime"`
	IsDelete       uint8  `gorm:"column:isDelete" json:"isDelete"`
}
