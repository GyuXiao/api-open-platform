// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id         uint64 `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AvatarUrl  string `json:"avatar_url"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	UserRole   uint8  `json:"user_role"`
	IsDelete   uint8  `json:"is_delete"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type RegisterReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterResp struct {
	Username string `json:"username"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id          uint64 `json:"id"`
	Username    string `json:"username"`
	AvatarUrl   string `json:"avatarUrl"`
	UserRole    uint8  `json:"userRole"`
	Token       string `json:"token"`
	TokenExpire int64  `json:"tokenExpire"`
}

type CurrentUserReq struct {
	Authorization string `header:"authorization"`
}

type CurrentUserResp struct {
	Id          uint64 `json:"id"`
	Username    string `json:"username"`
	AvatarUrl   string `json:"avatarUrl"`
	UserRole    uint8  `json:"userRole"`
	Token       string `json:"token"`
	TokenExpire int64  `json:"tokenExpire"`
}

type LogoutReq struct {
	Authorization string `header:"authorization"`
}

type LogoutResp struct {
	IsLogouted bool `json:"isLogouted"`
}

type InterfaceInfo struct {
	Id             uint64 `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	RequestParams  string `json:"requestParams"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
	Status         uint8  `json:"status"`
	Method         string `json:"method"`
	UserId         uint64 `json:"userId"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
	IsDelete       uint8  `json:"isDelete"`
}

type InvokeInterfaceInfo struct {
	InterfaceInfoName string `json:"name"`
	TotalNum          uint64 `json:"totalNum"`
}

type AddInterfaceInfoReq struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	RequestParams  string `json:"requestParams"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
	Method         string `json:"method"`
	UserId         uint64 `json:"userId"`
}

type AddInterfaceInfoResp struct {
	Name string `json:"name"`
}

type UpdateInterfaceInfoReq struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	RequestParams  string `json:"requestParams"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
	Method         string `json:"method"`
	Id             uint64 `json:"id"`
}

type UpdateInterfaceInfoResp struct {
	IsUpdated bool `json:"isUpdated"`
}

type DeleteInterfaceInfoReq struct {
	Id uint64 `json:"id"`
}

type DeleteInterfaceInfoResp struct {
	IsDeleted bool `json:"isDeleted"`
}

type GetInterfaceInfoReq struct {
	Id uint64 `form:"id"`
}

type GetInterfaceInfoResp struct {
	Description    string `json:"description"`
	Url            string `json:"url"`
	RequestParams  string `json:"requestParams"`
	RequestHeader  string `json:"requestHeader"`
	ResponseHeader string `json:"responseHeader"`
	Status         uint8  `json:"status"`
	Method         string `json:"method"`
	CreateTime     string `json:"createTime"`
	UpdateTime     string `json:"updateTime"`
}

type PageListReq struct {
	Keyword  string `form:"keyword,optional"`
	Current  uint64 `form:"current"`
	PageSize uint64 `form:"pageSize"`
}

type PageListResp struct {
	Total   uint64          `json:"total"`
	Records []InterfaceInfo `json:"records"`
}

type OnlineInterfaceInfoReq struct {
	Id            uint64 `json:"id"`
	Authorization string `header:"authorization"`
}

type OnlineInterfaceInfoResp struct {
	IsOnline bool `json:"isOnline"`
}

type OfflineInterfaceInfoReq struct {
	Id            uint64 `json:"id"`
	Authorization string `header:"authorization"`
}

type OfflineInterfaceInfoResp struct {
	IsOffline bool `json:"isOffline"`
}

type InvokeInterfaceInfoReq struct {
	Id            uint64 `json:"id"`
	RequestParams string `json:"requestParams"`
	Authorization string `header:"authorization"`
}

type InvokeInterfaceInfoResp struct {
	ResponseObject interface{} `json:"responseObject"`
}

type GetTopNInterfaceInfoReq struct {
	Limit         uint64 `form:"limit"`
	Authorization string `header:"authorization"`
}

type GetTopNInterfaceInfoResp struct {
	Records []InvokeInterfaceInfo `json:"records"`
}
