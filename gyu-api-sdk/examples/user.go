package main

import (
	"encoding/json"
	"gyu-api-sdk/sdk"
	"gyu-api-sdk/sdk/request"
	"gyu-api-sdk/sdk/response"
	"gyu-api-sdk/service/user"
)

var accessKey string = "gyu"
var secretKey string = "abcdefg"

func main() {
	// 使用 config 里的配置创建 client
	config := sdk.NewConfig(accessKey, secretKey)
	client, err := user.NewClient(config)
	if err != nil {
		client.Logger.Errorf("客户端创建错误: %v", err)
		panic(err)
	}
	userTest := NewUser("user_test1")
	userJson, _ := json.Marshal(userTest)
	// 创建请求
	req := &request.BaseRequest{
		URL:    "http://127.0.0.1:8123/api/user",
		Method: "POST",
		Header: nil,
		Body:   string(userJson),
	}
	// 通过 client 发起请求
	err = client.Send(req, &response.BaseResponse{})
	if err != nil {
		client.Logger.Errorf("客户端请求错误: %v", err)
	}
}

type User struct {
	Username string `json:"username"`
}

func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}
