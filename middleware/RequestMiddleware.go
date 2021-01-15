package middleware

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/xinlianit/kit-scaffold/model"
	"github.com/xinlianit/kit-scaffold/server"
)

// 请求中间件
func RequestMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// 请求体
		server.Request, _ = request.(model.Request)
		return next(ctx, server.Request.GetRequestBody())
	}
}
