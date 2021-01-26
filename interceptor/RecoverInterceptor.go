package interceptor

import (
	"context"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"github.com/xinlianit/kit-scaffold/common/enum"
	"github.com/xinlianit/kit-scaffold/common/exception"
	"google.golang.org/grpc"
	"log"
)

// 宕机恢复拦截器
func RecoverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 处理请求前
	log.Printf("[before] 宕机恢复拦截器: %v", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout))

	defer func() {
		if rec := recover(); rec != nil {
			// 系统异常
			err = exception.NewException(enum.CodeErrorServer.Value(), enum.CodeErrorServer.Name())
		}
	}()

	// 处理请求
	resp, err = handler(ctx, req)

	// 处理请求后
	log.Printf("[after] 宕机恢复拦截器: %v", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout))

	return resp, err
}
