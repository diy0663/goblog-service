package middleware

import (
	"github.com/gin-gonic/gin"
)

// 上下文数据设置传递的中间件, 可以把这个中间件应用到特定的api上面去

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog_service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
