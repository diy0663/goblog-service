package middleware

import (
	"github.com/diy0663/goblog-service/pkg/app"
	"github.com/diy0663/goblog-service/pkg/errcode"
	"github.com/diy0663/goblog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

// 限流中间件

//
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {

		key := l.Key(c)
		bucket, ok := l.GetBucket(key)

		if ok {

			// 存在针对本uri的对应的限流规则的时候
			// 请求那一个令牌,返回 0 说明用完
			count := bucket.TakeAvailable(1)
			// fmt.Println(count) 调试时  这个count 返回了1
			if count == 0 {
				// 没可用的了,就直接报错说太多请求了
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
