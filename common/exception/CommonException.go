package exception

import (
	"fmt"
	scaffold "github.com/xinlianit/kit-scaffold"
)

func NewCommonException(code int, message string) CommonException {
	return CommonException{
		code:    code,
		message: message,
	}
}

func NewCommonExceptionByEnum(enum scaffold.EnumType) CommonException {
	return NewCommonException(enum.Value(), enum.Name())
}

// 公共错误
type CommonException struct {
	code    int
	message string
}

func (e CommonException) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.code, e.message)
}

func (e CommonException) GetCode() int {
	return e.code
}

func (e CommonException) GetMessage() string {
	return e.message
}
