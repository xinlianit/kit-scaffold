package exception

import (
	"fmt"
	scaffold "github.com/xinlianit/kit-scaffold"
)

// NewException 创建全局异常
func NewException(code int, message string) Exception {
	return Exception{
		code:    code,
		message: message,
	}
}

// NewExceptionByEnum - 创建全局异常(通过枚举)
func NewExceptionByEnum(enum scaffold.EnumType) Exception {
	return NewException(enum.Value(), enum.Name())
}

// Exception 全局异常
type Exception struct {
	code    int
	message string
}

// 实现错误类型接口
func (e Exception) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.code, e.message)
}

// SetCode 设置异常状态码
func (e *Exception) SetCode(code int) {
	e.code = code
}

// GetCode 获取异常状态码
func (e Exception) GetCode() int {
	return e.code
}

// SetMessage 设置异常信息
func (e *Exception) SetMessage(message string) {
	e.message = message
}

// GetMessage 获取异常信息
func (e Exception) GetMessage() string {
	return e.message
}
