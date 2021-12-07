package middleware

import (
	"github.com/diy0663/goblog-service/pkg/app"
	"github.com/diy0663/goblog-service/pkg/errcode"
	"github.com/diy0663/goblog-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 自定义 Recovery 中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.ZapLog.Panic("panic recover err", zap.Any("panic", err))
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}

}
