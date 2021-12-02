package errcode

import (
	"fmt"
	"net/http"
)

type ErrorResp struct {
	// 需要大写,不然会提示  struct field code has json tag but is not exportedstructtag
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Detail []string `json:"detail"`
}

// 用于加载存放现有已定义的错误码
var codes = map[int]string{}

func NewErrorResp(code int, msg string) *ErrorResp {
	// 排重校验
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &ErrorResp{Code: code, Msg: msg}
}

func (e *ErrorResp) ErrorRespToString() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.GetCode(), e.GetMsg())
}

func (e *ErrorResp) GetCode() int {
	return e.Code
}
func (e *ErrorResp) GetMsg() string {
	return e.Msg
}

func (e *ErrorResp) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *ErrorResp) GetDetail() []string {
	return e.Detail
}

func (e *ErrorResp) WithDetails(details ...string) *ErrorResp {
	newErr := *e
	//初始化
	newErr.Detail = []string{}
	newErr.Detail = append(newErr.Detail, details...)
	return &newErr
}

// 返回自定义的错误码对应的http status code
func (e *ErrorResp) StatusCode() int {
	switch e.Code {
	case Success.GetCode():
		return http.StatusOK
	case ServerError.GetCode():
		return http.StatusInternalServerError
	case InvalidParams.GetCode():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.GetCode():
		fallthrough
	case UnauthorizedTokenError.GetCode():
		fallthrough
	case UnauthorizedTokenGenerate.GetCode():
		fallthrough
	case UnauthorizedTokenTimeout.GetCode():
		return http.StatusUnauthorized
	case TooManyRequests.GetCode():
		return http.StatusTooManyRequests

	}

	return http.StatusInternalServerError

}
