package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/pflag"
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/examples/http/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/http/middleware"
	"github.com/xinlianit/kit-scaffold/examples/http/transport"
	"github.com/xinlianit/kit-scaffold/handler"
	"net/http"
)

func main() {
	// 初始化命令行
	//commandLineInit()

	httpHandler := NewHttpHandler()

	// 运行服务
	scaffold.RunHTTPServer(httpHandler)
}

// 命令行初始化
func commandLineInit() {
	// 命令行参数
	pflag.String("test.app.id", "","测试APP_ID")

	// 命令行解析
	scaffold.CommandLineParse()
}

func NewHttpHandler() http.Handler {
	httpHandler := handler.NewHttpHandler()
	httpHandler.Use(middleware.TestMiddleware, middleware.Test2Middleware)

	indexEndpoint := endpoint.NewIndexEndpoint()
	indexTransport := transport.NewIndexTransport()
	helloHandler := httpHandler.Server(indexEndpoint.Hello, indexTransport.HelloDecode)
	testHandler := httpHandler.Server(indexEndpoint.Test, indexTransport.HelloDecode)

	route := mux.NewRouter()

	route.Methods(http.MethodGet).Path("/index/hello").Handler(helloHandler)
	route.Methods(http.MethodGet).Path("/index/test").Handler(testHandler)

	return route
}
