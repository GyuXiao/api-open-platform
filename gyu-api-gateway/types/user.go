package types

type UserRequest struct {
	UserName string `form:"username"`
}

type UserResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
