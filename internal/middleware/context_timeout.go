package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// 超时处理的中间件
func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置当前 context 的超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		// 重新赋予给了 gin.Context
		c.Request = c.Request.WithContext(ctx)
		c.Next()

	}
}
