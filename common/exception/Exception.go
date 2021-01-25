package exception

import (
	"fmt"
	scaffold "github.com/xinlianit/kit-scaffold"
)

func NewException(code int, message string) Exception {
	return Exception{
		code:    code,
		message: message,
	}
}

func NewExceptionByEnum(enum scaffold.EnumType) Exception {
	return NewException(enum.Value(), enum.Name())
}

// 异常
type Exception struct {
	code    int
	message string
}

func (e Exception) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.code, e.message)
}

func (e *Exception) SetCode(code int) {
	e.code = code
}

func (e Exception) GetCode() int {
	return e.code
}

func (e *Exception) SetMessage(message string) {
	e.message = message
}

func (e Exception) GetMessage() string {
	return e.message
}
