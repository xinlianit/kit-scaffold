package interceptor

import (
	"github.com/xinlianit/kit-scaffold/interceptor"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor 一元拦截器
func UnaryServerInterceptor() []grpc.UnaryServerInterceptor {
	// 默认拦截器
	defaultInterceptors := interceptor.DefaultUnaryServerInterceptor()
	// 应用拦截器
	interceptors := []grpc.UnaryServerInterceptor{
	}
	return append(defaultInterceptors, interceptors...)
}
