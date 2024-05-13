package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"net/url"
)

func GetRequestIP(c *gin.Context) string {
	reqIp := c.ClientIP()
	if reqIp == "::1" {
		reqIp = "127.0.0.1"
	}
	return reqIp
}

func GetDomainFromReferer(referer string) (string, error) {
	parsedURL, err := url.Parse(referer)
	if err != nil {
		return "", err
	}
	return parsedURL.Hostname(), nil
}

func GetRequestBody(c *gin.Context) map[string]interface{} {
	// gin 读取 request body 的方式
	rawData, err := c.GetRawData()
	if err != nil {
		logc.Errorf(c.Request.Context(), "读取请求内容错误: %v", err)
		return nil
	}
	body := map[string]interface{}{}
	err = json.Unmarshal(rawData, &body)
	if err != nil {
		logc.Errorf(c.Request.Context(), "解析请求内容错误: %v", err)
		return nil
	}
	return body
}
