package middleware

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"log"
)

// 宕机恢复中间件
func RecoverMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// 处理请求前
		log.Printf("[before] 宕机恢复中间件")
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("宕机恢复: %v", rec)
			}
		}()

		// 处理请求
		response, err = next(ctx, request)

		// 处理请求后
		log.Printf("[after] 宕机恢复中间件")
		return response, err
	}
}
