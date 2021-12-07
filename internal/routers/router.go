package routers

import (
	// 引入 docs,解决swagger页面 http://127.0.0.1:8080/swagger/index.html 访问后报错 Failed to load spec.
	// 写完注解之后,使用 swag init 命令生成文档
	_ "github.com/diy0663/goblog-service/docs"
	"github.com/diy0663/goblog-service/global"
	"github.com/diy0663/goblog-service/internal/middleware"
	v1 "github.com/diy0663/goblog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		// 非调试模式就启用中间件记录访问日志
		r.Use(middleware.AccessLog())
		// todo 非调试模式,调用自定义的 Recovery处理(例如新增 邮件通知)
		r.Use(middleware.Recovery())
	}

	//注册swagger的路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 使用中间件处理语言包,方便后面验证报错的时候根据header头传的 locale 的取值来决定翻译为哪一类语言
	// 从而支持 错误提示多语言的功能
	r.Use(middleware.Translations())

	tag := v1.NewTag()
	article := v1.NewArticle()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.AppInfo())
	apiv1.Use(middleware.JWT())
	{
		//新增
		apiv1.POST("/tags", tag.Create)
		// 删除 todo 这种带 :id 的 url 在参数验证的时候要提前把id赋值上去
		apiv1.DELETE("/tags/:id", tag.Delete)
		//更新
		apiv1.PUT("/tags/:id", tag.Update)
		// 获取列表
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Detail)

	}
	r.POST("/auth", v1.GetAuth)

	// r.Use(middleware.AccessLog())
	// {
	// r.POST("/auth", v1.GetAuth)
	// }

	// auth token相关路由
	//r.POST("/auth", v1.GetAuth)

	return r
}
