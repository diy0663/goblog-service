package errcode

// 基础错误码
var (
	Success                   = NewErrorResp(0, "成功")
	ServerError               = NewErrorResp(10000000, "服务内部错误")
	InvalidParams             = NewErrorResp(10000001, "入参错误")
	NotFound                  = NewErrorResp(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewErrorResp(10000003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewErrorResp(10000004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewErrorResp(10000005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewErrorResp(10000006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewErrorResp(10000007, "请求过多")
)
