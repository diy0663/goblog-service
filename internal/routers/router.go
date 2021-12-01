package routers

import (
	v1 "github.com/diy0663/goblog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
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
