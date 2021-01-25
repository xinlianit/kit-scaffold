package interceptor

import (
	"context"
	"github.com/xinlianit/kit-scaffold/common/exception"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 错误拦截器
func ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 处理请求
	resp, err := handler(ctx, req)

	if err != nil {
		// 错误类型
		switch errException := err.(type) {
		// 公共错误
		case exception.CommonException:
			return resp, status.Error(codes.Code(errException.GetCode()), errException.GetMessage())
		default:
			return resp, status.Error(codes.Unknown, errException.Error())
		}
	}
	return resp, err
}
