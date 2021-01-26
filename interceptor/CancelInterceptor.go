package interceptor

import (
	"context"
	"github.com/xinlianit/go-util"
	"github.com/xinlianit/kit-scaffold/common/constant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// 请求取消拦截器
func CancelInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 处理请求前
	log.Printf("[before] 请求取消拦截器: %v", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout))

	// 请求取消
	if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Canceled, "请求已取消")
	}

	// 处理请求
	resp, err := handler(ctx, req)

	// 处理请求后
	log.Printf("[after] 请求取消拦截器: %v", util.TimeUtil().GetCurrentDateTime(constant.DefaultTimeLayout))
	return resp, err
}
