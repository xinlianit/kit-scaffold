package main

import (
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/xinlianit/kit-scaffold/test/endpoint"
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

	indexEndpoint := endpoint.NewIndexEndpoint()
	indexTransport := transport.NewIndexTransport()
	indexHandler := httpTransport.NewServer(indexEndpoint.Hello(), indexTransport.HelloDecode(), indexTransport.HelloEncode())

	route := mux.NewRouter()

	//route.Use() // 中间件


	route.Methods(http.MethodGet).Path("/index/hello").Handler(indexHandler)

	return route
}
