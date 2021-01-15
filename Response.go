package scaffold

import "github.com/xinlianit/kit-scaffold/model"

func Response() response {
	return response{}
}

// 响应
type response struct {
}

// 成功
func (r response) Success(data interface{}) model.Response {
	return model.Response{Code: 0, Data: data}
}

// 失败
func Fail(msg string) model.Response {
	return model.Response{Code: 1, Msg: msg}
}

// 错误
func Error(code int, msg string) model.Response {
	return model.Response{Code: code, Msg: msg}
}
