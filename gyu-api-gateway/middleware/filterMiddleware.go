package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-gateway/global"
	"gyu-api-gateway/utils"
)

// WhiteIPList 白名单
var WhiteIPList = []string{"127.0.0.1"}

func FilterWithAccessControlInWhiteIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestIP := utils.GetRequestIP(c)

		// 先简单一点，后面改为动态加载 config 文件
		ok := false
		for _, ip := range WhiteIPList {
			if ip == requestIP {
				ok = true
				break
			}
		}

		filterInWhiteIPLog := logx.ContextWithFields(
			c.Request.Context(),
			logx.Field("requestIP", requestIP),
			logx.Field("isPass", ok),
		)
		logc.Info(filterInWhiteIPLog, "gyu-api-gateway ip 白名单")

		if ok {
			c.Next()
		} else {
			global.HandlerForbidden(c)
			return
		}
	}
}
