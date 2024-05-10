package utils

import (
	"github.com/gin-gonic/gin"
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
