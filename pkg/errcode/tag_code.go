package errcode

var (
	ErrorGetTagListFail = NewErrorResp(20010001, "获取标签列表失败")
	ErrorCreateTagFail  = NewErrorResp(20010002, "创建标签失败")
	ErrorUpdateTagFail  = NewErrorResp(20010003, "更新标签失败")
	ErrorDeleteTagFail  = NewErrorResp(20010004, "删除标签失败")
	ErrorCountTagFail   = NewErrorResp(20010005, "统计标签失败")
)
