package interceptor

import (
	"context"
	"github.com/xinlianit/kit-scaffold/common/exception"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorInterceptor 错误拦截器
func ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 处理请求
	resp, err = handler(ctx, req)

	if err != nil {
		// 错误类型
		switch errException := err.(type) {
		// 异常错误
		case exception.Exception:
			return resp, status.Error(codes.Code(errException.GetCode()), errException.GetMessage())
		// Service 异常
		case exception.ServiceException:
			return resp, status.Error(codes.Code(errException.GetCode()), errException.GetMessage())
		default:
			return resp, status.Error(codes.Unknown, errException.Error())
		}
	}

	return resp, err
}
