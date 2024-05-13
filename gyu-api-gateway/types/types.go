package types

type Request struct {
	Username string `json:"username"`
}

type UserResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
