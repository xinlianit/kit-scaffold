package interceptor

import (
	"google.golang.org/grpc"
)

// 默认一元服务拦截器
func DefaultUnaryServerInterceptor() []grpc.UnaryServerInterceptor {
	// 拦截器
	return []grpc.UnaryServerInterceptor{
		// 错误拦截器
		ErrorInterceptor,
		// 宕机恢复拦截器
		RecoverInterceptor,
		// 请求取消
		CancelInterceptor,
	}
}
