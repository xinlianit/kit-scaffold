package interceptor

import (
	"google.golang.org/grpc"
)

// DefaultUnaryServerInterceptor 默认一元服务拦截器
func DefaultUnaryServerInterceptor() []grpc.UnaryServerInterceptor {
	// 拦截器
	return []grpc.UnaryServerInterceptor{
		// 访问拦截器
		AccessInterceptor,
		// 宕机恢复拦截器
		RecoverInterceptor,
		// 错误拦截器
		ErrorInterceptor,
		// 请求取消
		CancelInterceptor,
		// 认证拦截器
		AuthInterceptor,
	}
}
