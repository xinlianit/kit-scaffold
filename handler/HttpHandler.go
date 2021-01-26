package handler

import (
	"github.com/go-kit/kit/endpoint"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/xinlianit/kit-scaffold/common"
	"github.com/xinlianit/kit-scaffold/middleware"
)

// 创建 Http 处理器
func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

// http 处理器
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

func (h HttpHandler) Server(e endpoint.Endpoint, dec httpTransport.DecodeRequestFunc) *httpTransport.Server {
	// 业务中间件
	if h.middlewares != nil && len(h.middlewares) > 0 {
		for _, middleware := range h.middlewares {
			e = middleware(e)
		}
	}

	// 请求中间件
	e = middleware.RequestMiddleware(e)
	// 日志中间件
	e = middleware.LoggerMiddleware(e)
	// 宕机恢复中间件
	e = middleware.RecoverMiddleware(e)

	return httpTransport.NewServer(e, common.RequestDecode(dec), common.ResponseEncode, h.options...)
}
