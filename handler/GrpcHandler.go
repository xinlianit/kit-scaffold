package handler

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/xinlianit/kit-scaffold/middleware"
)

type GrpcHandler struct {
	options     []grpc.ServerOption
	middlewares []endpoint.Middleware
}

func (h *GrpcHandler) Options(options ...grpc.ServerOption) *GrpcHandler {
	h.options = options
	return h
}

func (h *GrpcHandler) Use(middlewares ...endpoint.Middleware) *GrpcHandler {
	for _, m := range middlewares {
		h.middlewares = append([]endpoint.Middleware{m}, h.middlewares...)
	}
	return h
}

func (h GrpcHandler) Server(e endpoint.Endpoint, dec grpc.DecodeRequestFunc, enc grpc.EncodeResponseFunc) *grpc.Server {
	// 日志中间件
	h.middlewares = append(h.middlewares, middleware.LoggerMiddleware)

	// 业务中间件
	if h.middlewares != nil && len(h.middlewares) > 0 {
		for _, middleware := range h.middlewares {
			e = middleware(e)
		}
	}

	// 日志中间件
	e = middleware.LoggerMiddleware(e)

	server := grpc.NewServer(e, dec, enc, h.options...)

	return server
}
