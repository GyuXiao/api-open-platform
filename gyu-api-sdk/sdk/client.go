package sdk

import (
	"github.com/duke-git/lancet/v2/random"
	"gyu-api-sdk/sdk/logs"
	"gyu-api-sdk/sdk/request"
	"gyu-api-sdk/sdk/response"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Config *Config
	Logger logs.Logger
}

func (c *Client) Init() *Client {
	c.Logger = logs.New()
	return c
}

func (c *Client) WithConfig(config *Config) *Client {
	c.Config = config
	return c
}

func (c *Client) Send(req request.Request, resp response.Response) error {
	method := req.GetMethod()
	url := req.GetURL()
	body := req.GetBody()
	rawResponse, err := c.doSend(method, url, body)
	if err != nil {
		c.Logger.Errorf("客户端发起请求错误: %s", err.Error())
		return err
	}
	return response.ParseFromHttpResponse(rawResponse, resp)
}

func (c *Client) doSend(method, url, body string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		c.Logger.Errorf("发起请求错误: %s", err.Error())
		return nil, err
	}
	c.SetHeaders(c.Config.AccessKey, c.Config.SecretKey, body, req)
	return client.Do(req)
}

func (c *Client) SetHeaders(accessKey, secretKey, requestBody string, req *http.Request) {
	// 随机数: 一个长度为 100 的随机数字的字符串
	nonce := random.RandNumeral(100)
	// 当前时间戳（秒级别）
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	// 签名
	paramsMap := map[string]string{
		"title0": accessKey,
		"title1": secretKey,
		"title2": nonce,
		"title3": timestamp,
		"title4": requestBody,
	}
	signature := GenSign(paramsMap, secretKey)
	// 设置请求头
	req.Header.Set("accessKey", accessKey)
	req.Header.Set("nonce", nonce)
	req.Header.Set("timestamp", timestamp)
	req.Header.Set("sign", signature)
	req.Header.Set("body", requestBody)
	req.Header.Set("Content-Type", "application/json")
}
