package middleware

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/xinlianit/kit-scaffold/model"
)

// 请求中间件
func RequestMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// 请求体
		requestBody, ok := request.(model.Request)
		if !ok {
			return nil, errors.New("Request 错误")
		}

		return next(ctx, requestBody.GetData())
	}
}
