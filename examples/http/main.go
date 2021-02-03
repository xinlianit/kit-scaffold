package main

import (
	"github.com/gorilla/mux"
	"github.com/spf13/pflag"
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/examples/http/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/http/middleware"
	"github.com/xinlianit/kit-scaffold/examples/http/transport"
	"github.com/xinlianit/kit-scaffold/handler"
	"net/http"
)

func main() {
	commandLineParse()

	httpHandler := NewHttpHandler()

	// 运行服务
	scaffold.RunHttpServer(httpHandler)
}

// 命令行解析
// TODO: 封装到框架层
func commandLineParse() {
	// 解析命令行参数
	pflag.String("env", "PRD","环境名称")
	pflag.String("server.host", "0.0.0.0", "服务地址")
	pflag.Int("server.port", 80, "服务端口")
	pflag.String("server.gateway.host", "0.0.0.0", "网关地址")
	pflag.Int("server.gateway.port", 8080, "网关端口")
	pflag.String("nacos.address", "", "Nacos地址") // TODO: 待启用
	pflag.String("nacos.namespace", "", "Nacos名称空间") // TODO: 待启用
	pflag.String("consul.host", "", "Consul主机") // TODO: 待启用
	pflag.Int("consul.port", 0, "Consul端口") // TODO: 待启用
	pflag.Parse()
	config.Config().BindPFlags(pflag.CommandLine)
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
