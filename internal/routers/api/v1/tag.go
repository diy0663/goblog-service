package v1

import (
	"github.com/diy0663/goblog-service/internal/requests"
	"github.com/diy0663/goblog-service/internal/service"
	"github.com/diy0663/goblog-service/pkg/app"
	"github.com/diy0663/goblog-service/pkg/convert"
	"github.com/diy0663/goblog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// 感觉更像是控制器
type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.ErrorResp "请求错误"
// @Failure 500 {object} errcode.ErrorResp "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {

	// 初始化空的列表请求 参数结构体
	param := requests.TagListRequest{}

	// 初始化一个响应对象
	response := app.NewResponse(c)

	// 假如参数验证通过, 就会把正确的入参回写到 param
	valid, errs := requests.BindAndValid(c, &param)

	if !valid {
		//返回 false ,验证不通过
		//global.Logger.Errorf("ValidTagListRequest errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Error()))
		return
	}
	// 求总数 (有了总数才好分页)

	svc := service.New(c.Request.Context())
	pager := app.Pager{
		Page:     app.GetPage(c),
		PageSize: app.GetPageSize(c),
		//TotalRows: 0,
	}
	TotalRows, err := svc.CountTag(&requests.CountTagRequest{
		Name:  param.Name,
		State: param.State,
	})
	if err != nil {
		// 求总数出错
		// 写错误日志
		//global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
	}
	pager.TotalRows = int(TotalRows)
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		//global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, int(TotalRows))
	// 到这里暂时算正常返回, H is a shortcut for map[string]interface{}

}

func (t Tag) Create(c *gin.Context) {
	// 初始化空的列表请求 参数结构体
	param := requests.CreateTagRequest{}

	// 初始化一个响应对象
	response := app.NewResponse(c)

	// 假如参数验证通过, 就会把正确的入参回写到 param
	valid, errs := requests.BindAndValid(c, &param)

	if !valid {
		//返回 false ,验证不通过
		//	global.Logger.Errorf("ValidCreateTagRequest errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Error()))
		return
	}

	// 这里有问题
	svc := service.New(c.Request.Context())

	// todo 在这里面有问题
	err := svc.CreateTag(&param)

	if err != nil {

		//global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})

}
func (t Tag) Update(c *gin.Context) {
	// curl -X PUT http://127.0.0.1:8080/api/v1/tags/3 -F state=0 -F modified_by=eddycjy
	// id 是附加在url上的,需要通过 c.Param("id") 去获取
	param := requests.UpdateTagRequest{
		ID: uint64(convert.StrTo(c.Param("id")).MustUInt32()),
	}
	response := app.NewResponse(c)
	valid, errs := requests.BindAndValid(c, &param)
	if !valid {
		//返回 false ,验证不通过
		//	global.Logger.Errorf("ValidUpdateTagRequest errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Error()))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		//global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})

}
func (t Tag) Delete(c *gin.Context) {
	param := requests.DeleteTagRequest{
		ID: uint64(convert.StrTo(c.Param("id")).MustUInt32()),
	}
	response := app.NewResponse(c)
	valid, errs := requests.BindAndValid(c, &param)
	if !valid {
		// global.Logger.Errorf("ValidDeleteTagRequest errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Error()))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		//global.Logger.Errorf("svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})

}
