package app

import (
	"github.com/diy0663/goblog-service/global"
	"github.com/diy0663/goblog-service/pkg/convert"
	"github.com/gin-gonic/gin"
)

// 分页相关,从上下文中获取 page参数,没有的话就默认1
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}

	return page
}

// 从上下文中获取 page_size参数,没有的话就默认用系统配置的
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pageSize
}

// 计算偏移量
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
