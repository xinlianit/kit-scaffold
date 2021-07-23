package interceptor

import (
	"context"
	"github.com/xinlianit/kit-scaffold/common/enum"
	"github.com/xinlianit/kit-scaffold/common/exception"
	"github.com/xinlianit/kit-scaffold/logger"
	"google.golang.org/grpc"
)

// RecoverInterceptor 宕机恢复拦截器
func RecoverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			// 记录错误日志
			logger.ZapLogger.Sugar().Error(rec)
			// 系统异常
			err = exception.NewException(enum.CodeErrorServer.Value(), enum.CodeErrorServer.Name())
		}
	}()

	// 处理请求
	resp, err = handler(ctx, req)

	return resp, err
}
