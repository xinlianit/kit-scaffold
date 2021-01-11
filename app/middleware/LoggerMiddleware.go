package middleware

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

// 日志中间件
func LoggerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("=========== 日志中间件")
		// todo 记录日志
		return next(ctx, request)
	}
}
