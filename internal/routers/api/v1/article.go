package v1

import (
	"github.com/diy0663/goblog-service/pkg/app"
	"github.com/diy0663/goblog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// 不需要返回值??
func (a Article) Get(c *gin.Context) {}
func (a Article) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)

}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
func (a Article) Detail(c *gin.Context) {}
