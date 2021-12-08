package v1

import (
	"github.com/diy0663/goblog-service/internal/requests"
	"github.com/diy0663/goblog-service/internal/service"
	"github.com/diy0663/goblog-service/pkg/app"
	"github.com/diy0663/goblog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 一个单独的对应post: /auth路由  的处理方法
func GetAuth(c *gin.Context) {

	//  手动验证panic 是否会被recovery中间件处理并且触发发邮件的动作
	//panic("验证是否发了邮件")
	// 校验参数
	param := requests.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := requests.BindAndValid(c, &param)
	if !valid {
		// 参数验证不通过
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	//去数据库查有没有对应记录
	err := svc.CheckAuth(&param)
	if err != nil {
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		//生成token的时候出错
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})

}
