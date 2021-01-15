package handler

import (
	"github.com/go-kit/kit/endpoint"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/xinlianit/kit-scaffold/middleware"
)

type HttpHandler struct {
	options     []httpTransport.ServerOption
	middlewares []endpoint.Middleware
}

func (h *HttpHandler) Options(options ...httpTransport.ServerOption) *HttpHandler {
	h.options = options
	return h
}

func (h *HttpHandler) Use(middlewares ...endpoint.Middleware) *HttpHandler {
	for _, m := range middlewares {
		h.middlewares = append([]endpoint.Middleware{m}, h.middlewares...)
	}
	return h
}

func (h HttpHandler) Server(e endpoint.Endpoint, dec httpTransport.DecodeRequestFunc, enc httpTransport.EncodeResponseFunc) *httpTransport.Server {
	// 业务中间件
	if h.middlewares != nil && len(h.middlewares) > 0 {
		for _, middleware := range h.middlewares {
			e = middleware(e)
		}
	}

	// 日志中间件
	e = middleware.LoggerMiddleware(e)

	// 请求编码
	server := httpTransport.NewServer(e, dec, enc, h.options...)

	return server
}
