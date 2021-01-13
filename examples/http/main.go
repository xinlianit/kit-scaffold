package main

import (
	"fmt"
	"github.com/gorilla/mux"
	scaffold "github.com/xinlianit/kit-scaffold"
	"github.com/xinlianit/kit-scaffold/config"
	"github.com/xinlianit/kit-scaffold/examples/http/endpoint"
	"github.com/xinlianit/kit-scaffold/examples/http/middleware"
	"github.com/xinlianit/kit-scaffold/examples/http/transport"
	"net/http"
)

func main() {
	httpHandler := NewHttpHandler()

	// 服务地址
	serverAddress := fmt.Sprintf("%s:%d", config.Config().GetString("server.host"), config.Config().GetInt("server.port"))

	// 运行服务
	scaffold.RunHttpServer(serverAddress, httpHandler)
}

func NewHttpHandler() http.Handler {
	httpHandler := scaffold.NewHttpHandler()
	httpHandler.Use(middleware.TestMiddleware, middleware.Test2Middleware)

	indexEndpoint := endpoint.NewIndexEndpoint()
	indexTransport := transport.NewIndexTransport()
	helloHandler := httpHandler.Server(indexEndpoint.Hello, indexTransport.HelloDecode, indexTransport.HelloEncode)
	testHandler := httpHandler.Server(indexEndpoint.Test, indexTransport.HelloDecode, indexTransport.HelloEncode)

	route := mux.NewRouter()

	route.Methods(http.MethodGet).Path("/index/hello").Handler(helloHandler)
	route.Methods(http.MethodGet).Path("/index/test").Handler(testHandler)

	return route
}
