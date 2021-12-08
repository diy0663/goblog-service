package middleware

import (
	"fmt"
	"time"

	"github.com/diy0663/goblog-service/global"
	"github.com/diy0663/goblog-service/pkg/app"
	"github.com/diy0663/goblog-service/pkg/email"
	"github.com/diy0663/goblog-service/pkg/errcode"
	"github.com/diy0663/goblog-service/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 自定义 Recovery 中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {

		defaultMailer := email.NewEmail(&email.SMTPInfo{
			Host:     global.EmailSetting.Host,
			Port:     global.EmailSetting.Port,
			IsSSL:    global.EmailSetting.IsSSL,
			UserName: global.EmailSetting.UserName,
			Password: global.EmailSetting.Password,
			From:     global.EmailSetting.From,
		})

		defer func() {
			if err := recover(); err != nil {
				logger.ZapLog.Panic("panic recover err", zap.Any("panic", err))

				// 发邮件
				err := defaultMailer.SendMail(global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err))

				if err != nil {
					logger.ZapLog.Error("email send err", zap.Any("error", err))
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}

}
