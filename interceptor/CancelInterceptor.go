package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 请求取消拦截器
func CancelInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 请求取消
	if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
		return nil, status.Error(codes.Canceled, "请求已取消")
	}

	// 处理请求
	resp, err := handler(ctx, req)

	return resp, err
}
