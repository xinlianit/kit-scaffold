package interceptor

import (
	"context"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"github.com/xinlianit/kit-scaffold/common/exception"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// 错误拦截器
func ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 处理请求前
	log.Printf("[before] 错误拦截器: %v", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout))

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

	// 处理请求后
	log.Printf("[after] 错误拦截器: %v", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout))

	return resp, err
}
