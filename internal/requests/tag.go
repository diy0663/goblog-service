package requests

import "github.com/gin-gonic/gin"

// 标签相关的自己的表单验证

// 定义一系列对应接口的参数验证规则
// 统计之类的请求
type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

//列表接口请求参数
type TagListRequest struct {
	// 绑定请求的参数名以及验证规则 ()
	// gin 的 模型绑定和验证默认使用的是 go-playground/validator
	Name string `form:"name" binding:"max=100"`
	// 参数集内的其中之一
	State uint8 `form:"state,default=1"  binding:"oneof= 0 1"`
}

type CreateTagRequest struct {
	Name string `form:"name" binding:"required,min=3,max=100"`
	// required 代表必传
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint64 `form:"id" binding:"gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	// gte代表 大于等于
	ID uint64 `form:"id" binding:"required,gte=1"`
}

// // 用来验证请求tag列表的API 的参数是否正确
// func ValidTagListRequest(c *gin.Context) (valid bool, errs ValidErrors) {
// 	param := new(TagListRequest)

// 	valid, errs = BindAndValid(c, param)
// 	return
// }

func ValidTagCreateRequest(c *gin.Context) (valid bool, errs ValidErrors) {
	param := new(CreateTagRequest)

	valid, errs = BindAndValid(c, param)
	return
}
