package main

import (
	"github.com/gorilla/mux"
	"github.com/xinlianit/kit-scaffold/core"
	"github.com/xinlianit/kit-scaffold/test/endpoint"
	"github.com/xinlianit/kit-scaffold/test/middleware"
	"github.com/xinlianit/kit-scaffold/test/transport"
	"log"
	"net/http"
)

func main() {
	httpHandler := NewHttpHandler()

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: httpHandler,
	}

	log.Println("test server address: 8080")

	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}

func NewHttpHandler() http.Handler {
	httpHandler := core.NewHttpHandler()
	httpHandler.Use(middleware.TestMiddleware, middleware.Test2Middleware)

	indexEndpoint := endpoint.NewIndexEndpoint()
	indexTransport := transport.NewIndexTransport()
	helloHandler := httpHandler.Server(indexEndpoint.Hello, indexTransport.HelloDecode, indexTransport.HelloEncode)
	testHandler := httpHandler.Server(indexEndpoint.Test, indexTransport.HelloDecode, indexTransport.HelloEncode)

	route := mux.NewRouter()

	//route.Use() // 中间件

	route.Methods(http.MethodGet).Path("/index/hello").Handler(helloHandler)
	route.Methods(http.MethodGet).Path("/index/test").Handler(testHandler)

	return route
}
