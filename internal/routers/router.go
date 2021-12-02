package routers

import (
	// 引入 docs,解决swagger页面 http://127.0.0.1:8080/swagger/index.html 访问后报错 Failed to load spec.
	// 写完注解之后,使用 swag init 命令生成文档
	_ "github.com/diy0663/goblog-service/docs"
	"github.com/diy0663/goblog-service/internal/middleware"
	v1 "github.com/diy0663/goblog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//注册swagger的路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 使用中间件处理语言包,方便后面验证报错的时候根据header头传的 locale 的取值来决定翻译为哪一类语言
	// 从而支持 错误提示多语言的功能
	r.Use(middleware.Translations())

	tag := v1.NewTag()
	article := v1.NewArticle()

	apiv1 := r.Group("/api/v1")
	{
		//新增
		apiv1.POST("/tags", tag.Create)
		// 删除
		apiv1.DELETE("/tags/:id", tag.Delete)
		//更新
		apiv1.PUT("/tags/:id", tag.Update)
		//更新局部
		apiv1.PATCH("/tags/:id/state", tag.Update)
		// 获取列表
		apiv1.GET("/tags", tag.List)
		// 获取详情
		apiv1.GET("/tags/:id", tag.Detail)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Detail)

	}
	return r
}
