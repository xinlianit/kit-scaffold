package transport

import (
	"context"
	"log"
)

// 错误处理器
type ErrorHandler struct {

}

// 错误处理
func (h ErrorHandler) Handle(ctx context.Context, err error)  {
	log.Printf("异常错误: %#v", err)
}

// NewHealthTransport 健康检查传输
func NewHealthTransport() HealthTransport {
	return HealthTransport{}
}