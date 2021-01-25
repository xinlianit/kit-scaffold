package exception

import (
	scaffold "github.com/xinlianit/kit-scaffold"
)

// 服务异常
type ServiceException struct {
	Exception
}

func NewServiceException(enum scaffold.EnumType) ServiceException {
	e := ServiceException{}
	e.SetCode(enum.Value())
	e.SetMessage(enum.Name())
	return e
}
