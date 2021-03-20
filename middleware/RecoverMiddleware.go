package middleware

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/xinlianit/kit-scaffold/logger"
)

// 宕机恢复中间件
func RecoverMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("宕机恢复: %v", rec)

				logger.ZapLogger.Error(fmt.Sprint(rec))
			}
		}()

		// 处理请求
		response, err = next(ctx, request)

		return response, err
	}
}
