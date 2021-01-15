package common

import "github.com/xinlianit/kit-scaffold/model"

var Response response

func init() {
	Response = response{}
}

// 响应
type response struct {
}

// 成功
func (r response) Success(data interface{}) model.Response {
	return model.Response{Code: 0, Data: data}
}

// 失败
func (r response) Fail(msg string) model.Response {
	return model.Response{Code: 1, Msg: msg}
}

// 错误
func (r response) Error(code int, msg string) model.Response {
	return model.Response{Code: code, Msg: msg}
}
