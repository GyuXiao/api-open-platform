package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"gyu-api-gateway/config"
	"gyu-api-gateway/global"
	"sync"
)

var IPLimiter *RateLimiter

type RateLimiter struct {
	// sync.Map 比较适合读多写少的场景
	limiter sync.Map
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{limiter: sync.Map{}}
}

func (l *RateLimiter) GetLimiter(ip string, conf config.Config) *rate.Limiter {
	limiter, ok := l.limiter.Load(ip)
	if !ok {
		limiter = rate.NewLimiter(rate.Limit(conf.RateLimit.Request), conf.RateLimit.BucketSize)
		l.limiter.Store(ip, limiter)
	}
	return limiter.(*rate.Limiter)
}

func RateLimiterMiddleware(conf config.Config) gin.HandlerFunc {

	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := IPLimiter.GetLimiter(ip, conf)

		if !limiter.Allow() {
			global.HandlerExceedLimit(c)
			return
		}

		c.Next()
	}
}
