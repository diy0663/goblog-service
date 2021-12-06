package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/diy0663/goblog-service/pkg/app"
	"github.com/diy0663/goblog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		var token string
		var ecode = errcode.Success

		// query参数有就拿query的 ,否则拿header 的
		if s, exists := c.GetQuery("token"); exists {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			// 校验token
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				// 过期
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					//认为token不对
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		// token 验证通过,继续往下走
		c.Next()

	}
}
