package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
)

// 恢复中间件
func RecoverMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// 处理器
		//defer handler(ctx, request)

		response, err = next(ctx, request)

		if err != nil {
			rsp := map[string]interface{}{
				"code": 1,
				"msg": err.Error(),
			}
			return rsp, nil
		}

		log.Printf("--------response: %v, err: %v", response, err)

		return
	}
}

//恢复处理器
func handler(ctx context.Context, request interface{}) {
	if err := recover(); err != nil {
		log.Printf("程序崩溃了: %v", err)
	}
}
